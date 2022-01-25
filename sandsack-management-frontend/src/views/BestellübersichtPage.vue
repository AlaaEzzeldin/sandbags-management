<template>
  <div ref="content">
    <v-row>
      <v-col class="pt-13 justify-center align-center" sm="3">
        <h1 style="font-weight: bolder;">Bestellübersicht</h1>
      </v-col>
      <v-spacer></v-spacer>
      <v-col class="pt-15 justify-center align-center" sm="5">
        <v-menu
            v-model="menu"
            :close-on-content-click="false"
            :nudge-right="50"
            min-width="auto"
            offset-y
            transition="scale-transition"
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
    <v-row no-gutters style="text-align: center; background-color: #F1F2F6; border-radius: 8px; padding: 10px">
      <v-col>
        <v-list-item
            v-for="(item, i) in getStatisticschart"
            :key="i"
        >
          <h1 style="text-align: center ;color: red">
            <div v-if="item.type ==='Hauptabschnitten'">
              {{ item.general_statistics.total_number_of_orders }}
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
              {{ item.general_statistics.total_number_of_accepted_orders }}%
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
              {{ item.general_statistics.average_processing_time }}
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
<!--      <v-col cols="6">
        <v-select
            v-if="this.getCurrentUserRole==='Einsatzleiter'"
            v-model="selectedHaupt"
            :items="getHauptdata"
            filled
            item-text="state"
            label="Hauptabschnitt"
            outlined
            persistent-hint
            return-object
            single-line
        ></v-select>
        <v-select
            v-if="this.getCurrentUserRole==='Einsatzleiter'"
            v-model="selectedEinz"
            :items="getEinsatzdata"
            filled
            item-text="state"
            label="Einsatzabschnitt"
            outlined
            persistent-hint
            return-object
            single-line
        ></v-select>
        <v-select
            v-if="this.getCurrentUserRole==='Hauptabschnitt'"
            v-model="selectedEinzforHaupt"
            :items="getEinsatzforHaupt"
            filled
            item-text="state"
            label="Einsatzabschnitt"
            outlined
            persistent-hint
            return-object
            single-line
        ></v-select>
      </v-col>-->
    </v-row>
    <div ref="content">
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
    </div>
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

  computed: {
    getStatisticschart() {
      return this.$store.getters.getStatisticschart
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

    getHauptdata() {
      var HauptselectOptions = []
      for (let i = 0; i < this.getStatisticschart.length; i++) {
        if (this.getStatisticschart[i].type === "Hauptabschnitten") {
          for (let j = 0; j < this.getStatisticschart[i].statistics_per_hauptabschnitt.length; j++) {
            HauptselectOptions.push(this.getStatisticschart[i].statistics_per_hauptabschnitt[j].name)
          }
        }
      }
      return HauptselectOptions;
    },

    getEinsatzdata() {
      var EinsatzselectOptions = []
      for (let i = 0; i < this.getHauptdata.length; i++) {
        if (this.$data.selectedHaupt === this.getHauptdata[i]) {
          for (let j = 0; j < this.getStatisticschart[0].statistics_per_hauptabschnitt[i].statistics_per_Einsatzabschnitt.length; j++) {
            EinsatzselectOptions.push(this.getStatisticschart[0].statistics_per_hauptabschnitt[i].statistics_per_Einsatzabschnitt[j].name)
          }
        }
      }
      return EinsatzselectOptions;
    },

    getEinsatzforHaupt() {
      var EinsatzforHaupt = []
      for (let i = 0; i < this.getStatisticschart.length; i++) {
        for (let j = 0; j < this.getStatisticschart[i].statistics_per_hauptabschnitt.length; j++) {
          for (let k = 0; k < this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt.length; k++) {
            EinsatzforHaupt.push(this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt[k].name)
          }
          //console.log(this.getStatisticschart[0].statistics_per_hauptabschnitt[i].statistics_per_Einsatzabschnitt[j].name)

        }
      }
      return EinsatzforHaupt;
    },

    getUnterdata() {
      var UnterselectOptions = []
      UnterselectOptions.push(["Abschnitt", "Bestellungen"], ['unterabschnitt', 10], ['unterabschnitt', 15]) //Delete the hardcoded values while integrating
      for (let i = 0; i < this.getHauptdata.length; i++) {
        for (let j = 0; j < this.getEinsatzdata.length; j++) {
          if (this.$data.selectedHaupt === this.getHauptdata[i] &&
              this.$data.selectedEinz === this.getEinsatzdata[j]) {
            UnterselectOptions.pop() // Delete this during integrtation
            UnterselectOptions.pop() // Delete this during integrtation
            for (let k = 0; k < this.getStatisticschart[0].statistics_per_hauptabschnitt[i].statistics_per_Einsatzabschnitt[j].statistics_per_unterabschnitt.length; k++) {
              UnterselectOptions.push([this.getStatisticschart[0].statistics_per_hauptabschnitt[i].statistics_per_Einsatzabschnitt[j].statistics_per_unterabschnitt[k].name,
                parseInt(this.getStatisticschart[0].statistics_per_hauptabschnitt[i].statistics_per_Einsatzabschnitt[j].statistics_per_unterabschnitt[k].total_number_of_orders)])
            }

          }
        }

      }
      return UnterselectOptions;
    },
    getUnterforhaupt() {
      var UnterforHaupt = []
      UnterforHaupt.push(["Abschnitt", "Bestellungen"], ['unterabschnitt', 10], ['unterabschnitt', 15]) //Delete the hardcoded values while integrating
      for (let i = 0; i < this.getStatisticschart.length; i++) {
        for (let j = 0; j < this.getStatisticschart[i].statistics_per_hauptabschnitt.length; j++) {
          for (let k = 0; k < this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt.length; k++) {
            if (this.$data.selectedEinzforHaupt === this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt[k].name) {
              UnterforHaupt.pop() // Delete this during integrtation
              UnterforHaupt.pop() // Delete this during integrtation
              for (let l = 0; l < this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt[k].statistics_per_unterabschnitt.length; l++) {
                UnterforHaupt.push([this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt[k].statistics_per_unterabschnitt[l].name,
                  parseInt(this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt[k].statistics_per_unterabschnitt[l].total_number_of_orders)])
              }
            }
          }
        }
      }
      return UnterforHaupt
    },

    getUnterforEinz() {
      var UnterforEinz = []
      UnterforEinz.push(["Abschnitt", "Bestellungen"])

      for (let i = 0; i < this.getStatisticschart.length; i++) {
        for (let j = 0; j < this.getStatisticschart[i].statistics_per_hauptabschnitt.length; j++) {
          for (let k = 0; k < this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt.length; k++) {
            for (let l = 0; l < this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt[k].statistics_per_unterabschnitt.length; l++) {
              UnterforEinz.push([this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt[k].statistics_per_unterabschnitt[l].name,
                parseInt(this.getStatisticschart[i].statistics_per_hauptabschnitt[j].statistics_per_Einsatzabschnitt[k].statistics_per_unterabschnitt[l].total_number_of_orders)])
            }
          }
        }
      }
      return UnterforEinz
    },

    getCurrentUserRole() {
      return this.$store.getters.getCurrentUserRole
    }

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
      console.log("dates",date_from, date_to )
      //router.push({name: 'BestellübersichtPage', query: {date1: date_from, date2: date_to}})
     // this.$store.dispatch("loadStatisticschart", {'date_from': date_from, 'date_to': date_to})
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
