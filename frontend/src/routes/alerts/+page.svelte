<script lang="ts">
	const summaryCards = [
		{ label: 'Critical Alerts', value: '04', delta: '+2 since 1h', accent: '#ffb4ab' },
		{ label: 'Warnings', value: '12', delta: '-5 since 1h', accent: '#ffd1af' },
		{ label: 'System Health', value: 'NOMINAL', delta: 'Stability holding', accent: '#6feee1' }
	];

	const alertRows = [
		{
			time: '10:45 AM',
			severity: 'CRITICAL',
			metric: 'CPU',
			description: 'Kernel panic on srv-01',
			status: 'UNRESOLVED',
			dot: '#ffb4ab'
		},
		{
			time: '10:30 AM',
			severity: 'WARNING',
			metric: 'RAM',
			description: 'Swap usage > 80% on primary node',
			status: 'PENDING',
			dot: '#ffd1af'
		},
		{
			time: '10:15 AM',
			severity: 'RESOLVED',
			metric: 'DISK',
			description: 'Cleaned temporary logs from staging cluster',
			status: 'ARCHIVED',
			dot: '#6feee1'
		}
	];
</script>

<svelte:head>
	<title>AIOps Control | Alerts</title>
</svelte:head>

<section class="space-y-8">
	<div class="flex flex-col gap-4 md:flex-row md:items-end md:justify-between">
		<div>
			<p class="font-label text-xs uppercase tracking-[0.28em] text-primary/80">
				Monitoring real-time infrastructure pulse
			</p>
			<h1 class="mt-2 font-headline text-4xl font-black tracking-tight">Event Logs</h1>
		</div>

		<div class="flex flex-wrap gap-3">
			<button
				class="flex items-center gap-2 rounded-md border border-outline-variant/25 bg-surface-container-high px-4 py-2 font-label text-[0.68rem] font-semibold uppercase tracking-[0.22em] hover:bg-surface-container-highest"
			>
				<span class="material-symbols-outlined text-base">filter_alt</span>
				Filter Logs
			</button>
			<button
				class="flex items-center gap-2 rounded-md bg-gradient-to-br from-primary to-primary-container px-4 py-2 font-label text-[0.68rem] font-bold uppercase tracking-[0.22em] text-on-primary hover:opacity-90"
			>
				<span class="material-symbols-outlined text-base">sync</span>
				Refresh Stream
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
						<div class="flex items-end gap-1">
							{#each [40, 72, 28, 58, 86, 48] as height}
								<span
									class="block w-1.5 rounded-t-full bg-primary/40"
									style={`height:${height}px`}
								></span>
							{/each}
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
					Console Output :: Alerts_Stream_v4.0
				</span>
			</div>

			<div class="flex gap-4 font-label text-[0.62rem] uppercase tracking-[0.18em] text-primary/60">
				<span>Rows: 003</span>
				<span>Status: Connected</span>
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
					{#each alertRows as row}
						<tr>
							<td class="px-5 py-5 font-label text-sm text-on-surface/80">{row.time}</td>
							<td class="px-5 py-5">
								<span
									class="rounded-md border px-2 py-1 font-label text-[0.62rem] font-bold uppercase tracking-[0.18em]"
									style={`color:${row.dot};border-color:${row.dot}33;background:${row.dot}12`}
								>
									{row.severity}
								</span>
							</td>
							<td class="px-5 py-5 font-label text-sm font-bold">{row.metric}</td>
							<td class="px-5 py-5 text-sm text-on-surface-variant">
								<div class="flex items-center gap-3">
									<span
										class="h-2 w-2 rounded-full"
										style={`background:${row.dot}`}
									></span>
									<span>{row.description}</span>
								</div>
							</td>
							<td class="px-5 py-5">
								<span
									class="font-label text-[0.62rem] font-bold uppercase tracking-[0.18em]"
									style:color={row.dot}
								>
									{row.status}
								</span>
							</td>
							<td class="px-5 py-5 text-right">
								<div class="flex justify-end gap-2">
									<button
										class="rounded-md bg-outline-variant/20 px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface transition-colors hover:bg-outline-variant/30"
									>
										Resolve
									</button>
									<button
										class="rounded-md bg-primary/10 px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.18em] text-primary transition-colors hover:bg-primary/20"
									>
										AI-Explain
									</button>
								</div>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>

		<div
			class="flex flex-col gap-3 border-t border-outline-variant/5 bg-surface-container-lowest/50 px-5 py-4 sm:flex-row sm:items-center sm:justify-between"
		>
			<div class="flex flex-wrap gap-6 font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
				<div class="flex items-center gap-2">
					<span class="status-dot"></span>
					<span>Stream: Active</span>
				</div>
				<div class="flex items-center gap-2">
					<span class="inline-block h-2 w-2 rounded-full bg-outline"></span>
					<span>Latency: 12ms</span>
				</div>
			</div>

			<div class="flex items-center gap-3">
				<button class="material-symbols-outlined text-on-surface-variant transition-colors hover:text-primary">
					chevron_left
				</button>
				<span class="rounded-md bg-primary/10 px-3 py-1 font-label text-[0.62rem] uppercase tracking-[0.18em] text-primary">
					Page 01
				</span>
				<button class="material-symbols-outlined text-on-surface-variant transition-colors hover:text-primary">
					chevron_right
				</button>
			</div>
		</div>
	</section>

	<div class="grid grid-cols-1 gap-5 lg:grid-cols-3">
		<article class="panel p-5">
			<h2 class="font-label text-[0.65rem] uppercase tracking-[0.24em] text-primary/80">
				Neural Diagnostics
			</h2>
			<div class="mt-6">
				<div class="h-1.5 overflow-hidden rounded-full bg-surface-container-high">
					<div
						class="h-full rounded-full bg-primary"
						style="width:85%"
					></div>
				</div>
				<div class="mt-3 flex items-center justify-between font-label text-[0.68rem] uppercase tracking-[0.18em] text-on-surface-variant">
					<span>Pattern Match Accuracy</span>
					<span>98.2%</span>
				</div>
			</div>
		</article>

		<article class="panel p-5">
			<h2 class="font-label text-[0.65rem] uppercase tracking-[0.24em] text-primary/80">Uptime Vector</h2>
			<div class="mt-6 flex h-[4.5rem] items-end justify-between gap-2 px-2">
				{#each [32, 40, 24, 48, 40, 36, 44, 28] as height}
					<span class="w-1.5 rounded-t-full bg-primary/35" style={`height:${height}px`}></span>
				{/each}
			</div>
		</article>

		<article class="panel flex items-center gap-4 p-5">
			<div
				class="flex h-14 w-14 items-center justify-center rounded-xl border border-primary/20 bg-primary/10 text-primary"
			>
				<span class="material-symbols-outlined text-2xl">auto_awesome</span>
			</div>
			<div>
				<div class="font-label text-[0.65rem] uppercase tracking-[0.24em] text-primary/80">AI Agent</div>
				<p class="mt-1 text-sm text-on-surface-variant">
					Optimization script running in background.
				</p>
			</div>
		</article>
	</div>
</section>
