<template>
  <div>
    <Header></Header>
    
    <main class="main">
      <h1>Enter Text or Drag and Drop a File</h1>
      <Editor></Editor>

      <button class="upload" v-on:click="upload">Upload</button>
    </main>
  </div>
</template>

<script>
import axios from 'axios';

import Header from './Header.vue';
import Editor from './Editor.vue';

export default {
  methods: {
    reverseMessage: function() {
      this.message = this.msg;
    },
    upload: function() {
      console.log(API_URL);
      
      // const text = document.getElementById('editor').value;
      console.log(text);
    }
  },
  components: {
    Editor,
    Header
  }
};
</script>

<style scoped>
.main {
  margin: 16px 24px;
}

.upload {
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
