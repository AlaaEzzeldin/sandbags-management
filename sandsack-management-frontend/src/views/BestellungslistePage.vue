<template>
  <div>
    <v-row no-gutters>
      <v-col sm="3" class="pt-13 justify-center align-center">
        <h1 style="font-weight: bolder;">Bestellungsliste</h1>
      </v-col>
      <v-spacer></v-spacer>
      <v-col sm="2" class="pt-15 justify-center align-center">
        <v-btn
            v-if="this.getLoggedInUserRole()===1 || this.getLoggedInUserRole()===2"
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            color="primary"
            dark
            block
        >
          exportieren
        </v-btn>
        <v-btn
            v-if="this.getLoggedInUserRole() === 4"
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            color="primary"
            dark
            block
        >
          Lieferschein drucken
        </v-btn>
      </v-col>
    </v-row>
    <v-row no-gutters>
      <v-col cols="12">
        <Bestelltabelle class="mt-10"></Bestelltabelle>
      </v-col>
    </v-row>
  </div>

</template>

<script>

import Bestelltabelle from "@/components/Bestelltabelle";

export default {
  name: 'BestellungslistePage',
  components: {
    Bestelltabelle
  },
  computed:{

  },
  methods:{

    // hard coding the users roles
    getLoggedInUserRole() {
      if (this.$route.params.userRole === '1') // Hauptabschintt
        return 1
      else if (this.$route.params.userRole === '2') // Einzatsabschnitt
        return 2
      else if (this.$route.params.userRole === '3') //Unterabschnitt
        return 3
      else if (this.$route.params.userRole === '4') // Mollhof
        return 4
    }
  }
}
</script>
