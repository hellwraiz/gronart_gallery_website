<script lang="ts">
    import type { Painting } from "$lib/types"
    let { data } = $props()
    let currentPainting: Painting | null = $state(null)
    const perPage = 20
    let currentPage = $state(1)

    // Calculate start and end indices
    let startIndex = $derived((currentPage - 1) * perPage)
    let endIndex = $derived(startIndex + perPage)

    // Get only the paintings for current page
    let paintings = $derived(data.paintings.slice(startIndex, endIndex))

    // Optional: total pages for pagination controls
    let totalPages = $derived(Math.ceil(data.paintings.length / perPage))
</script>

{#if currentPainting != null}
    <div class="fixed inset-0 z-50 bg-white px-surround-phone pt-6 lg:px-surround">
        <div class="mx-auto flex max-w-content flex-col gap-8">
            <button
                class="inline-flex w-fit cursor-pointer gap-5 p-3"
                onclick={() => {
                    currentPainting = null
                }}
            >
                <img src="/assets/leftArrow.png" class="aspect-square h-10" alt="left arrow" />
                <p class="text-xl">Back to gallery</p>
            </button>
            <div class="flex items-center gap-5">
                <img
                    class="aspect-square w-0 flex-5 object-contain"
                    src={"/images/" + currentPainting.img_url}
                    alt=""
                />
                <div class="flex flex-4 flex-col">
                    <h1>{currentPainting.name}</h1>
                    <h2>{currentPainting.author}</h2>
                    <h3>{currentPainting.size}</h3>
                    <h3>{currentPainting.technique}</h3>
                    <h3>{currentPainting.description}</h3>
                    {#if currentPainting.sold}
                        <h2>Price: Sold!</h2>
                    {:else}
                        <h2>Price: {currentPainting.price} €</h2>
                    {/if}
                    {#if currentPainting.printable}
                        <h4>Prints available!</h4>
                    {/if}
                    {#if currentPainting.copiable}
                        <h4>Copies available!</h4>
                    {/if}
                </div>
            </div>
        </div>
    </div>
{/if}

<div class="relative flex h-screen w-full items-center justify-center">
    <img
        src="/images/cover.jpg"
        alt="cover"
        class="mg-auto absolute inset-0 h-screen w-full object-cover object-center"
    />
    <p class="mg-auto z-10 text-9xl font-bold text-amber-300">GRONA gallery</p>
</div>

<div class="w-full px-surround">
    <div class="mx-auto flex max-w-content justify-center">
        <p class="py-5 text-4xl font-bold">Here are our gallery's paintings!</p>
    </div>
    <div class="mx-auto grid max-w-content grid-cols-1 gap-2.5 pt-6 lg:grid-cols-4">
        {#each paintings as painting}
            <div class="flex flex-col justify-between">
                <button
                    onclick={() => {
                        currentPainting = painting
                    }}
                >
                    <img
                        src={"/images/" + painting.img_url}
                        class="aspect-square cursor-pointer object-contain transition-opacity hover:brightness-90"
                        alt={"Painting: " + painting.name}
                    />
                </button>
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
    <div class="mx-auto flex max-w-content justify-center py-5">
        <button
            disabled={currentPage === 1}
            aria-label="Go to prev page"
            class="pagination-button"
            onclick={() => {
                currentPage -= 1
            }}
        >
            <span class="left-arrow"></span>
        </button>
        {#each Array(totalPages) as _, i}
            <button
                class:active={i + 1 === currentPage}
                class="pagination-button"
                onclick={() => (currentPage = i + 1)}
            >
                {i + 1}
            </button>
        {/each}
        <button
            disabled={currentPage === totalPages}
            aria-label="Go to next page"
            class="pagination-button"
            onclick={() => {
                currentPage += 1
            }}
        >
            <span class="right-arrow"></span>
        </button>
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
    .left-arrow {
        border-bottom: 6px solid transparent;
        border-top: 6px solid transparent;
        border-right: 7px solid black;
    }
    .pagination-button {
        margin-left: -1px;
        height: 25px;
        width: 30px;
        border: 1px solid gray;
        display: flex; /* Change from inline-block */
        align-items: center; /* Center vertically */
        justify-content: center; /* Center horizontally */
    }
    .right-arrow {
        border-bottom: 6px solid transparent;
        border-top: 6px solid transparent;
        border-left: 7px solid black;
    }
    button.active {
        background: gray;
        color: white;
    }
</style>
