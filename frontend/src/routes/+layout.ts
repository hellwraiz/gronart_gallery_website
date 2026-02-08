import axios from "axios"
import type { Painting } from "$lib/types"

export const ssr = false
export const load = async () => {
    const paintings = await axios.get<Painting[]>("/api/paintings")
    return {
        paintings: paintings.data
    }
}
