<template style="background-color: white">
  <div>
    <v-navigation-drawer
        :permanent="$vuetify.breakpoint.mdAndUp"
        absolute
        left
        style="min-width: 350px;"
        v-model="drawer"
    >
      <v-toolbar
          flat
          class="hidden-lg-and-up"
          style="width: 350px; padding-left: 0;"
          @click.stop="drawer = !drawer">
        <h3>Feuerwehr Passau ({{ getLoggedInBranchName() }})</h3>
      </v-toolbar>
      <div class="hidden-md-and-down">
        <v-row class="pt-10 pb-10 pl-5">
          <v-col cols="3">
            <img class="mr-10" src="@/assets/images/logo.png" height="60"/>
          </v-col>
          <v-col cols="9">
            <h1 style="color: red;font-weight: bolder;  font-size: x-large" class="justify-center"> Feuerwehr Passau </h1>
            <h1 style="color: black;font-weight: bolder;  font-size: large;" class="justify-center">
              {{ getLoggedInBranchName() }} </h1>
          </v-col>
        </v-row>
      </div>

      <v-list class="mb-16 mt-16">
          <v-list-item
              v-for="(item, i) in getNavListForLoggedInUserRoll()"
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
        @click.stop="drawer = !drawer">
      <v-app-bar-nav-icon></v-app-bar-nav-icon>
      <h3>Feuerwehr Passau ({{ getLoggedInBranchName() }})</h3>
    </v-toolbar>
  </div>

</template>

<script>
import EquipmentQuantityTable from '@/components/EquipmentQuantityTable.vue'

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
          title: 'Bestellungliste',
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
        }
      ],
    }
  },
  computed:{
    getCurrentUserRole(){
      return this.$store.getters.getCurrentUserRole
    }
  },
  methods:
      {
        // hard coding for the branch name
        getLoggedInBranchName() {
          return this.$store.getters.getCurrentUserName
        },
        getNavListForLoggedInUserRoll() {
          if (['Einsatzabschnitt','Hauptabschnitt','Einsatzleiter'].includes(this.getCurrentUserRole))
            return this.navItemsEinsatzabschnittAndHauptabschnitt;
          else if (this.getCurrentUserRole === 'Unterabschnitt')
            return this.navItemsUnterabschintt;
          else if (this.getCurrentUserRole === 'Mollnhof')
            return this.navItemsMollnhof;
        },
      }

}
</script>
