<script>
    import { goto } from '$app/navigation';
    import axios from 'axios';
    import { onMount } from 'svelte';

    let email = localStorage.getItem('email')
    let password = localStorage.getItem('pass')

    onMount( async () => {
        const res = await axios.post('/api/login', { email: email, password: password })

        if (res.data.isLogged === "true") {
            console.log("successfully logged in")
        } else {
            console.log("login failed. Retreating")
            goto('/login')
        }
    })

</script>
