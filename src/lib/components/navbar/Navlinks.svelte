<span>
  <a href="/">Home</a><span /><a href="/blog">Blog</a>
</span>

<style>
  span {
    padding: 3px;
  }

  a {
    border-radius: 5px;
    padding: 2px;
    border: 2px solid transparent;
    background-color: var(--color-dark-component);
    color: var(--color-dark-heading);
  }

  :global(body.light-mode) a {
    background-color: var(--color-light-component);
    color: var(--color-light-heading);
  }

  a:hover {
    text-decoration: none;
    border: 2px solid var(--color-dark-anchor);
  }

  :global(body.light-mode) a:hover {
    border: 2px solid var(--color-light-anchor);
  }
</style>
