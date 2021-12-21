<template>
  <div ref="content">
    <v-row>
      <v-col sm="3" class="pt-13 justify-center align-center">
        <h1 style="font-weight: bolder;">Bestellübersicht</h1>
      </v-col>
      <v-spacer></v-spacer>
      <v-col sm="4" class="pt-15 justify-center align-center">
        <v-menu
            v-model="menu"
            :close-on-content-click="false"
            :nudge-right="50"
            transition="scale-transition"
            offset-y
            min-width="auto"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-text-field
                v-model="dateRangeText"
                label="Tage"
                filled
                prepend-icon="mdi-calendar"
                readonly
                v-bind="attrs"
                v-on="on"
            ></v-text-field>
          </template>
          <v-date-picker
              v-model="dates"
              range
              @input="menu = false"
          ></v-date-picker>
        </v-menu>
      </v-col>
    </v-row>
    <v-row no-gutters style="text-align: center; background-color: #F1F2F6; border-radius: 8px; padding: 10px">
      <v-col>
        <v-list-item
            v-for="(item, i) in getStatisticschart"
            :key="i"
        >
          <h1 style="text-align: center ;color: red">
            <div v-if="item.type ==='Hauptabschnitten'">
              {{item.general_statistics.total_number_of_orders}}
            </div>
          </h1>
        </v-list-item>
        <h3 style="text-align: left">Bestellungen</h3>
      </v-col>
      <v-col>
        <v-list-item
            v-for="(item, i) in getStatisticschart"
            :key="i"
        >
          <h1 style="color: red">
            <div v-if="item.type ==='Hauptabschnitten'">
              {{item.general_statistics.total_number_of_accepted_orders}}%
            </div>
          </h1>
        </v-list-item>
        <h3 style="text-align: left">Bestellungen bestatigt</h3>
      </v-col>
      <v-col>
        <v-list-item
            v-for="(item, i) in getStatisticschart"
            :key="i"
        >
          <h1 style="color: red">
            <div v-if="item.type ==='Hauptabschnitten'">
              {{item.general_statistics.average_processing_time}}
            </div>
          </h1>
        </v-list-item>
        <h3 style="text-align: left">Bestellungen/Uhr</h3>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="6">
        <h2>
          Bestellungen
        </h2>
      </v-col>
      <v-col cols="6">
        <v-select
            v-if="this.getLoggedInUserRole()===1 || this.getLoggedInUserRole()===2"
            v-model="select"
            :items="getHauptdata"
            item-text="state"
            label="Hauptabschnitt"
            filled
            outlined
            persistent-hint
            return-object
            single-line
        ></v-select>
        <v-select
            v-if="this.getLoggedInUserRole()===1 || this.getLoggedInUserRole()===2"
            v-model="select"
            :items="getEinsatzdata"
            item-text="state"
            label="Einsatzabschnitt"
            filled
            outlined
            persistent-hint
            return-object
            single-line
        ></v-select>
      </v-col>
    </v-row>
    <div ref="content">
      <GChart
          type="ColumnChart"
          :data="getUnterdata"
          :options="chartOptions"/>
    </div>
    <br>
    <v-row>
      <v-col cols="12" sm="3" offset="9">
        <v-btn
            v-if="this.getLoggedInUserRole()===1 || this.getLoggedInUserRole()===2"
            style="text-transform: capitalize; font-weight: bolder;"
            block
            rounded
            color="red"
            dark
        >

          <button v-on:click="download">Export pdf</button>
        </v-btn>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import jsPDF from "jspdf";
import domtoimage from "dom-to-image";
import { GChart } from "vue-google-charts";
export default {

  name: 'BestellübersichtPage',

  components: {
    GChart,
  },

  data() {
    return {

      dates:
          [(new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10),
            (new Date(Date.now() - (new Date()).getTimezoneOffset() * 60000)).toISOString().substr(0, 10)],
      modal: false,
      menu: false,

      // Array will be automatically processed with visualization.arrayToDataTable function
      /*chartData: [
        ["Abschnitt", "Bestellungen"],
        ["Unterabschnitt-1", 5],
        ["Unterabschnitt-2", 12],
        ["Unterabschnitt-3", 9],
        ["Unterabschnitt-4", 18],
        ["Unterabschnitt-5", 25],
      ],*/

      chartOptions: {
        hAxis: {
          title: "Abschnitt",
        },
        vAxis: {
          title: "Bestellungen",
          ticks: [0,10,20,30,40,50,60,70,80,90,100]
        },
        chart: {
          title: "order status",
        },
        colors: ['#D7201F'],
        lineWidth: 4,
        smoothLine: true,
      },
    };
  },
  created() {
    this.$store.dispatch("loadStatisticschart")
  },


  computed: {
    getStatisticschart() {
      return this.$store.getters.getStatisticschart
    },
    dateRangeText () {
      return this.dates.join(' / ')
    },
    getHauptdata(){
      var HauptselectOptions = []
      for (let i=0; i<this.getStatisticschart.length; i++){
        if(this.getStatisticschart[i].type == "Hauptabschnitten") {
          for (let j = 0; j < this.getStatisticschart[i].statistics_per_hauptabschnitt.length; j++) {
            HauptselectOptions.push(this.getStatisticschart[i].statistics_per_hauptabschnitt[j].name)
          }
        }
      }
      return HauptselectOptions;
    },

    getEinsatzdata(){
      var EinsatzselectOptions = []
      for(let i=0; i<this.getStatisticschart.length; i++) {
        if(this.getStatisticschart[i].type == "Einsatzabschnitten") {
          for(let j=0; j<this.getStatisticschart[i].statistics_per_Einsatzabschnitt.length; j++){
            EinsatzselectOptions.push(this.getStatisticschart[i].statistics_per_Einsatzabschnitt[j].name)
          }
        }
      }
      return EinsatzselectOptions;
    },

    getUnterdata(){
      var UnterselectOptions = []
      UnterselectOptions.push(["Abschnitt","Bestellungen"])
      for(let i=0; i<this.getStatisticschart.length; i++) {
        if(this.getStatisticschart[i].type == "unterabschnitten") {
          for(let j=0; j<this.getStatisticschart[i].statistics_per_unterabschnitt.length; j++){
            UnterselectOptions.push([this.getStatisticschart[i].statistics_per_unterabschnitt[j].name,
              parseInt(this.getStatisticschart[i].statistics_per_unterabschnitt[j].total_number_of_orders)])


          }
        }
      }
      return UnterselectOptions;
    }
  },

  props: {
    msg: String
  },
  methods: {

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
                "StatusReport_" +
                date.getFullYear() +
                ("0" + (date.getMonth() + 1)).slice(-2) +
                ("0" + date.getDate()).slice(-2) +
                ("0" + date.getHours()).slice(-2) +
                ".pdf";
            doc.save(filename)
            window.URL.revokeObjectURL(url);
            alert("Export File has downloaded!");
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

};
</script>

&lt;!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40pt 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #D7201F;
}
</style>