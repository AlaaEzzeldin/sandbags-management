<template>
  <v-container>
    <v-row>
      <v-col cols="4">
        <v-form ref="form" v-model="valid">
          <h2>Anmelden bei der</h2>
          <h2>Feuerwehr Passau</h2>
          <v-text-field
              v-model="email"
              :rules="emailRules"
              label="E-mail"
              required
          ></v-text-field>
          <v-text-field
              v-model="password"
              type="password"
              :rules="passwordRules"
              label="Password"
              required
          ></v-text-field>
          <v-text-field
              v-model="passwordRepeat"
              type="password"
              :rules="[passwordConfirmationRule()]"
              label="Repeat password"
              required
          ></v-text-field>
          <v-select
              v-model="rolle"
              :items="rollen"
              :rules="[v => !!v || 'Rolle is required']"
              label="Rolle"
              required
          ></v-select>
          <v-select
              v-model="abschnitt"
              :items="abschnittnameList"
              :rules="[v => !!v || 'Abschnittname is required']"
              label="Abschnittname"
              required
          ></v-select>
          <v-checkbox
              v-model="checkbox"
              :label="`Angemeldet bleiben`"
          ></v-checkbox>
          <v-btn text link :to="{name: links.logInPage}">
            <u>Zur Login-Seite</u>
          </v-btn>
          <v-btn
              :disabled="!valid"
              color="success"
              class="mr-4"
              @click="validate"
              :to="`/orders-list/`+getUserRole()"
          >
            Anmelden
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

export default {
  name: 'SignupPage',

  data: () => ({
    valid: true,
    checkbox: false,
    rolle: '',
    abschnitt: '',
    email: '',
    emailRules: [
      v => !!v || 'E-mail is required',
      v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
    password: '',
    passwordRepeat: '',
    passwordRules: [
      v => !!v || 'Password is required',
    ],
    abschnittnameList: [
      'Hauptabschintt-Mitte',
      'EA 1-Altstadt',
      'EA 1.1 Altstadt- Ost',
      'Mollnhof',
    ],
    rollen: [
      'Hauptabschnitt',
      'Einsatzabschnitt',
      'Unterabschnitt',
      'Mollnhof',
    ],
    links: {
      logInPage: "LoginPage",
    }
  }),

  methods: {
    validate() {
      this.$refs.form.validate()
    },
    passwordConfirmationRule() {
      return () => (this.password === this.passwordRepeat) || 'Password must match'
    },
    getUserRole() {
        if (this.rolle === 'Hauptabschnitt')
          return 1
        else if (this.rolle === 'Einsatzabschnitt')
          return 2
        else if (this.rolle === 'Unterabschnitt')
          return 3
        else if (this.rolle === 'Mollnhof')
          return 4
    }
  }

}
</script>
