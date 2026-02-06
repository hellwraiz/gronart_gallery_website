// +page.ts
import { redirect } from "@sveltejs/kit"
import axios from "axios"
let email = localStorage.getItem("email")
let password = localStorage.getItem("pass")
try {
    await axios.post("/api/login", { email: email, password: password })
} catch {
    throw redirect(303, "/login")
}
