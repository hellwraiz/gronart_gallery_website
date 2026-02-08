<script lang="ts">
    import type { PageData } from "./$types"
    export let data: PageData
    let paintings = data.paintings
</script>

<svelte:head>
    <title>Grona gallery</title>
</svelte:head>

<div class="relative flex h-screen w-full items-center justify-center">
    <img
        src="/images/cover.jpg"
        alt="cover"
        class="mg-auto absolute inset-0 h-screen w-full object-cover object-center"
    />
    <p class="mg-auto z-10 text-9xl font-bold text-amber-300">GRONA gallery</p>
</div>

<div class="w-full px-surround">
    <div class="mx-auto grid max-w-content grid-cols-1 gap-2.5 pt-6 lg:grid-cols-4">
        {#each paintings as painting}
            <div class="flex flex-col justify-between">
                <a href="/{painting.uuid}">
                    <img
                        src={"/images/" + painting.img_url}
                        class="aspect-square cursor-pointer object-contain transition-opacity hover:brightness-90"
                        alt={"Painting: " + painting.name}
                    />
                </a>
                <div class="flex flex-col p-3 pt-1">
                    <h1 class="pb-2">{painting.name}, {painting.size}</h1>
                    <h2>{painting.technique}</h2>
                    <h2>{painting.author}</h2>
                    {#if painting.sold}
                        <h3>Price: Sold!</h3>
                    {:else}
                        <h3>Price: {painting.price} €</h3>
                    {/if}
                    {#if painting.printable}
                        <h3>Prints available!</h3>
                    {/if}
                    {#if painting.copiable}
                        <h3>Copies available!</h3>
                    {/if}
                </div>
            </div>
        {/each}
    </div>
</div>

<style>
    h1 {
        font-size: 1.2rem;
        font-weight: 700;
    }
    h2 {
        font-size: 1.1rem;
    }
    h3 {
        font-size: 0.9rem;
    }
</style>
