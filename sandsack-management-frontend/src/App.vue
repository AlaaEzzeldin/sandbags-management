<template>
  <v-app>
    <Navigation/>
    <v-main>
      <v-container fluid>
        <router-view :key="$route.path"></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Navigation from '@/components/Nav.vue'
import EventBus from "./common/EventBus";

export default {
  name: 'App',

  components: {
    Navigation
  },
  data: () => ({
    //
  }),

  methods: {
    logOut() {
      this.$store.dispatch('auth/logout');
      this.$router.push('/login');
    }
  },
  mounted() {
    EventBus.on("logout", () => {
      this.logOut();
    });
  },
  beforeDestroy() {
    EventBus.remove("logout");
  }
};
</script>
