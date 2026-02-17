<script lang="ts">
    import type { Painting } from "$lib/types"
    import { onMount, tick } from "svelte"
    let { data } = $props()
    let currentPainting: Painting | null = $state(null)
    // HACK: Please return to normal value before letting it go to prod
    let currentPage = $state(1)
    let perPage = $state(0)
    updateVisibleCount()

    // Calculate start and end indices
    let startIndex = $derived((currentPage - 1) * perPage)
    let endIndex = $derived(startIndex + perPage)
    // Get only the paintings for current page
    let paintings = $derived(data.paintings.slice(startIndex, endIndex))
    let totalPages = $derived(Math.ceil(data.paintings.length / perPage))

    let paintingSection: HTMLElement

    async function changePage(newPage: number) {
        currentPage = newPage
        await tick()

        window.scrollTo({
            top: paintingSection.offsetTop - 80,
            behavior: "smooth"
        })
    }

    function getThumbnailUrl(imgUrl: string) {
        const ext = imgUrl.substring(imgUrl.lastIndexOf("."))
        const name = imgUrl.substring(0, imgUrl.lastIndexOf("."))
        return `/images/${name}_thumb${ext}`
    }

    function updateVisibleCount() {
        if (window.innerWidth >= 1024) {
            perPage = 20
        } else if (window.innerWidth >= 768) {
            perPage = 14
        } else {
            perPage = 8
        }
    }

    onMount(() => {
        updateVisibleCount()
        window.addEventListener("resize", updateVisibleCount)
        return () => window.removeEventListener("resize", updateVisibleCount)
    })
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
    <p
        class="mg-auto z-10 text-5xl font-bold text-amber-300 drop-shadow-[0px_0px_2px_rgba(0,0,0,0.45)] [-webkit-text-stroke:0.3px_gray] lg:text-9xl lg:[-webkit-text-stroke:0px]"
    >
        GRONA gallery
    </p>
</div>

<div bind:this={paintingSection} class="w-full px-surround-phone lg:px-surround">
    <div class="mx-auto flex max-w-content-phone justify-center lg:max-w-content">
        <p class="py-5 text-xl font-bold lg:text-4xl">Here are our gallery's paintings!</p>
    </div>
    <div
        class="mx-auto grid max-w-content-phone grid-cols-1 gap-2.5 pt-6 lg:max-w-content lg:grid-cols-4"
    >
        {#each paintings as painting}
            <div class="flex flex-col justify-between">
                <button
                    onclick={() => {
                        currentPainting = painting
                    }}
                    class="flex justify-center"
                >
                    <img
                        src={getThumbnailUrl(painting.img_url)}
                        class="aspect-square max-h-120 max-w-80 cursor-pointer self-center object-contain object-center transition-opacity hover:brightness-90 lg:max-h-full lg:max-w-full"
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
    {#if totalPages > 1}
        <div class="mx-auto flex max-w-content justify-center py-5">
            <button
                disabled={currentPage === 1}
                aria-label="Go to prev page"
                class="pagination-button"
                onclick={() => {
                    changePage(currentPage - 1)
                }}
            >
                <span class="left-arrow"></span>
            </button>
            {#each Array(totalPages) as _, i}
                <button
                    class:active={i + 1 === currentPage}
                    class="pagination-button"
                    onclick={() => changePage(i + 1)}
                >
                    {i + 1}
                </button>
            {/each}
            <button
                disabled={currentPage === totalPages}
                aria-label="Go to next page"
                class="pagination-button"
                onclick={() => {
                    changePage(currentPage + 1)
                }}
            >
                <span class="right-arrow"></span>
            </button>
        </div>
    {/if}
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
