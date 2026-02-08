import { error } from "@sveltejs/kit"
import type { PageLoad } from "./$types"
import type { Painting } from "$lib/types"

export const load: PageLoad = async ({ params, parent }) => {
    const { paintings } = await parent()

    const painting = paintings.find((p: Painting) => p.uuid === params.uuid)

    if (!painting) {
        error(404, "Painting not found")
    }

    return { painting }
}
