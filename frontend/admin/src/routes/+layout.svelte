<script>
	import '../app.css';
	import Sidebar from '$lib/resources/sidebar.svelte';
    import { page } from '$app/state';
	let { children } = $props();
</script>

{#if page.url.pathname.startsWith('/auth')}
	<main class="">
		{@render children()}
	</main>
{:else}
	<Sidebar />
	<main class="ml-[60px]">
		{@render children()}
	</main>
{/if}

