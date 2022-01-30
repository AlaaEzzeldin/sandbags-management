<template>
  <v-container fluid fill-height style="padding-top: 10%">
    <v-row align="center" justify="center" align-content="center">
      <v-col :cols="$vuetify.breakpoint.mdAndUp ? 4 : 12" align-self="center">
        <v-card elevation="0" class="justify-center align-center">
          <v-card-title>
            <h1 style="color: black;font-weight: bolder;  font-size: x-large">Einloggen bei der
              <br/>
              <span style="color: red;  font-size: x-large">Feuerwehr Passau</span>
            </h1>
          </v-card-title>
          <v-alert
              outlined
              text
              type="info"
          >Wenden Sie sich an den Systemadministrator, um Ihren Benutzernamen und Ihr Passwort zu erfahren
          </v-alert>
          <v-form ref="form" v-model="valid">
            <v-text-field
                v-model="user.email"
                :rules="emailRules"
                label="E-mail"
                required
                outlined
                prepend-inner-icon="mdi-email"
            ></v-text-field>
            <v-text-field
                v-model="user.password"
                type="password"
                :rules="passwordRules"
                label="Password"
                required
                outlined
                prepend-inner-icon="mdi-lock"
                @keyup.enter.native="submit"
            ></v-text-field>
            <v-btn
                :disabled="!valid"
                style="text-transform: capitalize; font-weight: bolder; color: white"
                color="red"
                block
                rounded
                @click="submit"
            >
              Einloggen
            </v-btn>
          </v-form>
        </v-card>
      </v-col>
    </v-row>

    <AuthFailureDialog v-if="authFailureDialog"
                       :dialog="authFailureDialog"
                       @close="authFailureDialog = false"
                       :message=" messageToDisplay"></AuthFailureDialog>
  </v-container>
</template>


<script>
import AuthFailureDialog from "../components/AuthFailureDialog";

export default {
  name: 'LoginPage',
  components: {AuthFailureDialog},
  data: () => ({
    user: {
      email: '',
      password: '',
    },
    valid: true,
    checkbox: false,
    emailRules: [
      v => !!v || 'E-mail is required',
      v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
    passwordRules: [
      v => !!v || 'Password is required',
    ],
    links: {
      recoverPasswordPage: "RecoverPasswordPage",
      signupPage: "SignupPage",
    },
    authFailureDialog: false,
    messageToDisplay: ''
  }),

  methods: {
    validate() {
      this.$refs.form.validate()
    },
    submit() {
      this.$store.dispatch('login', this.user).then(
          () => {
            this.$store.dispatch("getUserInfo").then(
                this.$router.push({name: 'BestellungslistePage'}))
          },
          error => {
           // this.message = (error.response && error.response) || error.message || error.toString();
            if (error && error.status === 404)
              this.messageToDisplay = 'Der von Ihnen angegebene Benutzername und das Passwort sind nicht korrekt, bitte versuchen Sie es erneut!'
            else this.messageToDisplay = 'Etwas ist schief gelaufen. Bitte versuchen Sie es erneut!'
            this.authFailureDialog = true
          }
      );
    }

  }

}
</script>
