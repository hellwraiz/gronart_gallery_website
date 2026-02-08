<script lang="ts">
    import "./app.css"

    let { children } = $props()
</script>

<svelte:head><link rel="icon" href="/favicon.png" /></svelte:head>
<header class="sticky top-0 z-50 w-full bg-gray-200 px-surround py-8 shadow-xl">
    <div class="mx-auto flex max-w-content items-center gap-10 text-2xl">
        <a href="/" class="w-28">
            <img src="/favicon.png" class="w-28" alt="LOGO" />
        </a>
        <a href="/">Home</a>
        <a href="/">Catalog</a>
        <!-- <a href="/paintings">Catalog</a> -->
        <!-- <a href="/paintors">Paintors</a> -->
        <!-- <a href="/projects">Projects</a> -->
        <a href="/">Contact</a>
    </div>
</header>

{@render children()}

<style>
</style>
