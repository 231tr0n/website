<script>
  import Theme from "$lib/components/navbar/Theme.svelte";
  import Logo from "$lib/components/navbar/Logo.svelte";
  import Navlinks from "$lib/components/navbar/Navlinks.svelte";
</script>

<header>
  <Logo />
  <div>
    <Navlinks />
    <Theme />
  </div>
</header>

<style>
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
    padding-left: 3px;
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
