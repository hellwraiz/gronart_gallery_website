<script>
    import { goto } from '$app/navigation';

    let password = localStorage.getItem('pass')
    if (!password) {
        goto('/login')
    }

    console.log(password)

</script>
