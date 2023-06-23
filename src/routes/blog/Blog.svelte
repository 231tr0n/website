<script>
  import Blogpost from "./Blogpost.svelte";

  let blogPosts = [];

  let blogPostCreator = (post, badges, description, url) => {
    return {
      post,
      badges,
      description,
      url,
    };
  };
  blogPosts.push(
    blogPostCreator(
      "TicTacToe",
      ["javscript", "html", "css", "browser"],
      "Creating TicTacToe game in browser with html, css and javascript.",
      "/blog/tictactoe"
    ),
    blogPostCreator(
      "TicTacToe",
      ["javscript", "html", "css", "browser"],
      "Creating TicTacToe game in browser with html, css and javascript.",
      "/blog/tictactoe"
    )
  );
</script>

<h1>Blogs</h1>

{#each blogPosts as blogPost}
  <Blogpost {...blogPost} />
{/each}
