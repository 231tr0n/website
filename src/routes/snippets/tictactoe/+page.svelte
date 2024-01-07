<script>
  import Page from '$lib/components/page/Page.svelte';
  import Tictactoe from './Tictactoe.svelte';
</script>

<Page scrollspy="true">
  <svelte:component this={Tictactoe} />
</Page>
