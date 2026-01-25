<script>
    export let data
    let paintings = data.paintings
</script>

<svelte:head>
    <title>Grona gallery</title>
</svelte:head>

<div class="grid grid-cols-4 gap-1.5">
    {#each paintings as painting}
        <div class="border p-1">
            <img src={"/images/" + painting.img_url} alt={painting.name} />
            <div class="flex items-center justify-center gap-1 pt-1">
                <h1>{painting.name}</h1>
                -
                <h2>{painting.author}</h2>
                -
                <h2>{painting.price}€</h2>
                -
                <h2>{painting.size}</h2>
                -
                <h2>{painting.technique}</h2>
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
