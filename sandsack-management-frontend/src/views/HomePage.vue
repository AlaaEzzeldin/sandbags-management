<template>
  <v-container class="ma-0 pa-0" style="background-color: white">
    <v-row no-gutters>
      <v-col cols="12">
        <Navigation/>
      </v-col>
      <v-col cols="10" offset="1">
        <router-view style="background: white;" :key="$route.path"></router-view>
      </v-col>
    </v-row>
  </v-container>
</template>


<script>
import Navigation from '@/components/Nav.vue'

export default {
  name: 'HomagePage',

  components: {
    Navigation
  },
  created() {
    this.$store.dispatch("loadLoggedIn", '0')
    console.log("user ", this.$store.getters.getLoggedInUser)
  },

};
</script>
