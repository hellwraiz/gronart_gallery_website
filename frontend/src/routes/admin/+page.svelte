<script lang="ts">
    /* ===================================
               Everything init
    ====================================*/
    import { invalidateAll } from "$app/navigation"
    import type { FormPainting, Painting } from "$lib/types"
    import axios from "axios"
    import type { PageData } from "./$types"
    export let data: PageData
    let showModal = false
    let editing = false
    let updated: Painting

    let form: FormPainting = {
        name: "",
        author: "",
        size: "",
        price: 0,
        image: undefined,
        img_url: "",
        technique: "",
        description: "",
        sold: false,
        printable: false,
        copiable: false
    }

    /* ===================================
               Helper functions
    ====================================*/

    /* function getChangedFields(original: Painting, modified: Painting) {
        const changes: Partial<Painting> = {};
        
        for (const key in modified) {
            if (modified[key] !== original[key]) {
                changes[key] = modified[key];
            }
        }
        
        return changes;
    } */
    function paintingToForm(painting: Painting): FormPainting {
        return {
            name: painting.name,
            author: painting.author,
            size: painting.size,
            price: painting.price,
            img_url: painting.img_url,
            technique: painting.technique,
            description: painting.description,
            sold: painting.sold,
            printable: painting.printable,
            copiable: painting.copiable,
            image: undefined
        }
    }

    function resetForm(): FormPainting {
        return {
            name: "",
            author: "",
            size: "",
            price: 0,
            image: undefined,
            img_url: "",
            technique: "",
            description: "",
            sold: false,
            printable: false,
            copiable: false
        }
    }

    /* ===================================
               UI logic
    ====================================*/

    async function showEditModal(painting: Painting) {
        form = paintingToForm(painting)
        editing = true
        updated = painting
        showModal = true
    }

    /* ===================================
               CRUD operations
    ====================================*/

    async function deletePainting(uuid: string) {
        try {
            await axios.delete(`api/paintings/${uuid}`, {
                headers: {
                    email: localStorage.getItem("email"),
                    pass: localStorage.getItem("pass")
                }
            })
            await invalidateAll()
        } catch (error) {
            alert("Couldn't delete painting, something went wrong.")
            console.log(error)
        }
    }

    async function modalSubmit() {
        // Uploading the image
        const photoData = new FormData()

        if (form.image && form.image.length > 0) {
            photoData.append("image", form.image[0])
            try {
                if (editing) {
                    let res = await axios.put(`api/upload/${updated.img_url}`, photoData, {
                        headers: {
                            "Content-Type": "multipart/form-data",
                            email: localStorage.getItem("email"),
                            pass: localStorage.getItem("pass")
                        }
                    })
                    form.img_url = res?.data?.img_url ?? null
                } else {
                    let res = await axios.post("api/upload/", photoData, {
                        headers: {
                            "Content-Type": "multipart/form-data",
                            email: localStorage.getItem("email"),
                            pass: localStorage.getItem("pass")
                        }
                    })
                    form.img_url = res?.data?.img_url ?? null
                }
                if (form.img_url === "cover.jpg") {
                    alert("successfully changed cover image")
                    return
                }
            } catch (error) {
                alert("Couldn't upload image for painting.")
                console.log(error)
                form.img_url = null
            }
        } else if (!editing) {
            alert("Please upload an image")
            return
        }
        form.image = undefined

        // uploading the painting
        try {
            if (editing) {
                await axios.patch(
                    `api/paintings/${updated.uuid}`,
                    { ...form },
                    {
                        headers: {
                            email: localStorage.getItem("email"),
                            pass: localStorage.getItem("pass")
                        }
                    }
                )
            } else {
                await axios.post(
                    "api/paintings",
                    { ...form },
                    {
                        headers: {
                            email: localStorage.getItem("email"),
                            pass: localStorage.getItem("pass")
                        }
                    }
                )
            }
        } catch (error) {
            console.log(error)
            alert("Failed to upload painting. Please try again.")
            await axios.delete(`api/upload/${form.img_url}`)
        }

        // cleanup
        // TODO: could make this more efficient
        await invalidateAll()
        alert("Uploaded the painting successfully!")
        form = resetForm()
        showModal = false
        editing = false
    }
</script>

