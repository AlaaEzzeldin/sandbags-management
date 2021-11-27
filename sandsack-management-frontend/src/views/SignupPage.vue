<template>
  <v-container>
    <v-row>
      <v-col cols="4">
        <v-form v-model="valid">
          <h2>Anmelden bei der</h2>
          <h2>Feuerwehr Passau</h2>
          <v-text-field
              v-model="email"
              :rules="[v => !!v || 'Email is required']"
              label="E-mail"
              required
          ></v-text-field>
          <v-text-field
              v-model="password"
              :rules="[v => !!v || 'Password is required']"
              label="Password"
              required
          ></v-text-field>
          <v-text-field
              v-model="password"
              :rules="[v => !!v || 'Password is required']"
              label="Repeat password"
              required
          ></v-text-field>
          <v-select
              v-model="select"
              :items="rolleList"
              :rules="[v => !!v || 'Rolle is required']"
              label="Rolle"
              required
          ></v-select>
          <v-select
              v-model="select"
              :items="abschnittnameList"
              :rules="[v => !!v || 'Abschnittname is required']"
              label="Abschnittname"
              required
          ></v-select>
          <v-checkbox
              v-model="checkbox"
              :label="`Angemeldet bleiben`"
          ></v-checkbox>
          <v-btn text>
            <u>Zur Login-Seite</u>
          </v-btn>
          <v-btn
              :disabled="!valid"
              color="success"
              class="mr-4"
              @click="validate"
          >
            Anmelden
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

export default {
  name: 'SignupPage',

  data: () => ({
    abschnittnameList: [
      'Hauptabschintt-Mitte',
      'EA 1-Altstadt',
      'EA 1.1 Altstadt- Ost',
      'Mollnhof',
    ],
    rolleList: [
        'Unterabschnittleiter',
        'Einsatzabschnittleiter',
        'Hauptabschnittleiter',
        'Mollnhofleiter',
    ]
  }),

}
</script>
