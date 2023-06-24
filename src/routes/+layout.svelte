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
  <!-- <link -->
  <!--   rel="preload" -->
  <!--   as="font" -->
  <!--   type="font/otf" -->
  <!--   href="/fonts/HurmitNerdFontMono-BoldItalic.otf" -->
  <!--   crossorigin -->
  <!-- /> -->
  <!-- <link -->
  <!--   rel="preload" -->
  <!--   as="font" -->
  <!--   type="font/otf" -->
  <!--   href="/fonts/HurmitNerdFontMono-Italic.otf" -->
  <!--   crossorigin -->
  <!-- /> -->
  <!-- <link -->
  <!--   rel="preload" -->
  <!--   as="font" -->
  <!--   type="font/otf" -->
  <!--   href="/fonts/HurmitNerdFontMono-LightItalic.otf" -->
  <!--   crossorigin -->
  <!-- /> -->
  <!-- <link -->
  <!--   rel="preload" -->
  <!--   as="font" -->
  <!--   type="font/otf" -->
  <!--   href="/fonts/HurmitNerdFontMono-Light.otf" -->
  <!--   crossorigin -->
  <!-- /> -->
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

<Navbar />
<div class="background">
  <div class="opacity" />
</div>
<main>
  <slot />
</main>
<Footer />

<style>
  .opacity {
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: var(--color-dark-background);
    opacity: 85%;
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
    box-sizing: border-box;
    overflow: hidden;
    z-index: 3;
  }
</style>
