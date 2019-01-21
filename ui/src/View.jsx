import React, { Component } from 'react';

import ViewText from './ViewText';
import ViewImage from './ViewImage';
import ViewVideo from './ViewVideo';

import axios from 'axios';

export default class View extends Component {
  constructor(props) {
    super(props);

    this.state = {
      loaded: false,
      mimeType: null,
      data: null
    };
  }

  componentDidMount() {
    const { objectId } = this.props.match.params;

    axios.get(`${API_URL}/${objectId}`)
      .then(res => {
        const mimeType = res.headers['content-type'];

        this.setState({
          mimeType: mimeType,
          data: res.data,
          loaded: true
        });
      })
      .catch(err => {
        console.error(err);
      });
  }

  render() {
    let content = null;
    if (this.state.loaded) {
      const { objectId } = this.props.match.params;
      const link = `${API_URL}/${objectId}`;

      if (this.state.mimeType.startsWith('image/')) {
        content = (
          <ViewImage link={link} />
        );
      } else if (this.state.mimeType.startsWith('video/')) {
        content = (
          <ViewVideo link={link} />
        );
      } else {
        content = (
          <ViewText text={this.state.data} />
        );
      }
    }

    return (
      <div>
        {content}
      </div>
    );
  };
}
