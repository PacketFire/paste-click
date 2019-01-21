import React, { Component } from 'react';

import ViewText from './ViewText';

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

    console.log('object ID is ' + objectId);
  }

  render() {
    let content = null;
    if (this.state.loaded) {
      content = (
        <ViewText text={this.state.data} />
      );
    }

    return (
      <div>
        {content}
      </div>
    );
  };
}
