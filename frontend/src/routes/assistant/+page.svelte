<script lang="ts">
	import { page } from '$app/state';
	import { explainAlert, getAlertById } from '$lib/api';
	import type { Alert, ExplainAlertRequest, ExplainAlertResponse } from '$lib/api/types';

	let selectedAlert = $state<Alert | null>(null);
	let explanation = $state<ExplainAlertResponse | null>(null);
	let isLoading = $state(false);
	let alertError = $state<string | null>(null);
	let explanationError = $state<string | null>(null);

	let lastLoadedAlertId = '';
	let requestVersion = 0;

	const selectedAlertId = $derived(page.url.searchParams.get('alert_id')?.trim() ?? '');
	const hasSelectedAlert = $derived(selectedAlertId.length > 0);

	$effect(() => {
		const alertId = selectedAlertId;

		if (!alertId) {
			lastLoadedAlertId = '';
			selectedAlert = null;
			explanation = null;
			alertError = null;
			explanationError = null;
			isLoading = false;
			return;
		}

		if (alertId === lastLoadedAlertId) {
			return;
		}

		lastLoadedAlertId = alertId;
		void loadAssistantState(alertId);
	});

	async function retryExplanation() {
		if (!selectedAlertId) {
			return;
		}

		lastLoadedAlertId = '';
		await loadAssistantState(selectedAlertId);
	}

	async function loadAssistantState(alertId: string) {
		const currentRequest = ++requestVersion;
		isLoading = true;
		alertError = null;
		explanationError = null;
		explanation = null;

		try {
			const alert = await getAlertById(alertId);
			if (currentRequest !== requestVersion) {
				return;
			}

			selectedAlert = alert;

			try {
				explanation = await explainAlert(buildExplainAlertRequest(alert));
			} catch (err) {
				if (currentRequest !== requestVersion) {
					return;
				}

				explanationError = err instanceof Error ? err.message : 'Failed to explain alert';
			}
		} catch (err) {
			if (currentRequest !== requestVersion) {
				return;
			}

			selectedAlert = null;
			alertError = err instanceof Error ? err.message : 'Failed to load alert';
		} finally {
			if (currentRequest === requestVersion) {
				isLoading = false;
			}
		}
	}

	function buildExplainAlertRequest(alert: Alert): ExplainAlertRequest {
		return {
			alert_id: alert.id,
			context: {
				additional_notes: `${formatSeverity(alert.severity)} alert currently ${formatStatus(alert.status)}.`
			}
		};
	}

	function formatMetricName(metricName: string) {
		return metricName
			.replace('_usage_', ' ')
			.replace(/_/g, ' ')
			.replace(/\b\w/g, (char) => char.toUpperCase());
	}

	function formatPercent(value: number) {
		return `${value.toFixed(1)}%`;
	}

	function formatSeverity(severity: Alert['severity']) {
		return severity.charAt(0).toUpperCase() + severity.slice(1);
	}

	function formatStatus(status: Alert['status']) {
		return status.replace(/_/g, ' ');
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

	function getSeverityClasses(severity: Alert['severity']) {
		return severity === 'critical'
			? 'border-error/20 bg-error/10 text-error'
			: 'border-tertiary/20 bg-tertiary/10 text-tertiary';
	}

	function getStatusClasses(status: Alert['status']) {
		switch (status) {
			case 'firing':
				return 'border-error/20 bg-error/10 text-error';
			case 'acknowledged':
				return 'border-primary/20 bg-primary/10 text-primary';
			case 'silenced':
				return 'border-outline-variant/20 bg-outline-variant/10 text-on-surface-variant';
			case 'resolved':
				return 'border-primary-container/20 bg-primary-container/10 text-primary-container';
		}
	}

	function getFallbackActions(alert: Alert | null) {
		if (!alert) {
			return [];
		}

		switch (alert.metric_name) {
			case 'cpu_usage_percent':
				return [
					'Inspect the processes or containers consuming the most CPU.',
					'Compare the spike time with recent deploys or workload changes.'
				];
			case 'memory_usage_bytes':
				return [
					'Review which process or container is holding the most memory.',
					'Check for recent deploys, cache growth, or memory leaks.'
				];
			case 'disk_usage_bytes':
				return [
					'Identify which directories or volumes grew most recently.',
					'Clear space safely before the filesystem reaches a hard limit.'
				];
			default:
				return [
					'Review the latest metric trend and confirm the alert is still active.',
					'Start with the least risky mitigation and record the result.'
				];
		}
	}

	const fallbackActions = $derived(getFallbackActions(selectedAlert));
</script>

<svelte:head>
	<title>AIOps Control | Assistant</title>
</svelte:head>

<section class="space-y-8">
	<header class="flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
		<div>
			<p class="font-label text-xs uppercase tracking-[0.28em] text-primary/75">
				Structured alert explanation backed by the Sprint 5 AI API
			</p>
			<h1 class="mt-2 font-headline text-4xl font-black tracking-tight">Assistant</h1>
		</div>

		<div class="flex flex-wrap gap-3">
			<a
				href="/alerts"
				class="rounded-md border border-outline-variant/20 bg-surface-container-high px-4 py-2 font-label text-[0.68rem] font-semibold uppercase tracking-[0.22em] text-on-surface transition-colors hover:bg-surface-container-highest"
			>
				Back to Alerts
			</a>
			<button
				type="button"
				class="rounded-md bg-gradient-to-br from-primary to-primary-container px-4 py-2 font-label text-[0.68rem] font-bold uppercase tracking-[0.22em] text-on-primary transition-opacity hover:opacity-90 disabled:cursor-not-allowed disabled:opacity-50"
				onclick={() => void retryExplanation()}
				disabled={!hasSelectedAlert || isLoading}
			>
				{isLoading ? 'Loading...' : 'Retry Explanation'}
			</button>
		</div>
	</header>

	{#if !hasSelectedAlert}
		<section class="panel p-8 text-center">
			<div class="mx-auto max-w-2xl space-y-4">
				<div class="mx-auto flex h-14 w-14 items-center justify-center rounded-2xl bg-primary/10 text-primary">
					<span class="material-symbols-outlined text-3xl">smart_toy</span>
				</div>
				<h2 class="font-headline text-2xl font-bold">No alert selected</h2>
				<p class="text-on-surface-variant">
					Open the assistant from the Alerts page to load a real alert explanation. The page expects an
					`alert_id` query parameter.
				</p>
				<a
					href="/alerts"
					class="inline-flex rounded-md bg-gradient-to-br from-primary to-primary-container px-5 py-3 font-label text-[0.68rem] font-bold uppercase tracking-[0.22em] text-on-primary transition-opacity hover:opacity-90"
				>
					Choose an Alert
				</a>
			</div>
		</section>
	{:else}
		<div class="grid grid-cols-1 gap-6 xl:grid-cols-12">
			<section class="panel p-6 xl:col-span-4">
				<div class="flex items-start justify-between gap-4">
					<div>
						<p class="font-label text-[0.65rem] uppercase tracking-[0.22em] text-primary/70">
							Selected Alert
						</p>
						<h2 class="mt-2 font-headline text-2xl font-bold">
							{selectedAlert ? formatMetricName(selectedAlert.metric_name) : 'Loading context'}
						</h2>
					</div>
					{#if selectedAlert}
						<span
							class={`rounded-md border px-3 py-1 font-label text-[0.62rem] font-bold uppercase tracking-[0.18em] ${getSeverityClasses(selectedAlert.severity)}`}
						>
							{selectedAlert.severity}
						</span>
					{/if}
				</div>

				{#if alertError}
					<div class="mt-6 rounded-2xl border border-error/20 bg-error/10 p-4">
						<p class="font-label text-[0.68rem] font-semibold uppercase tracking-[0.2em] text-error">
							Alert unavailable
						</p>
						<p class="mt-2 text-sm text-error">{alertError}</p>
					</div>
				{:else if selectedAlert}
					<div class="mt-6 space-y-5">
						<div class="rounded-2xl border border-outline-variant/10 bg-surface-container-low p-5">
							<p class="text-sm text-on-surface">{selectedAlert.description}</p>
						</div>

						<div class="grid grid-cols-2 gap-3">
							<div class="rounded-2xl border border-outline-variant/10 bg-surface-container-low p-4">
								<p class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
									Current Value
								</p>
								<p class="mt-2 font-label text-2xl font-bold text-on-surface">
									{formatPercent(selectedAlert.current_value)}
								</p>
							</div>
							<div class="rounded-2xl border border-outline-variant/10 bg-surface-container-low p-4">
								<p class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
									Threshold
								</p>
								<p class="mt-2 font-label text-2xl font-bold text-on-surface">
									{formatPercent(selectedAlert.threshold)}
								</p>
							</div>
						</div>

						<div class="flex flex-wrap gap-2">
							<span
								class={`rounded-md border px-3 py-2 font-label text-[0.62rem] font-bold uppercase tracking-[0.18em] ${getStatusClasses(selectedAlert.status)}`}
							>
								{formatStatus(selectedAlert.status)}
							</span>
							<span class="rounded-md border border-outline-variant/10 bg-surface-container-low px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
								Started {formatTimestamp(selectedAlert.started_at ?? selectedAlert.updated_at)}
							</span>
							<span class="rounded-md border border-outline-variant/10 bg-surface-container-low px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
								Alert ID {selectedAlert.id}
							</span>
						</div>
					</div>
				{:else}
					<div class="mt-6 rounded-2xl border border-outline-variant/10 bg-surface-container-low p-5 text-sm text-on-surface-variant">
						Loading alert context...
					</div>
				{/if}
			</section>

			<section class="panel p-6 xl:col-span-8">
				<div class="flex items-start justify-between gap-4">
					<div>
						<p class="font-label text-[0.65rem] uppercase tracking-[0.22em] text-primary/70">
							Explanation Output
						</p>
						<h2 class="mt-2 font-headline text-2xl font-bold">Structured Analysis</h2>
					</div>
					<div class="rounded-md border border-outline-variant/10 bg-surface-container-low px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
						{isLoading ? 'Loading' : explanation ? 'Ready' : explanationError ? 'Fallback' : 'Idle'}
					</div>
				</div>

				{#if isLoading}
					<div class="mt-8 flex items-center gap-4 rounded-2xl border border-outline-variant/10 bg-surface-container-low p-6">
						<div class="flex gap-1">
							<span class="h-2.5 w-2.5 animate-bounce rounded-full bg-primary"></span>
							<span class="h-2.5 w-2.5 animate-bounce rounded-full bg-primary [animation-delay:0.2s]"></span>
							<span class="h-2.5 w-2.5 animate-bounce rounded-full bg-primary [animation-delay:0.4s]"></span>
						</div>
						<div>
							<p class="font-label text-[0.68rem] font-semibold uppercase tracking-[0.2em] text-on-surface">
								Analyzing alert
							</p>
							<p class="mt-1 text-sm text-on-surface-variant">
								Loading the selected alert and requesting a structured explanation.
							</p>
						</div>
					</div>
				{:else if alertError}
					<div class="mt-8 rounded-2xl border border-error/20 bg-error/10 p-6">
						<p class="font-label text-[0.68rem] font-semibold uppercase tracking-[0.2em] text-error">
							Explanation unavailable
						</p>
						<p class="mt-2 text-sm text-error">
							The assistant cannot explain this alert because the alert record could not be loaded.
						</p>
					</div>
				{:else if explanation}
					<div class="mt-8 space-y-6">
						<article class="rounded-3xl border border-primary/15 bg-gradient-to-br from-primary/12 via-primary/5 to-transparent p-6">
							<p class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-primary/75">Summary</p>
							<p class="mt-3 text-lg leading-8 text-on-surface">{explanation.summary}</p>
						</article>

						<article class="rounded-2xl border border-outline-variant/10 bg-surface-container-low p-6">
							<p class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
								Probable Cause
							</p>
							<p class="mt-3 text-base leading-7 text-on-surface">{explanation.probable_cause}</p>
						</article>

						<article class="rounded-2xl border border-outline-variant/10 bg-surface-container-low p-6">
							<p class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
								Suggested Actions
							</p>
							<ul class="mt-4 space-y-3">
								{#each explanation.suggested_actions as action}
									<li class="flex gap-3 rounded-2xl border border-outline-variant/10 bg-surface px-4 py-4 text-sm text-on-surface">
										<span class="mt-0.5 material-symbols-outlined text-primary">arrow_forward</span>
										<span>{action}</span>
									</li>
								{/each}
							</ul>
						</article>
					</div>
				{:else}
					<div class="mt-8 space-y-6">
						<div class="rounded-2xl border border-tertiary/20 bg-tertiary/10 p-6">
							<p class="font-label text-[0.68rem] font-semibold uppercase tracking-[0.2em] text-tertiary">
								AI fallback mode
							</p>
							<p class="mt-2 text-sm text-on-surface">
								{explanationError ??
									'No explanation was returned, but the alert context is still available below.'}
							</p>
						</div>

						{#if selectedAlert}
							<article class="rounded-2xl border border-outline-variant/10 bg-surface-container-low p-6">
								<p class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
									Manual Next Steps
								</p>
								<ul class="mt-4 space-y-3">
									{#each fallbackActions as action}
										<li class="flex gap-3 rounded-2xl border border-outline-variant/10 bg-surface px-4 py-4 text-sm text-on-surface">
											<span class="mt-0.5 material-symbols-outlined text-primary">build_circle</span>
											<span>{action}</span>
										</li>
									{/each}
								</ul>
							</article>
						{/if}
					</div>
				{/if}
			</section>
		</div>
	{/if}
</section>
