<script lang="ts">
	const historyRows = [
		{
			timestamp: '2023.10.24 14:22:01',
			source: 'K8S_CLUSTER_MAIN',
			alert: '#ERR_MEM_LEAK_04',
			action: 'Node Restarted via AI-AutoScaling Engine',
			latency: '142',
			status: 'RESOLVED'
		},
		{
			timestamp: '2023.10.24 13:45:12',
			source: 'DB_INSTANCE_SQL_09',
			alert: '#WAR_SLOW_QUERY',
			action: 'Auto-Index Optimization Applied',
			latency: '890',
			status: 'RESOLVED'
		},
		{
			timestamp: '2023.10.24 11:02:44',
			source: 'API_GATEWAY_V2',
			alert: '#SEC_AUTH_FLOOD',
			action: 'IP Range Blocked (Temporal Quarantine)',
			latency: '12',
			status: 'RESOLVED'
		},
		{
			timestamp: '2023.10.24 09:15:33',
			source: 'STORAGE_FS_PROD',
			alert: '#IO_THROUGHPUT_LOW',
			action: 'Traffic Rerouted to Backup Volume',
			latency: '2140',
			status: 'RESOLVED'
		}
	];
</script>

<svelte:head>
	<title>AIOps Control | History</title>
</svelte:head>

