<template>

  <v-row align="right" class="">
    <v-col cols="12">
      <h1 style="font-weight: bolder;">Neue Bestellung</h1>
      <input v-model="message" placeholder="was bestellen sie">

      <v-select
          :items="items"
          :menu-props="{ top: true, offsetY: true }"
          label="Sandsack"
      ></v-select>
      <v-select
          :items="items"
          :menu-props="{ top: true, offsetY: true }"
          label="50"
      ></v-select>

      <input v-model="message" placeholder="wohim liefern wir?">

      <v-select
          :items="items"
          :menu-props="{ top: true, offsetY: true }"
          label="Kircenplatz 1, 94032 Passau"
      ></v-select>
      <input v-model="message" placeholder="Andere Infomation">

      <v-select
          :items="items"
          :menu-props="{ top: true, offsetY: true }"
          label="Prioritat Sehr Wichtig....."
      ></v-select>


      <v-alert
          dense
          text
          type="success"
      >
        We need them as soon as possible <strong>MAX 50</strong>
      </v-alert>
      <v-btn
          rounded
          color="red"
          dark
      >
        Bestellen
      </v-btn>


    </v-col>
  </v-row>

</template>
