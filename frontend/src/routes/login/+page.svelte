<script>
    import axios from "axios"
    import { goto } from "$app/navigation"
    let email = ""
    let password = ""

    async function login() {
        try {
            await axios.post("/api/login", { email: email, password: password })
            localStorage.setItem("email", email)
            localStorage.setItem("pass", password)
            console.log("successfully logged in")
            await goto("/admin")
        } catch (error) {
            console.log("login failed. Retreating")
            alert("Login failed")
        }
    }
</script>

<div class="flex min-h-screen items-center justify-center">
    <div>
        <div class="mb-4 flex flex-col">
            <label for="email" class="mb-1">Email</label>
            <input
                id="email"
                name="email"
                type="email"
                bind:value={email}
                class="border px-2 py-1"
            />
        </div>
        <div class="mb-8 flex flex-col">
            <label for="pass" class="mb-1">Password</label>
            <input
                id="pass"
                name="password"
                type="password"
                bind:value={password}
                class="border px-2 py-1"
            />
        </div>
        <button on:click={login} class="w-full bg-blue-500 py-2 text-white">Login</button>
    </div>
</div>
