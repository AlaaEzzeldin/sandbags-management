<template>
  <!--  <div>
      <v-row no-gutters>
        <v-col sm="10" class="pt-lg-16 justify-center align-center">
          <h1 style="font-weight:bolder;">Neue Bestellung</h1>

          <input v-model="message" placeholder="was bestellen sie?">
          <v-container fluid >
            <v-row align="left">
              <v-col
                  class="align-center"
                  cols="25"
                  sm="6"
              >
                <v-select
                    :items="items"
                    filled
                    :menu-props="{ top: true, offsetY: true }"
                    prepend-icon="mdi-home"
                    label="Sandsack"
                ></v-select>
              </v-col>

              <v-col
                  class="align-center"
                  cols="25"
                  sm="6"
              >
                <v-select
                    :items="items"
                    filled
                    :menu-props="{ top: true, offsetY: true }"
                    prepend-icon="mdi-star"
                    label="50"
                ></v-select>
                <v-row
                    align="right"
                    justify="space-around"
                >
                </v-row>
              </v-col>
            </v-row>
          </v-container>


          <input v-model="message" placeholder="wohim Liefern wir?">

          <v-select
              :items="items"
              filled
              prepend-icon="mdi-map-marker"
              :menu-props="{ top: true, offsetY: true }"
              label="Kircenplatz 1, 94032 Passau"
          ></v-select>

          <v-card
              class="pa-lg-1"
              flat
              height="200px"
              width="500"
              img="https://cdn.vuetifyjs.com/images/toolbar/map.jpg"
          >
            <v-toolbar
                dense
                floating
            >
              <v-text-field
                  hide-details
                  prepend-icon="mdi-magnify"
                  single-line
              ></v-text-field>

              <v-btn icon>
                <v-icon>mdi-crosshairs-gps</v-icon>
              </v-btn>

              <v-btn icon>
                <v-icon>mdi-dots-horizontal</v-icon>
              </v-btn>
            </v-toolbar>
          </v-card>
          <v-col cols="12">
            <neuebestellung class="mt-10"></neuebestellung>
          </v-col>
          <input v-model="message" placeholder="Andere Infomation">

          <v-select
              :items="items"
              filled
              prepend-icon="mdi-tools"
              :menu-props="{ top: true, offsetY: true }"
              label="Prioritize"
          ></v-select>
        </v-col>
      </v-row>

      <v-alert
          dense
          text
          type="success"
      >
        We need them as soon as possible please<strong>MAX 50</strong>
      </v-alert>


      <div class="text-right">
        <v-btn
            block
            color="red"
            dark
        >
          Bestellen
        </v-btn>
      </div>

    </div>-->
  <v-card elevation="0" class="pt-10">
    <v-card-title class="pt-10">
      <h1 style="font-weight: bolder; ">Neue Bestellung</h1>
    </v-card-title>

    <v-card-text class="pt-16 ">

      <v-row align="left">
        <v-col cols="12">
          <v-select
              v-model="abschnitt"
              :items="abschnittnameList"
              prepend-icon="mdi-account"
              :rules="[v => !!v || 'Bitte geben Sie den Namen des Abschnitts ein']"
              label="Abschnittname"
              required
              filled
              outlined
          ></v-select>
        </v-col>
      </v-row>

      <v-row align="left">
        <v-col sm="6">
          <v-select
              :items="types"
              filled
              outlined
              :menu-props="{ top: true, offsetY: true }"
              prepend-icon="mdi-home"
              label="was möchten Sie bestellen?"
              :rules="[v => !!v || 'Bitte geben Sie ein, was genau Sie bestellen möchten?']"
          ></v-select>
        </v-col>
        <v-col sm="6">
          <v-text-field
              filled
              outlined
              :menu-props="{ top: true, offsetY: true }"
              :rules="[v => !!v || 'Die Menge ist erforderlich']"
              prepend-icon="mdi-pound"
              label="Anzahl"
          ></v-text-field>
        </v-col>
      </v-row>

      <v-row align="left">
        <v-col cols="12">
          <v-text-field
              filled
              outlined
              prepend-icon="mdi-map-marker"
              :rules="[v => !!v || 'Die Adresse ist erforderlich']"
              :menu-props="{ top: true, offsetY: true }"
              label="Wohim Liefern wir?"
          ></v-text-field>
        </v-col>
      </v-row>

      <v-row align="left">
        <v-col cols="12">
          <v-select
              v-model="priority"
              :items="priorities"
              :rules="[v => !!v || 'Die Priorität ist erforderlich']"
              label="Priorität"
              prepend-icon="mdi-alert"
              required
              filled
              outlined

          ></v-select>
        </v-col>
      </v-row>

      <v-row align="left">
        <v-col cols="12">
          <v-textarea
              outlined
              filled
              prepend-icon="mdi-alert"
              name="input-7-4"
              label="Irgendwelche zusätzlichen Anmerkungen?"

          ></v-textarea>
        </v-col>
      </v-row>

    </v-card-text>

    <v-card-actions>
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              block
              rounded
              color="red"
              dark
          >
            Bestellen
          </v-btn>
        </v-col>
      </v-row>

    </v-card-actions>

  </v-card>
</template>

<script>

export default {
  name: 'BestelldetailsPage',


  data: () => ({
    priority: '',
    abschnitt: '',
    types: [
      'Sandsäcke',
    ],
    abschnittnameList: [
      'Hauptabschintt-Mitte',
      'EA 1-Altstadt',
      'EA 1.1 Altstadt- Ost',
      'Mollnhof',
    ],
    priorities: [
      'Niedrig',
      'Mittle',
      'Hohe',
    ],

  }),

}
</script>