<template>
  <div id="editor"></div>
</template>

<script>
import CodeMirror from 'codemirror';

export default {
  mounted() {
    console.log(this.text);
    this.editor = CodeMirror(document.getElementById('editor'), {
      value: this.text,
      lineNumbers: true,
      readOnly: 'nocursor'
    });
  },
  props: ['text', 'mimeType']
};
</script>

<style scoped>
#editor {
  width: 50%;

  border: 1px solid #ddd;

  cursor: text;
}

#editor >>> .CodeMirror {
  height: auto;
}

@media only screen and (max-width: 900px) {
  #editor {
    width: 100%;
  }
}
</style>
