<script lang="ts">
    import axios from "axios"

    let form = {
        name: "",
        author: "",
        size: "",
        price: "",
        technique: ""
    }
    let image: FileList | undefined

    const fields: Array<{
        key: keyof typeof form // ← This tells TS the key must be a valid form property
        label: string
    }> = [
        { key: "name", label: "Painting's name" },
        { key: "author", label: "Your name" },
        { key: "size", label: "Painting's size" },
        { key: "price", label: "Price" },
        { key: "technique", label: "Techniques used" }
    ]

    async function submit() {
        console.log(form)
        console.log(image)

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

            console.log(res)
        } catch (error) {
            alert("Couldn't upload painting. Please try again later!")
            console.log(error)
        }
    }
</script>

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

<div class="flex h-screen flex-col items-center justify-center gap-2">
    {#each fields as field}
        <label for={field.key}>{field.label}</label>
        <input bind:value={form[field.key]} name={field.key} id={field.key} />
    {/each}

    <label for="photo">Photo used for the cover</label>
    <input bind:files={image} type="file" name="photo" id="photo" />

    <button on:click={submit}>Upload painting</button>
</div>

<style>
    input {
        margin-bottom: 4px;
        width: 20rem;
        border: 1px solid;
    }
</style>
