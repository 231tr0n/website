<script>
  import highlightjs from "highlight.js/lib/core";
  import "highlight.js/styles/base16/solarized-dark.css";

  export let code = null;
  export let langDef = null;
  export let langName = null;
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

<div class="code-block">
  {#if code && langDef && langName}
    <div class="code-context">
      <span>{fileName}</span><span>{langName}</span>
    </div>
    <hr />
    <pre><code>{@html codeHTML}</code></pre>
    {#if output !== ""}
      <hr />
      <div class="code-context"><span>Output</span></div>
      <hr />
      <pre>{output}</pre>
    {/if}
  {:else}
    This component requires attributes "code", "langDef", "langName".
  {/if}
</div>

<style>
  pre {
    margin: 0px;
    font-size: 13px;
    padding: 10px;
    overflow-x: auto;
  }

  .code-block {
    margin-top: 1em;
    margin-bottom: 1em;
    padding: 5px;
    border-radius: 5px;
    background-color: var(--color-dark-component);
  }

  :global(body.light-mode) .code-block {
    background-color: var(--color-light-component);
  }

  .code-context {
    background-color: var(--color-dark-background);
    color: var(--color-dark-keyword);
    border-radius: 5px;
    padding: 5px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  :global(body.light-mode) .code-context {
    background-color: var(--color-light-background);
    color: var(--color-light-keyword);
  }
</style>
