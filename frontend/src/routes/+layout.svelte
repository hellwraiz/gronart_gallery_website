<script lang="ts">
    import "./app.css"
    import { slide } from "svelte/transition"

    let { children } = $props()
    let open = $state(false)
</script>

<svelte:head>
    <link rel="icon" href="/favicon.png" />
    <title>Grona gallery</title>
</svelte:head>
<header>
    <div class="top-0 z-40 hidden w-full bg-gray-200 px-surround py-8 shadow-xl lg:sticky">
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
    </div>
    <div class="fixed z-40 lg:hidden">
        <button
            onclick={() => {
                open = !open
            }}
            class="p-4 lg:hidden"
            aria-label="burger menu"
            disabled={open}
        >
            <svg
                class="h-12 w-12 text-gray-800 backdrop-blur-2xl"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                viewBox="0 0 24 24"
            >
                <path stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h16M4 18h16" />
            </svg>
        </button>
    </div>

    <button
        class="fixed inset-0 z-40 bg-black/40 backdrop-blur-sm transition-opacity duration-300"
        class:opacity-100={open}
        class:opacity-0={!open}
        class:pointer-events-none={!open}
        onclick={() => (open = false)}
        aria-label="backdrop for burger menu"
    ></button>

    <div
        class="fixed top-0 right-0 z-50 h-screen w-64 transform bg-white
         shadow-xl transition-transform duration-300 ease-out will-change-transform"
        class:translate-x-0={open}
        class:translate-x-full={!open}
    >
        <div class="p-6">
            <button onclick={() => (open = false)}>Close</button>

            <nav class="mt-8 flex flex-col gap-4">
                <a href="/">Home</a>
                <a href="/">Catalog</a>
                <a href="/">Contact</a>
            </nav>
        </div>
    </div>
</header>

{@render children()}
