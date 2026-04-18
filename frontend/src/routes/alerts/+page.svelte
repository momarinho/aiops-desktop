<script lang="ts">
	import { onMount } from 'svelte';
	import { acknowledgeAlert, getAlerts, silenceAlert } from '$lib/api';
	import type { Alert, AlertSeverity, AlertStatus } from '$lib/api/types';

	type AlertAction = 'acknowledge' | 'silence';

	const POLL_INTERVAL_MS = 5000;

	let alerts = $state<Alert[]>([]);
	let isLoading = $state(true);
	let isRefreshing = $state(false);
	let error = $state<string | null>(null);
	let lastUpdated = $state<string | null>(null);
	let pendingActions = $state<Record<string, AlertAction | undefined>>({});

	const activeAlerts = $derived(alerts.filter((alert) => alert.status !== 'resolved'));
	const criticalCount = $derived(activeAlerts.filter((alert) => alert.severity === 'critical').length);
	const warningCount = $derived(activeAlerts.filter((alert) => alert.severity === 'warning').length);
	const systemHealth = $derived(getSystemHealth(criticalCount, warningCount));

	onMount(() => {
		let isDisposed = false;

		void loadAlerts({ background: false });

		const interval = window.setInterval(() => {
			if (!isDisposed) {
				void loadAlerts({ background: true });
			}
		}, POLL_INTERVAL_MS);

		return () => {
			isDisposed = true;
			window.clearInterval(interval);
		};
	});

	async function loadAlerts({ background }: { background: boolean }) {
		if (background) {
			isRefreshing = true;
		} else {
			isLoading = true;
		}

		try {
			alerts = await getAlerts();
			error = null;
			lastUpdated = new Date().toISOString();
		} catch (err) {
			error = err instanceof Error ? err.message : 'Failed to load alerts';
		} finally {
			isLoading = false;
			isRefreshing = false;
		}
	}

	async function handleAction(alertId: string, action: AlertAction) {
		pendingActions = { ...pendingActions, [alertId]: action };

		try {
			const updatedAlert =
				action === 'acknowledge' ? await acknowledgeAlert(alertId) : await silenceAlert(alertId);

			alerts = alerts.map((alert) => (alert.id === alertId ? updatedAlert : alert));
			error = null;
			lastUpdated = new Date().toISOString();
		} catch (err) {
			error = err instanceof Error ? err.message : `Failed to ${action} alert`;
		} finally {
			const nextPending = { ...pendingActions };
			delete nextPending[alertId];
			pendingActions = nextPending;
		}
	}

	function getSystemHealth(critical: number, warning: number) {
		if (critical > 0) {
			return {
				label: 'DEGRADED',
				detail: `${critical} critical alert${critical === 1 ? '' : 's'} active`,
				accent: '#ffb4ab'
			};
		}

		if (warning > 0) {
			return {
				label: 'WATCH',
				detail: `${warning} warning alert${warning === 1 ? '' : 's'} active`,
				accent: '#ffd1af'
			};
		}

		return {
			label: 'STABLE',
			detail: 'No active alert conditions',
			accent: '#6feee1'
		};
	}

	function getSeverityColor(severity: AlertSeverity) {
		return severity === 'critical' ? '#ffb4ab' : '#ffd1af';
	}

	function getStatusColor(status: AlertStatus) {
		switch (status) {
			case 'firing':
				return '#ffb4ab';
			case 'acknowledged':
				return '#6feee1';
			case 'silenced':
				return '#c3c6d1';
			case 'resolved':
				return '#5adace';
		}
	}

	function formatStatus(status: AlertStatus) {
		return status.replace('_', ' ').toUpperCase();
	}

	function formatMetricName(metricName: string) {
		return metricName
			.replace('_usage_', ' ')
			.replace(/_/g, ' ')
			.replace(/\b\w/g, (char) => char.toUpperCase());
	}

	function formatAlertTime(alert: Alert) {
		return formatTimestamp(alert.started_at ?? alert.updated_at);
	}

	function formatTimestamp(value?: string) {
		if (!value) {
			return 'N/A';
		}

		return new Intl.DateTimeFormat(undefined, {
			month: 'short',
			day: '2-digit',
			hour: '2-digit',
			minute: '2-digit'
		}).format(new Date(value));
	}

	function formatPercent(value: number) {
		return `${value.toFixed(1)}%`;
	}

	function canAcknowledge(alert: Alert) {
		return alert.status === 'firing';
	}

	function canSilence(alert: Alert) {
		return alert.status === 'firing' || alert.status === 'acknowledged';
	}

	function isActionPending(alertId: string, action: AlertAction) {
		return pendingActions[alertId] === action;
	}

	const summaryCards = $derived([
		{
			label: 'Critical Alerts',
			value: String(criticalCount).padStart(2, '0'),
			delta:
				criticalCount > 0
					? 'Immediate attention required'
					: 'No critical conditions active',
			accent: '#ffb4ab'
		},
		{
			label: 'Warnings',
			value: String(warningCount).padStart(2, '0'),
			delta: warningCount > 0 ? 'Monitoring non-critical pressure' : 'No warning conditions active',
			accent: '#ffd1af'
		},
		{
			label: 'System Health',
			value: systemHealth.label,
			delta: systemHealth.detail,
			accent: systemHealth.accent
		}
	]);
