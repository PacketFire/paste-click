<template>
  <div class="app">
    <Header></Header>
    <div>{{ this.$route.meta.header }}</div>
    <router-view></router-view>
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
</style>
