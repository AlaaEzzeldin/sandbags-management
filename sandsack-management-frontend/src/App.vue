<template>
  <v-app>
    <v-main>
      <v-container style="background-color: white">
            <router-view style="background: white;" :key="$route.path"></router-view>
      </v-container>
    </v-main>
    <v-footer padless>
      <v-list-item
          link
          :to="{name:'HelpPage'}"
      >
        Hilfe
      </v-list-item>
    </v-footer>
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
