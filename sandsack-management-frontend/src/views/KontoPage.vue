<!-- TODO: Add link to login page on Ausloggen button after authentication feature has been merged -->
<template>
  <div>
    <v-row no-gutters>
      <v-col sm="3" class="pt-13 justify-center align-center">
        <h1 style="font-weight: bolder;">Konto</h1>
      </v-col>
      <v-spacer></v-spacer>
      <v-col sm="2" class="pt-15">
        <v-btn
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            outlined
            color="primary"
            dark
            block
            :to="{name: ''}"
        >
          Ausloggen
        </v-btn>
      </v-col>
      <v-col sm="2" class="pt-15">
        <KontoEditDialog/>
      </v-col>
    </v-row>
    <v-row no-gutters>
      <v-card color="#F1F2F6" outlined min-width="600" class="pa-5 mt-5">
        <v-row>
          <v-col cols="2">
            <v-avatar
              color="white"
              size="60">
              <v-icon color="black" large>mdi-home</v-icon>
            </v-avatar>
          </v-col>
          <v-col cols="10">
            <h2>{{getLoggedInBranchDetails().name}}</h2>
            <h3>{{getLoggedInBranchDetails().roleName}}</h3>
          </v-col>
        </v-row>
        <v-row class="pt-2">
          <v-list-item>
            <v-list-item-icon>
              <v-icon color="black">mdi-map-marker</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{getLoggedInBranchDetails().address}}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item>
            <v-list-item-icon>
              <v-icon color="black">mdi-phone</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{getLoggedInBranchDetails().phone}}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item>
            <v-list-item-icon>
              <v-icon color="black">mdi-email</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{getLoggedInBranchDetails().email}}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-row>
      </v-card>
    </v-row>
  </div>
</template>

<script>

import KontoEditDialog from "../components/KontoEditDialog";

export default {
  name: 'KontoPage',

  components: {KontoEditDialog},

  methods:
      {
        // hard coding for the branch name
        getLoggedInBranchDetails() {
          if (this.getLoggedInUserRole() === 1)
            return {
              name: "Hauptabschnitt-Mitte",
              roleName: "Hauptabschnitt",
              address: "Leonhard-Paminger-Str. 20",
              email: "info@ffpassau.de",
              phone: "+49 123 45 67 8 9012"
            };
          else if (this.getLoggedInUserRole() === 2)
            return {
              name: "EA 1-Altstadt",
              roleName: "Einsatzabschnitt",
              address: "Leonhard-Paminger-Str. 21",
              email: "ea-altstadt@ffpassau.de",
              phone: "+49 123 45 67 8 9013"
            };
          else if (this.getLoggedInUserRole() === 3)
            return {
              name: "EA 1.1 Altstadt- Ost",
              roleName: "Unterabschnitt",
              address: "Leonhard-Paminger-Str. 22",
              email: "ea11-altstadt@ffpassau.de",
              phone: "+49 123 45 67 8 9014"
            };
          else if (this.getLoggedInUserRole() === 4)
            return {
              name: "Mollnhof",
              roleName: "Mollnhof",
              address: "Leonhard-Paminger-Str. 23",
              email: "mollnhof@ffpassau.de",
              phone: "+49 123 45 67 8 9015"
            };
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
