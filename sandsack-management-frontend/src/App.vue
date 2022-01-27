<template>
  <v-app>
    <Navigation/>
    <v-main>
      <v-container fluid fill-height>
        <router-view :key="$route.path"></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Navigation from '@/components/Nav.vue'

export default {
  name: 'App',

  components: {
    Navigation
  },
  data: () => ({
    //
  }),
};
</script>
