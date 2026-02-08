<script lang="ts">
    import type { PageData } from "./$types"
    export let data: PageData
</script>

<div class="w-full px-surround pt-6">
    <div class="mx-auto flex max-w-content flex-col gap-8">
        <a href="/" class="inline-flex w-fit gap-5 p-3">
            <img src="/assets/leftArrow.png" class="aspect-square h-10" alt="left arrow" />
            <p class="text-xl">Back to gallery</p>
        </a>
        <div class="flex items-center gap-5">
            <img
                class="aspect-square w-0 flex-5 object-contain"
                src={"/images/" + data.painting.img_url}
                alt=""
            />
            <div class="flex flex-7 flex-col">
                <h1>{data.painting.name}</h1>
                <h2>{data.painting.author}</h2>
                <h3>{data.painting.size}</h3>
                <h3>{data.painting.technique}</h3>
                <h3>{data.painting.description}</h3>
                {#if data.painting.sold}
                    <h2>Price: Sold!</h2>
                {:else}
                    <h2>Price: {data.painting.price} €</h2>
                {/if}
                {#if data.painting.printable}
                    <h4>Prints available!</h4>
                {/if}
                {#if data.painting.copiable}
                    <h4>Copies available!</h4>
                {/if}
            </div>
        </div>
    </div>
</div>

<style>
    h1 {
        font-size: 2.4rem;
        font-weight: 700;
    }
    h2 {
        font-size: 1.8rem;
    }
    h3 {
        font-size: 1.2rem;
    }
    h4 {
        font-size: 1rem;
        font-weight: 300;
    }
</style>
