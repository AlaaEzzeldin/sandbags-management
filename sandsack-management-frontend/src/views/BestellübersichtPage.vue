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
            label="Beim Abschnitt"
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
        ["EA 1.1 Altstadt - Ost", 5],
        ["EA 1.2 Altstadt - Mitte", 12],
        ["EA 1.3 Altstadt - West", 9],
        ["EA 2.1 Neumarkt - Nord", 25],
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
          ticks: [0,5,10,15,20,25,30]
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
  color: #42b983;
}
</style>
