<template>
  <div v-if="this.loaded">
    <ViewText v-bind:text="this.data" v-bind:mime-type="this.mimeType"></ViewText>
  </div>
</template>

<script>
import ViewText from './ViewText.vue';

import axios from 'axios';

export default {
  data() {
    return {
      loaded: false,
      mimeType: null,
      data: null
    };
  },
  created() {
    const { objectId } = this.$route.params;

    axios.get(`${API_URL}/${objectId}`)
      .then(res => {
        const mimeType = res.headers['content-type'];

        this.mimeType = mimeType;
        this.data = res.data;
        this.loaded = true;
      })
      .catch(err => {
        console.error(err);
      });

    console.log('object ID is ' + objectId);
  },
  components: {
    ViewText
  }
}
</script>

<style scoped>

</style>
