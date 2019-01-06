<template>
  <div id="editor">

  </div>    
</template>

<script>
import CodeMirror from 'codemirror';

export default {
  mounted: function() {
    this.editor = CodeMirror(document.getElementById('editor'), {
      lineNumbers: true
    });
  }
};
</script>

<style scoped>
#editor {
  width: 50%;

  border: 1px solid #ddd;

  cursor: text;
}
</style>
