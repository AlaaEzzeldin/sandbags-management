<template>
  <router-view :key="$route.path"></router-view>
</template>


<script>

export default {
  name: 'HomagePage',

  components: {
    Navigation
  },

};
</script>
