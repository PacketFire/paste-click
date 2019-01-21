import styles from './viewtext.css';

import React, { Component } from 'react';

import CodeMirror from 'codemirror';

export default class ViewText extends Component {
  componentDidMount() {
    console.log(this.props.text);
    this.editor = CodeMirror(document.getElementById('editor'), {
      value: this.props.text,
      lineNumbers: true,
      readOnly: 'nocursor'
    });
  }

  render() {
    return (
      <div className={styles.editor} id="editor"></div>
    );
  }
}
