<script>
  import Theme from "./Theme.svelte";
  import Logo from "./Logo.svelte";
  import Navlinks from "./Navlinks.svelte";
</script>

<header>
  <div>
    <Logo />
    <div class="padder">
      <Navlinks />
    </div>
  </div>
  <div>
    <Theme />
  </div>
</header>

<style>
  .padder {
    padding: 3px;
  }

  div {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  header {
    position: fixed;
    top: 0px;
    left: 0px;
    right: 0px;
    height: 45px;
    display: flex;
    padding-left: 5px;
    padding-right: 5px;
    box-sizing: border-box;
    justify-content: space-between;
    z-index: 100;
    border-bottom: 1px solid var(--color-dark-foreground);
    background-color: var(--color-dark-background);
  }

  :global(body.light-mode) header {
    border-bottom: 1px solid var(--color-light-foreground);
    background-color: var(--color-light-background);
  }
</style>
