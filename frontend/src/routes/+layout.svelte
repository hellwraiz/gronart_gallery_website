<script lang="ts">
    import "./app.css"
    import favicon from "$lib/assets/favicon.ico"

    let { children } = $props()
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<header class="sticky w-full bg-gray-400 px-24 py-8">
    <div class="mx-auto flex max-w-content">
        <a href="/" class="w-28">
            <img src="./src/lib/assets/favicon.png" class="w-28" alt="LOGO" />
        </a>
    </div>
</header>
{@render children()}

<style>
</style>
