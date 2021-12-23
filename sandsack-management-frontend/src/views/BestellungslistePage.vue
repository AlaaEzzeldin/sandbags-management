<template>
  <div ref="content">
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
            color="red"
            dark
            block
        >
          <button  @click="pdfgen">Exportiern</button>
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

  created() {
    this.$store.dispatch("loadOrders")
  },
  computed: {
    getOrders() {
      return this.$store.getters.getOrders
    }
  },
  methods:{
    pdfgen: function () {
      var pdfMake = require('pdfmake/build/pdfmake.js')
      if (pdfMake.vfs === undefined){
        var pdfFonts = require('pdfmake/build/vfs_fonts.js')
        pdfMake.vfs = pdfFonts.pdfMake.vfs;
      }
      let body =  [[ 'id', 'Zeit', 'Von', 'Priorität', 'Status', 'Anschrift' ]];
      for (let order of this.getOrders) {
        body.push(
            [
              order.id,
              order.created_at,
              order.from,
              order.priority,
              order.status,
              order.deliveryAddress
            ]
        );
      }
      var docDefinition = {
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
      pdfMake.createPdf(docDefinition).download('Bestellungen.pdf')
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
