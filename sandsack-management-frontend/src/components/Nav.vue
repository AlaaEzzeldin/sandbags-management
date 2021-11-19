<template style="background-color: white">
  <v-navigation-drawer
      :permanent="$vuetify.breakpoint.smAndUp"
      app
      style="position: sticky; z-index: 0"
      width="300"
      class="pa-0 ma-0"
  >
    <v-list rounded>
      <v-list-item-group
          color="blue"
      >
        <v-list-item
            v-for="(item, i) in getNavListForLoggedInUserRoll"
            :key="i"
            link
            :to="{name:item.name}"
        >
          <v-icon style="color: red" v-text="item.icon"></v-icon>
          <v-list-item-title
              v-if="$vuetify.breakpoint.lgAndUp"
              class="text-left ml-4"
              style="font-weight: bolder; font-family: 'Arial Rounded MT Bold'; font-size: large;"
              v-text="item.title"
          ></v-list-item-title>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </v-navigation-drawer>
</template>

<script>

export default {
  name: "Navigation",

  data() {
    return {
      navItemsUnterabschintt: [
        {
          title: 'NeueBestellung',
          name: 'NeueBestellung',
          icon: 'mdi-playlist-plus',
        },
        {
          title: 'Bestellungliste',
          name: 'Bestellungliste',
          icon: 'mdi-format-list-bulleted',
        },
        {
          title: 'Konto',
          name: 'Konto',
          icon: 'mdi-account',
        },
      ],
      navItemsEinsatzabschnittAndHauptabschnitt: [
        {
          title: 'NeueBestellung',
          name: 'NeueBestellung',
          icon: 'mdi-playlist-plus',
        },
        {
          title: 'Bestellungliste',
          name: 'Bestellungliste',
          icon: 'mdi-format-list-bulleted',
        },
        {
          title: 'Konto',
          name: 'Konto',
          icon: 'mdi-account',
        },
      ],
    }
  },
  computed:{
    getNavListForLoggedInUserRoll(){
        return this.navItemsUnterabschintt;
    }
  }

}
</script>