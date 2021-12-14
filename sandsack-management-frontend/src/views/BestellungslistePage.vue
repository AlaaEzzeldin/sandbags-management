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
          <button  @click="download">Exportiern</button>
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
import jsPDF from "jspdf";
import domtoimage from "dom-to-image";
import Bestelltabelle from "@/components/Bestelltabelle";

export default {
  name: 'BestellungslistePage',
  components: {
    Bestelltabelle
  },
  computed:{

  },
  methods:{

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
