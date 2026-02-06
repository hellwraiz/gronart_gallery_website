<script lang="ts">
    /* ===================================
               Everything init
    ====================================*/
    import { invalidateAll } from "$app/navigation"
    import type { FormPainting, Painting } from "$lib/types"
    import axios from "axios"
    export let data
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

    const fields: Array<{
        key: keyof typeof form // This tells TS the key must be a valid form property
        label: string
    }> = [
        { key: "name", label: "Painting's name" },
        { key: "author", label: "Your name" },
        { key: "size", label: "Painting's size" },
        { key: "price", label: "Price" },
        { key: "technique", label: "Techniques used" },
        { key: "description", label: "Description for the paniting" }
    ]
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
            sold: painting.sold === 1,
            printable: painting.printable === 1,
            copiable: painting.copiable === 1,
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
                let res = await axios.put(`api/upload/${updated.img_url}`, photoData, {
                    headers: {
                        "Content-Type": "multipart/form-data",
                        email: localStorage.getItem("email"),
                        pass: localStorage.getItem("pass")
                    }
                })
                form.img_url = res?.data?.img_url ?? null
                if (form.img_url === "cover.jpg") {
                    alert("successfully changed cover image")
                    return
                }
            } catch (error) {
                alert("Couldn't upload image for painting.")
                console.log(error)
                form.img_url = null
            }
        }
        form.image = undefined

        //// changes

        try {
            if (editing) {
                await axios.patch(
                    `api/paintings/${updated.uuid}`,
                    { form },
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
                    { form },
                    {
                        headers: {
                            email: localStorage.getItem("email"),
                            pass: localStorage.getItem("pass")
                        }
                    }
                )
            }
            alert("Uploaded the painting successfully!")
        } catch (error) {
            console.log(error)
            alert("Failed to upload painting. Please try again.")
            await axios.delete(`api/upload/${form.img_url}`)
        }

        //// cleanup
        // TODO: could make this more efficient
        await invalidateAll()
        alert("Uploaded the painting successfully!")
        form = resetForm()
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
                    <p>Price: {painting.price}</p>
                    <p>Technique: {painting.technique}</p>
                </div>
                <div class="grow"></div>
                <div class="mr-5 flex items-center gap-3">
                    <button
                        on:click={() => showEditModal(painting)}
                        class="min-w-20 rounded-xl bg-blue-400 p-3">Edit</button
                    >
                    <button
                        on:click={() => deletePainting(painting.uuid)}
                        class="min-w-20 rounded-xl bg-red-400 p-3">Delete</button
                    >
                </div>
            </div>
        {/each}
        <button
            on:click={() => (showModal = true)}
            class="mt-6 mb-10 w-40 self-center rounded-xl bg-green-300 p-3">Upload Painting</button
        >
    </div>
</div>
{#if showModal}
    <div
        class="fixed top-0 right-0 bottom-0 left-0 z-60 m-auto flex flex-col items-center justify-center gap-2 bg-white"
    >
        <button
            on:click={() => (showModal = false)}
            class="absolute top-8 right-24 h-8 w-8 rounded-full bg-red-600 text-white">X</button
        >
        {#each fields as field}
            <label for={field.key}>{field.label}</label>
            <input bind:value={form[field.key]} name={field.key} id={field.key} />
        {/each}

        <label for="photo">Photo used for the cover</label>
        <input bind:files={form.image} type="file" name="photo" id="photo" />

        <button on:click={() => modalSubmit()} class="rounded-xl bg-green-300 p-3"
            >Upload painting</button
        >
    </div>
{/if}

<!--
type Painting struct {
	ID				int			`db:"id" json:"-"`
    UUID        	string    	`db:"uuid" json:"uuid"`
    Name        	string    	`db:"name" json:"name"`
    Author			string		`db:"author" json:"author"`
    Size        	string    	`db:"size" json:"size"`
	Price			int			`db:"price" json:"price"`
    ImgURL      	string    	`db:"img_url" json:"img_url"`
    Technique   	string    	`db:"technique" json:"technique"`
    UploadedAt  	time.Time   `db:"uploaded_at" json:"uploaded_at"`
    LastEditedAt 	time.Time   `db:"last_edited_at" json:"last_edited_at"`
-->

<style>
    input {
        margin-bottom: 4px;
        width: 20rem;
        border: 1px solid;
    }

    h1 {
        font-size: 4rem;
        padding-top: 20px;
        padding-bottom: 40px;
    }
</style>
