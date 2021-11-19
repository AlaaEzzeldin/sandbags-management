<template>
  <v-container class="mt-0 pt-0" style="background-color: white">
    <v-row no-gutters>
      <v-col cols="12" class="hidden-sm-and-down" md="1" lg="3" xl="2" offset-xl="2" z-index="0">
        <Navigation/>
      </v-col>
      <v-col cols="12" xs="11" sm="11" md="11" lg="9" xl="6" class="mb-10">
        <router-view style="background: white;" :key="$route.path"></router-view>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Navigation from './components/Nav'

export default {
  name: 'App',

  components: {
    Navigation,
  },
  data: () => ({
    //
  }),
};
</script>
