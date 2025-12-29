// +page.ts
import { redirect } from "@sveltejs/kit"
import axios from "axios"

export async function load() {
    let email = localStorage.getItem("email")
    let password = localStorage.getItem("pass")
    try {
        console.log("successfully logged in")
        await axios.post("/api/login", { email: email, password: password })
    } catch {
        console.log("login failed. Retreating")
        throw redirect(303, "/login")
    }
}
