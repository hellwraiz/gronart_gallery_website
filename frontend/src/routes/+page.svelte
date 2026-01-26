<script>
    export let data
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
    <div class="mx-auto grid max-w-content grid-cols-4 gap-2.5 pt-6">
        {#each paintings as painting}
            <div class="flex flex-col justify-between">
                <img src={"/images/" + painting.img_url} alt={painting.name} />
                <div class="flex flex-col gap-1 p-3 pt-1">
                    <h1 class="pb-2">{painting.name}, {painting.size}, {painting.technique}</h1>
                    <h2>{painting.author}</h2>
                    <h2>Price: {painting.price} €</h2>
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
</style>
