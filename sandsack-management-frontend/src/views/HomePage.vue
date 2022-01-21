<template>
  <router-view :key="$route.path"></router-view>
</template>


<script>

export default {
  name: 'HomagePage',
  created() {
    this.$store.dispatch("loadLoggedIn", '0')
  },

};
</script>
