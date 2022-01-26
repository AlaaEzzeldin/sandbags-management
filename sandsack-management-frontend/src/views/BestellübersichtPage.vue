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
    <v-row no-gutters style="text-align: center; background-color: #F1F2F6; border-radius: 8px; padding: 10px"
           v-if="getGeneralStatisticsForCurrentRole">
      <v-col cols="4">
        <h1 style="text-align: center; color: red">
          {{ getGeneralStatisticsForCurrentRole.general_statistics.total_number_of_orders }}</h1>
        <h3 style="text-align: center">Bestellungen</h3>
      </v-col>

      <v-col cols="4">
        <h1 style="text-align: center;color: red">
          {{ getGeneralStatisticsForCurrentRole.general_statistics.total_number_of_accepted_orders }}</h1>
        <h3 style="text-align: center;">Bestellungen bestätigt</h3>
      </v-col>

      <v-col cols="4">
        <h1 style="text-align: center;color: red">
          {{ getGeneralStatisticsForCurrentRole.general_statistics.average_processing_time }}</h1>
        <h3 style="text-align: center">Durchsch. Bearbeitungszeit</h3>
      </v-col>

    </v-row>
    <v-row>
      <v-col cols="6">
        <h2>
          Bestellungen
        </h2>
      </v-col>
    </v-row>
    <div id="content" class="mt-10 pt-10">
      <GChart
          v-if="getStatisticsForCurrentRole"
          :data="getStatisticsForCurrentRole"
          type="ColumnChart"
          :options="chartOptions"
      />
    </div>

    <v-alert
        v-if="!getStatisticsForCurrentRole"
        text
        type="error"
        icon="mdi-cloud-alert"
    >
      Ups, ein Fehler! Für das von Ihnen gewählte Zeitintervall sind keine Bestellungen vorhanden! Bitte ändern Sie die
      Daten.
    </v-alert>

    <v-spacer></v-spacer>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-btn
            id="pdf"
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            @click="download"
            block
            color="red"
            outlined
        >
          PDF exportieren
        </v-btn>
      </v-col>
    </v-row>


  </div>
</template>

<script>
import jsPDF from "jspdf";
import domtoimage from "dom-to-image";
import {GChart} from "vue-google-charts";

export default {

  name: 'BestellübersichtPage',

  components: {
    GChart,
  },

  data() {
    return {

      dates: [new Date(new Date().setDate(new Date().getDate() - 7)).toISOString().slice(0, 10),
        new Date(new Date().setDate(new Date().getDate())).toISOString().slice(0, 10)],

      menu: false,
      chartOptions: {
        chart: {
          title: "Company Performance",
          subtitle: "Sales, Expenses, and Profit: 2014-2017"
        },
        height: 400,
        vAxis: {
          title: "Bestellungen",
        },
      }
    };
  },

  created() {
    this.loadStatsByDates(this.dates[0], this.dates[1])
  },

  computed: {
    getStatistics() {
      return this.$store.getters.getStatistics.statistics
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
    getGeneralStatisticsForCurrentRole() {
      if (this.getStatistics) {
        if (this.getCurrentUserRole === 'Einsatzabschnitt')
          return this.getStatistics.find(data => data.type === "Unterabschnitten")
        else if (this.getCurrentUserRole === 'Hauptabschnitt')
          return this.getStatistics.find(data => data.type === 'Einsatzabschnitten')
        else if (this.getCurrentUserRole === 'Einsatzleiter')
          return this.getStatistics.find(data => data.type === 'Hauptabschnitten')
        else return null
      } else return null
    },
    getStatisticsForCurrentRole() {
      if (this.getStatistics) {
        let statistics = []
        let result = []
        result.push(["Abschnitten", "Bestellungen", "Bestellungen bestätigt"])
        //console.log('all statitics', this.getStatistics)
        if (this.getCurrentUserRole === 'Einsatzabschnitt')
          statistics = this.getStatistics.find(data => data.type === "Unterabschnitten").statistics_per_unterabschnitt
        else if (this.getCurrentUserRole === 'Hauptabschnitt')
          statistics = this.getStatistics.find(data => data.type === 'Einsatzabschnitten').statistics_per_Einsatzabschnitt
        else if (this.getCurrentUserRole === 'Einsatzleiter')
          statistics = this.getStatistics.find(data => data.type === 'Hauptabschnitten').statistics_per_hauptabschnitt
        if (statistics) {
          statistics.forEach(d => result.push([d.name, parseInt(d.total_number_of_orders), parseInt(d.total_number_of_accepted_orders)]))
          // console.log("visualized data", result)
          return result
        } else return null
      } else return null
    },
  },
  methods: {
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
      //  console.log("dates", date_from, date_to)
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
