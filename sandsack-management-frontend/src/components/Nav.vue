<template style="background-color: white">
  <div>
    <v-navigation-drawer
        :permanent="$vuetify.breakpoint.mdAndUp"
        left
        app
        v-model="drawer"
        width="350"
        v-if="isLoggedIn"
    >

    <v-row class="pt-lg-8 pt-3">
      <v-col cols="3">
        <img src="@/assets/images/logo.png" height="60"/>
      </v-col>
      <v-col cols="9">
        <h1 style="color: red;font-weight: bolder;  font-size: x-large; white-space: nowrap;"> Feuerwehr Passau </h1>
        <h1 style="color: black;font-weight: bolder;  font-size: large;">{{ getLoggedInUser.name }}</h1>
      </v-col>
    </v-row>

      <v-list class="mb-16 mt-16">
          <v-list-item
              v-for="(item, i) in getNavListForLoggedInUserRoll"
              :key="i"
              link
              :to="{name:item.component}"
          >
            <v-list-item-icon>
              <v-icon style="color: red" v-text="item.icon"></v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title
                  class="text-left ml-2"
                  style="font-weight: bolder;  font-size: large;"
                  v-text="item.title"
              ></v-list-item-title>
            </v-list-item-content>
          </v-list-item>
      </v-list>
      <EquipmentQuantityTable/>
    </v-navigation-drawer>
    <v-toolbar
        flat
        class="hidden-md-and-up"
        style="width: 350px; padding-left: 0;"
        @click.stop="drawer = !drawer"
        v-if="isLoggedIn"
    >
      <v-app-bar-nav-icon></v-app-bar-nav-icon>
      <h3>Feuerwehr Passau</h3>
    </v-toolbar>
  </div>
</template>

<script>
import EquipmentQuantityTable from '@/components/EquipmentQuantityTable.vue'
import EventBus from "../common/EventBus";

export default {
  name: "Navigation",

  components: {
    EquipmentQuantityTable
  },

  data() {
    return {
      drawer: null,
      navItemsUnterabschintt: [
        {
          title: 'Bestellungsliste',
          component: 'BestellungslistePage',
          icon: 'mdi-format-list-bulleted',
        },
        {
          title: 'Neue Bestellung',
          component: 'NeueBestellungPage',
          icon: 'mdi-plus',
        },
        {
          title: 'Konto',
          component: 'KontoPage',
          icon: 'mdi-account',
        },
        {
          title: 'Hilfe',
          component: 'HelpPage',
          icon: 'mdi-help-circle-outline',
        },
      ],
      navItemsEinsatzabschnittAndHauptabschnitt: [

        {
          title: 'Bestellungsliste',
          component: 'BestellungslistePage',
          icon: 'mdi-format-list-bulleted',
        },
        {
          title: 'Bestellübersicht',
          component: 'BestellübersichtPage',
          icon: 'mdi-chart-line',
        },
        {
          title: 'Konto',
          component: 'KontoPage',
          icon: 'mdi-account',
        },
        {
          title: 'Hilfe',
          component: 'HelpPage',
          icon: 'mdi-help-circle-outline',
        },
      ],
      navItemsMollnhof: [
        {
          title: 'Bestellungsliste',
          component: 'BestellungslistePage',
          icon: 'mdi-format-list-bulleted',
        },
        {
          title: 'Ausrüstungverwaltung',
          component: 'ManageEquipmentPage',
          icon: 'mdi-cog'
        },
        {
          title: 'Konto',
          component: 'KontoPage',
          icon: 'mdi-account',
        },
        {
          title: 'Hilfe',
          component: 'HelpPage',
          icon: 'mdi-help-circle-outline',
        },
      ],
    }
  },
  created() {
    if(this.isLoggedIn)
    {
      this.$store.dispatch("getUserInfo").then(
          response => {
            console.log(response)
          },
          error => {
            this.content =
                (error.response && error.response.data && error.response.data.message) ||
                error.message ||
                error.toString();

            if (error.response && error.response.status === 403) {
              EventBus.dispatch("logout");
            }
          }
      );
      this.$store.dispatch("loadEquipment");
      this.$store.dispatch("loadPriorities");
    }
  },
  computed: {
    getCurrentUserRole() {
      return this.$store.getters.getCurrentUserRole
    },
    getLoggedInUser() {
      return this.$store.getters.getLoggedInUser
    },
    isLoggedIn() {
      return this.$store.getters.isLoggedIn
    },
    getNavListForLoggedInUserRoll() {
      if (['Einsatzabschnitt','Hauptabschnitt','Einsatzleiter'].includes(this.getCurrentUserRole))
        return this.navItemsEinsatzabschnittAndHauptabschnitt;
      else if (this.getCurrentUserRole === 'Unterabschnitt')
        return this.navItemsUnterabschintt;
      else if (this.getCurrentUserRole === 'Mollnhof')
        return this.navItemsMollnhof;
      else return null
    }
  },


}
</script>
