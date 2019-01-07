require('../node_modules/normalize.css/normalize.css');
require('../node_modules/codemirror/lib/codemirror.css');
require('./global.css');

import Vue from 'vue';
import VueRouter from 'vue-router';

import App from './App.vue';
import Up from './Up.vue';
import View from './View.vue';

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  base: '/beta',
  routes: [
    { path: '/up', component: Up },
    { path: '/s', component: View }
  ]
});

new Vue({
  el: '#app',
  router: router,
  render: h => h(App)
});
