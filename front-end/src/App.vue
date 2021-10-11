<template>
  <div id="app">
    <div class="header">
      <div>
        <form
          action="http://localhost:3333/files"
          method="POST"
          enctype="multipart/form-data"
          target="request"
        >
          <label for="file">Subir proyecto</label>
          <input type="file" name="file" />
          <input type="submit" value="Enviar" />
        </form>

        <br />

        <label for="project">Cargar Proyecto</label>
        <select name="project" autocomplete="">
          <option value="0" disabled selected>-------</option>
          <option v-for="item in object" v-bind:key="item">{{ item }}</option>
        </select>
      </div>

      <div>
        <h3>Codigo Phyton</h3>
        <textarea name="code" id="code" v-model="content"></textarea>
      </div>
    </div>
    <hr />

    <Rete />
  </div>
</template>

<script>
import Rete from "./components/Rete.vue";

export default {
  name: "App",
  components: {
    Rete,
  },
  data() {
    return {
      object: {
        project1: "Uno",
        project2: "Dos",
        project3: "Tres",
        project4: "Cuatro",
        project5: "Cinco",
        project6: "Seis",
      },
      content: "",
    };
  },
  methods: {},
  mounted() {
    const axios = require("axios").default;
    axios.get("http://localhost:3333/files").then((response) => {
      console.log(response.data);
    });

    axios.get("http://localhost:3333/files/notes.txt").then((response) => {
      this.content = response.data;
    });
  },
};
</script>

<style>
.header {
  display: flex;
  gap: 2rem;
  align-items: center;
}
h3 {
  margin: 0;
  padding: 0;
  text-align: center;
}
select {
  width: 75%;
}
textarea {
  width: 60rem;
  height: 5rem;
}
</style>
