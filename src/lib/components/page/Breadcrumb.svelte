<script>
  import { onMount } from "svelte";

  let name = "";
  let activeSections = {};
  let io = null;
  let sections = [];
  let page = null;
  let selection = null;

  function scroller() {
    for (const section of sections) {
      if (section.innerText == selection.value) {
        section.scrollIntoView();
        page.scrollBy(0, -40);
      }
    }
  }

  function selectionUpdate() {
    let prev = null;
    for (const key in activeSections) {
      if (activeSections[key]) {
        selection.value = prev;
        break;
      }
      prev = key;
    }
    selection.value = prev;
  }

  onMount(() => {
    name = document.querySelector("div.page div.content h1");
    sections = document.querySelectorAll("div.page div.content h2");
    page = document.querySelector("div.page");
    selection = document.querySelector("div.page div.content h4 select");

    for (const [key, section] of sections.entries()) {
      if (key == 0) {
        activeSections[section.innerText] = true;
      }
      activeSections[section.innerText] = false;
    }

    io = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          activeSections[entry.target.innerText] = false;
          if (entry.isIntersecting) {
            activeSections[entry.target.innerText] = true;
          }
        });
        selectionUpdate();
      },
      {
        root: page,
        rootMargin: "0px",
        threshold: 1,
      }
    );

    for (const section of sections) {
      io.observe(section);
    }
  });
</script>

<h4>
  <span class="name">{name.innerText}</span>
  <span>>></span>
  <span class="section">
    <select on:change={scroller}>
      {#each sections as section}
        <option value={section.innerText}>{section.innerText}</option>
      {:else}
        <option>None</option>
      {/each}
    </select>
  </span>
</h4>

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
