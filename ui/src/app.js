require('../node_modules/normalize.css/normalize.css');
require('../node_modules/codemirror/lib/codemirror.css');
require('./global.css');

import Vue from 'vue';
import App from './App.vue';

new Vue({
  el: '#app',
  render: h => h(App)
});
