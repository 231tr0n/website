<script>
  export let post = "";
  export let badges = [];
  export let description = "";
  export let url = "";
  let visible = false;

  let toggle = () => {
    visible = !visible;
  };
</script>

{#if post && badges && description}
  <div class="underliner">
    <div class="flex-middle">
      <button on:click={toggle}
        ><span class="flex-middle"
          >{#if visible}▲{:else}▼{/if}</span
        ></button
      >
      <h2>
        {post}
      </h2>
      <a href={url}><button>Open</button></a>
    </div>
  </div>
  <div class="accordion-content flex-center">
    <div class="size-fixer">
      {#if visible}
        <div class="center">
          {#each badges as badge}
            <span class="badge">{badge}</span>
          {/each}
        </div>
        <div class="flex-center">
          <div class="justify">
            {description}
          </div>
        </div>
      {/if}
    </div>
  </div>
{/if}

<style>
  .underliner {
    border-bottom: 1px solid var(--color-dark-foreground);
  }

  :global(body.light-mode) .underliner {
    border-bottom: 1px solid var(--color-light-foreground);
  }

  .size-fixer {
    width: 800px;
    max-width: 90vw;
  }

  .accordion-content {
    margin-top: 1em;
    padding: 5px;
  }

  h2 {
    width: 100%;
    border: 0px;
    margin: 0px;
  }

  :global(body.light-mode) h2 {
    border: 0px;
  }

  div.accordion-content {
    margin-bottom: 1em;
  }

  div.justify {
    margin-top: 1em;
  }
</style>
