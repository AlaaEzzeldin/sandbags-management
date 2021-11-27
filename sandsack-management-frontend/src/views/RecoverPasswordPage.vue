<template>
  <v-container>
    <v-row>
      <v-col cols="4">
        <v-form v-model="valid">
          <h2>Passwort</h2>
          <h2>vergessen?</h2>
          <v-text-field
              v-model="email"
              :rules="emailRules"
              label="E-mail"
              required
          ></v-text-field>
          <v-btn
              :disabled="!valid"
              color="success"
              class="mr-4"
              @click="validate"
          >
            Passwort wiederherstellen
          </v-btn>
          <v-btn text link :to="{name: links.logInPage}">
            <u>Zur Login-Seite</u>
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

export default {
  name: 'RecoverPasswordPage',

  data: () => ({
    valid: true,
    email: '',
    emailRules: [
      v => !!v || 'E-mail is required',
      v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
    links: {
      logInPage: "LoginPage",
    }
  }),

}
</script>
