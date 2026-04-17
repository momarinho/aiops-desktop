<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { executeAction, getActions, getProcesses } from '$lib/api';
	import type { Action, ActionType, ExecuteActionRequest, ProcessInfo } from '$lib/api/types';
	import ActionModal from '$lib/components/ActionModal.svelte';

	let actions: Action[] = $state([]);
	let isLoading = $state(true);
	let error = $state<string | null>(null);
	let showModal = $state(false);
	let selectedAction = $state<ExecuteActionRequest | null>(null);

	// Process browser state
	let processes: ProcessInfo[] = $state([]);
	let isProcessesLoading = $state(false);
	let processSearch = $state('');
	let showProcessBrowser = $state(false);

	// Form state
	let actionType: ActionType = $state('kill_process');
	let actionTarget = $state('');
	let actionParameters = $state<Record<string, any>>({});
	let replicas = $state('1');

	// Polling for action updates
	let pollInterval: number;

	onMount(async () => {
		await loadActions();
		// Poll every 2 seconds for action status updates
		pollInterval = window.setInterval(loadActions, 2000);
	});

	onDestroy(() => {
		if (pollInterval) {
			clearInterval(pollInterval);
		}
	});

	async function loadActions() {
		try {
			actions = await getActions();
			error = null;
		} catch (err) {
			error = err instanceof Error ? err.message : 'Failed to load actions';
		} finally {
			isLoading = false;
		}
	}

	function handleExecuteClick() {
		if (!actionTarget) {
			error = 'Please provide a target';
			return;
		}

		// Build parameters based on action type
		let params: Record<string, any> = {};
		if (actionType === 'scale_container') {
			if (!replicas) {
				error = 'Please specify number of replicas';
				return;
			}
			params = { replicas: parseInt(replicas, 10) };
		}

		selectedAction = {
			type: actionType,
			target: actionTarget,
			parameters: params,
			user: 'current_user' // In real app, this would come from auth
		};

		showModal = true;
	}

	async function confirmAction() {
		if (!selectedAction) return;

		try {
			await executeAction(selectedAction);
			await loadActions();
			resetForm();
		} catch (err) {
			error = err instanceof Error ? err.message : 'Failed to execute action';
		}
	}

	function cancelAction() {
		showModal = false;
		selectedAction = null;
	}

	function resetForm() {
		actionTarget = '';
		actionParameters = {};
		replicas = '1';
		showModal = false;
		selectedAction = null;
	}

	function getStatusColor(status: Action['status']) {
		switch (status) {
			case 'pending':
				return 'text-tertiary border-tertiary/20 bg-tertiary/10';
			case 'success':
				return 'text-primary border-primary/20 bg-primary/10';
			case 'failed':
				return 'text-error border-error/20 bg-error/10';
		}
	}

	function getStatusIcon(status: Action['status']) {
		switch (status) {
			case 'pending':
				return 'hourglass_empty';
			case 'success':
				return 'check_circle';
			case 'failed':
				return 'cancel';
		}
	}

	function getActionLabel(type: ActionType) {
		switch (type) {
			case 'kill_process':
				return 'Kill Process';
			case 'restart_container':
				return 'Restart Container';
			case 'scale_container':
				return 'Scale Container';
		}
	}

	function getActionIcon(type: ActionType) {
		switch (type) {
			case 'kill_process':
				return 'cancel';
			case 'restart_container':
				return 'refresh';
			case 'scale_container':
				return 'aspect_ratio';
		}
	}

	function formatTimestamp(timestamp: string): string {
		return new Date(timestamp).toLocaleString();
	}

	function getDuration(start?: string, end?: string): string {
		if (!start) return '-';
		const startTime = new Date(start).getTime();
		const endTime = end ? new Date(end).getTime() : Date.now();
		const duration = Math.round((endTime - startTime) / 1000);
		return `${duration}s`;
	}

	async function loadProcesses() {
		try {
			isProcessesLoading = true;
			processes = await getProcesses();
		} catch (err) {
			error = err instanceof Error ? err.message : 'Failed to load processes';
		} finally {
			isProcessesLoading = false;
		}
	}

	function selectProcess(process: ProcessInfo) {
		actionTarget = process.pid.toString();
		showProcessBrowser = false;
	}

	const filteredProcesses = $derived(
		processes.filter((p) => {
			if (!processSearch) return true;
			const search = processSearch.toLowerCase();
			return (
				p.name.toLowerCase().includes(search) ||
				p.command.toLowerCase().includes(search) ||
				p.user.toLowerCase().includes(search) ||
				p.pid.toString().includes(search)
			);
		})
	);

	const processStatusColor = $derived({
		critical: 'text-error border-error/20 bg-error/10',
		system: 'text-tertiary border-tertiary/20 bg-tertiary/10',
		user: 'text-primary border-primary/20 bg-primary/10'
	});

	function getProcessStatusIcon(status: ProcessInfo['status']) {
		switch (status) {
			case 'critical':
				return 'warning';
			case 'system':
				return 'settings';
			case 'user':
				return 'person';
		}
	}
