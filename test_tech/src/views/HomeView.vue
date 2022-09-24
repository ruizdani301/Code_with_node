<script>
import theheader from "../components/theheader.vue";
import DrawFlow from "../components/DrawFlow.vue";
export default {
  name: "App",
  components: { DrawFlow, theheader },

  data() {
    return {
      dateToDraw: "",
    };
  },

  methods: {
    toDraw(value) {
      this.dateToDraw = value;
      console.log("ste array llega al padre", this.dateToDraw);
      alert("dateToDraw");
    },
  },
};
</script>

<template>
  <theheader @listen="toDraw" />
  <DrawFlow :valor="dateToDraw" />
</template>

<style></style>
