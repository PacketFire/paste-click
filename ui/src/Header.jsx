import styles from './header.css';

import React, { Component } from 'react';

export default class Header extends Component {
  render() {
    return (
      <header className={styles.header}>
        <div className={styles.title}>
          <a className={styles.title_link} href="/beta">paste.click</a>
        </div>
        <div className={styles.links}>
          <a className={styles.paste} href="/beta/up">Paste</a>
        </div>
      </header>
    );
  }
}
