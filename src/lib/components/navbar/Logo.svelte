<script>
  import Heading from "./Heading.svelte";
</script>

<svelte:head>
  <link
    rel="preload"
    as="image"
    type="image/png"
    href="/images/zeltron.png"
    crossorigin
  />
</svelte:head>

<div>
  <a href="https://231tr0n.github.io" target="_blank">
    <img alt="Zeltron" src="/images/zeltron.png" />
  </a>
  <div class="padder" />
  <a href="https://231tr0n.github.io" target="_blank">
    <div class="heading">
      <Heading />
    </div>
  </a>
</div>

<style>
  .padder {
    padding: 3px;
  }

  div {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .heading {
    padding: 2px;
    background-color: var(--color-dark-component);
    border: 2px solid transparent;
    border-radius: 5px;
  }

  :global(body.light-mode) .heading {
    background-color: var(--color-light-component);
  }

  .heading:hover {
    border: 2px solid var(--color-dark-anchor);
  }

  :global(body.light-mode) .heading:hover {
    border: 2px solid var(--color-light-anchor);
  }

  a {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  a:hover {
    text-decoration: none;
  }

  img {
    vertical-align: middle;
    height: 30px;
    width: 30px;
    border-radius: 50%;
  }

  a img {
    border: 2px solid var(--color-dark-keyword);
  }

  :global(body.light-mode) a img {
    border: 2px solid var(--color-black);
  }

  a img:hover {
    border: 2px solid var(--color-dark-anchor);
  }

  :global(body.light-mode) a img:hover {
    border: 2px solid var(--color-light-anchor);
  }
</style>
