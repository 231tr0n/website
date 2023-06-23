<script>
  export let title = "";
  export let src = "";
</script>

{#if title && src}
  <div class="iframe body">
    <div class="iframe-context component">
      <span>{title}</span><span>sandbox</span>
    </div>
    <div class="iframe-block flex-center">
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
    padding: 5px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .iframe-block {
    padding: 10px;
  }

  .iframe {
    border-radius: 5px;
    border: 2px solid var(--color-dark-component-background);
  }

  :global(body.light-mode) .iframe {
    border: 2px solid var(--color-light-component-background);
  }

  iframe {
    border: 0px;
    overflow: auto;
  }
</style>
