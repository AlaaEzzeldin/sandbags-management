<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-col cols="4 ">
        <v-card elevation="0" class="justify-center align-center mt-16 pt-16 ">
          <v-card-title class=" ml-10 pl-3">
            <h1 style="color: black;font-weight: bolder;  font-size: x-large">Passwort vergessen?</h1>
          </v-card-title>
          <v-card-text>
            <v-form v-model="valid">
              <v-text-field
                  v-model="email"
                  :rules="emailRules"
                  label="E-mail"
                  required
                  filled
                  outlined
                  prepend-icon="mdi-email"
              ></v-text-field>
              <v-row no-gutters>
                <v-col >
                  <v-btn
                      style="text-transform: capitalize; font-weight: bolder;"
                      block
                      dark
                      color="red"
                      @click="validate"
                  >
                    Passwort wiederherstellen
                  </v-btn>
                </v-col>
              </v-row>
              <v-row >
                <v-col class="text-right">
                  <v-btn text small @click="goToLogin ">
                    <u>Zur Login-Seite</u>
                  </v-btn>
                </v-col>

              </v-row>

            </v-form>

          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-row>

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

  methods:
      {
        goToLogin(){
          this.$router.push({name: 'LoginPage'})
        }

      }


}
</script>
