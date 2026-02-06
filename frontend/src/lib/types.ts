export type Painting = {
    uuid: string
    name: string
    author: string
    size: string
    price: number
    img_url: string
    technique: string
    description: string
    position: string
    sold: number
    printable: number
    copiable: number
    uploaded_at: string
    last_edited_at: string
}

export type FormPainting = {
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
