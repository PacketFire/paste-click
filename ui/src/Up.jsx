import styles from './up.css';

import React, { Component } from 'react';

import Editor from './Editor';

import axios from 'axios';

export default class Up extends Component {
  constructor(props) {
    super(props);

    this.state = {
      text: ''
    };

    this.onDrop = this.onDrop.bind(this);
    this.onDragOver = this.onDragOver.bind(this);
    this.onMouseOut = this.onMouseOut.bind(this);
    this.onTextUpdate = this.onTextUpdate.bind(this);
    this.upload = this.upload.bind(this);
  }

  onDrop(event) {
    event.preventDefault();
    
    const files = event.dataTransfer.files;
    
    if (files.length === 1) {
      const file = files[0];
      axios.post(`${API_URL}`, file)
        .then(res => {
          const components = res.data.split('/');
          const objectId = components[components.length - 1];

          window.location = `/beta/s/${objectId}`;
        })
        .catch(err => {
          console.error(err);
        });
    }
  }

  onDragOver(event) {
    event.preventDefault();
  }

  onMouseOut(event) {
    event.preventDefault();
  }

  onTextUpdate(text) {
    this.setState({
      text
    });
  }

  upload(event) {
    axios.post(`${API_URL}`, this.state.text)
      .then(res => {
        window.location = res.data;
      })
      .catch(err => {
        console.error(err);
      });
  }

  render() {
    return (
      <div onDrop={this.onDrop} onDragOver={this.onDragOver} onMouseOut={this.onMouseOut}>
        <h2>Enter Text or Drag and Drop a File</h2>
        <Editor onTextUpdate={this.onTextUpdate} />

        <button className={styles.paste} onClick={this.upload}>Paste</button>
      </div>
    );
  };
}
