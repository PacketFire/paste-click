import styles from './viewimage.css';

import React, { Component } from 'react';

export default class ViewImage extends Component {
  render() {
    return (
      <div>
        <a href={this.props.link}>
          <img className={styles.image} src={this.props.link} />
        </a>
      </div>
    );
  }
}
