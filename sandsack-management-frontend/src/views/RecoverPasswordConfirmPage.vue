<template>
  <v-container>
    <v-row>
      <v-col cols="4">
        <v-form v-model="valid">
          <h2>Überprüfen Sie Ihren E-Mail</h2>
          <h2>Posteingang</h2>
          Der Link zur Passwortwiederherstellung wurde an Ihre E-Mail gesendet.
          Noch kein E-Mail?
          <v-btn text link :to="{name: links.recoverPasswordPage}">
            <u>Link zur Passwortwiederherstellung erneut senden</u>
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

export default {
  name: 'RecoverPasswordConfirmPage',

  data: () => ({
    links: {
      recoverPasswordPage: "RecoverPasswordPage",
    }
  }),

}
</script>
