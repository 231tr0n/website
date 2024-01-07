<script>
  import Page from '$lib/components/page/Page.svelte';
  import Home from './Home.svelte';
</script>

<Page scrollspy="true">
  <svelte:component this={Home} />
</Page>
