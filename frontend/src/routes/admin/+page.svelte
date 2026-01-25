<script lang="ts">
    import { invalidateAll } from "$app/navigation"
    import axios from "axios"
    export let data
    console.log(data)
    let showCreate = false

    let form = {
        name: "",
        author: "",
        size: "",
        price: "",
        technique: ""
    }
    let image: FileList | undefined

    const fields: Array<{
        key: keyof typeof form // This tells TS the key must be a valid form property
        label: string
    }> = [
        { key: "name", label: "Painting's name" },
        { key: "author", label: "Your name" },
        { key: "size", label: "Painting's size" },
        { key: "price", label: "Price" },
        { key: "technique", label: "Techniques used" }
    ]

    async function editPainting(uuid: string) {
        alert("Not working at the moment." + uuid)
    }

    async function deletePainting(uuid: string) {
        try {
            let res = await axios.delete(`api/paintings/${uuid}`, {
                headers: {
                    email: localStorage.getItem("email"),
                    pass: localStorage.getItem("pass")
                }
            })
            console.log(res)
            await invalidateAll()
            alert("Successfully deleted the painting!")
        } catch (error) {
            alert("Couldn't delete painting, something went wrong.")
            console.log(error)
        }
    }

    async function createPainting() {
        const photoData = new FormData()

        if (image && image.length > 0) {
            photoData.append("image", image[0])
        }

        try {
            let res = await axios.post("api/upload", photoData, {
                headers: {
                    "Content-Type": "multipart/form-data",
                    email: localStorage.getItem("email"),
                    pass: localStorage.getItem("pass")
                }
            })
            let img_url = res.data.img_url
            if (img_url == "cover.jpg") {
                alert("successfully changed cover image")
            } else {
                let intPrice = parseInt(form.price, 10)

                res = await axios.post(
                    "api/paintings",
                    {
                        name: form.name,
                        author: form.author,
                        size: form.size,
                        price: intPrice,
                        img_url: img_url,
                        technique: form.technique
                    },
                    {
                        headers: {
                            email: localStorage.getItem("email"),
                            pass: localStorage.getItem("pass")
                        }
                    }
                )

                await invalidateAll()
                alert("Uploaded the painting successfully!")
            }
        } catch (error) {
            alert("Couldn't upload painting. Please try again later!")
        }
    }
</script>

<!-- TODO: remove the padding from top once header is there -->
<div class="flex min-h-screen flex-col items-center p-8">
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
                </div>
                <div class="grow"></div>
                <div class="mr-5 flex items-center gap-3">
                    <button
                        on:click={() => editPainting(painting.uuid)}
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
            on:click={() => (showCreate = true)}
            class="mt-3 w-40 self-center rounded-xl bg-green-300 p-3">Upload Painting</button
        >
    </div>
    {#if showCreate}
        <div
            class="absolute inset-0 m-auto flex flex-col items-center justify-center gap-2 bg-white"
        >
            {#each fields as field}
                <label for={field.key}>{field.label}</label>
                <input bind:value={form[field.key]} name={field.key} id={field.key} />
            {/each}

            <label for="photo">Photo used for the cover</label>
            <input bind:files={image} type="file" name="photo" id="photo" />

            <button on:click={() => createPainting()} class="rounded-xl bg-green-300 p-3"
                >Upload painting</button
            >
        </div>
        <button
            on:click={() => (showCreate = false)}
            class="absolute top-8 right-24 h-8 w-8 rounded-full bg-red-600 text-white">X</button
        >
    {/if}
</div>

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
