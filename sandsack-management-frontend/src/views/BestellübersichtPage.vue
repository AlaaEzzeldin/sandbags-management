<template>
  <div>
    <v-row no-gutters>
      <v-col sm="3" class="pt-13 justify-center align-center">
        <h1 style="font-weight: bolder;">Bestellübersicht</h1>
      </v-col>
      <v-spacer></v-spacer>
      <v-col sm="2" class="pt-15 justify-center align-center">

        <v-select
            :items="items"
            label="Tage"
            solo
        ></v-select>

      </v-col>
    </v-row>
    <v-row no-gutters>
      <v-text-field
          v-model="first"
          value="10.00"
          label="Bestellungen"
          filled
      ></v-text-field>

      <v-text-field
          v-model="second"
          value="10.00"
          label="Bestellungen bestatigt"
          filled
      ></v-text-field>

      <v-text-field
          v-model="third"
          value="10.00"
          label="Bestellungen/Uhr"
          filled
      ></v-text-field>
      <v-col cols="12">
        <Bestellübersicht class="mt-10"></Bestellübersicht>
      </v-col>

      <input v-model="message" placeholder="Bestellungen">
      <v-spacer></v-spacer>
      <v-select
          :Abschnitt="Abschnitt"
          label="Beim Abschnitt"
          solo
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
    </v-row>

    <div class="text-right">
      <v-btn
          rounded
          color="red"
          dark
      >
        PDF Export
      </v-btn>
    </div>

  </div>

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

<<<<<<< HEAD
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





=======

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
    items: ['Heute', 'Gestern',],
    width: 2,
    radius: 10,
    padding: 8,
    lineCap: 'round',
    gradient: gradients[5],
    value: [0, 5, 8, 20, 5, 10, 3, 5, 1, 8, 2, 9, 5, 0,],
    gradientDirection: 'top',
    gradients,
    fill: false,
    type: 'trend',
    autoLineWidth: false,
  }),
}
</script>


>>>>>>> ffe0b7b... report generato, status and graph
