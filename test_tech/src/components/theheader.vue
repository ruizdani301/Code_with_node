<script>
import axios from "axios";

export default {
  name: "theheader",
  //variables used in the template
  data: function () {
    return {
      items: [],
      itemSelected: "",
    };
  },
  methods: {
    async getOneItem() {
      console.log("item seleccionado", this.itemSelected);
      this.$emit("listen", this.itemSelected);
      /*try {
        let response = await axios.get("http://localhost:5000/programs/uid");
        this.items = response.data.results;
      } catch (error) {
          console.log(error);
      }*/
      //console.log(this.items);
    },

    async getAllItems() {
      try {
        let response = await axios.get("http://localhost:5000/programs");
        this.items = response.data.results;
      } catch (error) {
        console.log(error);
      }
      console.log(this.items);
    },
  },
};
</script>

<template>
  <h2>Code Nodes</h2>
  <nav>
    <form class="flex-forms">
      <input type="submit" value="Guardar" class="name" />
      <input
        type="text"
        id="getone"
        class="getone"
        placeholder="nombre del programa"
      /><br />
    </form>

    <form class="flex-list">
      <label class="list" form="programs" @click="getAllItems"
        >Listar programas</label
      >
      <select
        name="programs"
        id="programs"
        v-model="itemSelected"
        @change="getOneItem()"
      >
        <option v-for="item in items" :value="item">
          {{ item.name }}
        </option>
      </select>
    </form>
  </nav>
</template>

<style scoped>
h2 {
  font-size: 35px;
  font-style: oblique;
  color: #10a07b;
  text-align: center;
  animation-name: turn;
  animation-duration: 1s;
}
@keyframes turn {
  0% {
    color: red;
  }

  50% {
    color: lightblue;
  }
  100% {
    color: green;
  }
}
nav {
  display: flex;
  box-sizing: border-box;
  place-items: center;
  width: 30%;
}

.flex-forms input {
  display: flex;
  background-color: #d4efdf;
  border-radius: 5px;
  border: 0;
  box-sizing: border-box;
  justify-content: center;
  font-size: 20px;
  width: 90%;
}

.flex-forms input.name {
  background: red;
}
.flex-forms input.name:hover {
  background: gold;
}

.flex-forms input #getone {
  width: 70%;
}

.listar:hover {
  background: gold;
}
.flex-list {
  display: flex;
  background-color: #d4efdf;
  font-size: 20px;
  flex-direction: column;
  width: 60%;
  margin-top: -16px;
}
.flex-list #programs {
  height: 26px;
  border-radius: 5px;
}
.flex-list label {
  padding-left: 10px;
  border-radius: 5px;
  background: red;
}
.flex-list label.list:hover {
  background: gold;
}
</style>
