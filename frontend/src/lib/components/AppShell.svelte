<script lang="ts">
	import { page } from '$app/state';

	type NavItem = {
		href: string;
		label: string;
		icon: string;
	};

	const navigation: NavItem[] = [
		{ href: '/', label: 'Dashboard', icon: 'dashboard' },
		{ href: '/alerts', label: 'Alerts', icon: 'notifications' },
		{ href: '/settings', label: 'Settings', icon: 'settings' },
		{ href: '/history', label: 'History', icon: 'history' },
		{ href: '/assistant', label: 'Assistant', icon: 'smart_toy' }
	];

	const metaByPath: Record<string, { title: string; subtitle: string }> = {
		'/': {
			title: 'AIOps Control',
			subtitle: 'Connectivity'
		},
		'/alerts': {
			title: 'Alerts Stream',
			subtitle: 'Event Monitoring'
		},
		'/settings': {
			title: 'Automation Settings',
			subtitle: 'Config Control'
		},
		'/history': {
			title: 'Temporal Logs',
			subtitle: 'Historical Analysis'
		},
		'/assistant': {
			title: 'Synthetic Architect',
			subtitle: 'AI Operations'
		}
	};

	let { children } = $props();
	let mobileNavOpen = $state(false);

	const pathname = $derived(page.url.pathname);
	const currentMeta = $derived(metaByPath[pathname] ?? metaByPath['/']);

	function isActive(href: string, currentPath: string) {
		return href === '/' ? currentPath === '/' : currentPath.startsWith(href);
	}
</script>

<div class="min-h-screen bg-background text-on-surface shell-noise">
	{#if mobileNavOpen}
		<button
			class="fixed inset-0 z-40 bg-black/50 lg:hidden"
			type="button"
			aria-label="Close navigation"
			onclick={() => (mobileNavOpen = false)}
		></button>
	{/if}

	<aside
		class={`fixed inset-y-0 left-0 z-50 flex w-72 flex-col border-r border-outline-variant/15 bg-surface-container-low px-5 py-6 transition-transform duration-300 lg:translate-x-0 ${
			mobileNavOpen ? 'translate-x-0' : '-translate-x-full'
		}`}
	>
		<div class="px-3">
			<div class="flex items-center gap-3">
				<div
					class="flex h-10 w-10 items-center justify-center rounded-lg bg-primary-container text-on-primary"
				>
					<span class="material-symbols-outlined" style:font-variation-settings="'FILL' 1">
						architecture
					</span>
				</div>
				<div>
					<h1 class="font-headline text-sm font-black uppercase tracking-[0.26em] text-primary">
						Synthetic Architect
					</h1>
					<p class="mt-1 font-label text-[0.62rem] uppercase tracking-[0.26em] text-on-surface-variant/70">
						AI-Ops Terminal
					</p>
				</div>
			</div>
		</div>

		<nav class="mt-10 flex-1 space-y-1.5">
			{#each navigation as item}
				<a
					href={item.href}
					class={`flex items-center gap-3 rounded-xl px-4 py-3 font-body text-sm tracking-tight transition-all duration-200 ${
						isActive(item.href, pathname)
							? 'border-l-2 border-primary bg-surface-container-high font-semibold text-primary'
							: 'text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface'
					}`}
					onclick={() => (mobileNavOpen = false)}
				>
					<span class="material-symbols-outlined text-[1.15rem]">{item.icon}</span>
					<span>{item.label}</span>
				</a>
			{/each}
		</nav>

		<div class="mt-6 border-t border-outline-variant/10 pt-6">
			<div class="rounded-2xl border border-outline-variant/10 bg-surface-container p-4">
				<div class="flex items-center gap-3">
					<div
						class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary/10 text-primary"
					>
						<span class="material-symbols-outlined">admin_panel_settings</span>
					</div>
					<div class="min-w-0">
						<p class="truncate text-xs font-semibold text-on-surface">System Administrator</p>
						<p
							class="truncate font-label text-[0.62rem] uppercase tracking-[0.2em] text-on-surface-variant"
						>
							Level 04 Auth
						</p>
					</div>
				</div>
			</div>
		</div>
	</aside>

	<div class="lg:pl-72">
		<header
			class="sticky top-0 z-30 border-b border-outline-variant/10 bg-background/75 backdrop-blur-xl"
		>
			<div class="mx-auto flex h-[4.5rem] max-w-7xl items-center justify-between gap-4 px-4 sm:px-6 lg:px-8">
				<div class="flex items-center gap-4 sm:gap-8">
					<button
						class="flex h-10 w-10 items-center justify-center rounded-xl border border-outline-variant/20 bg-surface-container text-on-surface lg:hidden"
						type="button"
						aria-label="Open navigation"
						onclick={() => (mobileNavOpen = true)}
					>
						<span class="material-symbols-outlined">menu</span>
					</button>

					<div>
						<p class="font-headline text-lg font-black tracking-tight text-on-surface">
							{currentMeta.title}
						</p>
						<p
							class="font-label text-[0.62rem] uppercase tracking-[0.28em] text-on-surface-variant/80"
						>
							{currentMeta.subtitle}
						</p>
					</div>

					<nav class="hidden gap-6 lg:flex">
						<a
							href="/"
							class="font-label text-xs uppercase tracking-[0.22em] text-primary transition-opacity"
						>
							Connectivity
						</a>
						<a
							href="/history"
							class="font-label text-xs uppercase tracking-[0.22em] text-on-surface-variant/70 transition-colors hover:text-primary"
						>
							Active Containers
						</a>
					</nav>
				</div>

				<div class="flex items-center gap-3 sm:gap-4">
					<label class="relative hidden md:block">
						<span
							class="material-symbols-outlined pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-sm text-on-surface-variant"
						>
							search
						</span>
						<input
							class="w-56 rounded-xl border border-outline-variant/15 bg-surface-container-lowest py-2 pl-9 pr-4 font-label text-[0.68rem] uppercase tracking-[0.2em] text-on-surface placeholder:text-on-surface-variant/45 focus:border-primary focus:ring-primary/20"
							placeholder="Query system..."
							type="text"
						/>
					</label>

					<button
						class="flex h-10 w-10 items-center justify-center rounded-xl border border-outline-variant/15 bg-surface-container text-on-surface-variant transition-colors hover:text-primary"
						type="button"
					>
						<span class="material-symbols-outlined">account_tree</span>
					</button>
					<button
						class="flex h-10 w-10 items-center justify-center rounded-xl border border-outline-variant/15 bg-surface-container text-on-surface-variant transition-colors hover:text-primary"
						type="button"
					>
						<span class="material-symbols-outlined">sensors</span>
					</button>
				</div>
			</div>
		</header>

		<main class="px-4 py-6 sm:px-6 lg:px-8 lg:py-8">
			<div class="mx-auto max-w-7xl">
				{@render children()}
			</div>
		</main>
	</div>
</div>
