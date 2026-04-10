<script lang="ts">
    import Paintings from "$lib/components/admin/Paintings.svelte"
    import Projects from "$lib/components/admin/Projects.svelte"
    import Users from "$lib/components/admin/Users.svelte"
    let { data } = $props()
    let activeTab = $state("paintings")
</script>

<div class="mx-auto flex w-full max-w-4xl flex-col px-4">
    <h1>Admin Panel</h1>
    <div class="flex gap-4">
        <button onclick={() => (activeTab = "paintings")}>Paintings</button>
        <button onclick={() => (activeTab = "users")}>Users</button>
        <button onclick={() => (activeTab = "projects")}>Projects</button>
    </div>
    {#if activeTab === "paintings"}
        <Paintings {data} />
    {:else if activeTab === "users"}
        <Users />
    {:else if activeTab === "projects"}
        <Projects />
    {/if}
</div>

<style>
    h1 {
        font-size: 4rem;
        padding-top: 20px;
        padding-bottom: 40px;
    }
</style>
