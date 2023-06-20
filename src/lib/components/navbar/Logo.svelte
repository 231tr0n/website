<script>
  import Heading from "$lib/components/navbar/Heading.svelte";
</script>

<div>
  <a href="https://231tr0n.github.io" target="_blank">
    <img alt="Zeltron" src="/images/zeltron.png" />
    <Heading />
  </a>
</div>

<style>
  div {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  div a {
    display: flex;
    background-color: var(--color-dark-component);
    border: 2px solid transparent;
  }

  :global(body.light-mode) div a {
    background-color: var(--color-light-component);
  }

  div a:hover {
    border: 2px solid var(--color-dark-anchor);
  }

  :global(body.light-mode) div a:hover {
    border: 2px solid var(--color-light-anchor);
  }

  img {
    vertical-align: middle;
    height: 30px;
    width: 30px;
    border-radius: 50%;
  }
</style>
