<!-- TODO: Add link to login page on Ausloggen button after authentication feature has been merged -->
<template>
  <div v-if="getLoggedInUser">
    <v-row no-gutters>
      <v-col sm="3" class="pt-13 justify-center align-center">
        <h1 style="font-weight: bolder;">Konto</h1>
      </v-col>
      <v-col cols="2" class="pt-15">
        <v-btn
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            outlined
            color="red"
            dark
            block
            @click="logout"
        >
          Ausloggen
        </v-btn>
      </v-col>
      <v-col cols="2" class="pt-15">
        <v-btn
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            color="orange"
            @click="dialog=true"
            dark
            block
        >
          Bearbeiten
        </v-btn>
      </v-col>
      <v-col sm="2" class=" pl-3 pt-15">
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
            <h2>{{ getLoggedInUser.name }}</h2>
            <h3>{{ getLoggedInUser.branch_name }}</h3>
          </v-col>
        </v-row>
        <v-row class="pt-2">
<!--
          <v-list-item>
            <v-list-item-icon>
              <v-icon color="black">mdi-map-marker</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{ getLoggedInUser.address }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
-->
          <v-list-item>
            <v-list-item-icon>
              <v-icon color="black">mdi-phone</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{ getLoggedInUser.phone }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-list-item>
            <v-list-item-icon>
              <v-icon color="black">mdi-email</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{ getLoggedInUser.email }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-row>
      </v-card>
    </v-row>
    <konto-edit-dialog :dialog="dialog"
                       @close="dialog=false"
    ></konto-edit-dialog>
  </div>
</template>

<script>

import KontoEditDialog from "../components/KontoEditDialog";

export default {
  name: 'KontoPage',

  components: {KontoEditDialog},
  data: () => ({
    dialog: null
  }),

  computed: {
    getLoggedInUser() {
      return this.$store.getters.getLoggedInUser
    }
  },
  methods:
      {
        logout(){
            this.$store.dispatch('logout').then(
                () => {
                  this.$router.push({name: 'LoginPage'})
                },
                error => {
                  this.message =
                      (error.response && error.response) ||
                      error.message ||
                      error.toString();
                }
            );
          }

      }
}
</script>
