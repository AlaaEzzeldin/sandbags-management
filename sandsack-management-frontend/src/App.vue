<template>
  <v-app>
    <v-main>
      <v-container :style="{ backgroundColor: 'white', marginLeft: $vuetify.breakpoint.mdAndUp ? '350px' : '0'}">
            <router-view style="background: white;" :key="$route.path"></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  components: {
  },
  data: () => ({
    //
  }),
};
</script>
