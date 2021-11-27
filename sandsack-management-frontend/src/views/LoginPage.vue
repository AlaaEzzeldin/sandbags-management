<template>
  <v-container>
    <v-row>
      <v-col cols="4">
        <v-form ref="form" v-model="valid">
          <h2>Einloggen bei der</h2>
          <h2>Feuerwehr Passau</h2>
          <v-text-field
              v-model="email"
              :rules="emailRules"
              label="E-mail"
              required
          ></v-text-field>
          <v-text-field
              v-model="password"
              :rules="passwordRules"
              label="Password"
              required
          ></v-text-field>
          <v-checkbox
              v-model="checkbox"
              :label="`Angemeldet bleiben`"
          ></v-checkbox>
          <v-btn text link :to="{name: links.recoverPasswordPage}">
            <u>Passwort vergessen?</u>
          </v-btn>
          <v-btn
              :disabled="!valid"
              color="success"
              class="mr-4"
              @click="validate"
              :to="`/orders-list/`+getUserRole()"
          >
            Einloggen
          </v-btn>
            Sie haben noch kein Konto?
            <v-btn text link :to="{name: links.signupPage}">
              <u>Registrieren</u>
            </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

export default {
  name: 'LoginPage',

  data: () => ({
    valid: true,
    checkbox: false,
    email: '',
    emailRules: [
      v => !!v || 'E-mail is required',
      v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
    password: '',
    passwordRules: [
      v => !!v || 'Password is required',
    ],
    links: {
      recoverPasswordPage: "RecoverPasswordPage",
      signupPage: "SignupPage",
    }
  }),

  methods: {
    validate() {
      this.$refs.form.validate()
    },
    getUserRole() {
      if (this.email === 'haupt@gmail.com')
        return 1
      else if (this.email === 'einsatz@gmail.com')
        return 2
      else if (this.email === 'unter@gmail.com')
        return 3
      else if (this.email === 'mollnhof@gmail.com')
        return 4
    }
  }

}
</script>
