<script>
  import { onMount, afterUpdate } from "svelte";
  import "$lib/css/page.css";

  export let blog;

  let name = "";
  let sections = [];
  let page = null;
  let selection = null;
  let breadcrumb = null;

  let scroller = () => {
    sections[selection.value].scrollIntoView();
    page.scrollBy(0, -breadcrumb.offsetHeight - 4);
  };

  let updateBreadcrumb = () => {
    let prev = null;
    for (const [index, section] of sections.entries()) {
      if (
        breadcrumb.offsetTop + breadcrumb.offsetHeight + 5 <
        section.offsetTop
      ) {
        if (prev) {
          selection.value = index - 1;
        } else {
          selection.value = index;
        }
        break;
      } else {
        selection.value = index;
      }
      prev = section;
    }
  };

  afterUpdate(() => {
    name = document.querySelector("div.page div.content h1");
    sections = document.querySelectorAll("div.page div.content h2");
    sections = Array.from(sections);
    sections.unshift(name);
  });

  onMount(() => {
    name = document.querySelector("div.page div.content h1");
    sections = document.querySelectorAll("div.page div.content h2");
    sections = Array.from(sections);
    sections.unshift(name);
    page = document.querySelector("div.page");
    breadcrumb = document.querySelector("div.page div.content h4");
    selection = document.querySelector("div.page div.content h4 select");
  });
</script>

<div class="page" on:scroll={updateBreadcrumb}>
  <div class="content">
    <h4>
      <span class="name">{name.innerText}</span>
      <span>>></span>
      <span class="section">
        <select on:change={scroller}>
          {#each sections as section, index}
            <option value={index}>{section.innerText}</option>
          {:else}
            <option value="" hidden />
          {/each}
        </select>
      </span>
    </h4>
    <svelte:component this={blog} />
  </div>
</div>

<style>
  .name {
    color: var(--color-dark-heading);
  }

  :global(body.light-mode) .name {
    color: var(--color-light-heading);
  }

  h4 {
    vertical-align: middle;
    padding: 5px;
    position: sticky;
    top: 0px;
    border-radius: 5px;
    background-color: var(--color-dark-component);
    text-align: center;
  }

  :global(body.light-mode) h4 {
    background-color: var(--color-light-component);
  }
</style>
