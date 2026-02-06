import axios from "axios"
import type { Painting } from "$lib/types"

export const ssr = false
export async function load() {
    const paintings = await axios.get<Painting[]>("/api/paintings")
    return {
        paintings: paintings.data
    }
}
