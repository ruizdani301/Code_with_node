<script>
import axios from "axios";

 export default {
  name: 'theheader',
  // pongo todas las variables que se usen en el render del template
  data: function() { 
  return {
    items: [],
    itemSelected: ''
  };
},
 methods:{
    async getOneItem() {
      console.log('item seleccionado', this.itemSelected)
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
  }
 }
</script>

<template>
  <!--<nav class="flex-container"> 
    <button class="flex-item" @click="guardar">Guardar</button>
    <button class="flex-item" @click="getOneItem">un programa <input type="text" id="getone" placeholder="nombre del programa"></button>
    <button class="flex-item" v-on:click="getAllItems">Listar programas</button>
    <button class="flex-item" @click="">Eliminar</button>
  </nav>-->
  <nav>
    <form class="flex-forms">
      <input type="button" @click="getOneItem" value="Obtener programa" class="name">
      <input type="text" id="getone" class="getone" placeholder="nombre del programa"><br>
    </form>
     <form class="flex-forms">
      <input type="submit" value="Guardar" class="name">
      <input type="text" id="getone" class="getone" placeholder="nombre del programa"><br>
    </form>

    <form>
      <label form="programs"  @click="getAllItems">listar programas</label>
      <select name="programs" id="programs" v-model="itemSelected" @change="getOneItem()">
        <option v-for="item in items" :value="item">{{ item.name }}</option>
      </select>
       <input type="submit" value="traer" class="name">
    </form>
    <!--<button class="listar" @click="getAllItems">Listar programas
      <ul v-for="item in items">
        <li>{{ item.name }}</li>
      </ul>
    </button>-->
  </nav>
</template>

<style scoped>
nav{
display:flex;
box-sizing: border-box;
place-items: center;
width:75%;
}

.flex-forms input{
  display: flex;
  background-color: #D4EFDF;
  border-radius: 5px;
  border: 0;
  box-sizing: border-box;
  justify-content: center;
  font-size:20px;
  width:90%
}

.flex-forms input.name{
background:red;
}
.flex-forms input.name:hover{
background:gold;
}

.flex-forms input #getone {
width:70%;
}
.listar{
  border-radius: 5px;
  border: 0;
  box-sizing: border-box;
  justify-content: center;
  font-size:20px;
  width:30%;
  background:red;
  
}
.listar:hover{
  background:gold;
}
/*.flex-item {
  width:22%;
  margin:1px;
  font-size:20px;
  line-height: 40px;
  border-top-left-radius:30px;
  background-color: #D4EFDF;
  border: PowderBlue 5px double;
}*/
</style>
