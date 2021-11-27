<template>
  <div>
    <v-row no-gutters>
      <v-col sm="10" class="pt-lg-16 justify-center align-center">
        <h1 style="font-weight:bolder;">Neue Bestellung</h1>
        <neuebestellung class="mt-10"></neuebestellung>
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
      We need them as soon as possible <strong>MAX 50</strong>
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

  </div>

</template>
