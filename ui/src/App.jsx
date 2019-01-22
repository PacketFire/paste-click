import styles from './app.css';

import React, { Component } from 'react';
import { BrowserRouter, Route } from 'react-router-dom';

import Header from './Header';
import Up from './Up';
import View from './View';

import axios from 'axios';

export default class App extends Component {
  componentDidMount() {

  }

  render() {
    return (
      <BrowserRouter basename="/beta">
        <div className={styles.app}>
          <Header />

          <div className={styles.main}>
            <Route path="/up" component={Up} />
            <Route path="/s/:objectId" component={View} />
          </div>
        </div>
      </BrowserRouter>
    )
  }
}
