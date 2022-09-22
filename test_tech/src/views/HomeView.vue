<script>
import theheader from "../components/theheader.vue";
import DrawFlow from "../components/DrawFlow.vue";
export default {
  name: "App",
  components: { DrawFlow },
}
</script>

<template>
<DrawFlow :numberB="numberB" />
<h2>hola gordis</h2>
<theheader />
  
</template>

<style>

</style>