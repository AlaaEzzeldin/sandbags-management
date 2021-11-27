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
          <v-btn text>
            <u>Passwort vergessen?</u>
          </v-btn>
          <v-btn
              :disabled="!valid"
              color="success"
              class="mr-4"
              @click="validate"
          >
            Einloggen
          </v-btn>
          <v-btn text>
            Sie haben noch kein Konto? <u>Registrieren</u>
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
    email: '',
    emailRules: [
      v => !!v || 'E-mail is required',
      v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
    password: '',
    passwordRules: [
      v => !!v || 'Password is required',
    ]
  }),

  methods: {
    validate() {
      this.$refs.form.validate()
    },
  }

}
</script>
