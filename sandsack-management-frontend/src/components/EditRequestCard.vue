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
              :value="getOrder.from"
              outlined
          ></v-text-field>
        </v-col>
      </v-row>
            <v-row no-gutters>
              <v-col cols="12" sm="2">
                <h3 style="font-weight: bolder; color: black">Typ:</h3>
              </v-col>
              <v-col cols="12" sm="3">
                <v-text-field
                    disabled
                    :value="getOrder.type"
                    outlined
                ></v-text-field>
              </v-col>
            </v-row>
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
          <v-select
              v-model="getOrder.priority"
              :items="priorities"
              :label="getOrder.priority"
              outlined
          ></v-select>
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
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row no-gutters v-if="getOrder.notesByUnterabschnitt">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Anmerkungen des Anforderers:</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              class="mt-3"
              v-model="getOrder.notesByUnterabschnitt"
              outlined
              :disabled="getCurrentUserRole!==3"
          ></v-textarea>
        </v-col>
      </v-row>

      <v-row v-if="getOrder.notesByEinsatzabschnitt">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Notizen aus dem Einsatzabschnitt</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              outlined
              v-model="getOrder.notesByEinsatzabschnitt"
          ></v-textarea>
        </v-col>
      </v-row>
      <v-row v-if="getOrder.notesByHauptabschnitt">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Notizen aus dem hauptabschnitt</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              outlined
              v-model="getOrder.notesByHauptabschnitt"
          ></v-textarea>
        </v-col>
      </v-row>

    </v-card-text>

    <!------------------------------------------------- Actions ------------------------------------------->
    <v-card-actions v-if="['Einsatzabschnitt','Hauptabschnitt','Einsatzleiter', 'Unterabschnitt'].includes(getCurrentUserRole)">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              outlined
              @click="submitUpdatedOrder"

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
              @click="gotToOrderDetails"
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
    priorities: [
      'Niedrig',
      'Mittle',
      'Hohe',
    ],

  }),
  created() {
    this.$store.dispatch("loadOrder", this.$route.params.orderId)
  },
  computed: {
    getOrder() {
      return this.$store.getters.getOrder
    },
    getCurrentUserRole(){
      return this.$store.getters.getCurrentUserRole
    },
  },
  methods: {
    getColor(status) {
      if (status === 'akzeptiert') return 'blue'
      if (status === 'geliefert') return 'green'
      else if (status === 'abgelehnt') return 'red'
      else if (status === 'storniert') return 'red'
      else if (status === 'Auf dem Weg') return 'orange'
      else if (status === 'anstehend') return 'grey'
      else if (status === 'weitergeleitet') return 'black'
    },
    goBack() {
      this.$router.go(-1)
    },
    submitUpdatedOrder(){
      let data={
        "quantity": this.getOrder.quantity,
        "deliveryAddress" : this.getOrder.deliveryAddress,
        "priority": this.getOrder.priority,
        "notesByUnterabschnitt": this.getOrder.notesByUnterabschnitt,
        "notesByEinsatzabschnitt": this.getOrder.notesByEinsatzabschnitt,
        "notesByHauptabschnitt": this.getOrder.notesByHauptabschnitt
      }
      let id= this.getOrder.id
      this.$store.dispatch("updateOrder",  {id, data}  )
      this.gotToOrderDetails()
    },
    gotToOrderDetails(){
      const orderId = this.getOrder.id;
      this.$router.push({name: 'BestelldetailsPage', params: {orderId}})
    }

  }
}
</script>