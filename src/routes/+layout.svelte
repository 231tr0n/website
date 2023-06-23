<script>
  import Navbar from "$lib/components/navbar/Navbar.svelte";
  import Footer from "$lib/components/footer/Footer.svelte";
  import "$lib/css/main.css";
</script>

<svelte:head>
  <link
    rel="preload"
    as="font"
    type="font/otf"
    href="/fonts/HurmitNerdFontMono-Bold.otf"
    crossorigin
  />
  <link
    rel="preload"
    as="font"
    type="font/otf"
    href="/fonts/HurmitNerdFontMono-BoldItalic.otf"
    crossorigin
  />
  <link
    rel="preload"
    as="font"
    type="font/otf"
    href="/fonts/HurmitNerdFontMono-Italic.otf"
    crossorigin
  />
  <link
    rel="preload"
    as="font"
    type="font/otf"
    href="/fonts/HurmitNerdFontMono-LightItalic.otf"
    crossorigin
  />
  <link
    rel="preload"
    as="font"
    type="font/otf"
    href="/fonts/HurmitNerdFontMono-Light.otf"
    crossorigin
  />
  <link
    rel="preload"
    as="font"
    type="font/otf"
    href="/fonts/HurmitNerdFontMono-Regular.otf"
    crossorigin
  />
  <link
    rel="preload"
    as="image"
    type="image/png"
    href="/images/background.png"
    crossorigin
  />
</svelte:head>

<div class="body">
  <Navbar />
  <div class="background">
    <div class="opacity" />
  </div>
  <main>
    <slot />
  </main>
  <Footer />
</div>

<style>
  .opacity {
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: var(--color-dark-background);
    opacity: 90%;
    z-index: 2;
  }

  :global(body.light-mode) .opacity {
    background-color: var(--color-light-background);
  }

  .background {
    background-image: url("/images/background.png");
    background-position: center;
    background-repeat: no-repeat;
    filter: blur(5px);
    position: fixed;
    top: 45px;
    left: 0px;
    right: 0px;
    bottom: 45px;
    z-index: 1;
  }

  main {
    position: fixed;
    top: 45px;
    left: 0px;
    right: 0px;
    bottom: 45px;
    padding: 0px;
    margin: 0px;
    overflow: hidden;
    z-index: 3;
  }

  .body {
    overflow: hidden;
    position: fixed;
    top: 0px;
    left: 0px;
    right: 0px;
    bottom: 0px;
    margin: 0px;
    padding: 0px;
  }
</style>
