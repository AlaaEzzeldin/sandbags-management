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
        <h1 style="color: red">35</h1>
        <h3>Bestellungen</h3>
      </v-col>
      <v-col>
        <h1 style="color: red">94%</h1>
        <h3>Bestellungen bestatigt</h3>
      </v-col>
      <v-col>
        <h1 style="color: red">5</h1>
        <h3>Bestellungen/Uhr</h3>
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
            v-model="select"
            :items="selectOptions"
            item-text="state"
            label="Hauptabschnitt"
            filled
            outlined
            persistent-hint
            return-object
            single-line
        ></v-select>

        <v-select
            v-model="select"
            :items="selectOptions"
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
      <GChart type="ColumnChart" :data="chartData" :options="chartOptions"/>
    </div>
    <br>

    <v-row>
      <v-col cols="12" sm="3" offset="9">
        <v-btn
            align="right"
            style="text-transform: capitalize; font-weight: bolder;"
            block
            rounded
            color="red"
            dark
        >
          <button  @click="download">Export pdf</button>
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
  name: "App",
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
      chartData: [
        ["Abschnitt", "Bestellungen"],
        ["Unterabschnitt-1", 5],
        ["Unterabschnitt-2", 12],
        ["Unterabschnitt-3", 9],
        ["Unterabschnitt-4", 18],
        ["Unterabschnitt-5", 25],
      ],

      selectOptions: [
        { state: 'Beim Abschnitt',},
        { state: 'Beim Tag',},
      ],
      chartOptions: {
        hAxis: {
          title: "Abschnitt",
        },
        vAxis: {
          title: "Bestellungen",
          ticks: [0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30]
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
    dateRangeText () {
      return this.dates.join(' / ')
    },
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
                "OrderStatusReport_" +
                date.getFullYear() +
                ("0" + (date.getMonth() + 1)).slice(-2) +
                ("0" + date.getDate()).slice(-2) +
                ("0" + date.getHours()).slice(-2) +
                ".pdf";
            doc.save(filename)
            window.URL.revokeObjectURL(url);
            alert("Export File has downloaded!"); // or you know, something with better UX...
          })
          .catch(function(error) {
            console.error("oops, something went wrong!", error);
          });
    },
  }
};
</script>