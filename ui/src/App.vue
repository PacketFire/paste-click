<template>
  <div>
    <Header></Header>
    <h1>Hello World!</h1>
    <div v-on:click="reverseMessage">
      {{ message }}
    </div>
    <input type="text" v-model="msg" />
    <input type="checkbox" v-model="isChecked" />
  </div>
</template>

<script>
  import Header from './Header.vue';

  export default {
    data: function() {
      return {
        message: 'Hello',
        msg: 'initial',
        isChecked: false
      }
    },
    methods: {
      reverseMessage: function() {
        this.message = this.msg;
      }
    },
    components: {
      Header
    }
  };
</script>

<style scoped>

</style>
