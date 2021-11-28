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
        >
          Ausloggen
        </v-btn>
      </v-col>
      <v-col sm="2" class="pt-15">
        <v-btn
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            color="primary"
            dark
            block
        >
          Bearbeiten
        </v-btn>
      </v-col>
    </v-row>
    <v-row no-gutters>
      <v-card color="#F1F2F6" outlined min-width="600" class="pa-5 mt-5">
        <v-row>
          <v-col cols="1">
            <v-avatar
              color="white"
              size="60">
              <v-icon color="black" large>mdi-home</v-icon>
            </v-avatar>
          </v-col>
          <v-col cols="11">
            <h2>{{getLoggedInBranchName()}}</h2>
            <h3>{{getLoggedInUserRoleName()}}</h3>
          </v-col>
        </v-row>
        <v-row class="pt-2">
          <v-list-item>
            <v-list-item-icon>
              <v-icon color="black">mdi-map-marker</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>Leonhard Paminger Straße 20, 94032</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item>
            <v-list-item-icon>
              <v-icon color="black">mdi-phone</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>+49 123 45 67 8 9012</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item>
            <v-list-item-icon>
              <v-icon color="black">mdi-email</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>abschnitt11@feuerwehrpassau.de</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-row>
      </v-card>
    </v-row>
  </div>
</template>

<script>

export default {
  name: 'KontoPage',

  components: {

  },

  methods:
      {
        // hard coding for the branch name
        getLoggedInBranchName() {
          if (this.getLoggedInUserRole() === 1)
            return "Hauptabschnitt-Mitte";
          else if (this.getLoggedInUserRole() === 2)
            return "EA 1-Altstadt";
          else if (this.getLoggedInUserRole() === 3)
            return "EA 1.1 Altstadt- Ost";
          else if (this.getLoggedInUserRole() === 4)
            return "Mollnhof";
        },

        getLoggedInUserRoleName() {
          if (this.getLoggedInUserRole() === 1)
            return "Hauptabschnitt";
          else if (this.getLoggedInUserRole() === 2)
            return "Einsatzabschnitt";
          else if (this.getLoggedInUserRole() === 3)
            return "Unterabschnitt";
          else if (this.getLoggedInUserRole() === 4)
            return "Mollnhof";
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
