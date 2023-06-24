<script>
  let dark = "/images/dark.svg";
  let light = "/images/light.png";
  let src = dark;

  let toggletheme = () => {
    window.document.body.classList.toggle("light-mode");
    src = window.document.body.classList.contains("light-mode") ? light : dark;
    localStorage.setItem(
      "theme",
      window.document.body.classList.contains("light-mode") ? "light" : "dark"
    );
  };

  let theme = localStorage.getItem("theme");
  if (theme) {
    if (theme == "light") {
      toggletheme();
    }
  } else {
    localStorage.setItem("theme", "dark");
  }
</script>

<svelte:head>
  <link
    rel="preload"
    as="image"
    type="image/svg+xml"
    href="/images/dark.svg"
    crossorigin
  />
  <link
    rel="preload"
    as="image"
    type="image/svg+xml"
    href="/images/light.png"
    crossorigin
  />
</svelte:head>

<button on:click={toggletheme}>
  <img alt={src == dark ? "Dark Theme" : "Light Theme"} {src} />
</button>

<style>
  button {
    height: 35px;
    width: 35px;
    border-radius: 50%;
  }

  img {
    height: 20px;
    width: 20px;
  }
</style>
