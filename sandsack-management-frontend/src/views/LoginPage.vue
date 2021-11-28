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
                    v-model="email"
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
                    v-model="password"
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
                    style="text-transform: capitalize; font-weight: bolder;"
                    color="red"
                    dark
                    block
                    class="mr-4"
                    @click="validate"
                    :to="`/orders-list/`+getUserRole()"
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
