<script lang="ts">
	let collectors = $state([
		{ label: 'CPU_LOAD', icon: 'memory', enabled: true },
		{ label: 'RAM_USAGE', icon: 'memory_alt', enabled: true },
		{ label: 'DISK_IO', icon: 'database', enabled: true },
		{ label: 'NET_TRAFFIC', icon: 'public', enabled: false },
		{ label: 'CONTAINER_MON', icon: 'token', enabled: true }
	]);

	let warningThreshold = $state(75);
	let criticalThreshold = $state(92);
	let provider = $state('Anthropic (Claude 3.5 Sonnet)');
	let apiKey = $state('sk-••••••••••••••••••••••••');
	let maskKey = $state(true);

	const providers = [
		'Anthropic (Claude 3.5 Sonnet)',
		'OpenAI (GPT-4o)',
		'Local (Llama 3 70B)',
		'Custom Endpoint'
	];
</script>

<svelte:head>
	<title>AIOps Control | Settings</title>
</svelte:head>

<section class="mx-auto max-w-5xl space-y-10">
	<header class="space-y-3">
		<p class="font-label text-xs uppercase tracking-[0.3em] text-primary/80">
			ARCHITECT_v4.2 // AUTOMATION_CORE_ADJUSTMENT
		</p>
		<h1 class="font-headline text-4xl font-black uppercase tracking-tight">System Configuration</h1>
	</header>

	<section class="space-y-5">
		<div class="flex items-center gap-3">
			<span class="h-6 w-1 rounded-full bg-primary"></span>
			<h2 class="font-headline text-2xl font-bold uppercase tracking-wide">01. Metrics Collectors</h2>
		</div>

		<div class="grid grid-cols-1 gap-4 md:grid-cols-5">
			{#each collectors as collector}
				<button
					class="panel flex h-36 flex-col justify-between p-5 text-left transition-colors hover:bg-surface-container-high"
					type="button"
					onclick={() => (collector.enabled = !collector.enabled)}
				>
					<div class="flex items-start justify-between gap-3">
						<span class="material-symbols-outlined text-2xl text-primary">{collector.icon}</span>
						<span
							class={`relative inline-flex h-5 w-10 items-center rounded-full transition-colors ${
								collector.enabled ? 'bg-primary' : 'bg-surface-container-highest'
							}`}
						>
							<span
								class={`absolute h-4 w-4 rounded-full bg-on-surface transition-transform ${
									collector.enabled ? 'translate-x-5' : 'translate-x-0.5'
								}`}
							></span>
						</span>
					</div>
					<span class="font-label text-xs font-medium uppercase tracking-[0.2em]">
						{collector.label}
					</span>
				</button>
			{/each}
		</div>
	</section>

	<section class="space-y-5">
		<div class="flex items-center gap-3">
			<span class="h-6 w-1 rounded-full bg-primary"></span>
			<h2 class="font-headline text-2xl font-bold uppercase tracking-wide">02. Alert Thresholds</h2>
		</div>

		<div class="panel overflow-hidden">
			<div class="space-y-10 p-6 sm:p-8">
				<div class="space-y-4">
					<div class="flex items-end justify-between gap-4">
						<label
							for="warning-threshold"
							class="font-label text-xs uppercase tracking-[0.24em] text-on-surface-variant"
						>
							Warning Threshold
						</label>
						<span class="font-label text-lg font-bold text-primary">{warningThreshold}%</span>
					</div>
					<input
						id="warning-threshold"
						class="w-full accent-primary"
						type="range"
						min="0"
						max="100"
						bind:value={warningThreshold}
					/>
					<p class="text-xs italic text-on-surface-variant/60">
						System starts logging anomaly patterns but maintains full operation.
					</p>
				</div>

				<div class="space-y-4">
					<div class="flex items-end justify-between gap-4">
						<label
							for="critical-threshold"
							class="font-label text-xs uppercase tracking-[0.24em] text-on-surface-variant"
						>
							Critical Threshold
						</label>
						<span class="font-label text-lg font-bold text-error">{criticalThreshold}%</span>
					</div>
					<input
						id="critical-threshold"
						class="w-full accent-primary"
						type="range"
						min="0"
						max="100"
						bind:value={criticalThreshold}
					/>
					<p class="text-xs italic text-on-surface-variant/60">
						Automated container restart and scaling protocol initiated immediately.
					</p>
				</div>
			</div>
		</div>
	</section>

	<section class="space-y-5">
		<div class="flex items-center gap-3">
			<span class="h-6 w-1 rounded-full bg-primary"></span>
			<h2 class="font-headline text-2xl font-bold uppercase tracking-wide">
				03. AI Intelligence Tuning
			</h2>
		</div>

		<div class="grid grid-cols-1 gap-6 md:grid-cols-2">
			<div class="space-y-2">
				<label
					for="llm-provider"
					class="ml-1 font-label text-xs uppercase tracking-[0.24em] text-on-surface-variant"
				>
					LLM Provider
				</label>
				<div class="relative">
					<select
						id="llm-provider"
						class="w-full border-0 border-b-2 border-surface-container-highest bg-surface-container-low py-4 px-4 text-on-surface focus:border-primary focus:ring-0"
						bind:value={provider}
					>
						{#each providers as option}
							<option>{option}</option>
						{/each}
					</select>
					<span
						class="material-symbols-outlined pointer-events-none absolute right-4 top-1/2 -translate-y-1/2 text-on-surface-variant"
					>
						expand_more
					</span>
				</div>
			</div>

			<div class="space-y-2">
				<label
					for="api-auth-key"
					class="ml-1 font-label text-xs uppercase tracking-[0.24em] text-on-surface-variant"
				>
					API Authentication Key
				</label>
				<div class="relative">
					<input
						id="api-auth-key"
						class="w-full border-0 border-b-2 border-surface-container-highest bg-surface-container-low py-4 px-4 pr-12 font-label tracking-tight text-on-surface focus:border-primary focus:ring-0"
						type={maskKey ? 'password' : 'text'}
						bind:value={apiKey}
					/>
					<button
						class="material-symbols-outlined absolute right-4 top-1/2 -translate-y-1/2 text-on-surface-variant transition-colors hover:text-primary"
						type="button"
						onclick={() => (maskKey = !maskKey)}
					>
						{maskKey ? 'visibility_off' : 'visibility'}
					</button>
				</div>
			</div>
		</div>

		<div class="panel-deep overflow-hidden">
			<div
				class="flex items-center justify-between border-b border-outline-variant/10 px-4 py-4 sm:px-6"
			>
				<span class="font-label text-[0.68rem] uppercase tracking-[0.24em] text-primary">
					Intelligence_Sandbox.log
				</span>
				<div class="flex gap-2">
					<span class="h-2 w-2 rounded-full bg-error/35"></span>
					<span class="h-2 w-2 rounded-full bg-tertiary/35"></span>
					<span class="h-2 w-2 rounded-full bg-primary/35"></span>
				</div>
			</div>

			<div class="space-y-4 p-6 font-label text-xs">
				<div class="flex gap-4">
					<span class="shrink-0 text-on-surface-variant/40">14:02:11</span>
					<span class="text-on-surface">
						Prompt: How should I scale container 'nginx-prod-01' during a CPU spike?
					</span>
				</div>

				<div class="rounded-md border-l-2 border-primary bg-surface-container-low p-4">
					<div class="flex gap-4">
						<span class="shrink-0 text-on-surface-variant/40">14:02:12</span>
						<div class="space-y-2">
							<span class="font-bold uppercase tracking-[0.2em] text-primary">
								AI_RESPONSE_CORE:
							</span>
							<p class="leading-relaxed text-on-surface-variant">
								Based on current thresholds ({criticalThreshold}% Critical), I recommend
								horizontal autoscaling to 4 instances. Memory headroom is sufficient for
								expansion.
							</p>
							<pre class="overflow-x-auto rounded-md bg-surface-container-lowest p-3 text-[0.68rem] text-tertiary-container thin-scrollbar">kubectl scale deployment/nginx-prod-01 --replicas=4</pre>
						</div>
					</div>
				</div>
			</div>
		</div>
	</section>

	<div
		class="sticky bottom-4 flex flex-col gap-3 rounded-2xl border border-outline-variant/15 bg-background/80 p-4 shadow-[0_14px_36px_rgba(0,0,0,0.22)] backdrop-blur sm:flex-row sm:items-center sm:justify-between"
	>
		<div class="flex flex-wrap items-center gap-5 font-label text-[0.65rem] uppercase tracking-[0.22em] text-on-surface-variant/70">
			<span class="flex items-center gap-2">
				<span class="status-dot pulse-ring"></span>
				System Ready
			</span>
			<span>Version 4.2.0-Alpha</span>
		</div>
		<div class="flex flex-wrap gap-3">
			<button
				class="rounded-md border border-outline-variant/20 px-5 py-3 font-label text-xs font-bold uppercase tracking-[0.22em] transition-colors hover:bg-surface-container-highest"
			>
				Discard Changes
			</button>
			<button
				class="rounded-md bg-gradient-to-br from-primary to-primary-container px-6 py-3 font-label text-xs font-bold uppercase tracking-[0.22em] text-on-primary shadow-[0_14px_28px_rgba(79,209,197,0.18)] hover:opacity-90"
			>
				Save & Apply Config
			</button>
		</div>
	</div>
</section>
