<script>
  import Blogpost from "./Blogpost.svelte";

  let blogPosts = [];

  let blogPostCreator = (name, badges, description, url) => {
    return {
      name,
      badges,
      description,
      url,
    };
  };
  blogPosts.push(
    blogPostCreator(
      "TicTacToe",
      ["javascript", "html", "css", "browser"],
      "Creating TicTacToe game in browser with html, css and javascript.",
      "/blog/tictactoe"
    ),
    blogPostCreator(
      "HTML Playground",
      ["html", "css", "javascript", "browser"],
      "A simple playground to play with html, css and javascript.",
      "/blog/playground"
    )
  );
</script>

<h1>Blogs</h1>

{#each blogPosts as blogPost}
  <Blogpost {...blogPost} />
{/each}
