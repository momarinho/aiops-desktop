<svelte:head>
	<title>AIOps Control | Assistant</title>
</svelte:head>

<section class="flex min-h-[calc(100vh-9rem)] flex-col overflow-hidden rounded-[1.75rem] border border-outline-variant/10 bg-surface shadow-[0_24px_56px_rgba(0,0,0,0.22)]">
	<div
		class="flex flex-col gap-4 border-b border-outline-variant/10 bg-surface-container-low px-5 py-4 sm:flex-row sm:items-center sm:justify-between sm:px-8"
	>
		<div class="flex items-center gap-3">
			<span class="h-2.5 w-2.5 rounded-full bg-error animate-pulse"></span>
			<span class="font-label text-xs font-bold uppercase tracking-[0.24em] text-error">
				Current Context:
			</span>
			<h1 class="text-sm italic text-on-surface-variant">"Alert: High CPU on Node-04"</h1>
		</div>
		<div class="flex flex-wrap gap-2">
			<button
				class="rounded-md bg-surface-container-high px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.2em] text-on-surface-variant transition-colors hover:bg-surface-variant"
			>
				Details
			</button>
			<button
				class="rounded-md bg-error-container/20 px-3 py-2 font-label text-[0.62rem] uppercase tracking-[0.2em] text-error transition-colors hover:bg-error-container/30"
			>
				Mute Alert
			</button>
		</div>
	</div>

	<div class="thin-scrollbar flex-1 space-y-10 overflow-y-auto px-5 py-8 sm:px-8">
		<div class="mx-auto flex max-w-4xl gap-5">
			<div
				class="flex h-11 w-11 shrink-0 items-center justify-center rounded-md bg-gradient-to-br from-primary to-primary-container text-on-primary"
			>
				<span class="material-symbols-outlined" style:font-variation-settings="'FILL' 1">
					smart_toy
				</span>
			</div>
			<div class="flex-1 space-y-6">
				<div class="border-l-2 border-primary pl-5">
					<p class="text-lg leading-relaxed text-on-surface">
						I've analyzed the logs for
						<span class="font-label font-bold text-primary">Node-04</span>. The process
						<code class="rounded bg-surface-container-highest px-1.5 py-0.5 text-sm text-primary-fixed-dim">
							db_worker
						</code>
						is consuming 95% CPU due to an unindexed query. You might want to run the
						optimization script below.
					</p>
				</div>

				<div class="panel-deep overflow-hidden">
					<div
						class="flex items-center justify-between bg-surface-container-high px-4 py-3 sm:px-5"
					>
						<span class="font-label text-[0.65rem] uppercase tracking-[0.24em] text-on-surface-variant">
							SQL Index Optimization
						</span>
						<span class="font-label text-[0.62rem] uppercase tracking-[0.18em] text-primary/60">
							PostgreSQL
						</span>
					</div>
					<div class="bg-surface-container-low/50 p-6 font-label leading-7 text-primary-fixed-dim">
						<pre class="whitespace-pre-wrap">CREATE INDEX idx_user_id ON orders(user_id);</pre>
					</div>
					<div
						class="flex flex-col gap-3 border-t border-outline-variant/10 bg-surface-container-high/30 p-4 sm:flex-row"
					>
						<button
							class="flex items-center justify-center gap-2 rounded-md bg-surface-container-highest px-4 py-3 font-label text-xs uppercase tracking-[0.2em] transition-colors hover:bg-surface-variant"
						>
							<span class="material-symbols-outlined text-base">content_copy</span>
							Copy fix to clipboard
						</button>
						<button
							class="flex items-center justify-center gap-2 rounded-md bg-gradient-to-br from-primary to-primary-container px-5 py-3 font-label text-xs font-bold uppercase tracking-[0.22em] text-on-primary transition-opacity hover:opacity-90"
						>
							<span class="material-symbols-outlined text-base" style:font-variation-settings="'FILL' 1">
								bolt
							</span>
							Execute Fix
						</button>
					</div>
				</div>
			</div>
		</div>

		<div class="mx-auto flex max-w-4xl gap-5 opacity-65">
			<div
				class="flex h-11 w-11 shrink-0 items-center justify-center rounded-md bg-surface-container-high text-on-surface-variant"
			>
				<span class="material-symbols-outlined">terminal</span>
			</div>
			<div class="flex flex-col justify-center gap-2">
				<div class="flex items-center gap-3">
					<span class="font-label text-xs font-medium uppercase tracking-[0.2em]">
						Architect is thinking
					</span>
					<div class="flex gap-1">
						<span class="h-1.5 w-1.5 animate-bounce rounded-full bg-primary"></span>
						<span class="h-1.5 w-1.5 animate-bounce rounded-full bg-primary [animation-delay:0.2s]"></span>
						<span class="h-1.5 w-1.5 animate-bounce rounded-full bg-primary [animation-delay:0.4s]"></span>
					</div>
				</div>
				<p class="font-label text-[0.65rem] uppercase tracking-[0.18em] text-on-surface-variant/70">
					Querying active transaction logs...
				</p>
			</div>
		</div>
	</div>

	<div class="border-t border-outline-variant/10 bg-gradient-to-t from-background to-transparent px-5 py-6 sm:px-8">
		<div class="mx-auto max-w-4xl">
			<label class="group relative block">
				<span
					class="material-symbols-outlined pointer-events-none absolute left-5 top-1/2 -translate-y-1/2 text-primary/60 transition-colors group-focus-within:text-primary"
				>
					auto_awesome
				</span>
				<input
					class="w-full rounded-t-2xl border-0 border-b-2 border-surface-container-highest bg-surface-container-high py-5 pl-14 pr-32 text-lg text-on-surface placeholder:text-on-surface-variant/40 focus:border-primary focus:bg-surface-container focus:ring-0"
					placeholder="Ask AI"
					type="text"
				/>
				<div class="absolute inset-y-0 right-4 flex items-center gap-3">
					<kbd
						class="hidden items-center gap-1 rounded bg-surface-container-highest px-2 py-1 font-label text-[0.62rem] text-on-surface-variant md:flex"
					>
						<span>CTRL</span>
						<span>+</span>
						<span>ENTER</span>
					</kbd>
					<button
						class="rounded-lg bg-primary/10 p-2 text-primary transition-colors hover:bg-primary/20"
						type="button"
					>
						<span class="material-symbols-outlined">send</span>
					</button>
				</div>
			</label>
			<p class="mt-4 text-center font-label text-[0.62rem] uppercase tracking-[0.28em] text-on-surface-variant/45">
				Synthetic Intelligence Node: Active
			</p>
		</div>
	</div>
</section>
