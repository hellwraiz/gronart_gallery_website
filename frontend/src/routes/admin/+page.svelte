<script>
    import { goto } from '$app/navigation';

    let password = localStorage.getItem('password')
    if (!password) {
        goto('/login')
    }

    console.log(password)

</script>
