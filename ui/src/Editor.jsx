import styles from './editor.css';

import React, { Component } from 'react';

import CodeMirror from 'codemirror';

export default class Editor extends Component {
  componentDidMount() {
    this.editor = CodeMirror(document.getElementById('editor'), {
      lineNumbers: true
    });

    this.editor.on('change', cm => {
      const text = cm.getValue();
      this.props.onTextUpdate(text);
    });
  }

  render() {
    return (
      <div className={styles.editor} id="editor"></div>
    );
  }
}
