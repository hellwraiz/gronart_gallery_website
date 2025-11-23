<script>
  async function fetchData() {
    const res = await fetch('/api/paintings');
    const data = await res.json();
    console.log(data);
  }
</script>

<button on:click={fetchData}>Test API</button>
