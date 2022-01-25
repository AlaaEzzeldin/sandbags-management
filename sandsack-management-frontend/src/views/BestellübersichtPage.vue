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
                v-model="dateSelected"
                filled
                label="Tage"
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
    <v-row no-gutters style="text-align: center; background-color: #F1F2F6; border-radius: 8px; padding: 10px" v-if="getStatistics">
      <v-col>
        <h1 style="color: red">{{ getGeneralStatisticsForCurrentRole().general_statistics.total_number_of_orders }}</h1>
        <h3 style="text-align: left">Bestellungen</h3>
      </v-col>

      <v-col>
        <h1 style="color: red">{{getGeneralStatisticsForCurrentRole().general_statistics.total_number_of_accepted_orders}}</h1>
        <h3 style="text-align: left">Bestellungen bestätigt</h3>
      </v-col>

      <v-col>
        <h1 style="color: red">{{ getGeneralStatisticsForCurrentRole().general_statistics.average_processing_time }}</h1>
        <h3 style="text-align: left">Bestellungen/Uhr</h3>
      </v-col>

    </v-row>
    <v-row>
      <v-col cols="6">
        <h2>
          Bestellungen
        </h2>
      </v-col>
    </v-row>
    <!--    <div ref="content">
          <GChart
              v-if="this.getCurrentUserRole==='Einsatzleiter'"
              :data="getUnterdata"
              :options="chartOptions"
              type="ColumnChart"/>
        </div>
        <div ref="content">
          <GChart
              v-if="this.getCurrentUserRole==='Hauptabschnitt'"
              :data="getUnterforhaupt"
              :options="chartOptions"
              type="ColumnChart"/>
        </div>
        <div ref="content">
          <GChart
              v-if="this.getCurrentUserRole==='Einsatzabschnitt'"
              :data="getUnterforEinz"
              :options="chartOptions"
              type="ColumnChart"/>
        </div>-->
    <br>
    <v-spacer></v-spacer>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3" class="mt-10">
        <v-btn
            id="pdf"
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            @click="download"
            block
            color="red"
            outlined
        >
          pdf exportieren
        </v-btn>
      </v-col>
    </v-row>


  </div>
</template>

<script>
import jsPDF from "jspdf";
import domtoimage from "dom-to-image";
//import {GChart} from "vue-google-charts";

export default {

  name: 'BestellübersichtPage',

  components: {
    // GChart,
  },

  data() {
    return {

      dates: [new Date(new Date().setDate(new Date().getDate() - 7)).toISOString().slice(0, 10),
        new Date(new Date().setDate(new Date().getDate())).toISOString().slice(0, 10)],

      modal: false,
      menu: false,
      selectedHaupt: '',
      selectedEinz: '',
      selectedEinzforHaupt: '',

      chartOptions: {
        hAxis: {
          title: "Abschnitt",
        },
        vAxis: {
          title: "Bestellungen",
          ticks: [0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50]
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
  created(){
    this.loadStatsByDates(this.dates[0], this.dates[1])
  },

  computed: {
    getStatistics() {
      return this.$store.getters.getStatistics
    },
    dateSelected() {
      if (this.dates[1] < this.dates[0]) {
        window.alert('Bitte wählen Sie einen gültigen Zeitraum aus!')
        return ('Gültiges Datum auswählen')
      } else if ((typeof (this.dates[1]) !== "undefined")) {
        this.loadStatsByDates(this.dates[0], this.dates[1]);
      } else {
        this.loadStatsByDates(this.dates[0], this.dates[0]);
      }
      return this.dates.join(' / ')
    },

    getCurrentUserRole() {
      return this.$store.getters.getCurrentUserRole
    },
  },
  methods: {
    getGeneralStatisticsForCurrentRole() {
      if(this.getStatistics){
        console.log(this.getStatistics)
        if (this.getCurrentUserRole === 'Einsatzabschnitt')
          return this.getStatistics.find(data => data.type === 'Unterabschnitten')
        else if (this.getCurrentUserRole === 'Hauptabschnitt')
          return this.getStatistics.find(data => data.type === 'Einsatzabschnitten')
        else if (this.getCurrentUserRole === 'Einsatzleiter')
          return this.getStatistics.find(data => data.type === 'Hauptabschnitten')
        else return null
      }
      else return null
    },
    download() {
      /** To Block the Button */
      document.getElementById('pdf').style.display = 'none';
      domtoimage
          .toPng(this.$refs.content)
          .then(function (dataUrl) {
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
                "Statistics-Report_" +
                date.getFullYear() +
                ("0" + (date.getMonth() + 1)).slice(-2) +
                ("0" + date.getDate()).slice(-2) +
                ("0" + date.getHours()).slice(-2) +
                ".pdf";
            doc.save(filename)
            window.URL.revokeObjectURL(url);
            /** To Un-Block the Button after download */
            if (!alert('Die Datei wird heruntergeladen!')) {
              document.getElementById('pdf').style.display = 'block';
            }
          })
          .catch(function (error) {
            console.error("oops, Fehler!", error);
          });
    },
    loadStatsByDates(date_from, date_to) {
      console.log("dates", date_from, date_to)
      //router.push({name: 'BestellübersichtPage', query: {date1: date_from, date2: date_to}})
      let data = {
        "end_date": date_to,
        "start_date": date_from
      }
      this.$store.dispatch("loadStatistics", data)
    },
  }

};
</script>

<style>


ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
  color: #D7201F;
}
</style>
