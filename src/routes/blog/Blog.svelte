<script>
	import Blogpost from './Blogpost.svelte';

	let blogPosts = [];

	let blogPostCreator = (name, badges, description, url) => {
	  return {
	    name,
	    badges,
	    description,
	    url
	  };
	};
	blogPosts.push(
	  blogPostCreator('Me', ['introduction', 'about', 'me'], 'Introduction of who I am.', '/')
	);
</script>

<h1>Blogs</h1>

{#each blogPosts as blogPost}
	<Blogpost {...blogPost} />
{/each}
