<template>
  <div v-on:drop="drop" v-on:dragover="dragOver" v-on:mouseout="mouseOut">
    <main class="main">
      <h2>Enter Text or Drag and Drop a File</h2>
      <Editor></Editor>

      <button class="paste" v-on:click="upload">Paste</button>
    </main>
  </div>
</template>

<script>
import axios from 'axios';

import Editor from './Editor.vue';

export default {
  data() {
    return {
      text: ''
    }
  },
  created() {
    this.$root.$on('editor-change', text => {
      this.text = text;
    });
  },
  methods: {
    drop(e) {
      e.preventDefault();
      
      const files = e.dataTransfer.files;
      
      if (files.length === 1) {
        const file = files[0];
        axios.post(`${API_URL}`, file)
          .then(res => {
            window.location = res.data;
          })
          .catch(err => {
            console.error(err);
          });
      }
    },
    dragOver(e) {
      e.preventDefault();
    },
    mouseOut(e) {
      e.preventDefault();
    },
    upload() {
      axios.post(`${API_URL}`, this.text)
        .then(res => {
          window.location = res.data;
        })
        .catch(err => {
          console.error(err);
        });
    }
  },
  components: {
    Editor
  }
};
</script>

<style scoped>
.main {
  margin: 16px 24px;
}

.paste {
  margin-top: 16px;
  padding: 10px 18px;

  color: #fff;
  font-size: 16px;
  font-family: 'Roboto', sans-serif;
  font-weight: 700;

  background-color: #0389ff;

  outline: 0;
  border: none;
  border-radius: 2px;

  cursor: pointer;
}
</style>
