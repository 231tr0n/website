<script>
  import { onMount } from "svelte";
  import Codehighlighter from "$lib/components/page/Codehighlighter.svelte";
  import Sandbox from "$lib/components/page/Sandbox.svelte";
  import highlightJs from "highlight.js/lib/languages/javascript";
  import highlightCss from "highlight.js/lib/languages/css";
  import highlightXml from "highlight.js/lib/languages/xml";

  let script = null;
  let css = null;
  let html = null;
  let render = false;

  onMount(async () => {
    try {
      script = await fetch("/resources/blog/tictactoe/index.js");
      css = await fetch("/resources/blog/tictactoe/style.css");
      html = await fetch("/resources/blog/tictactoe/index.html");
      script = await script.text();
      css = await css.text();
      html = await html.text();
      render = true;
    } catch (error) {
      console.log(error);
    }
  });
</script>

<svelte:head>
  <!-- <link -->
  <!--   rel="preload" -->
  <!--   as="document" -->
  <!--   type="text/html" -->
  <!--   src="/resources/blog/tictactoe/index.html" -->
  <!--   crossorigin -->
  <!-- /> -->
  <!-- <link -->
  <!--   rel="preload" -->
  <!--   as="script" -->
  <!--   type="text/javascript" -->
  <!--   src="/resources/blog/tictactoe/index.js" -->
  <!--   crossorigin -->
  <!-- /> -->
  <!-- <link -->
  <!--   rel="preload" -->
  <!--   as="style" -->
  <!--   type="text/css" -->
  <!--   src="/resources/blog/tictactoe/style.css" -->
  <!--   crossorigin -->
  <!-- /> -->
</svelte:head>

<h1>TicTacToe</h1>

<h2>Source Code</h2>
{#if render}
  <Codehighlighter
    code={script}
    langDef={highlightJs}
    langName={"javascript"}
    filePath={"./"}
    fileName={"index.js"}
  />
  <Codehighlighter
    code={css}
    langDef={highlightCss}
    langName={"css"}
    filePath={"./"}
    fileName={"style.css"}
  />
  <Codehighlighter
    code={html}
    langDef={highlightXml}
    langName={"html"}
    filePath={"./"}
    fileName={"index.html"}
  />
{/if}

<h2>Output</h2>
{#if render}
  <Sandbox title="tic-tac-toe" src="/resources/blog/tictactoe/index.html" />
{/if}
