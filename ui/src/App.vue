<template>
  <div class="app">
    <Header></Header>
    <div class="main">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import Header from './Header.vue';

import axios from 'axios';

export default {
  components: {
    Header
  }
}
</script>

<style scoped>
.app {
  width: 100%;
}

.main {
  margin: 16px 24px;
}
</style>
