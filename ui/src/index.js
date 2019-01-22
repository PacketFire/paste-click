require('../node_modules/normalize.css/normalize.css');
require('../node_modules/codemirror/lib/codemirror.css');
require('./global.css');

import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';

ReactDOM.render(<App />, document.getElementById('app-root'));
