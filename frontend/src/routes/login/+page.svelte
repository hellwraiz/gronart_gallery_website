<script>
    import axios from 'axios'
    import { goto } from '$app/navigation';
    let email = "";
    let password = "";

    async function login() {
        const res = await axios.post('/api/login', { email: email, password: password })

        if (res.data.isLogged === "true") {
            localStorage.setItem('email', email)
            localStorage.setItem('pass', password)
            console.log("successfully logged in")
            goto('/admin')
        } else {
            console.log("login failed. Retreating")
            alert('Login failed')
        }
    }
</script>

<input bind:value={email} />
<input type="password" bind:value={password} />
<button on:click={login}>Login</button>