<div class="flex min-h-screen flex-col items-center">
    <div class="flex w-225 flex-col">
        <h1>Admin Panel</h1>
        {#each data.paintings as painting}
            <div class="m-1.5 flex w-full gap-4 border p-0.5">
                <img
                    src={"/images/" + painting.img_url}
                    alt={painting.name}
                    class="h-72 w-72 object-cover object-center"
                />
                <div class="flex flex-col justify-center gap-1 pt-1">
                    <p>Name: {painting.name}</p>
                    <p>Author: {painting.author}</p>
                    <p>Size: {painting.size}</p>
                    <p>Price: {painting.sold ? "n/a" : painting.price}</p>
                    <p>Technique: {painting.technique}</p>
                    <div class="flex gap-2">
                        <p>Description:</p>
                        <p>{painting.description}</p>
                    </div>
                    <div class="flex flex-col">
                        {#if painting.sold}
                            <p style="color: green;">This painting was sold!</p>
                        {/if}
                        {#if painting.copiable}
                            <p style="color: green;">This painting can be copied!</p>
                        {/if}
                        {#if painting.printable}
                            <p style="color: green;">This painting can be printed!</p>
                        {/if}
                    </div>
                </div>
                <div class="grow"></div>
                <div class="mr-5 flex items-center gap-3">
                    <button
                        on:click={() => showEditModal(painting)}
                        class="min-w-20 cursor-pointer rounded-xl bg-blue-400 p-3">Edit</button
                    >
                    <button
                        on:click={() => deletePainting(painting.uuid)}
                        class="min-w-20 cursor-pointer rounded-xl bg-red-400 p-3">Delete</button
                    >
                </div>
            </div>
        {/each}
        <button
            on:click={() => (showModal = true)}
            class="mt-6 mb-10 w-40 cursor-pointer self-center rounded-xl bg-green-300 p-3"
            >Upload Painting</button
        >
    </div>
</div>

<!-- edit/create modal here -->

{#if showModal}
    <div
        class="fixed top-0 right-0 bottom-0 left-0 z-60 m-auto flex flex-col items-center justify-center gap-2 bg-white"
    >
        <button
            on:click={() => (showModal = false)}
            class="absolute top-8 right-24 h-8 w-8 cursor-pointer rounded-full bg-red-600 text-white"
            >X</button
        >
        <label for="name">Painting's name</label>
        <input bind:value={form["name"]} required class="text-input" name="name" id="name" />

        <label for="author">Painting's author</label>
        <input bind:value={form["author"]} required class="text-input" name="author" id="author" />

        <label for="size">Painting's size</label>
        <input bind:value={form["size"]} required class="text-input" name="size" id="size" />

        <label for="price">Painting's price</label>
        <input
            type="number"
            min="0"
            bind:value={form["price"]}
            name="price"
            id="price"
            class="no-spinner text-input"
            disabled={form.sold}
            required={!form.sold}
        />

        <label for="technique">Painting's technique</label>
        <input
            bind:value={form["technique"]}
            required
            class="text-input"
            name="technique"
            id="technique"
        />

        <label for="description">A description for the painting!</label>
        <textarea
            bind:value={form["description"]}
            required
            name="description"
            id="description"
            class="text-input"
        ></textarea>

        <div class="flex gap-5">
            <div class="flex">
                <label for="sold">Sold: </label>
                <input
                    bind:checked={form["sold"]}
                    class="m-1"
                    type="checkbox"
                    name="sold"
                    id="sold"
                />
            </div>
            <div class="flex">
                <label for="copiable">Can copy: </label>
                <input
                    bind:checked={form["copiable"]}
                    class="m-1"
                    type="checkbox"
                    name="copiable"
                    id="copiable"
                />
            </div>
            <div class="flex">
                <label for="printable">Can print: </label>
                <input
                    bind:checked={form["printable"]}
                    class="m-1"
                    type="checkbox"
                    name="printable"
                    id="printable"
                />
            </div>
        </div>

        <label for="photo">Photo used for the cover</label>
        <input
            bind:files={form.image}
            required
            class="text-input"
            type="file"
            name="photo"
            id="photo"
        />

        <button on:click={() => modalSubmit()} class="cursor-pointer rounded-xl bg-green-300 p-3"
            >Upload painting</button
        >
    </div>
{/if}

<style>
    .text-input {
        margin-bottom: 4px;
        width: 20rem;
        border: 1px solid;
    }

    h1 {
        font-size: 4rem;
        padding-top: 20px;
        padding-bottom: 40px;
    }

    .no-spinner::-webkit-outer-spin-button,
    .no-spinner::-webkit-inner-spin-button {
        -webkit-appearance: none;
        margin: 0;
    }
    .no-spinner {
        -moz-appearance: textfield;
        appearance: textfield;
    }

    input:disabled {
        opacity: 0.5;
        cursor: not-allowed;
    }
</style>
