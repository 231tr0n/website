<script>
  export let post = "";
  export let badges = [];
  export let description = "";
  export let url = "";
</script>

{#if post && badges && description}
  <h2>{post}</h2>
  <div class="center">
    {#each badges as badge}
      <span class="badge">{badge}</span>
    {/each}
  </div>
  <br />
  {description}
  <div class="center">
    <a href={url} class="button">Open</a>
  </div>
{/if}
