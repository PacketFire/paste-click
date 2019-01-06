<template>
  <div id="editor">

  </div>    
</template>

<script>
import CodeMirror from 'codemirror';

export default {
  mounted() {
    this.editor = CodeMirror(document.getElementById('editor'), {
      lineNumbers: true
    });

    this.editor.on('change', cm => {
      const text = cm.getValue();
      this.$root.$emit('editor-change', text);
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

@media only screen and (max-width: 900px) {
  #editor {
    width: 100%;
  }
}
</style>