</script>

<svelte:head>
	<title>AIOps Control | Actions</title>
</svelte:head>

<section class="space-y-8">
	<header class="flex flex-col gap-5 md:flex-row md:items-end md:justify-between">
		<div>
			<h1 class="font-headline text-4xl font-black tracking-tight">System Actions</h1>
			<p class="mt-2 font-label text-xs uppercase tracking-[0.28em] text-on-surface-variant/70">
				Safe Automation & Operational Commands
			</p>
		</div>

		<div class="flex items-center gap-3">
			{#if error}
				<div class="rounded-lg bg-error/10 border border-error/20 px-4 py-2">
					<p class="font-label text-xs font-semibold uppercase tracking-[0.18em] text-error">{error}</p>
				</div>
			{/if}
		</div>
	</header>

	<div class="grid grid-cols-1 gap-6 xl:grid-cols-3">
		<!-- Action Form -->
		<article class="panel p-6 xl:col-span-1">
			<div class="mb-6 flex items-center gap-3">
				<span class="material-symbols-outlined text-primary">terminal</span>
				<h2 class="font-headline text-xl font-bold">Execute Action</h2>
			</div>

			<div class="space-y-5">
				<div>
					<label for="action-type" class="mb-2 block font-label text-xs font-semibold uppercase tracking-[0.18em] text-on-surface-variant">
						Action Type
					</label>
					<select
						id="action-type"
						bind:value={actionType}
						class="w-full rounded-lg border border-outline-variant/20 bg-surface-container-low px-4 py-3 font-label text-sm text-on-surface transition-colors focus:border-primary focus:outline-none"
					>
						<option value="kill_process">Kill Process</option>
						<option value="restart_container">Restart Container</option>
						<option value="scale_container">Scale Container</option>
					</select>
				</div>

				<div>
					<label for="action-target" class="mb-2 block font-label text-xs font-semibold uppercase tracking-[0.18em] text-on-surface-variant">
						Target
					</label>
					<div class="flex gap-2">
						<input
							id="action-target"
							type="text"
							bind:value={actionTarget}
							placeholder={actionType === 'kill_process' ? 'Process ID (e.g., 1234)' : 'Container ID/Name'}
							disabled={actionType === 'kill_process'}
							class="flex-1 rounded-lg border border-outline-variant/20 bg-surface-container-low px-4 py-3 font-label text-sm text-on-surface transition-colors focus:border-primary focus:outline-none disabled:opacity-50"
						/>
						{#if actionType === 'kill_process'}
							<button
								onclick={() => {
									showProcessBrowser = !showProcessBrowser;
									if (showProcessBrowser) {
										loadProcesses();
									}
								}}
								class="rounded-lg border border-outline-variant/20 bg-surface-container-low px-4 py-3 font-label text-sm text-on-surface transition-colors hover:bg-surface-container-high"
							>
								<span class="material-symbols-outlined">list</span>
							</button>
						{/if}
					</div>
				</div>

				{#if actionType === 'scale_container'}
					<div>
						<label for="replicas" class="mb-2 block font-label text-xs font-semibold uppercase tracking-[0.18em] text-on-surface-variant">
							Replicas
						</label>
						<input
							id="replicas"
							type="number"
							min="1"
							bind:value={replicas}
							class="w-full rounded-lg border border-outline-variant/20 bg-surface-container-low px-4 py-3 font-label text-sm text-on-surface transition-colors focus:border-primary focus:outline-none"
						/>
					</div>
				{/if}

				<button
					onclick={handleExecuteClick}
					class="w-full rounded-lg bg-gradient-to-br from-primary to-primary-container px-4 py-3 font-label text-xs font-bold uppercase tracking-[0.22em] text-on-primary transition-opacity hover:opacity-90"
				>
					Execute Action
				</button>

				<!-- Process Browser -->
				{#if showProcessBrowser}
					<div class="mt-4 rounded-lg border border-outline-variant/10 bg-surface-container-low p-4">
						<div class="mb-4 flex items-center justify-between">
							<h3 class="font-label text-sm font-semibold uppercase tracking-[0.18em] text-on-surface-variant">
								Process Browser
							</h3>
							<div class="flex items-center gap-2">
								<input
									type="text"
									bind:value={processSearch}
									placeholder="Search processes..."
									class="rounded-lg border border-outline-variant/20 bg-surface-container-lowest px-3 py-2 font-label text-xs text-on-surface transition-colors focus:border-primary focus:outline-none"
								/>
								<button
									onclick={loadProcesses}
									class="rounded-lg border border-outline-variant/20 p-2 text-on-surface-variant transition-colors hover:bg-surface-container-high hover:text-primary"
								>
									<span class="material-symbols-outlined text-base">refresh</span>
								</button>
							</div>
						</div>

						{#if isProcessesLoading}
							<div class="flex min-h-32 items-center justify-center">
								<p class="text-on-surface-variant">Loading processes...</p>
							</div>
						{:else if filteredProcesses.length === 0}
							<div class="flex min-h-32 items-center justify-center">
								<p class="text-on-surface-variant">No processes found</p>
							</div>
						{:else}
							<div class="max-h-80 overflow-y-auto space-y-2">
								{#each filteredProcesses as process (process.pid)}
									<div
										onclick={() => selectProcess(process)}
										class="cursor-pointer rounded-lg border border-outline-variant/10 p-3 transition-colors hover:bg-surface-container-high {process.is_critical ? 'opacity-50' : ''}"
									>
										<div class="flex items-start justify-between gap-3">
											<div class="flex-1">
												<div class="flex items-center gap-2">
													<p class="font-label text-sm font-semibold">{process.name}</p>
													<span
														class={`rounded-md border px-2 py-0.5 font-label text-[0.6rem] font-bold uppercase tracking-[0.16em] ${processStatusColor[process.status]}`}
													>
														<span class="material-symbols-outlined text-xs">
															{getProcessStatusIcon(process.status)}
														</span>
														{process.status}
													</span>
													{#if process.is_critical}
														<span class="material-symbols-outlined text-sm text-error">warning</span>
													{/if}
												</div>
												<p class="font-label text-xs text-on-surface-variant">
													PID: <span class="font-medium text-on-surface">{process.pid}</span>
													| User: <span class="font-medium text-on-surface">{process.user}</span>
												</p>
												<p class="mt-1 truncate text-xs text-on-surface-variant/70">
													{process.command}
												</p>
											</div>
											<div class="text-right">
												<p class="font-label text-xs text-on-surface-variant">
													CPU: <span class="font-medium text-on-surface">{process.cpu_percent.toFixed(1)}%</span>
												</p>
												<p class="font-label text-xs text-on-surface-variant">
													Mem: <span class="font-medium text-on-surface">{process.memory_mb.toFixed(1)} MB</span>
												</p>
											</div>
										</div>
									</div>
								{/each}
							</div>
						{/if}
					</div>
				{/if}
			</div>
		</article>

		<!-- Action History -->
		<article class="panel p-6 xl:col-span-2">
			<div class="mb-6 flex items-start justify-between">
				<div class="flex items-center gap-3">
					<span class="material-symbols-outlined text-primary">history</span>
					<div>
						<h2 class="font-headline text-xl font-bold">Action History</h2>
						<p class="font-label text-xs uppercase tracking-[0.18em] text-on-surface-variant/70">
							Recent Executions
						</p>
					</div>
				</div>
				<button
					onclick={loadActions}
					class="rounded-lg border border-outline-variant/20 p-2 text-on-surface-variant transition-colors hover:bg-surface-container-high hover:text-primary"
				>
					<span class="material-symbols-outlined">refresh</span>
				</button>
			</div>

			{#if isLoading}
				<div class="flex min-h-48 items-center justify-center">
					<p class="text-on-surface-variant">Loading actions...</p>
				</div>
			{:else if actions.length === 0}
				<div class="flex min-h-48 items-center justify-center">
					<p class="text-on-surface-variant">No actions executed yet</p>
				</div>
			{:else}
				<div class="space-y-3">
					{#each actions.slice().reverse() as action (action.id)}
						<div class="rounded-lg border border-outline-variant/10 bg-surface-container-low/50 p-4 transition-colors hover:bg-surface-container-low">
							<div class="flex items-start justify-between gap-4">
								<div class="flex items-start gap-3">
									<div class="mt-1 rounded-md bg-surface-container-high p-2">
										<span class="material-symbols-outlined text-primary">
											{getActionIcon(action.type)}
										</span>
									</div>
									<div class="flex-1">
										<div class="flex items-center gap-2">
											<p class="font-label text-sm font-semibold">
												{getActionLabel(action.type)}
											</p>
											{#if action.risky}
												<span class="material-symbols-outlined text-sm text-tertiary">warning</span>
											{/if}
										</div>
										<p class="font-label text-xs text-on-surface-variant">
											Target: <span class="font-medium text-on-surface">{action.target}</span>
										</p>
										<p class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant/70">
											{formatTimestamp(action.request_time)}
										</p>
									</div>
								</div>

								<div class="flex flex-col items-end gap-2">
									<span
										class={`rounded-md border px-3 py-1.5 font-label text-[0.62rem] font-bold uppercase tracking-[0.18em] ${getStatusColor(action.status)}`}
									>
										<span class="material-symbols-outlined text-sm">
											{getStatusIcon(action.status)}
										</span>
										{action.status}
									</span>
									<span class="font-label text-[0.62rem] text-on-surface-variant/70">
										{getDuration(action.start_time, action.end_time)}
									</span>
								</div>
							</div>

							{#if action.error}
								<div class="mt-3 rounded-md bg-error/10 border border-error/20 p-3">
									<p class="font-label text-xs font-semibold uppercase tracking-[0.18em] text-error">
										Error
									</p>
									<p class="mt-1 text-sm text-error">{action.error}</p>
								</div>
							{/if}

							{#if action.output && action.status === 'success'}
								<div class="mt-3 rounded-md bg-primary/10 border border-primary/20 p-3">
									<p class="font-label text-xs font-semibold uppercase tracking-[0.18em] text-primary">
										Output
									</p>
									<pre class="mt-1 text-sm text-on-surface-variant">{action.output}</pre>
								</div>
							{/if}
						</div>
					{/each}
				</div>
			{/if}
		</article>
	</div>

	<!-- Action Modal -->
	<ActionModal
		isOpen={showModal}
		actionType={selectedAction?.type || 'kill_process'}
		actionTarget={selectedAction?.target || ''}
		onConfirm={confirmAction}
		onCancel={cancelAction}
	/>
</section>
