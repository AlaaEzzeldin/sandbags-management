<!-- TODO: Add link to login page on Ausloggen button after authentication feature has been merged -->
<template>
  <div>
    <v-row v-if="isLoggedIn" no-gutters>
      <v-col sm="3" class="pt-13 justify-center align-center">
        <h1 style="font-weight: bolder; ">Konto</h1>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="12">
        <v-card color="#F1F2F6" outlined class="pa-5 mt-5">
          <v-row>
            <v-col cols="2" md="1" >
              <v-avatar
                  color="white"
                  size="60">
                <v-icon color="black" large>mdi-home</v-icon>
              </v-avatar>
            </v-col>
            <v-col cols="9" sm="10" offset="1" offset-sm="0">
              <h2>{{ getLoggedInUser.name }}</h2>
              <h3>{{ getLoggedInUser.branch_name }}</h3>
            </v-col>
          </v-row>
          <v-row class="pt-2">
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

      </v-col>
    </v-row>


    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
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
      <v-col cols="12" sm="6" offset-sm="3">
        <v-btn
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            outlined
            color="orange"
            @click="dialogKonto=true"
            dark
            block
        >
          Konto bearbeiten
        </v-btn>
      </v-col>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-btn
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            @click="dialogPassword=true"
            block
            outlined
        >
          Kennwort ändern
        </v-btn>
      </v-col>
    </v-row>

    <konto-edit-dialog :dialog="dialogKonto"
                       @close="dialogKonto=false"
    ></konto-edit-dialog>
    <password-edit-dialog :dialog="dialogPassword"
                          @close="dialogPassword=false"
    ></password-edit-dialog>
  </div>



</template>

<script>

import KontoEditDialog from "../components/KontoEditDialog";
import PasswordEditDialog from "../components/PasswordEditDialog";

export default {
  name: 'KontoPage',

  components: {KontoEditDialog, PasswordEditDialog},
  data: () => ({
    dialogKonto: null,
    dialogPassword: null
  }),
  computed: {
    isLoggedIn() {
      return this.$store.getters.isLoggedIn
    },
    getLoggedInUser() {
      return this.$store.getters.getLoggedInUser
    },

  },
  methods:
      {
        logout() {
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