</script>

<svelte:head>
	<title>AIOps Control | Alerts</title>
</svelte:head>

<section class="space-y-8">
	<div class="flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
		<div>
			<p class="font-label text-xs uppercase tracking-[0.28em] text-primary/80">
				Stateful alert view backed by live backend data
			</p>
			<h1 class="mt-2 font-headline text-4xl font-black tracking-tight">Alerts</h1>
		</div>

		<div class="flex flex-wrap gap-3">
			<button
				type="button"
				class="flex items-center gap-2 rounded-md bg-gradient-to-br from-primary to-primary-container px-4 py-2 font-label text-[0.68rem] font-bold uppercase tracking-[0.22em] text-on-primary hover:opacity-90"
				onclick={() => void loadAlerts({ background: true })}
				disabled={isLoading || isRefreshing}
			>
				<span class="material-symbols-outlined text-base">sync</span>
				{isRefreshing ? 'Refreshing' : 'Refresh Alerts'}
			</button>
		</div>
	</div>

	<div class="grid grid-cols-1 gap-4 lg:grid-cols-12">
		{#each summaryCards as card, index}
			<article
				class={`panel p-5 ${index < 2 ? 'lg:col-span-3' : 'lg:col-span-6'}`}
				style={`border-left:3px solid ${card.accent}`}
			>
				<p class="font-label text-[0.65rem] uppercase tracking-[0.22em] text-on-surface-variant">
					{card.label}
				</p>
				<div class="mt-3 flex items-end justify-between gap-3">
					<div>
						<div class="font-label text-4xl font-bold" style:color={card.accent}>{card.value}</div>
						<div class="mt-1 font-label text-[0.65rem] uppercase tracking-[0.18em] text-on-surface-variant/70">
							{card.delta}
						</div>
					</div>

					{#if index === 2}
						<div
							class="rounded-md border px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.18em]"
							style={`color:${card.accent};border-color:${card.accent}33;background:${card.accent}12`}
						>
							Live Polling
						</div>
					{/if}
				</div>
			</article>
		{/each}
	</div>

	<section class="terminal-window overflow-hidden">
		<div
			class="flex flex-col gap-3 border-b border-outline-variant/10 bg-surface-container-high px-5 py-4 sm:flex-row sm:items-center sm:justify-between"
		>
			<div class="flex items-center gap-3">
				<div class="flex gap-1.5">
					<span class="h-2.5 w-2.5 rounded-full bg-error/35"></span>
					<span class="h-2.5 w-2.5 rounded-full bg-tertiary/35"></span>
					<span class="h-2.5 w-2.5 rounded-full bg-primary/35"></span>
				</div>
				<span class="font-label text-[0.68rem] uppercase tracking-[0.24em] text-on-surface-variant">
					Console Output :: Alerts_API_v3.0
				</span>
			</div>

			<div class="flex gap-4 font-label text-[0.62rem] uppercase tracking-[0.18em] text-primary/60">
				<span>Rows: {alerts.length.toString().padStart(3, '0')}</span>
				<span>{error ? 'Status: Error' : isLoading ? 'Status: Loading' : 'Status: Connected'}</span>
			</div>
		</div>

		<div class="thin-scrollbar overflow-x-auto">
			<table class="min-w-full text-left">
				<thead class="bg-surface-container-lowest/80">
					<tr>
						<th class="px-5 py-4">Time</th>
						<th class="px-5 py-4">Severity</th>
						<th class="px-5 py-4">Metric</th>
						<th class="px-5 py-4">Description</th>
						<th class="px-5 py-4">Status</th>
						<th class="px-5 py-4 text-right">Actions</th>
					</tr>
				</thead>
				<tbody>
					{#if isLoading}
						<tr>
							<td colspan="6" class="px-5 py-10 text-center text-sm text-on-surface-variant">
								Loading alerts from the backend...
							</td>
						</tr>
					{:else if error && alerts.length === 0}
						<tr>
							<td colspan="6" class="px-5 py-10 text-center text-sm text-error">{error}</td>
						</tr>
					{:else if alerts.length === 0}
						<tr>
							<td colspan="6" class="px-5 py-10 text-center text-sm text-on-surface-variant">
								No alerts yet. Stress CPU, memory, or disk to trigger Sprint 3 rules.
							</td>
						</tr>
					{:else}
						{#each alerts as alert}
							<tr>
								<td class="px-5 py-5 font-label text-sm text-on-surface/80">{formatAlertTime(alert)}</td>
								<td class="px-5 py-5">
									<span
										class="rounded-md border px-2 py-1 font-label text-[0.62rem] font-bold uppercase tracking-[0.18em]"
										style={`color:${getSeverityColor(alert.severity)};border-color:${getSeverityColor(alert.severity)}33;background:${getSeverityColor(alert.severity)}12`}
									>
										{alert.severity}
									</span>
								</td>
								<td class="px-5 py-5 font-label text-sm font-bold">
									{formatMetricName(alert.metric_name)}
								</td>
								<td class="px-5 py-5 text-sm text-on-surface-variant">
									<div class="flex items-start gap-3">
										<span
											class="mt-1 h-2 w-2 rounded-full"
											style={`background:${getStatusColor(alert.status)}`}
										></span>
										<div class="space-y-1">
											<div>{alert.description}</div>
											<div class="font-label text-[0.62rem] uppercase tracking-[0.16em] text-on-surface-variant/70">
												Current {formatPercent(alert.current_value)} / Threshold {formatPercent(alert.threshold)}
											</div>
										</div>
									</div>
								</td>
								<td class="px-5 py-5">
									<span
										class="font-label text-[0.62rem] font-bold uppercase tracking-[0.18em]"
										style={`color:${getStatusColor(alert.status)}`}
									>
										{formatStatus(alert.status)}
									</span>
								</td>
								<td class="px-5 py-5 text-right">
										<div class="flex justify-end gap-2">
											<a
												href={`/assistant?alert_id=${alert.id}`}
												class="rounded-md bg-tertiary/10 px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.18em] text-tertiary transition-colors hover:bg-tertiary/20"
											>
												Explain
											</a>
											<button
												type="button"
												class="rounded-md bg-outline-variant/20 px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface transition-colors hover:bg-outline-variant/30 disabled:cursor-not-allowed disabled:opacity-40"
												onclick={() => void handleAction(alert.id, 'acknowledge')}
											disabled={!canAcknowledge(alert) || !!pendingActions[alert.id]}
										>
											{isActionPending(alert.id, 'acknowledge') ? 'Working...' : 'Acknowledge'}
										</button>
										<button
											type="button"
											class="rounded-md bg-primary/10 px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.18em] text-primary transition-colors hover:bg-primary/20 disabled:cursor-not-allowed disabled:opacity-40"
											onclick={() => void handleAction(alert.id, 'silence')}
											disabled={!canSilence(alert) || !!pendingActions[alert.id]}
										>
											{isActionPending(alert.id, 'silence') ? 'Working...' : 'Silence'}
										</button>
									</div>
								</td>
							</tr>
						{/each}
					{/if}
				</tbody>
			</table>
		</div>

		<div
			class="flex flex-col gap-3 border-t border-outline-variant/5 bg-surface-container-lowest/50 px-5 py-4 sm:flex-row sm:items-center sm:justify-between"
		>
			<div class="flex flex-wrap gap-6 font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
				<div class="flex items-center gap-2">
					<span class={`inline-block h-2 w-2 rounded-full ${error ? 'bg-error' : 'bg-primary'}`}></span>
					<span>{error ? 'Backend: Unreachable' : 'Backend: Synced'}</span>
				</div>
				<div class="flex items-center gap-2">
					<span class="inline-block h-2 w-2 rounded-full bg-outline"></span>
					<span>Poll Interval: {POLL_INTERVAL_MS / 1000}s</span>
				</div>
			</div>

			<span class="rounded-md bg-primary/10 px-3 py-1 font-label text-[0.62rem] uppercase tracking-[0.18em] text-primary">
				Last Sync: {lastUpdated ? formatTimestamp(lastUpdated) : 'Pending'}
			</span>
		</div>
	</section>

	{#if error && alerts.length > 0}
		<div class="rounded-xl border border-error/30 bg-error/8 px-4 py-3 text-sm text-error">
			{error}
		</div>
	{/if}

	<div class="grid grid-cols-1 gap-5 lg:grid-cols-3">
		<article class="panel p-5">
			<h2 class="font-label text-[0.65rem] uppercase tracking-[0.24em] text-primary/80">
				Alert Lifecycle
			</h2>
			<div class="mt-6 grid grid-cols-2 gap-3">
				<div class="rounded-xl border border-outline-variant/10 bg-surface-container p-4">
					<div class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
						Firing
					</div>
					<div class="mt-2 font-label text-2xl font-bold text-error">
						{alerts.filter((alert) => alert.status === 'firing').length}
					</div>
				</div>
				<div class="rounded-xl border border-outline-variant/10 bg-surface-container p-4">
					<div class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
						Acknowledged
					</div>
					<div class="mt-2 font-label text-2xl font-bold text-primary">
						{alerts.filter((alert) => alert.status === 'acknowledged').length}
					</div>
				</div>
			</div>
		</article>

		<article class="panel p-5">
			<h2 class="font-label text-[0.65rem] uppercase tracking-[0.24em] text-primary/80">Resolution State</h2>
			<div class="mt-6 grid grid-cols-2 gap-3">
				<div class="rounded-xl border border-outline-variant/10 bg-surface-container p-4">
					<div class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
						Silenced
					</div>
					<div class="mt-2 font-label text-2xl font-bold text-secondary">
						{alerts.filter((alert) => alert.status === 'silenced').length}
					</div>
				</div>
				<div class="rounded-xl border border-outline-variant/10 bg-surface-container p-4">
					<div class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
						Resolved
					</div>
					<div class="mt-2 font-label text-2xl font-bold text-primary-fixed-dim">
						{alerts.filter((alert) => alert.status === 'resolved').length}
					</div>
				</div>
			</div>
		</article>

		<article class="panel flex items-center gap-4 p-5">
			<div
				class="flex h-14 w-14 items-center justify-center rounded-xl border border-primary/20 bg-primary/10 text-primary"
			>
				<span class="material-symbols-outlined text-2xl">monitoring</span>
			</div>
			<div>
				<div class="font-label text-[0.65rem] uppercase tracking-[0.24em] text-primary/80">Sprint 3</div>
				<p class="mt-1 text-sm text-on-surface-variant">
					Alerts now read from the backend contract and support acknowledge and silence actions.
				</p>
			</div>
		</article>
	</div>
</section>
