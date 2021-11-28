<template>
  <v-app>
    <v-main>
      <v-container class="ma-0 pa-0" style="background-color: white">
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
