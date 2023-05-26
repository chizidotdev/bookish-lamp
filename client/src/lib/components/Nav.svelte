<script lang="ts">
	import { enhance, type SubmitFunction } from '$app/forms';
	import { page } from '$app/stores';
	import '$lib/app.css';

	const submitUpdateTheme: SubmitFunction = ({ action }) => {
		const theme = action.searchParams.get('theme');

		if (theme) {
			document.documentElement.setAttribute('data-theme', theme);
		}
	};

	const themes = ['light', 'dark', 'cupcake', 'dracula', 'business', 'night'];

    const handleLogin = () => {
        fetch(`http://localhost:8080/login`)
    }
</script>

<div class="bg-base-100 sticky top-0 shadow-lg z-40">
	<div class="navbar container mx-auto px-2 sm:px-5 flex items-center justify-between p-2">
		<div class="navbar bg-base-100">
			<div class="navbar-start">
				<div class="dropdown">
					<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
					<!-- svelte-ignore a11y-label-has-associated-control -->
					<label tabindex="0" class="btn btn-ghost lg:hidden">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							class="h-5 w-5"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
							><path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M4 6h16M4 12h8m-8 6h16"
							/></svg
						>
					</label>
					<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
					<ul
						tabindex="0"
						class="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52"
					>
						<li><a href="/items">Items</a></li>
						<li><a href="/">Community</a></li>
						<li><a href="/">Pricing</a></li>
					</ul>
				</div>
				<a href="/" class="normal-case text-xl">Copia</a>
			</div>
			<div class="navbar-center hidden lg:flex">
				<ul class="menu menu-horizontal px-1">
					<li><a href="/items">Items</a></li>
					<li><a href="/">Community</a></li>
					<li><a href="/">Pricing</a></li>
				</ul>
			</div>
			<div class="navbar-end gap-2">
				<btn class="btn btn-outline" on:click|preventDefault={handleLogin}>Login</btn>
				<a href="/signup" class="btn">Get started</a>
			</div>
		</div>

		<div class="">
			<div class="dropdown dropdown-hover dropdown-end">
				<!-- svelte-ignore a11y-label-has-associated-control -->
				<label
					><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
						><path
							fill="none"
							stroke="currentColor"
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M9 21h6M12 3a6 6 0 0 0-5.019 9.29c.954 1.452 1.43 2.178 1.493 2.286c.55.965.449.625.518 1.734c.008.124.008.313.008.69a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1c0-.377 0-.566.008-.69c.07-1.11-.033-.769.518-1.734c.062-.108.54-.834 1.493-2.287A6 6 0 0 0 12 3Z"
						/></svg
					></label
				>
				<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
				<ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 h-max overflow-scroll">
					<form method="POST" use:enhance={submitUpdateTheme}>
						{#each themes as theme}
							<li>
								<button formaction="/?/setTheme&theme={theme}&redirectTo={$page.url.pathname}"
									>{theme}</button
								>
							</li>
						{/each}
					</form>
				</ul>
			</div>
		</div>
	</div>
</div>
