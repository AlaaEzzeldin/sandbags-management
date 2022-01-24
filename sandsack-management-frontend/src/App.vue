<template>
  <v-app>
    <v-main>
      <v-container style="background-color: white">
            <router-view style="background: white;" :key="$route.path"></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import EventBus from "./common/EventBus";

export default {
  name: 'App',

  components: {
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
