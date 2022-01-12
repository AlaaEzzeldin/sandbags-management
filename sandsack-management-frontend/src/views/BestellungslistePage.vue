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
            @click="exportAllOrders"
        >
<<<<<<< HEAD
          Exportieren
=======
          <button  @click="download">Exportiern</button>
>>>>>>> 285d454f5d7d37370b314a62a759391e4489de7c
        </v-btn>
        <v-btn
            v-if="this.getLoggedInUserRole() === 4"
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            color="primary"
            dark
            block
            @click="lieferscheinDruecken"
        >
          Lieferschein drücken
        </v-btn>
      </v-col>
    </v-row>
    <v-row no-gutters>
      <v-col cols="12">
        <Bestelltabelle
            class="mt-10"
            :orders="getOrders"
        ></Bestelltabelle>
      </v-col>
    </v-row>

  </div>

</template>

<script>
<<<<<<< HEAD
=======
import jsPDF from "jspdf";
import domtoimage from "dom-to-image";
>>>>>>> 285d454f5d7d37370b314a62a759391e4489de7c
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
    getOrders() {
      return this.$store.getters.getOrders
    }
  },

  methods: {
    exportAllOrders () {
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
      this.printPDF(body, "Bestellungen");
    },

    lieferscheinDruecken: function () {
      let body =  [[ 'id', 'Von', 'Priorität', 'Anschrift', 'Menge' ]];
      for (let order of this.getOrders) {
        let dateCreated = moment(order.created_at, 'DD.MM.yyyy HH:mm');
        if (order.status === 'akzeptiert' &&
            dateCreated.isSame(moment(), "day")
        ) {
          body.push(
              [
                order.id,
                order.from,
                order.priority,
                order.deliveryAddress,
                order.quantity
              ]
          );
        }
      }
      this.printPDF(body, "Lieferschein");
    },

    printPDF(body, name) {
      let pdfMake = require('pdfmake/build/pdfmake.js')
      if (pdfMake.vfs === undefined){
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
      pdfMake.createPdf(docDefinition).download(name+'.pdf')
    },

    download() {
      /** WITH CSS */
      domtoimage
          .toPng(this.$refs.content)
          .then(function(dataUrl) {
            var img = new Image();
            img.src = dataUrl;
            const doc = new jsPDF({
              orientation: "portrait",
              unit: "pt",
              format: [900, 1500]
            });
            doc.addImage(img, "JPEG", 100, 100);
            const date = new Date();
            const url = window.URL.createObjectURL;
            const filename =
                "Exportiern_" +
                date.getFullYear() +
                ("0" + (date.getMonth() + 1)).slice(-2) +
                ("0" + date.getDate()).slice(-2) +
                ("0" + date.getHours()).slice(-2) +
                ".pdf";
            doc.save(filename)
            window.URL.revokeObjectURL(url);
            alert("Exportiern Downloaded!"); // or you know, something with better UX...
          })
          .catch(function(error) {
            console.error("oops, something went wrong!", error);
          });
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
