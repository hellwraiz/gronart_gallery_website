<script>
    export let data
    let paintings = data.paintings
</script>

<svelte:head>
    <title>Grona gallery</title>
</svelte:head>

<div class="relative">
    <p class="absolute top-1/4 left-1/4 text-9xl font-bold text-white">GRONA gallery</p>
    <img src="/images/cover.jpg" alt="cover" class="h-screen w-full object-cover object-center" />
</div>

<div class="mx-auto grid max-w-content grid-cols-4 gap-1.5 pt-6">
    {#each paintings as painting}
        <div class="flex flex-col justify-between border p-1">
            <img src={"/images/" + painting.img_url} alt={painting.name} />
            <div class="flex flex-col gap-1 p-3 pt-1">
                <h1 class="pb-2">{painting.name}, {painting.size}, {painting.technique}</h1>
                <h2>{painting.author}</h2>
                <h2>Price: {painting.price} €</h2>
            </div>
        </div>
    {/each}
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
