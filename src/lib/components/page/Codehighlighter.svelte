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
  <div class="code-block">
    <div class="code-context">
      <span>{fileName}</span><span>{langName}</span>
    </div>
    <pre><code>{@html codeHTML}</code></pre>
    {#if output !== ""}
      <div class="code-context"><span>Output</span></div>
      <pre>{output}</pre>
    {/if}
  </div>
{/if}

<style>
  pre {
    margin: 0px;
    font-size: 15px;
    padding: 5px;
    overflow-x: auto;
  }

  .code-block {
    margin-top: 1em;
    margin-bottom: 1em;
    padding: 5px;
    border-radius: 5px;
    background-color: var(--color-solarized-base02);
  }

  :global(body.light-mode) .code-block {
    background-color: var(--color-solarized-base2);
  }

  .code-context {
    background-color: var(--color-dark-component);
    color: var(--color-dark-keyword);
    border-radius: 5px;
    padding: 5px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  :global(body.light-mode) .code-context {
    background-color: var(--color-light-component);
    color: var(--color-light-keyword);
  }
</style>