<section class="space-y-8">
	<header class="flex flex-col gap-5 md:flex-row md:items-end md:justify-between">
		<div>
			<h1 class="font-headline text-4xl font-black tracking-tight">System History</h1>
			<p class="mt-2 font-label text-xs uppercase tracking-[0.28em] text-on-surface-variant/70">
				Temporal Intelligence Analysis
			</p>
		</div>

		<div class="flex flex-wrap items-center gap-2 rounded-2xl bg-surface-container-low p-1">
			{#each ['1H', '6H', '24H', '7D'] as range, index}
				<button
					class={`rounded-xl px-4 py-2 font-label text-xs font-medium uppercase tracking-[0.18em] transition-colors ${
						index === 0
							? 'border border-primary/20 bg-surface-container-high text-primary'
							: 'text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface'
					}`}
				>
					{range}
				</button>
			{/each}
			<button
				class="ml-1 flex h-9 w-9 items-center justify-center rounded-xl text-on-surface-variant transition-colors hover:bg-surface-container-high hover:text-primary"
			>
				<span class="material-symbols-outlined text-base">calendar_today</span>
			</button>
		</div>
	</header>

	<div class="grid grid-cols-1 gap-5 xl:grid-cols-12">
		<article class="panel grid-overlay relative overflow-hidden p-6 xl:col-span-8">
			<div class="relative z-10 flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
				<div>
					<h2 class="font-headline text-2xl font-bold">Resource Utilization</h2>
					<div class="mt-3 flex flex-wrap gap-4">
						<div class="flex items-center gap-2 font-label text-[0.62rem] uppercase tracking-[0.2em] text-on-surface-variant">
							<span class="h-2 w-2 rounded-full bg-primary"></span>
							CPU Core Alpha
						</div>
						<div class="flex items-center gap-2 font-label text-[0.62rem] uppercase tracking-[0.2em] text-on-surface-variant">
							<span class="h-2 w-2 rounded-full bg-tertiary"></span>
							Memory Stack
						</div>
					</div>
				</div>
				<div class="text-left sm:text-right">
					<div class="font-label text-3xl font-bold text-primary">42.8%</div>
					<div class="font-label text-[0.65rem] uppercase tracking-[0.18em] text-on-surface-variant">
						Peak Load Avg
					</div>
				</div>
			</div>

			<div class="relative mt-8 h-72">
				<svg class="h-full w-full overflow-visible" viewBox="0 0 800 200" preserveAspectRatio="none">
					<defs>
						<linearGradient id="cpuArea" x1="0%" x2="0%" y1="0%" y2="100%">
							<stop offset="0%" stop-color="#6feee1" stop-opacity="0.26"></stop>
							<stop offset="100%" stop-color="#6feee1" stop-opacity="0"></stop>
						</linearGradient>
					</defs>
					<path
						d="M0,150 Q50,140 100,160 T200,120 T300,140 T400,100 T500,130 T600,110 T700,150 T800,130"
						fill="none"
						stroke="#6feee1"
						stroke-width="2.5"
					></path>
					<path
						d="M0,100 Q100,110 200,90 T400,100 T600,80 T800,90"
						fill="none"
						stroke="#ffd1af"
						stroke-dasharray="4 2"
						stroke-width="2.5"
					></path>
					<path
						d="M0,150 Q50,140 100,160 T200,120 T300,140 T400,100 T500,130 T600,110 T700,150 T800,130 L800,200 L0,200 Z"
						fill="url(#cpuArea)"
					></path>
				</svg>
				<div
					class="pointer-events-none absolute inset-x-0 bottom-2 flex justify-between font-label text-[0.58rem] uppercase tracking-[0.16em] text-on-surface-variant/35"
				>
					<span>00:00</span>
					<span>02:00</span>
					<span>04:00</span>
					<span>06:00</span>
					<span>08:00</span>
					<span>10:00</span>
					<span>12:00</span>
				</div>
			</div>
		</article>

		<article class="panel p-6 xl:col-span-4">
			<div class="flex items-center justify-between">
				<h2 class="font-headline text-xl font-bold">Network I/O</h2>
				<span class="material-symbols-outlined text-primary">wifi_tethering</span>
			</div>

			<div class="mt-6 space-y-5">
				<div>
					<div
						class="mb-2 flex items-center justify-between font-label text-[0.68rem] uppercase tracking-[0.18em] text-on-surface-variant"
					>
						<span>Ingress</span>
						<span class="text-primary">1.2 GB/s</span>
					</div>
					<div class="h-1.5 overflow-hidden rounded-full bg-surface-container-highest">
						<div class="h-full w-2/3 rounded-full bg-primary"></div>
					</div>
				</div>
				<div>
					<div
						class="mb-2 flex items-center justify-between font-label text-[0.68rem] uppercase tracking-[0.18em] text-on-surface-variant"
					>
						<span>Egress</span>
						<span class="text-tertiary">480 MB/s</span>
					</div>
					<div class="h-1.5 overflow-hidden rounded-full bg-surface-container-highest">
						<div class="h-full w-1/3 rounded-full bg-tertiary"></div>
					</div>
				</div>
			</div>

			<div class="grid-overlay mt-8 overflow-hidden rounded-xl border border-outline-variant/10 p-3">
				<svg class="h-36 w-full" viewBox="0 0 200 100" preserveAspectRatio="none">
					<polyline
						fill="none"
						points="0,80 20,75 40,85 60,40 80,50 100,20 120,45 140,30 160,60 180,55 200,70"
						stroke="#6feee1"
						stroke-width="1.5"
					></polyline>
				</svg>
			</div>
		</article>
	</div>

	<section class="terminal-window overflow-hidden">
		<div
			class="flex flex-col gap-3 border-b border-outline-variant/10 px-5 py-5 sm:flex-row sm:items-center sm:justify-between"
		>
			<div class="flex items-center gap-3">
				<span class="material-symbols-outlined text-on-surface-variant">event_available</span>
				<h2 class="font-headline text-xl font-bold uppercase tracking-tight">
					Resolved Intelligence Logs
				</h2>
			</div>
			<div class="flex flex-wrap gap-3">
				<button
					class="flex items-center gap-2 rounded-md border border-outline-variant/25 px-4 py-2 font-label text-[0.65rem] font-bold uppercase tracking-[0.18em] text-on-surface transition-colors hover:bg-surface-container-high"
				>
					<span class="material-symbols-outlined text-base">download</span>
					Export CSV
				</button>
				<button
					class="flex items-center gap-2 rounded-md border border-outline-variant/25 px-4 py-2 font-label text-[0.65rem] font-bold uppercase tracking-[0.18em] text-on-surface transition-colors hover:bg-surface-container-high"
				>
					<span class="material-symbols-outlined text-base">code</span>
					Export JSON
				</button>
			</div>
		</div>

		<div class="thin-scrollbar overflow-x-auto">
			<table class="min-w-full text-left">
				<thead class="bg-surface-container-lowest/50">
					<tr>
						<th class="px-5 py-4 font-label text-[0.65rem] font-semibold uppercase tracking-[0.24em] text-on-surface-variant">
							Timestamp
						</th>
						<th class="px-5 py-4 font-label text-[0.65rem] font-semibold uppercase tracking-[0.24em] text-on-surface-variant">
							Event Source
						</th>
						<th class="px-5 py-4 font-label text-[0.65rem] font-semibold uppercase tracking-[0.24em] text-on-surface-variant">
							Alert ID
						</th>
						<th class="px-5 py-4 font-label text-[0.65rem] font-semibold uppercase tracking-[0.24em] text-on-surface-variant">
							Resolution Action
						</th>
						<th class="px-5 py-4 font-label text-[0.65rem] font-semibold uppercase tracking-[0.24em] text-on-surface-variant">
							Lat (ms)
						</th>
						<th class="px-5 py-4 text-right font-label text-[0.65rem] font-semibold uppercase tracking-[0.24em] text-on-surface-variant">
							Status
						</th>
					</tr>
				</thead>
				<tbody class="divide-y divide-outline-variant/10">
					{#each historyRows as row}
						<tr class="transition-colors hover:bg-surface-container-highest/20">
							<td class="px-5 py-4 font-label text-xs text-on-surface/80">{row.timestamp}</td>
							<td class="px-5 py-4 font-label text-xs">{row.source}</td>
							<td class="px-5 py-4 font-label text-xs text-primary-fixed-dim">{row.alert}</td>
							<td class="px-5 py-4 text-xs text-on-surface/60">{row.action}</td>
							<td class="px-5 py-4 font-label text-xs text-on-surface/40">{row.latency}</td>
							<td class="px-5 py-4 text-right">
								<span
									class="rounded-md border border-primary/20 bg-primary/10 px-2 py-1 font-label text-[0.62rem] font-bold uppercase tracking-[0.18em] text-primary"
								>
									{row.status}
								</span>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>

		<div
			class="flex flex-col gap-3 border-t border-outline-variant/5 bg-surface-container-lowest/35 px-5 py-4 sm:flex-row sm:items-center sm:justify-between"
		>
			<span class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-on-surface-variant">
				Displaying 1-4 of 1,248 Records
			</span>
			<div class="flex items-center gap-2">
				<button
					class="flex h-9 w-9 items-center justify-center rounded-lg border border-outline-variant/20 text-on-surface-variant transition-colors hover:text-primary"
				>
					<span class="material-symbols-outlined text-base">chevron_left</span>
				</button>
				<span class="px-3 font-label text-xs text-primary">1</span>
				<button
					class="flex h-9 w-9 items-center justify-center rounded-lg border border-outline-variant/20 text-on-surface-variant transition-colors hover:text-primary"
				>
					<span class="material-symbols-outlined text-base">chevron_right</span>
				</button>
			</div>
		</div>
	</section>

	<div class="flex justify-end">
		<div
			class="rounded-2xl border border-outline-variant/15 bg-surface-container-high/80 p-4 backdrop-blur-xl"
		>
			<div class="flex items-center gap-4">
				<div class="text-right">
					<div class="font-label text-[0.62rem] uppercase tracking-[0.24em] text-on-surface-variant">
						System Pulse
					</div>
					<div class="font-label text-sm font-bold text-primary">STABLE_O1</div>
				</div>
				<div class="relative h-3 w-3 rounded-full bg-primary">
					<div class="pulse-ring absolute inset-0 rounded-full"></div>
				</div>
			</div>
		</div>
	</div>
</section>
