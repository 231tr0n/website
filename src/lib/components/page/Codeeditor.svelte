<script>
  import { onMount } from "svelte";
  import ace from "ace-code";
  import solarized_light from "ace-code/src/theme/solarized_light";
  import solarized_dark from "ace-code/src/theme/solarized_dark";
  import vim from "ace-code/src/keyboard/vim";
  import { lightMode } from "$lib/stores/theme";
  import "$lib/css/editor.css";

  export let langName = "";
  export let vimMode = false;
  export let mode = "";
  export let readOnly = false;

  let editorDiv = "";
  let editor = "";
  let editorElement = "";
  let copied = false;
  let fullscreen = false;

  let copy = () => {
    navigator.clipboard.writeText(editor.session.getValue());
    copied = true;
    setTimeout(() => {
      copied = false;
    }, 2000);
  };

  let toggleFullscreen = () => {
    if (document.fullscreenElement) {
      document.exitFullscreen();
    } else {
      editorElement.requestFullscreen();
    }
  };

  onMount(() => {
    editorDiv.id = Math.random().toString(16);
    editor = ace.edit(editorDiv.id);
    editor.setTheme(solarized_dark);
    if (vimMode) {
      editor.setKeyboardHandler(vim.handler);
    }
    if (mode) {
      editor.session.setMode(new mode.Mode());
    }
    editor.setReadOnly(readOnly);
    editor.session.setTabSize(2);
    editor.session.setUseSoftTabs(true);
    editor.setShowPrintMargin(false);
    lightMode.subscribe((value) => {
      if (value) {
        editor.setTheme(solarized_light);
      } else {
        editor.setTheme(solarized_dark);
      }
    });

    editorElement.onfullscreenchange = () => {
      if (document.fullscreenElement) {
        fullscreen = true;
      } else {
        fullscreen = false;
      }
    };
  });
</script>

{#if mode && langName}
  <div class="flex-middle">
    <div bind:this={editorElement} class="editor-block body">
      <div class="filename editor-context component">
        <span>
          {"Editor"}
        </span>
      </div>
      <div class="editor-context component">
        <span>
          <button on:click={copy}>
            {#if !copied}
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-clipboard"
                viewBox="0 0 16 16"
              >
                <path
                  d="M4 1.5H3a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3.5a2 2 0 0 0-2-2h-1v1h1a1 1 0 0 1 1 1V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3.5a1 1 0 0 1 1-1h1v-1z"
                />
                <path
                  d="M9.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5h3zm-3-1A1.5 1.5 0 0 0 5 1.5v1A1.5 1.5 0 0 0 6.5 4h3A1.5 1.5 0 0 0 11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3z"
                />
              </svg>
            {:else}
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-clipboard-check"
                viewBox="0 0 16 16"
              >
                <path
                  fill-rule="evenodd"
                  d="M10.854 7.146a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 1 1 .708-.708L7.5 9.793l2.646-2.647a.5.5 0 0 1 .708 0z"
                />
                <path
                  d="M4 1.5H3a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3.5a2 2 0 0 0-2-2h-1v1h1a1 1 0 0 1 1 1V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3.5a1 1 0 0 1 1-1h1v-1z"
                />
                <path
                  d="M9.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5h3zm-3-1A1.5 1.5 0 0 0 5 1.5v1A1.5 1.5 0 0 0 6.5 4h3A1.5 1.5 0 0 0 11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3z"
                />
              </svg>
            {/if}
          </button><button on:click={toggleFullscreen}>
            {#if !fullscreen}
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-fullscreen"
                viewBox="0 0 16 16"
              >
                <path
                  d="M1.5 1a.5.5 0 0 0-.5.5v4a.5.5 0 0 1-1 0v-4A1.5 1.5 0 0 1 1.5 0h4a.5.5 0 0 1 0 1h-4zM10 .5a.5.5 0 0 1 .5-.5h4A1.5 1.5 0 0 1 16 1.5v4a.5.5 0 0 1-1 0v-4a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 1-.5-.5zM.5 10a.5.5 0 0 1 .5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 1 0 1h-4A1.5 1.5 0 0 1 0 14.5v-4a.5.5 0 0 1 .5-.5zm15 0a.5.5 0 0 1 .5.5v4a1.5 1.5 0 0 1-1.5 1.5h-4a.5.5 0 0 1 0-1h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 1 .5-.5z"
                />
              </svg>
            {:else}
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-fullscreen-exit"
                viewBox="0 0 16 16"
              >
                <path
                  d="M5.5 0a.5.5 0 0 1 .5.5v4A1.5 1.5 0 0 1 4.5 6h-4a.5.5 0 0 1 0-1h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 1 .5-.5zm5 0a.5.5 0 0 1 .5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 1 0 1h-4A1.5 1.5 0 0 1 10 4.5v-4a.5.5 0 0 1 .5-.5zM0 10.5a.5.5 0 0 1 .5-.5h4A1.5 1.5 0 0 1 6 11.5v4a.5.5 0 0 1-1 0v-4a.5.5 0 0 0-.5-.5h-4a.5.5 0 0 1-.5-.5zm10 1a1.5 1.5 0 0 1 1.5-1.5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 0-.5.5v4a.5.5 0 0 1-1 0v-4z"
                />
              </svg>
            {/if}
          </button>
        </span>
        <span>{langName}</span>
      </div>
      <div bind:this={editorDiv} />
    </div>
  </div>
{/if}

<style>
  button {
    padding: 3px;
    display: inline-flex;
    justify-content: center;
    align-items: center;
  }

  .filename {
    border-bottom: 1px solid var(--color-dark-component-foreground);
  }

  :global(body.light-mode) .filename {
    border-bottom: 1px solid var(--color-light-component-foreground);
  }

  .editor-block {
    margin-top: 1em;
    margin-bottom: 1em;
    border-radius: 5px;
    width: 100%;
    max-width: 85vw;
    box-sizing: border-box;
    border: 2px solid var(--color-dark-component-background);
    overflow: auto;
    height: calc(100vh - 10vh - 45px - 45px);
  }

  :global(body.light-mode) .editor-block {
    border: 2px solid var(--color-light-component-background);
  }

  .editor-context {
    padding: 5px;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
</style>
