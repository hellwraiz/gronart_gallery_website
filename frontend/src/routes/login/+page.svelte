<script>
    import axios from 'axios'
    let email = "";
    let password = "";

    async function login() {
        const res = await axios.post('/api/login', { email: email, password: password })

        console.log(res)

        if (res.data.isLogged === true) {
            localStorage.setItem('pass', password)
            window.location.href = '/admin'
        } else {
            alert('Login failed')
        }
    }
</script>

<input bind:value={email} />
<input type="password" bind:value={password} />
<button on:click={login}>Login</button>
