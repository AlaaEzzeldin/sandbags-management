<template>
  <v-dialog
      v-model="dialog"
      width="500"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-btn
          style="text-transform: capitalize; font-weight: bolder;"
          rounded
          color="orange"
          v-bind="attrs"
          v-on="on"
          dark
          block
      >
        Bearbeiten
      </v-btn>
    </template>

    <v-card outlined class="pa-5">
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
        <v-form
            ref="form"
            v-model="valid"
            lazy-validation
        >
          <v-text-field
              v-model="email"
              :rules="emailRules"
              label="E-mail"
              required
              filled
              outlined
          ></v-text-field>
          <v-text-field
              v-model="phone"
              :rules="phoneRules"
              label="Phone"
              required
              filled
              outlined
          ></v-text-field>
          <v-text-field
              v-model="password"
              :rules="passwordRules"
              label="Altes Passwort"
              required
              filled
              outlined
          ></v-text-field>
          <v-text-field
              v-model="password"
              :rules="passwordRules"
              label="Neues Passwort"
              required
              filled
              outlined
          ></v-text-field>
        </v-form>
      </v-row>
      <v-row class="pt-2">
        <v-card-actions>
          <v-btn
              color="primary"
              text
              @click="dialog = false"
          >
            Cancel
          </v-btn>
          <v-btn
              color="primary"
              text
              @click="dialog = false"
          >
            Speichern
          </v-btn>
        </v-card-actions>
      </v-row>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: 'KontoEditDialog',
  data() {
    let self = this;
    return {
      dialog: false,
      valid: true,
      email: self.getLoggedInBranchDetails().email,
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
      ],
      phone: self.getLoggedInBranchDetails().phone,
      phoneRules: [
        v => !!v || 'Phone is required',
        v => /\(?\+\(?49\)?[ ()]?([- ()]?\d[- ()]?){10}/.test(v) || 'Phone must be valid',
      ],
      password: '',
      passwordRules: [
        v => !!v || 'Password is required',
      ],
    }
  },

  computed: {

  },
  methods: {
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
