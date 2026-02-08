export type Painting = {
    uuid: string
    name: string
    author: string
    size: string
    price: number
    img_url: string
    technique: string
    description: string
    position: number
    sold: boolean
    printable: boolean
    copiable: boolean
    uploaded_at: string
    last_edited_at: string
}

export type FormPainting = Omit<
    Painting,
    "uuid" | "img_url" | "position" | "uploaded_at" | "last_edited_at"
> & {
    image: FileList | undefined
    img_url: string | null
}

export type FormPaintingg = {
    name: string
    author: string
    size: string
    price: number
    image: FileList | undefined
    img_url: string | null
    technique: string
    description: string
    sold: boolean
    printable: boolean
    copiable: boolean
}
