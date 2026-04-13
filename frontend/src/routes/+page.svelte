<script lang="ts">
	const metrics = [
		{
			label: 'Central Processor',
			value: '75.4%',
			icon: 'memory',
			color: '#ffb77f',
			left: 'CORE_LOAD_SPIKE',
			right: 'NODE-04 ACTIVE',
			bars: [24, 38, 32, 64, 78, 92]
		},
		{
			label: 'Memory Usage',
			value: '40.2%',
			icon: 'developer_board',
			color: '#6feee1',
			left: 'BUFFER_HEALTHY',
			right: '64GB ECC LPDDR5',
			bars: [52, 52, 58, 50, 48, 56]
		},
		{
			label: 'SSD Pool',
			value: '60.1%',
			icon: 'database',
			color: '#5adace',
			left: 'NVME_ARRAY_01',
			right: 'STABLE',
			bars: [62, 64, 64, 72, 80, 84]
		}
	];

	const intelligenceFeed = [
		{
			level: 'Critical',
			ago: '2M AGO',
			text: 'High CPU spike detected on Node-04. Intelligent throttling initiated to prevent cascade failure.',
			accent: 'text-error',
			border: 'border-error/50',
			button: 'RESOLVE'
		},
		{
			level: 'Warning',
			ago: '10M AGO',
			text: 'Memory leakage suspected in Container-X. Garbage collection cycle performed manually.',
			accent: 'text-tertiary',
			border: 'border-tertiary/50',
			button: 'INSPECT'
		},
		{
			level: 'Info',
			ago: '1H AGO',
			text: 'System backup completed successfully. Snapshot #8842 archived to primary-cold-storage.',
			accent: 'text-primary',
			border: 'border-primary/50',
			button: 'ARCHIVED'
		}
	];
</script>

<svelte:head>
	<title>AIOps Control | Dashboard</title>
</svelte:head>

