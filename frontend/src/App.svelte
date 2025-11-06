<script>
  async function fetchData() {
    const res = await fetch('/api/ping');
    const data = await res.json();
    console.log(data.message);
  }
</script>

<button on:click={fetchData}>Test API</button>
