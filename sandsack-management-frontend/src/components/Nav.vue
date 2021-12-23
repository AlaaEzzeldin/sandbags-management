<template style="background-color: white">
  <v-navigation-drawer
      :permanent="$vuetify.breakpoint.smAndUp"
      width="350"
      app
      class="logo pa-0 ma-0"
  >
    <div>
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

    <v-list rounded class="mb-16">
      <v-list-item-group>
        <v-list-item
            v-for="(item, i) in getNavListForLoggedInUserRoll()"
            :key="i"
            link
            :to="{name:item.component}"
        >
          <v-icon style="color: red" v-text="item.icon"></v-icon>
          <v-list-item-title
              v-if="$vuetify.breakpoint.lgAndUp"
              class="text-left ml-2"
              style="font-weight: bolder;  font-size: large;"
              v-text="item.title"
          ></v-list-item-title>
        </v-list-item>
      </v-list-item-group>
    </v-list>
    <EquipmentQuantityTable/>
  </v-navigation-drawer>
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
  computed: {},

  methods:
      {
        // hard coding for the branch name
        getLoggedInBranchName() {
          if (this.getLoggedInUserRole() === 1)
            return "Hauptabschintt-Mitte";
          else if (this.getLoggedInUserRole() === 2)
            return "EA 1-Altstadt";
          else if (this.getLoggedInUserRole() === 3)
            return "EA 1.1 Altstadt- Ost";
          else if (this.getLoggedInUserRole() === 4)
            return "Mollnhof";
        },
        getNavListForLoggedInUserRoll() {
          if (this.getLoggedInUserRole() === 1 || this.getLoggedInUserRole() === 2)
            return this.navItemsEinsatzabschnittAndHauptabschnitt;
          else if (this.getLoggedInUserRole() === 3)
            return this.navItemsUnterabschintt;
          else if (this.getLoggedInUserRole() === 4)
            return this.navItemsMollnhof;
        },

        // hard coding the users roles
        getLoggedInUserRole() {
          if (this.$route.params.userRole === '1') // Hauptabschintt
            return 1
          else if (this.$route.params.userRole === '2') // Einzatsabschnitt
            return 2
          else if (this.$route.params.userRole === '3') //Unterabschnitt
            return 3
          else if (this.$route.params.userRole === '4') // Mollhof
            return 4
        }
      }

}
</script>
