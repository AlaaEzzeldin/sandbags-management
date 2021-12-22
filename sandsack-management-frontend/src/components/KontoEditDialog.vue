<template>
  <v-dialog
      v-model="dialog"
      width="800"
  >
    <v-card outlined >
      <v-card-title>
        <v-col cols="2">
          <v-avatar
              color="white"
              size="60">
            <v-icon color="black" large>mdi-home</v-icon>
          </v-avatar>
        </v-col>
        <v-col cols="10">
          <h2>{{getLoggedInUser.name}}</h2>
          <h3>{{getLoggedInUser.roleName}}</h3>
        </v-col>
      </v-card-title>
      <v-card-text>
        <v-form
            ref="form"
            v-model="valid"
            lazy-validation
        >
          <v-text-field
              v-model="getLoggedInUser.email"
              :rules="emailRules"
              label="E-mail"
              required
              filled
              outlined
              prepend-icon="mdi-email"
          ></v-text-field>
          <v-text-field
              v-model="getLoggedInUser.phone"
              :rules="phoneRules"
              label="Phone"
              required
              filled
              outlined
              prepend-icon="mdi-phone"
          ></v-text-field>
          <v-text-field
              v-model="getLoggedInUser.password"
              :rules="passwordRules"
              label="Altes Passwort"
              prepend-icon="mdi-lock"
              required
              filled
              outlined
          ></v-text-field>
          <v-text-field
              v-model="getLoggedInUser.password"
              :rules="passwordRules"
              label="Neues Passwort"
              prepend-icon="mdi-lock"
              required
              filled
              outlined
          ></v-text-field>
        </v-form>
      </v-card-text>
        <v-card-actions>
          <v-row>
            <v-col class="align-center justify-center" cols="3" offset="3">
              <v-btn
                  style="text-transform: capitalize; font-weight: bolder;"
                  block
                  rounded
                  color="red"
                  dark
                  @click="dialog = false"
              >
                Abrechen
              </v-btn>

            </v-col>
            <v-col cols="3">

              <v-btn
                  style="text-transform: capitalize; font-weight: bolder;"
                  block
                  rounded
                  color="red"
                  dark
                  @click="submitUpdatedInfo"
              >
                Speichern
              </v-btn>
            </v-col>

          </v-row>

        </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>

export default {
  name: 'KontoEditDialog',
  props: ['dialog'],

  data() {
    return {
      valid: true,
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
      ],
      phoneRules: [
        v => !!v || 'Phone is required',
        v => /\(?\+\(?49\)?[ ()]?([- ()]?\d[- ()]?){10}/.test(v) || 'Phone must be valid',
      ],
      passwordRules: [
        v => !!v || 'Password is required',
      ],
    }
  },

  computed: {
    getLoggedInUser() {
      return this.$store.getters.getLoggedInUser
    }
    },
  methods: {
    submitUpdatedInfo(){
      let data={
        "address": this.getLoggedInUser.address,
        "email": this.getLoggedInUser.email,
        "phone": this.getLoggedInUser.phone,
        "password": this.getLoggedInUser.password,
      }
      let id= this.getLoggedInUser.id
      this.$store.dispatch("updateUserInfo",  {id, data} )
      this.$emit("close")
    },
/*    // hard coding for the branch name
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
    }*/
  }
}
</script>
