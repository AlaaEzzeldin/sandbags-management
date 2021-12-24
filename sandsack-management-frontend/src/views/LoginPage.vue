<template>
  <v-container fluid fill-height>
    <v-row align="center" justify="center">
      <v-col cols="5">
        <v-card elevation="0" class="justify-center align-center mt-16 pt-16 ">
          <v-card-title class=" ml-10 pl-3">
            <h1 style="color: black;font-weight: bolder;  font-size: x-large">Einloggen bei der</h1>
          </v-card-title>
          <v-card-subtitle class=" ml-10 pl-3">
            <h1 style="color: red;font-weight: bolder;  font-size: x-large"> Feuerwehr Passau </h1>
          </v-card-subtitle>

          <v-form ref="form" v-model="valid" class=" ma-10 pa-3">

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
                    :rules="passwordRules"
                    label="Password"
                    required
                    filled
                    outlined
                    prepend-icon="mdi-lock"
                ></v-text-field>
              </v-col>
            </v-row>

            <v-row align-content="center" justify="center">
              <v-col class="text-left" cols="6">
                <v-checkbox
                    v-model="checkbox"
                    :label="`Angemeldet bleiben`"
                    class="justify-center align-center mt-0"
                ></v-checkbox>
              </v-col>
              <v-col class="text-right" cols="6">
                <v-btn text link :to="{name: links.recoverPasswordPage}" small>
                  <u>Passwort vergessen?</u>
                </v-btn>
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

            <v-row class="text-right">
              <v-col>
                Sie haben noch kein Konto?
                <v-btn text link :to="{name: links.signupPage}"
                       small
                >
                  <u>Registrieren</u>
                </v-btn>
              </v-col>
            </v-row>

          </v-form>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

export default {
  name: 'LoginPage',

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
    }
  }),

  methods: {
    validate() {
      this.$refs.form.validate()
    },
    submit() {
      console.log("submit", this.user)
      this.$store.dispatch('login', this.user).then(
          () => {
            this.$router.push({name: 'BestellungslistePage'})
            console.log("user", this.$store.state.user)

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
