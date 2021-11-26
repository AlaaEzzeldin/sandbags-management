<template>



  <v-container class="grey lighten-5">
    <h1 style="font-weight: bolder;">Bestellübersicht</h1>


    <v-select
        :items="items"
        :menu-props="{ top: true, offsetY: true }"
        label="Heute"
    ></v-select>

    <v-row no-gutters>
      <v-col order="last">
        <v-card
            class="pa-2"
            outlined
            tile
        >
          Bestellungen/Uhr
        </v-card>
      </v-col>
      <v-col>
        <v-card
            class="pa-2"
            outlined
            tile
        >
          Bestullungen bestaigt
        </v-card>
      </v-col>
      <v-col order="first">
        <v-card
            class="pa-2"
            outlined
            tile
        >
          Bestallungen
        </v-card>

        <input v-model="message" placeholder="Bestellungen">

        <v-select
            :items="items"
            :menu-props="{ top: true, offsetY: true }"
            label="Beim Abschnitt"
        ></v-select>

        <v-sparkline
            :value="value"
            :gradient="gradient"
            :smooth="radius || false"
            :padding="padding"
            :line-width="width"
            :stroke-linecap="lineCap"
            :gradient-direction="gradientDirection"
            :fill="fill"
            :type="type"
            :auto-line-width="autoLineWidth"
            auto-draw
        ></v-sparkline>


        <v-btn

            rounded
            color="red"
            align="right"
            dark
        >
          PDF EXPORT
        </v-btn>



      </v-col>
    </v-row>
  </v-container>


</template>

<script>
const gradients = [
  ['#222'],
  ['#42b3f4'],
  ['red', 'orange', 'yellow'],
  ['purple', 'violet'],
  ['#00c6ff', '#F0F', '#FF0'],
  ['#f72047', '#ffd200', '#1feaea'],
]

export default {
  data: () => ({
    width: 2,
    radius: 10,
    padding: 8,
    lineCap: 'round',
    gradient: gradients[5],
    value: [0, 5, 8, 3, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 20, 0,],
    gradientDirection: 'top',
    gradients,
    fill: false,
    type: 'trend',
    autoLineWidth: false,
  }),
}
</script>





