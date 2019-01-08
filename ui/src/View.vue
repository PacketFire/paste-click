<template>
  <div>
    
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {

    };
  },
  created() {
    const { objectId } = this.$route.params;

    axios.get(`${API_URL}/${objectId}`)
      .then(res => {
        console.log(res.data);
      })
      .catch(err => {
        console.error(err);
      });

    console.log('object ID is ' + objectId);
  }
}
</script>

<style scoped>

</style>
