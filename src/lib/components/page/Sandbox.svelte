<script>
  export let title = "";
  export let src = "";
</script>

{#if title && src}
  <div class="flex-middle">
    <div class="iframe body">
      <div class="iframe-context component">
        <span>{title}</span><span>sandbox</span>
      </div>
      <iframe
        {title}
        allowfullscreen
        sandbox="allow-forms allow-modals allow-popups allow-scripts allow-downloads"
        {src}
      />
    </div>
  </div>
{/if}

<style>
  .iframe-context {
    padding: 5px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .iframe {
    width: 800px;
    max-width: 85vw;
    margin-top: 1em;
    margin-bottom: 1em;
    border-radius: 5px;
    border: 2px solid var(--color-dark-component-background);
    height: calc(100vh - 10vh - 45px - 45px - 35px);
    box-sizing: border-box;
  }

  :global(body.light-mode) .iframe {
    border: 2px solid var(--color-light-component-background);
  }

  iframe {
    border: 0px;
    border-radius: 4px;
    width: 100%;
    height: calc(100% - 33px);
    overflow: auto;
  }
</style>
