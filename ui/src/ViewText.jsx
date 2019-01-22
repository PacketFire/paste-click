import styles from './viewtext.css';

import React, { Component } from 'react';

import classNames from 'classnames';
import CodeMirror from 'codemirror';

export default class ViewText extends Component {
  componentDidMount() {
    this.editor = CodeMirror(document.getElementById('editor'), {
      value: this.props.text,
      lineNumbers: true,
      readOnly: 'nocursor'
    });
  }

  render() {
    const classes = classNames(styles.editor, 'codemirror-expand');

    return (
      <div className={classes} id="editor"></div>
    );
  }
}
