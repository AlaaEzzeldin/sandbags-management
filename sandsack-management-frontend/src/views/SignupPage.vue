<template>
  <v-container fluid fill-height>
    <v-row  justify="center">
      <v-col cols="5">
        <v-card elevation="0" class="justify-center align-center mt-16 pt-16 ">
          <v-card-title style="color: black;font-weight: bolder;  font-size: x-large" >Anmelden bei der</v-card-title>
          <v-card-subtitle style="color: red;font-weight: bolder;  font-size: x-large"> Feuerwehr Passau </v-card-subtitle>
          <v-card-text class="mt-3">
            <v-form ref="form" v-model="valid">
              <v-text-field
                  v-model="email"
                  :rules="emailRules"
                  label="E-mail"
                  required
                  filled
                  outlined
                  prepend-icon="mdi-email"
              ></v-text-field>
              <v-text-field
                  v-model="password"
                  type="password"
                  :rules="passwordRules"
                  label="Password"
                  required
                  filled
                  outlined
                  prepend-icon="mdi-lock"
              ></v-text-field>
              <v-text-field
                  v-model="passwordRepeat"
                  type="password"
                  :rules="[passwordConfirmationRule()]"
                  label="Repeat password"
                  required
                  filled
                  outlined
                  prepend-icon="mdi-lock"
              ></v-text-field>
              <v-select
                  v-model="rolle"
                  :items="rollen"
                  :rules="[v => !!v || 'Rolle is required']"
                  label="Rolle"
                  required
                  filled
                  outlined
                  prepend-icon="mdi-home-group"
              ></v-select>
              <v-select
                  v-model="abschnitt"
                  :items="abschnittnameList"
                  :rules="[v => !!v || 'Abschnittname is required']"
                  label="Abschnittname"
                  required
                  filled
                  outlined
                  prepend-icon="mdi-account"
              ></v-select>

              <v-row align-content="center" justify="center">
                <v-col class="text-left" cols="5" offset="1">
                  <v-checkbox
                      v-model="checkbox"
                      :label="`Angemeldet bleiben`"
                      class="justify-center align-center mt-0"
                  ></v-checkbox>
                </v-col>
                <v-col class="text-right" cols="6">
                  <v-btn text @click="goToLogin">
                    <u>Zur Login-Seite</u>
                  </v-btn>
                </v-col>
              </v-row>

              <v-row no-gutters>
                <v-col cols="10" offset="1">
                  <v-btn
                      style="text-transform: capitalize; font-weight: bolder;"
                      color="red"
                      block
                      dark
                      class="mr-4"
                      @click="validate"
                      :to="`/orders-list/`+getUserRole()"
                  >
                    Anmelden
                  </v-btn>
                </v-col>
              </v-row>

            </v-form>
          </v-card-text>
        </v-card>
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
    },
    goToLogin(){
      this.$router.push({name: 'LoginPage'})
    }
  }

}
</script>
