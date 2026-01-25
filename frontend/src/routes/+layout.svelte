<script lang="ts">
    import "./app.css"
    import favicon from "$lib/assets/favicon.ico"

    let { children } = $props()
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<header class="sticky top-0 z-50 w-full bg-gray-400 px-surround py-8 shadow-xl">
    <div class="mx-auto flex max-w-content items-center gap-4 text-3xl">
        <a href="/" class="w-28">
            <img src="./src/lib/assets/favicon.png" class="w-28" alt="LOGO" />
        </a>
        <a href="/">Home</a>
        <a href="/paintings">Catalog</a>
        <a href="/paintors">Paintors</a>
        <a href="/projects">Projects</a>
        <a href="/contact">Contact</a>
    </div>
</header>

{@render children()}

<style>
</style>
