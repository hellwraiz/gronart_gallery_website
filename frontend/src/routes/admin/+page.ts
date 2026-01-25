// +page.ts
import { redirect } from "@sveltejs/kit"
import axios from "axios"
let email = localStorage.getItem("email")
let password = localStorage.getItem("pass")
console.log("test" + email + " " + password)
try {
    console.log("hello")
    await axios.post("/api/login", { email: email, password: password })
} catch {
    console.log("no")
    throw redirect(303, "/login")
}
