<script>
  import Page from "$lib/components/page/Page.svelte";
  import Playground from "./HTMLplayground.svelte";
</script>

<Page scrollspy="true">
  <svelte:component this={Playground} />
</Page>
