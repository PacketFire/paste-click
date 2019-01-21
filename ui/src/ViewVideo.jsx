import styles from './viewvideo.css';

import React, { Component } from 'react';

export default class ViewImage extends Component {
  render() {
    return (
      <div>
        <video className={styles.video} src={this.props.link} controls />
      </div>
    );
  }
}
