<script>
  export let title = "";
  export let src = "";
</script>

{#if title && src}
  <div class="iframe">
    <div class="iframe-context">
      <span>{title}</span><span>sandbox</span>
    </div>
    <div class="iframe-block">
      <iframe
        {title}
        allowfullscreen
        sandbox="allow-forms allow-modals allow-popups allow-scripts allow-same-origin"
        {src}
        onload="this.style.height=(this.contentWindow.document.body.scrollHeight+20)+'px';"
      />
    </div>
  </div>
{/if}

<style>
  .iframe-context {
    background-color: var(--color-dark-component);
    color: var(--color-dark-keyword);
    border-radius: 5px;
    padding: 5px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  :global(body.light-mode) .iframe-context {
    background-color: var(--color-light-component);
    color: var(--color-light-keyword);
  }

  .iframe-block {
    display: flex;
    justify-content: center;
  }

  .iframe {
    padding: 5px;
    border-radius: 5px;
    background-color: var(--color-solarized-base02);
  }

  :global(body.light-mode) div {
    background-color: var(--color-solarized-base2);
  }

  iframe {
    display: block;
    border: 0px;
    display: flex;
    overflow: scroll;
    width: auto;
    height: auto;
  }
</style>
