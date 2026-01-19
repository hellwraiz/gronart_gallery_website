import axios from "axios"

export const ssr = false
export async function load() {
    const paintings = await axios.get("/api/paintings")
    return {
        paintings: paintings.data
    }
}
