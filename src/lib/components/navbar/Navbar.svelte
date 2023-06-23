<script>
  import Theme from "./Theme.svelte";
</script>

<header>
  <span>
    <a href="https://231tr0n.github.io" target="_blank">
      <button><strong>231tr0n</strong></button>
    </a>
    <a href="/"><button>Home</button></a>
    <a href="/blog"><button>Blog</button></a>
  </span>
  <Theme />
</header>

<style>
  strong {
    font-family: HurmitNerdFontMono-Bold;
    font-weight: bolder;
  }

  header {
    position: fixed;
    top: 0px;
    left: 0px;
    right: 0px;
    height: 45px;
    padding-top: 0px;
    padding-bottom: 0px;
    padding-left: 7px;
    padding-right: 7px;
    box-sizing: border-box;
    display: flex;
    justify-content: space-between;
    align-items: center;
    z-index: 100;
    border-bottom: 1px solid var(--color-dark-foreground);
    background-color: var(--color-dark-strong-background);
  }

  :global(body.light-mode) header {
    border-bottom: 1px solid var(--color-light-foreground);
    background-color: var(--color-light-strong-background);
  }
</style>
