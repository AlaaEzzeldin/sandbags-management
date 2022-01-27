<template>
  <v-container fluid fill-height>
    <v-row align="center" justify="center" align-content="center">
      <v-col :cols="$vuetify.breakpoint.mdAndUp ? 4 : 12" align-self="center">
        <v-card elevation="0" class="justify-center align-center">
          <v-card-title>
            <h1 style="color: black;font-weight: bolder;  font-size: x-large">Einloggen bei der</h1>
          </v-card-title>
          <v-card-subtitle>
            <h1 style="color: red;font-weight: bolder;  font-size: x-large"> Feuerwehr Passau </h1>
          </v-card-subtitle>
          <v-form ref="form" v-model="valid" class=" pa-3">
            <v-row>
              <v-col>
                <v-alert
                    outlined
                    text
                    type="info"
                >Wenden Sie sich an den Systemadministrator, um Ihren Benutzernamen und Ihr Passwort zu erfahren
                </v-alert>
              </v-col>
            </v-row>
            <v-row>
              <v-col>
                <v-text-field
                    v-model="user.email"
                    :rules="emailRules"
                    label="E-mail"
                    required
                    filled
                    outlined
                    prepend-icon="mdi-email"
                ></v-text-field>
              </v-col>
            </v-row>

            <v-row>
              <v-col>
                <v-text-field
                    v-model="user.password"
                    type="password"
                    :rules="passwordRules"
                    label="Password"
                    required
                    filled
                    outlined
                    prepend-icon="mdi-lock"
                ></v-text-field>
              </v-col>
            </v-row>

            <v-row no-gutters>
              <v-col>
                <v-btn
                    :disabled="!valid"
                    style="text-transform: capitalize; font-weight: bolder; color: white"
                    color="red"
                    block
                    class="mr-4"
                    @click="submit"
                >
                  Einloggen
                </v-btn>
              </v-col>
            </v-row>

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
            this.$router.push({name: 'BestellungslistePage'})
          },
          error => {
            this.message = (error.response && error.response) || error.message || error.toString();
            if (this.message.data.err_code === 404)
              this.messageToDisplay = 'Der von Ihnen angegebene Benutzername und das Passwort sind nicht korrekt, bitte versuchen Sie es erneut!'
            else this.messageToDisplay = 'Etwas ist schief gelaufen. Bitte versuchen Sie es erneut!'
            this.authFailureDialog=true
          }
      );
    }

  }

}
</script>
