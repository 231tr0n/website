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
  <div class="wrapper">
    <div class="post">
      <button on:click={toggle}
        ><span class="accordion-button"
          >{#if visible}▲{:else}▼{/if}</span
        ></button
      >
      <h2>
        {post}
      </h2>
      <a href={url}><button>Open</button></a>
    </div>
  </div>
  <div class="center-content-position">
    <div class="size-fixer">
      {#if visible}
        <div class="center">
          {#each badges as badge}
            <span class="badge">{badge}</span>
          {/each}
        </div>
        <div class="center-content-position">
          <div class="justify">
            {description}
          </div>
        </div>
      {/if}
    </div>
  </div>
{/if}

<style>
  .wrapper {
    border-bottom: 1px solid var(--color-dark-foreground);
  }

  :global(body.light-mode) .wrapper {
    border-bottom: 1px solid var(--color-light-foreground);
  }

  .size-fixer {
    width: 800px;
    max-width: 90vh;
  }

  .center-content-position {
    margin-top: 1em;
    display: flex;
    justify-content: center;
    padding: 5px;
  }

  div.post {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  h2 {
    display: inline-block;
    width: 100%;
    border: 0px;
    margin: 0px;
  }

  :global(body.light-mode) h2 {
    border: 0px;
  }

  a {
    font-size: 15px;
  }

  button {
    display: inline-block;
    margin-left: 0px;
    margin-right: 5px;
  }

  div.center {
    margin-bottom: 1em;
  }

  span.accordion-button {
    display: flex;
    justify-content: center;
    align-items: center;
    padding-left: 5px;
    padding-right: 5px;
  }
</style>
