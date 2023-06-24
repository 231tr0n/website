<script>
  import { onMount, afterUpdate } from "svelte";
  import "$lib/css/page.css";

  export let blog;

  let name = "";
  let sections = [];
  let page = null;
  let selection = null;
  let breadcrumb = null;
  let content = null;

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
    sections = document.querySelectorAll("div.page div.content h2");
    sections = Array.from(sections);
    sections.unshift(name);
    content.style.paddingBottom = `calc(100vh - ${
      document.querySelector("footer").offsetHeight
    }px - ${document.querySelector("header").offsetHeight}px - ${
      sections[sections.length - 1].offsetHeight
    }px - ${breadcrumb.offsetHeight}px)`;
  });

  onMount(() => {
    content = document.querySelector("div.page div.content");
    name = document.querySelector("div.page div.content h1");
    sections = document.querySelectorAll("div.page div.content h2");
    sections = Array.from(sections);
    sections.unshift(name);
    page = document.querySelector("div.page");
    breadcrumb = document.querySelector("div.page div.content h4");
    selection = document.querySelector("div.page div.content h4 select");
    content.style.paddingBottom = `calc(100vh - ${
      document.querySelector("footer").offsetHeight
    }px - ${document.querySelector("header").offsetHeight}px - ${
      sections[sections.length - 1].offsetHeight
    }px - ${breadcrumb.offsetHeight}px)`;
  });
</script>

<div class="page flex-center" on:scroll={updateBreadcrumb}>
  <div class="content justify">
    <h4 class="component center flex-middle">
      <select class="anchor body" on:change={scroller}>
        {#each sections as section, index}
          <option value={index}>{section.innerText}</option>
        {:else}
          <option value="" />
        {/each}
      </select>
    </h4>
    <svelte:component this={blog} />
  </div>
</div>

<style>
  .page {
    width: 100%;
    height: 100%;
    overflow: auto;
    box-sizing: border-box;
    padding: 0px;
    margin: 0px;
  }

  .content {
    padding: 0px;
    margin: 0px;
    max-width: 90vw;
    min-width: 50vw;
    height: max-content;
  }

  h4 {
    padding: 5px;
    position: sticky;
    top: 0px;
    border-radius: 5px;
  }

  select {
    border-radius: 5px;
  }
</style>
