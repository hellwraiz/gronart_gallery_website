import axios from "axios"

let paintings = await axios.get("/api/paintings")

export function load() {
    return {
        paintings: paintings.data
    }
}
