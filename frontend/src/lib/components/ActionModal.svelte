<script lang="ts">
	import type { ActionType } from '$lib/api/types';

	interface Props {
		isOpen?: boolean;
		actionType?: ActionType;
		actionTarget?: string;
		onConfirm?: () => void;
		onCancel?: () => void;
	}

	let {
		isOpen = false,
		actionType = 'kill_process',
		actionTarget = '',
		onConfirm = () => {},
		onCancel = () => {}
	}: Props = $props();

	const riskyActions: ActionType[] = ['kill_process', 'restart_container'];
	const isRisky = riskyActions.includes(actionType);

	const actionInfo: Record<ActionType, { title: string; description: string; icon: string; color: string }> = {
		kill_process: {
			title: 'Kill Process',
			description: 'This will terminate the specified process immediately. This action cannot be undone.',
			icon: 'cancel',
			color: 'text-error'
		},
		restart_container: {
			title: 'Restart Container',
			description: 'This will restart the specified container. The container will be stopped and started again.',
			icon: 'refresh',
			color: 'text-tertiary'
		},
		scale_container: {
			title: 'Scale Container',
			description: 'This will adjust the number of replicas for the specified service.',
			icon: 'aspect_ratio',
			color: 'text-primary'
		}
	};

	const info = actionInfo[actionType];
</script>

{#if isOpen}
	<div class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 backdrop-blur-sm">
		<div class="panel w-full max-w-lg rounded-xl p-6 shadow-2xl">
			<div class="mb-6 flex items-start justify-between">
				<div class="flex items-center gap-3">
					<div class="rounded-lg bg-surface-container-high p-3">
						<span class={`material-symbols-outlined text-2xl ${info.color}`}>{info.icon}</span>
					</div>
					<div>
						<h3 class="font-headline text-xl font-bold">{info.title}</h3>
						<p class="font-label text-xs uppercase tracking-[0.18em] text-on-surface-variant">
							Target: {actionTarget}
						</p>
					</div>
				</div>
				<button
					onclick={onCancel}
					class="rounded-lg p-2 text-on-surface-variant transition-colors hover:bg-surface-container-high hover:text-on-surface"
				>
					<span class="material-symbols-outlined">close</span>
				</button>
			</div>

			<div class="mb-6 space-y-4">
				<p class="text-sm text-on-surface-variant">{info.description}</p>

				{#if isRisky}
					<div class="rounded-lg bg-error/10 border border-error/20 p-4">
						<div class="flex items-start gap-3">
							<span class="material-symbols-outlined text-error">warning</span>
							<div>
								<p class="font-label text-xs font-semibold uppercase tracking-[0.18em] text-error">
									Risky Action
								</p>
								<p class="mt-1 text-sm text-on-surface-variant">
									This action may cause service disruption. Make sure you have a backup plan and
									understand the consequences.
								</p>
							</div>
						</div>
					</div>
				{/if}
			</div>

			<div class="flex justify-end gap-3">
				<button
					onclick={onCancel}
					class="rounded-lg border border-outline-variant/20 px-5 py-2.5 font-label text-xs font-semibold uppercase tracking-[0.22em] text-on-surface transition-colors hover:bg-surface-container-high"
				>
					Cancel
				</button>
				<button
					onclick={onConfirm}
					class={`rounded-lg px-5 py-2.5 font-label text-xs font-bold uppercase tracking-[0.22em] text-on-surface transition-opacity hover:opacity-90 ${
						isRisky
							? 'bg-gradient-to-br from-error to-error-container text-on-error'
							: 'bg-gradient-to-br from-primary to-primary-container text-on-primary'
					}`}
				>
					{isRisky ? 'Execute Anyway' : 'Execute Action'}
				</button>
			</div>
		</div>
	</div>
{/if}
