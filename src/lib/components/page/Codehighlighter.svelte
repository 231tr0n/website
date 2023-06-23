<script>
  import highlightjs from "highlight.js/lib/core";
  import "highlight.js/styles/base16/solarized-dark.css";

  export let code = "";
  export let langDef = "";
  export let langName = "";
  export let fileName = "";
  export let output = "";
  let codeHTML = "";

  if (code && langName && langDef) {
    highlightjs.registerLanguage(langName, langDef);
    codeHTML = highlightjs.highlight(code, {
      language: langName,
    }).value;
  }
</script>

{#if code && langDef && langName}
  <div class="code-block body">
    <div class="code-context component">
      <span>{fileName}</span><span>{langName}</span>
    </div>
    <pre><code>{@html codeHTML}</code></pre>
    {#if output !== ""}
      <div class="code-context component"><span>Output</span></div>
      <pre>{output}</pre>
    {/if}
  </div>
{/if}

<style>
  div.body {
    border: 2px solid var(--color-dark-component-background);
  }

  :global(body.light-mode) div.body {
    border: 2px solid var(--color-light-component-background);
  }

  pre {
    margin: 0px;
    font-size: 15px;
    padding: 10px;
    overflow: auto;
  }

  .code-block {
    margin-top: 1em;
    margin-bottom: 1em;
    border-radius: 5px;
  }

  .code-context {
    padding: 5px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
</style>
