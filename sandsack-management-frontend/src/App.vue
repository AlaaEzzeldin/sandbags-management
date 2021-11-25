<template>
  <v-app>
    <v-main>
      <v-container class="ma-0 pa-0" style="background-color: white">
        <v-row no-gutters v-if="isAuthPage()">
          <router-view style="background: white;" :key="$route.path"></router-view>
        </v-row>
        <v-row no-gutters v-if="!isAuthPage()">
          <v-col cols="12">
            <Navigation/>
          </v-col>
          <v-col cols="10" offset="1">
            <router-view style="background: white;" :key="$route.path"></router-view>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Navigation from './components/Nav'

export default {
  name: 'App',

  components: {
    Navigation,
  },

  methods: {
     isAuthPage() {
       return !!this.currentRoute.includes('/auth/');
     }
  },

  data: () => ({
    currentRoute: window.location.pathname
  }),
};
</script>
