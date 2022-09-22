<script>
import Drawflow from "drawflow";
import "drawflow/dist/drawflow.min.css";

export default {
  name: "draw-flow",
  async mounted() {
    var el = document.getElementById("drawflow");
    const editor = new Drawflow(el);

    editor.start();
    editor.reroute = true;
    editor.reroute_fix_curvature = true;
    editor.reroute_curvature_start_end = 3;
    editor.reroute_curvature = 3;

    editor.createCurvature = function (startPosX, startPosY, endPosX, endPosY) {
      var centerY = (endPosY - startPosY) / 2 + startPosY;
      return (
        " M " +
        startPosX +
        " " +
        startPosY +
        " L " +
        startPosX +
        " " +
        centerY +
        " L " +
        endPosX +
        " " +
        centerY +
        " L " +
        endPosX +
        " " +
        endPosY
      );
    };

    const data = {
      name: "",
    };
    // to configure node condition
    var suma = "<b>Total suma</b>"
    editor.addNode("foo", 2, 1, 700, 250, "action", data, suma, false);

    // to configure nodo numA
    var numA = `<label for="numero">numeroA:</label>
                  <input type="number" style="width : 40px" 
                  margin: 0;  id="a"><br>`;
    editor.addNode("bar", 0, 1, 300, 400, "action", data, numA, false);

    // to configure nodo numB
    var numB = `<label for="numero">numeroB:</label>
                  <input type="number" style="width : 40px" 
                  margin: 0;  id="b" v-model="numberB"><br>`;
    editor.addNode("bar", 0, 1, 300, 100, "action", data, numB, false);

    // to configure nodo Totalsuma
    var condition = "<b> If Condition</b>"
    editor.addNode("bar", 2, 1, 10, 10, "action", data, condition, false);
    
    // to configure nodo Totalresta
    var resta = "<b>Total resta</b>"
    editor.addNode("bar", 2, 1, 10, 90, "action", data, resta, false);

    // to configure nodo Totalresta
    var resta = "<b>Total Multiplicación</b>"
    editor.addNode("bar", 2, 1, 10, 180, "action", data, resta, false);

    editor.addConnection(1, 2, "output_1", "input_1");
    editor.addConnection(1, 3, "output_1", "input_1");

    editor.zoom = 0.5;
  },
};
</script>

<template>
  <div id="drawflow"></div>
</template>

<style scoped>
#drawflow {
  width: 100%;
  height: 600px;
  border: 1px solid #000;
}

.drawflow .drawflow-node {
  display: block;
 
}

.drawflow .drawflow-node .inputs,
.drawflow .drawflow-node .outputs {
  display: flex;
  width: auto;
  display: flex;
  justify-content: center;
  align-items: center;
  
}

.drawflow .drawflow-node .input {
  top: -12px;
  left: 0px;
}

.drawflow .drawflow-node .output {
  top: -8px;
  right: 0px;
}

.drawflow .drawflow-node.action {
  width: 200px;
  height: 50px;
  padding: 0;
}


</style>
