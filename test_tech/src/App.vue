<script setup>
import { RouterLink, RouterView } from "vue-router";
import theheader from "./components/theheader.vue"
</script>

<template>
  <header>
    <img
      alt="Vue logo"
      class="logo"
      src="@/assets/logo.svg"
      width="125"
      height="125"
    />
  
  </header>
  <section>
    <theheader />
  </section>
  <RouterView />
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

/*nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}*/

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
  section{
    display: flex;
    justify-content: center;
  }
  
}
</style>
