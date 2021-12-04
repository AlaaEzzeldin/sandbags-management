<template>
  <v-card elevation="0" class="pt-10">
    <v-card-title class="pt-10">
      <v-btn icon @click="goBack">
        <v-icon large color="black" class="pr-5">mdi-keyboard-backspace</v-icon>
      </v-btn>
      <h1 style="font-weight: bolder; ">Bestellung # {{ getOrder.id }}</h1>
      <v-chip
          class="ml-5"
          :color="getColor(getOrder.status)" outlined
          dark>
        {{ getOrder.status }}
      </v-chip>
    </v-card-title>

    <v-card-text class="pt-16 ">

      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Von:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <h3 style="font-weight: bolder; color: black">{{getOrder.from}}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Typ:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <h3 style="font-weight: bolder; color: black">sandsäcke</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Anzahl:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <h3 style="font-weight: bolder; color: black">{{getOrder.quantity}}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Priorität:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <h3 style="font-weight: bolder; color: black">{{getOrder.priority}}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Lieferadresse:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <h3 style="font-weight: bolder; color: black">{{getOrder.deliveryAddress}}</h3>
        </v-col>
      </v-row>
      <v-row v-if="getOrder.notesByUnterabschnitt">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Anmerkungen des Anforderers:</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              readonly
              :value="getOrder.notesByUnterabschnitt"
              outlined
          ></v-textarea>
        </v-col>
      </v-row>
      <v-row v-if="getOrder.notesByEinsatzORderHaupt">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Notizen aus dem hauptabschnitt/einsatzabschnitt</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              readonly
              outlined
              :value="getOrder.notesByEinsatzORderHaupt"
          ></v-textarea>
        </v-col>
      </v-row>
      <!--      <v-row>
              <v-col cols="12" sm="12">
                <h3 style="font-weight: bolder; color: black">Besteellverlauf</h3>
              </v-col>
            </v-row>-->

    </v-card-text>

    <!------------------------------------------------- Unterabschnitt ------------------------------------------->
    <v-card-actions v-if="getLoggedInUserRole() === 3">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="orange"
              dark
              block
              outlined
              :disabled="getOrder.status!=='anstehend'"
              @click="editItem"
          >
            Bestellung bearbeiten
          </v-btn>
        </v-col>

        <v-col cols="12" sm="6" offset="3" v-if="getOrder.status==='Auf dem Weg'">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="green"
              dark
              block
              outlined
          >
            Lieferung bestätigen
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              :disabled="getOrder.status!=='anstehend'"
          >
            Bestellung stornieren
          </v-btn>
        </v-col>
      </v-row>

    </v-card-actions>

    <!---------------------------------- Einsatzabschnitt & Hauptabschnitt -------------------------------->
    <v-card-actions v-if="getLoggedInUserRole() === 1 || this.getLoggedInUserRole() === 2">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="green"
              dark
              block
              outlined
              :disabled="getOrder.status!=='anstehend'"
          >
            Bestellung direkt annehmen
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              outlined
              :disabled="getOrder.status!=='anstehend'"
              @click="editItem"
          >
            Bestellung bearbeiten und dann akzeptieren
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              :disabled="getOrder.status!=='anstehend'"
          >
            Bestellung ablehnen
          </v-btn>
        </v-col>
      </v-row>
    </v-card-actions>

    <!------------------------------------------------- Mollhof ------------------------------------------->
    <v-card-actions v-if="getLoggedInUserRole() === 4">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="green"
              dark
              block
              outlined
              :disabled="getOrder.status!=='akzeptiert'"
          >
            Bestellung abgesendet
          </v-btn>
        </v-col>
      </v-row>
    </v-card-actions>
  </v-card>
</template>

<script>
export default {
  name: 'BestelldetailsCard',
  data: () => ({
    orders: [
      {
        'id': '0',
        'created_at': '10.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Ost',
        'priority': 'hohe',
        'status': 'anstehend',
        'quantity':'12',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau',
        'notesByUnterabschnitt':'please process ASAP'
      },
      {
        'id': '1',
        'created_at': '11.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Ost',
        'priority': 'hohe',
        'status': 'akzeptiert',
        'quantity':'5',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau',
        'notesByUnterabschnitt':'please process ASAP'

      },
      {
        'id': '2',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
        'quantity':'54',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau',
        'notesByUnterabschnitt':'please process ASAP',

      },
      {
        'id': '3',
        'created_at': '13.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Mitte',
        'priority': 'mittel',
        'status': 'geliefert',
        'quantity':'7',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau',
        'notesByEinsatzORderHaupt':'please process ASAP'
      },
      {
        'id': '4',
        'created_at': '14.11.2021 12:01',
        'from': 'EA 2.1 Nuemarkt- Nord',
        'priority': 'hohe',
        'status': 'Auf dem Weg',
        'quantity':'3',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau',
        'notesByEinsatzORderHaupt':'please process ASAP'

      },
      {
        'id': '5',
        'created_at': '15.11.2021 12:01',
        'from': 'EA 3.2 Universität-West',
        'priority': 'niedrige',
        'status': 'abgelehnt',
        'quantity':'9',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau',
        'notesByEinsatzORderHaupt':'There is no need for it, we don not have enough bags'
      },
      {
        'id': '6',
        'created_at': '11.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Ost',
        'priority': 'hohe',
        'status': 'akzeptiert',
        'quantity':'12',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '7',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
        'quantity':'2',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '8',
        'created_at': '13.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Mitte',
        'priority': 'mittel',
        'status': 'geliefert',
        'quantity':'26',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '9',
        'created_at': '14.11.2021 12:01',
        'from': 'EA 2.1 Nuemarkt- Nord',
        'priority': 'hohe',
        'status': 'Auf dem Weg',
        'quantity':'9',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '10',
        'created_at': '15.11.2021 12:01',
        'from': 'EA 3.2 Universität-West',
        'priority': 'niedrige',
        'status': 'abgelehnt',
        'quantity':'17',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '11',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
        'quantity':'14',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '12',
        'created_at': '13.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Mitte',
        'priority': 'mittel',
        'status': 'geliefert',
        'quantity':'20',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '13',
        'created_at': '14.11.2021 12:01',
        'from': 'EA 2.1 Nuemarkt- Nord',
        'priority': 'hohe',
        'status': 'Auf dem Weg',
        'quantity':'17',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '14',
        'created_at': '15.11.2021 12:01',
        'from': 'EA 3.2 Universität-West',
        'priority': 'niedrige',
        'status': 'abgelehnt',
        'quantity':'12',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '15',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
        'quantity':'12',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },

    ],
  }),

  computed: {
    getOrder() {
      return this.orders[this.$route.params.orderId]
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
    editItem() {
      const orderId = this.getOrder.id;
      this.$router.push({name: 'BestellBearbeitenPage', params: {orderId}})
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