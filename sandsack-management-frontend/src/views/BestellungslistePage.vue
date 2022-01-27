<template>
  <div ref="content" :class="$vuetify.breakpoint.mdAndUp ? 'pt-10 pl-3 pr-3' : ''">
    <v-row no-gutters>
      <v-col sm="3">
        <h1 style="font-weight: bolder;">Bestellungsliste</h1>
      </v-col>
      <v-col
          sm="2"
          v-if="this.getCurrentUserRole==='Einsatzleiter' || this.getCurrentUserRole==='Hauptabschnitt' ||this.getCurrentUserRole==='Einsatzabschnitt'"
      >
        <v-btn
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            color="red"
            dark
            block
            @click="exportAllOrders"
        >
          <v-icon left>
            mdi-file-export
          </v-icon>
          Exportieren
        </v-btn>
      </v-col>
      <v-col
          sm="2"
          v-if="this.getCurrentUserRole === 'Mollnhof' && IsWaitingForDispatchOrders"
      >
        <v-btn
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            color="primary"
            dark
            block
            @click="lieferscheinDruecken"
        >
          <v-icon left>
            mdi-file-export
          </v-icon>
          Lieferschein drücken
        </v-btn>
      </v-col>
    </v-row>
    <v-row no-gutters>
      <v-col cols="12">
        <Bestelltabelle
            class="mt-5"
            :orders="getOrders"
        ></Bestelltabelle>
      </v-col>
    </v-row>

  </div>

</template>

<script>
import Bestelltabelle from "@/components/Bestelltabelle";
import moment from 'moment';

export default {
  name: 'BestellungslistePage',

  components: {
    Bestelltabelle
  },

  created() {
    this.$store.dispatch("loadOrders")
  },

  computed: {
    getCurrentUserRole() {
      return this.$store.getters.getCurrentUserRole
    },
    getOrders() {
      return this.$store.getters.getOrders
    },
    getPriorityByID(){
      return this.$store.getters.getPriorityByID
    },
    IsWaitingForDispatchOrders(){
      return this.getOrders.find(order=> order.status_name==='AKZEPTIERT')
    }
  },

  methods: {
    exportAllOrders() {
      let body = [['id', 'Zeit', 'Von', 'Priorität', 'Status', 'Anschrift']];
      for (let order of this.getOrders) {
        body.push(
            [
              order.id,
              order.create_date,
              order.address_from,
              this.getPriorityByID(order.priority_id).name,
              order.status_name,
              order.address_to
            ]
        );
      }
      this.printPDF(body, "Bestellungen");
    },

    lieferscheinDruecken: function () {
      let body = [['id', 'Von', 'Priorität', 'Anschrift', 'Menge']];
      for (let order of this.getOrders) {
        let dateCreated = moment(order.create_date);
        if (order.status_name === 'AKZEPTIERT' &&
            dateCreated.isSame(moment(), "day")
        ) {
          body.push(
              [
                order.id,
                order.address_from,
                this.getPriorityByID(order.priority_id).name,
                order.address_to,
                order.equipments[0].quantity
              ]
          );
        }
      }
      this.printPDF(body, "Lieferschein");
    },

    printPDF(body, name) {
      let pdfMake = require('pdfmake/build/pdfmake.js')
      if (pdfMake.vfs === undefined) {
        let pdfFonts = require('pdfmake/build/vfs_fonts.js')
        pdfMake.vfs = pdfFonts.pdfMake.vfs;
      }
      let docDefinition = {
        content: [
          {
            layout: 'lightHorizontalLines',
            table: {
              headerRows: 1,
              body: body
            }
          }
        ]
      }
      pdfMake.createPdf(docDefinition).download(name + '.pdf')
    }
  }
}
</script>
