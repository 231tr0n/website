<script>
  import Utilitypost from './Utilitypost.svelte';

  let utilityPosts = [];

  let utilityPostCreator = (name, badges, description, url) => {
    return {
      name,
      badges,
      description,
      url
    };
  };
  utilityPosts.push(
    utilityPostCreator(
      'HTML Playground',
      ['html', 'css', 'javascript', 'browser'],
      'A simple playground to play with html, css and javascript.',
      '/utilities/htmlplayground'
    )
  );
</script>

<h1>Utilities</h1>

{#each utilityPosts as utilityPost}
  <Utilitypost {...utilityPost} />
{/each}
