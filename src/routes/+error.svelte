<script>
  import { page } from "$app/stores";
</script>

<div>{$page.status}: {$page.error.message}</div>

<style>
  div {
    font-size: 30px;
    padding: 1em;
    text-align: center;
    color: var(--color-dark-error);
  }

  :global(body.light-mode) div {
    color: var(--color-light-error);
  }
</style>
