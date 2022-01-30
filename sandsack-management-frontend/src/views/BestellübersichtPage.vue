<template>
  <div :class="$vuetify.breakpoint.mdAndUp ? 'pt-10 pl-3 pr-3' : ''">
      <v-row v-if="$vuetify.breakpoint.mdAndUp">
        <v-col class="justify-center align-center">
          <h1 style="font-weight: bolder;">Bestellübersicht</h1>
        </v-col>
        <v-col sm="4" class="justify-center align-center">
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
                scrollable
                @input="menu = false"
                locale="de"
                no-title
            ></v-date-picker>
          </v-menu>
        </v-col>
      </v-row>

      <v-row v-if="!$vuetify.breakpoint.mdAndUp">
        <v-col class="justify-center align-center">
          <h1 style="font-weight: bolder;">Bestellübersicht</h1>
        </v-col>
        <v-spacer></v-spacer>
      </v-row>
      <v-row no-gutters v-if="!$vuetify.breakpoint.mdAndUp" class="mt-5">
        <v-col sm="4" class="justify-center align-center">
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
             v-if="getGeneralStatisticsForCurrentRole && $vuetify.breakpoint.mdAndUp">
        <v-col cols="4">
          <h1 style="text-align: center; color: red">
            {{ getGeneralStatisticsForCurrentRole.general_statistics.total_number_of_orders }}</h1>
          <h3 style="text-align: center">Bestellungen</h3>
        </v-col>

        <v-col cols="4">
          <h1 style="text-align: center;color: red">
            {{ getGeneralStatisticsForCurrentRole.general_statistics.total_number_of_accepted_orders }}</h1>
          <h3 style="text-align: center;">Bestellungen geliefert</h3>
        </v-col>

        <v-col cols="4">
          <h1 style="text-align: center;color: red">
            {{ getGeneralStatisticsForCurrentRole.general_statistics.average_processing_time }}</h1>
          <h3 style="text-align: center">Durchsch. Lieferzeit</h3>
        </v-col>

      </v-row>
      <div
          v-if="getGeneralStatisticsForCurrentRole && !$vuetify.breakpoint.mdAndUp"
          style="background-color: #F1F2F6; border-radius: 8px"
      >
        <v-row>
          <v-col cols="3" style="text-align: center">
            <h3 style="color: red">
              {{ getGeneralStatisticsForCurrentRole.general_statistics.total_number_of_orders }}</h3>
          </v-col>
          <v-col>
            <h3>Bestellungen</h3>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="3" style="text-align: center">
            <h3 style="color: red">
              {{ getGeneralStatisticsForCurrentRole.general_statistics.total_number_of_accepted_orders }}</h3>
          </v-col>
          <v-col>
            <h3>Bestellungen geliefert</h3>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="3" style="text-align: center">
            <h3 style="color: red">
              {{ getGeneralStatisticsForCurrentRole.general_statistics.average_processing_time }}</h3>
          </v-col>
          <v-col>
            <h3>Durchschnittliche Lieferungszeit</h3>
          </v-col>
        </v-row>
      </div>

      <v-row class="mt-5" no-gutters>
        <v-col cols="6">
          <h2>
            Bestellungen
          </h2>
        </v-col>
      </v-row>
      <div id="content">
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
    <VueHtml2pdf
        :show-layout="false"
        :float-layout="true"
        :enable-download="true"
        :preview-modal="false"
        :paginate-elements-by-height="1400"
        :filename="getFileName()"
        :pdf-quality="2"
        :manual-pagination="false"
        pdf-format="a4"
        pdf-orientation="landscape"
        :pdf-content-width="$vuetify.breakpoint.xs ? '600px' : '1110px'"
        ref="html2Pdf"
    >
      <section slot="pdf-content" >
        <!-- PDF Content Here -->
        <div :class="$vuetify.breakpoint.mdAndUp ? 'pt-10 pl-3 pr-3' : ''">

          <v-row>
            <v-col class="justify-center align-center">
              <h1 style="font-weight: bolder;">Bestellübersicht</h1>
            </v-col>
            <v-col sm="4" class="justify-center align-center">
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
                    scrollable
                    @input="menu = false"
                    locale="de"
                    no-title
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
              <h3 style="text-align: center;">Bestellungen geliefert</h3>
            </v-col>

            <v-col cols="4">
              <h1 style="text-align: center;color: red">
                {{ getGeneralStatisticsForCurrentRole.general_statistics.average_processing_time }}</h1>
              <h3 style="text-align: center">Durchsch. Lieferzeit</h3>
            </v-col>

          </v-row>
          <v-row class="mt-5" no-gutters>
            <v-col cols="12">
              <h2>
                Bestellungen
              </h2>
            </v-col>

            <v-col cols="12">
                <GChart
                    id="chart"
                    v-if="getStatisticsForCurrentRole"
                    :data="getStatisticsForCurrentRole"
                    type="ColumnChart"
                    :options="chartOptions"
                />
            </v-col>
            <v-col cols="12">
              <v-alert
                  v-if="!getStatisticsForCurrentRole"
                  text
                  type="error"
                  icon="mdi-cloud-alert"
              >
                Ups, ein Fehler! Für das von Ihnen gewählte Zeitintervall sind keine Bestellungen vorhanden! Bitte ändern Sie die
                Daten.
              </v-alert>
            </v-col>
          </v-row>
        </div>

      </section>
    </VueHtml2pdf>

    <v-spacer></v-spacer>

    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-btn
            id="pdf"
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            @click="generateReport"
            block
            color="red"
            outlined
        >
          <v-icon left>
            mdi-file-export
          </v-icon>
          PDF exportieren
        </v-btn>
      </v-col>
    </v-row>


  </div>
</template>

<script>

import {GChart} from "vue-google-charts";
import moment from 'moment';
import VueHtml2pdf from 'vue-html2pdf'


export default {

  name: 'BestellübersichtPage',

  components: {
    GChart,
    VueHtml2pdf
  },

  data() {
    return {
      dates: [new Date(new Date().setDate(new Date().getDate() - 7)).toISOString().slice(0, 10),
        new Date(new Date().setDate(new Date().getDate())).toISOString().slice(0, 10)],
      menu: false,
      chartOptions: {
        chart: {
          title: "Bestellungen",
        },
        height: 400,
        vAxis: {
          title: "Bestellungen"
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
        return moment(this.dates[0]).format('DD-MM-YYYY') + '/' + moment(this.dates[1]).format('DD-MM-YYYY')
      } else {
        this.loadStatsByDates(this.dates[0], this.dates[0]);
        return moment(this.dates[0]).format('DD-MM-YYYY') + '/' + moment(this.dates[0]).format('DD-MM-YYYY')
      }

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
        result.push(["Abschnitten", "Bestellungen", "Bestellungen geliefert"])
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

    generateReport () {
      this.$refs.html2Pdf.generatePdf()
    },
    getFileName(){
      const date = new Date();
      return "Statistics-Report_" +
          ("0" + date.getDate()).slice(-2) + "." +
          ("0" + (date.getMonth() + 1)).slice(-2) + "." +
          date.getFullYear()
    },
    loadStatsByDates(date_from, date_to) {
      //  console.log("dates", date_from, date_to)
      let data = {
        "end_date": date_to,
        "start_date": date_from
      }
      this.$store.dispatch("loadStatistics", data)
    }
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
