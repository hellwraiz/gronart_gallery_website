<script lang="ts">
    /* ===================================
               Everything init
    ====================================*/
    import { invalidateAll } from "$app/navigation"
    import type { FormPainting, Painting } from "$lib/types"
    import axios from "axios"
    let { data } = $props()
    let showModal = $state(false)
    let editing = $state(false)
    let updated: Painting
    let paintingElements: HTMLDivElement[] = []
    let cursorY = $state(0)

    let form: FormPainting = $state({
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
    })

    // All the dragging stuff
    let sourcePosition = $state(0)
    let targetPosition = $state(0)
    let isDragging = $state(false)
    let shadowHeight = $state(0)

    /* ===================================
               Helper functions
    ====================================*/

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

    async function closeEditModal() {
        form = resetForm()
        editing = false
        showModal = false
    }

    /* ===================================
               CRUD operations
    ====================================*/

    async function movePainting() {
        if (sourcePosition === targetPosition - 1) {
            console.log("painting not moved. Skipping")
            return
        }
        if (sourcePosition > targetPosition) {
            console.log("Moving painting " + sourcePosition + " to position " + targetPosition)
        } else {
            console.log(
                "Moving painting " + sourcePosition + " to position " + (targetPosition - 1)
            )
        }
    }

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

        console.log(photoData)
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
                    // TODO: remove this crap and put it in a proper place
                    if (form.img_url === "cover.jpg") {
                        alert("successfully changed cover image")
                        return
                    }
                    let res = await axios.post("api/upload/", photoData, {
                        headers: {
                            "Content-Type": "multipart/form-data",
                            email: localStorage.getItem("email"),
                            pass: localStorage.getItem("pass")
                        }
                    })
                    form.img_url = res?.data?.img_url ?? null
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

<div
    onpointermove={(e) => {
        if (isDragging) {
            cursorY = e.clientY
            e.preventDefault()
            let bottomestBottom = 0
            for (let i = 0; i < paintingElements.length; i++) {
                const element = paintingElements[i].getBoundingClientRect()
                if (i == 0 && element.top > cursorY) {
                    targetPosition = 1
                    return
                }
                if (element.top < 0) {
                    continue
                }
                bottomestBottom = element.bottom
                if (
                    i !== sourcePosition - 1 &&
                    element.top < cursorY &&
                    bottomestBottom > cursorY
                ) {
                    if (targetPosition == i + 1) {
                        targetPosition++
                        return
                    } else {
                        targetPosition = i + 1
                        return
                    }
                }
            }
            if (cursorY > bottomestBottom) {
                targetPosition = paintingElements.length + 1
            }
        }
    }}
    onpointerup={() => {
        if (isDragging) {
            movePainting()
            isDragging = false
            sourcePosition = 0
            targetPosition = 0
            shadowHeight = 0
        }
    }}
    class="flex min-h-screen flex-col items-center"
>
    <div class="flex w-225 flex-col">
        <h1>Admin Panel</h1>
        {#each data.paintings as painting, i}
            <div
                hidden={painting.position !== targetPosition}
                class="my-1.5 w-full items-center justify-center bg-amber-500 p-0.5"
                style="height: {shadowHeight}px;"
            ></div>
            <div
                bind:this={paintingElements[i]}
                class="my-1.5 flex w-full items-center gap-4 border bg-white p-0.5"
                style={sourcePosition == painting.position
                    ? "top: " + (cursorY - shadowHeight / 2) + "px; position: fixed; width: 900px;"
                    : ""}
            >
                <div
                    class="h-40 w-20 bg-gray-300"
                    onpointerdown={(e) => {
                        e.preventDefault()
                        isDragging = true
                        sourcePosition = painting.position
                        targetPosition = painting.position + 1
                        shadowHeight = paintingElements[i].offsetHeight
                        cursorY = e.clientY
                    }}
                ></div>
                <img
                    src={"/images/" + painting.img_url}
                    alt={painting.name}
                    class="h-48 w-48 object-cover object-center"
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
                        onclick={() => showEditModal(painting)}
                        class="min-w-20 cursor-pointer rounded-xl bg-blue-400 p-3">Edit</button
                    >
                    <button
                        onclick={() => deletePainting(painting.uuid)}
                        class="min-w-20 cursor-pointer rounded-xl bg-red-400 p-3">Delete</button
                    >
                </div>
            </div>
        {/each}
        <div
            hidden={paintingElements.length + 1 !== targetPosition}
            class="my-1.5 w-full items-center justify-center bg-amber-500 p-0.5"
            style="height: {shadowHeight}px;"
        ></div>
        <button
            onclick={() => (showModal = true)}
            class="mt-6 mb-10 w-40 cursor-pointer self-center rounded-xl bg-green-300 p-3"
            >Upload Painting</button
        >
    </div>
</div>

<!-- edit/create modal here -->

{#if showModal}
    <form
        class="fixed top-0 right-0 bottom-0 left-0 z-60 m-auto flex flex-col items-center justify-center gap-2 bg-white"
        onsubmit={(e) => {
            e.preventDefault()
            modalSubmit()
        }}
    >
        <button
            type="button"
            onclick={() => closeEditModal()}
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
            required={!editing}
            class="text-input"
            type="file"
            name="photo"
            id="photo"
        />

        <button type="submit" class="cursor-pointer rounded-xl bg-green-300 p-3">
            Upload painting
        </button>
    </form>
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