<section class="space-y-8">
	<div class="flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
		<div>
			<p class="font-label text-xs uppercase tracking-[0.28em] text-primary/70">
				Real-time Telemetry & Neural Node Status
			</p>
			<h1 class="mt-2 font-headline text-4xl font-black tracking-tight text-on-surface">
				System Overview
			</h1>
		</div>

		<div class="flex flex-wrap gap-3">
			<button
				class="rounded-md border border-outline-variant/25 bg-surface-container-high px-4 py-2 font-label text-[0.7rem] font-semibold uppercase tracking-[0.26em] text-on-surface transition-colors hover:bg-surface-container-highest"
			>
				Export Logs
			</button>
			<button
				class="rounded-md bg-gradient-to-br from-primary to-primary-container px-4 py-2 font-label text-[0.7rem] font-bold uppercase tracking-[0.26em] text-on-primary shadow-[0_14px_28px_rgba(79,209,197,0.18)] transition-opacity hover:opacity-90"
			>
				Optimize Cluster
			</button>
		</div>
	</div>

	<div class="grid grid-cols-1 gap-5 xl:grid-cols-12">
		<div class="grid grid-cols-1 gap-5 md:grid-cols-3 xl:col-span-12">
			{#each metrics as metric}
				<article class="panel metric-card p-6" style={`--accent:${metric.color}`}>
					<div class="relative z-10 flex items-start justify-between gap-3">
						<div>
							<p class="font-label text-[0.65rem] uppercase tracking-[0.26em] text-on-surface-variant">
								{metric.label}
							</p>
							<h2 class="mt-2 font-label text-3xl font-bold" style:color={metric.color}>
								{metric.value}
							</h2>
						</div>
						<span class="material-symbols-outlined text-3xl" style:color={metric.color}>
							{metric.icon}
						</span>
					</div>

					<div class="metric-bars relative z-10 mt-6">
						{#each metric.bars as bar}
							<span style={`height:${bar}%`}></span>
						{/each}
					</div>

					<div
						class="relative z-10 mt-4 flex items-center justify-between font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant/70"
					>
						<span>{metric.left}</span>
						<span>{metric.right}</span>
					</div>
				</article>
			{/each}
		</div>

		<div class="panel relative xl:col-span-8">
			<div class="p-6 sm:p-8">
				<div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
					<div>
						<h2 class="font-headline text-2xl font-bold">Network Traffic Flow</h2>
						<p class="mt-1 font-label text-[0.65rem] uppercase tracking-[0.24em] text-primary/80">
							Global Distribution Edge
						</p>
					</div>

					<div class="flex flex-wrap gap-4 font-label text-[0.68rem] uppercase tracking-[0.18em]">
						<div class="flex items-center gap-2 text-on-surface-variant">
							<span class="status-dot pulse-ring"></span>
							<span>TX: 8.4 GB/s</span>
						</div>
						<div class="flex items-center gap-2 text-on-surface-variant">
							<span class="inline-block h-2 w-2 rounded-full bg-primary/35"></span>
							<span>RX: 2.1 GB/s</span>
						</div>
					</div>
				</div>

				<div class="network-map relative mt-8 overflow-hidden rounded-2xl border border-outline-variant/10 p-6">
					<div class="relative z-10 grid min-h-72 place-items-center gap-6 lg:grid-cols-[1fr_auto_1fr]">
						<div class="flex flex-col items-center gap-3">
							<div
								class="world-node flex h-[4.5rem] w-[4.5rem] items-center justify-center rounded-full border border-primary/25 bg-primary/10"
							>
								<span class="material-symbols-outlined text-3xl text-primary">hub</span>
							</div>
							<span class="font-label text-xs uppercase tracking-[0.24em] text-primary/80">US-EAST</span>
						</div>

						<div class="flex items-center gap-4">
							<div class="hidden h-px w-14 bg-gradient-to-r from-primary/0 via-primary to-primary/0 sm:block"></div>
							<div
								class="world-node flex h-24 w-24 items-center justify-center rounded-full border-2 border-primary/45 bg-primary/10 shadow-[0_0_36px_rgba(111,238,225,0.16)]"
							>
								<span class="material-symbols-outlined text-5xl text-primary">cloud</span>
							</div>
							<div class="hidden h-px w-14 bg-gradient-to-r from-primary/0 via-primary to-primary/0 sm:block"></div>
						</div>

						<div class="flex flex-col items-center gap-3">
							<div
								class="world-node flex h-[4.5rem] w-[4.5rem] items-center justify-center rounded-full border border-primary/25 bg-primary/10"
							>
								<span class="material-symbols-outlined text-3xl text-primary">dns</span>
							</div>
							<span class="font-label text-xs uppercase tracking-[0.24em] text-primary/80">EU-WEST</span>
						</div>
					</div>
				</div>
			</div>
		</div>

		<div class="panel-soft xl:col-span-4">
			<div class="flex h-full flex-col justify-between p-6 sm:p-8">
				<div>
					<div class="flex items-center gap-3">
						<div class="rounded-lg bg-primary/10 p-2">
							<span class="material-symbols-outlined text-primary">token</span>
						</div>
						<h2 class="font-headline text-xl font-bold">Active Containers</h2>
					</div>

					<div class="mt-8 space-y-5">
						<div class="flex items-center justify-between font-label text-sm">
							<span class="text-on-surface-variant">Orchestration</span>
							<span class="font-bold text-primary">KUBERNETES-V3</span>
						</div>
						<div class="flex items-center justify-between font-label text-sm">
							<span class="text-on-surface-variant">Live Instances</span>
							<span class="font-bold text-primary">20 / 20</span>
						</div>
						<div class="h-2 overflow-hidden rounded-full bg-surface-container-lowest">
							<div class="h-full w-full rounded-full bg-primary"></div>
						</div>
					</div>
				</div>

				<div class="mt-10 border-t border-outline-variant/10 pt-8">
					<div class="mb-3 flex items-center justify-between font-label text-[0.68rem] uppercase tracking-[0.18em]">
						<span class="text-on-surface-variant">Health Index</span>
						<span class="font-semibold text-primary">98.2% Optimal</span>
					</div>
					<button
						class="w-full rounded-lg bg-surface-container-highest px-4 py-3 font-label text-[0.68rem] font-semibold uppercase tracking-[0.22em] text-on-surface transition-colors hover:text-primary"
					>
						Manage Cluster
					</button>
				</div>
			</div>
		</div>
	</div>

	<section class="space-y-5">
		<div class="flex items-center gap-3">
			<span class="material-symbols-outlined text-primary">terminal</span>
			<h2 class="font-headline text-2xl font-bold tracking-tight">Intelligence Feed</h2>
		</div>

		<div class="space-y-3">
			{#each intelligenceFeed as item}
				<article
					class={`panel flex flex-col gap-4 border-l-2 p-4 sm:flex-row sm:items-center sm:gap-6 ${item.border}`}
				>
					<div class="min-w-16 font-label text-[0.65rem] uppercase tracking-[0.18em]">
						<div class={`font-bold ${item.accent}`}>{item.level}</div>
						<div class="mt-1 text-on-surface-variant">{item.ago}</div>
					</div>

					<span class={`material-symbols-outlined text-2xl ${item.accent}`}>smart_toy</span>

					<p class="flex-1 text-sm leading-relaxed text-on-surface-variant">{item.text}</p>

					<button
						class="rounded-md border border-outline-variant/20 px-3 py-2 font-label text-[0.62rem] font-semibold uppercase tracking-[0.22em] text-on-surface transition-colors hover:bg-surface-container-high"
					>
						{item.button}
					</button>
				</article>
			{/each}
		</div>
	</section>

	<div class="flex justify-end">
		<div
			class="rounded-xl border border-outline-variant/15 bg-surface-container-highest/70 px-5 py-3 shadow-[0_12px_28px_rgba(0,0,0,0.24)] backdrop-blur"
		>
			<div class="flex flex-wrap items-center gap-4 font-label text-[0.64rem] uppercase tracking-[0.22em]">
				<div class="flex items-center gap-2">
					<span class="status-dot pulse-ring"></span>
					<span>SYS_UPTIME: 14D:02H:44M</span>
				</div>
				<div class="h-4 w-px bg-outline-variant/30"></div>
				<div class="flex items-center gap-2">
					<span class="material-symbols-outlined text-sm text-primary">cloud_done</span>
					<span>SYNC: SECURE</span>
				</div>
			</div>
		</div>
	</div>
</section>
