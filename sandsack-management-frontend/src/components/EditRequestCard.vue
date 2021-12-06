<template>
  <v-card elevation="0" class="pt-10">
    <v-card-title class="pt-10">
      <v-btn icon @click="goBack">
        <v-icon large color="black" class="pr-5">mdi-keyboard-backspace</v-icon>
      </v-btn>
      <h1 style="font-weight: bolder; ">Bestellung bearbeiten # {{ getOrder.id }} </h1>
      <v-chip
          class="ml-5"
          :color="getColor(getOrder.status)" outlined
          dark>
        {{ getOrder.status }}
      </v-chip>
    </v-card-title>

    <v-card-text class="pt-16 ">

      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Von:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <v-text-field
              disabled
              v-model="getOrder.from"
              outlined
          ></v-text-field>
        </v-col>
      </v-row>
      <!--      <v-row no-gutters>
              <v-col cols="12" sm="2">
                <h3 style="font-weight: bolder; color: black">Typ:</h3>
              </v-col>
              <v-col cols="12" sm="3">
                <h3 style="font-weight: bolder; color: black">sandsäcke</h3>
              </v-col>
            </v-row>-->
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Anzahl:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <v-text-field
              v-model="getOrder.quantity"
              outlined
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Priorität:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <v-text-field
              v-model="getOrder.priority"
              outlined
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Lieferadresse:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <v-text-field
              v-model="getOrder.deliveryAddress"
              outlined
              disabled
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row no-gutters v-if="getOrder.notesByUnterabschnitt">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Anmerkungen des Anforderers:</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              v-model="getOrder.notesByUnterabschnitt"
              outlined
              :disabled="getLoggedInUserRole()!==3"
          ></v-textarea>
        </v-col>
      </v-row>
      <v-row no-gutters v-if="getOrder.notesByEinsatzORderHaupt">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Notizen aus dem hauptabschnitt/einsatzabschnitt</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              readonly
              outlined
              v-model="getOrder.notesByEinsatzORderHaupt"
              :disabled="getLoggedInUserRole()!==1 || getLoggedInUserRole()!==2"

          ></v-textarea>
        </v-col>
      </v-row>

    </v-card-text>

    <!------------------------------------------------- Actions ------------------------------------------->
    <v-card-actions v-if="getLoggedInUserRole()===1   || getLoggedInUserRole()===2 ||getLoggedInUserRole()===3 ">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              outlined
          >
            speichern
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
          >
            Abbrechen
          </v-btn>
        </v-col>
      </v-row>

    </v-card-actions>

  </v-card>
</template>

<script>
export default {
  name: 'EditRequestCard',
  data: () => ({

  }),
  created() {
    this.$store.dispatch("loadOrder", this.$route.params.orderId)
  },
  computed: {
    getOrder() {
      return this.$store.getters.getOrder
    }
  },
  methods: {
    getColor(status) {
      if (status === 'akzeptiert') return 'blue'
      if (status === 'geliefert') return 'green'
      else if (status === 'abgelehnt') return 'red'
      else if (status === 'Auf dem Weg') return 'orange'
      else if (status === 'anstehend') return 'grey'
    },
    goBack() {
      this.$router.go(-1)
    },

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