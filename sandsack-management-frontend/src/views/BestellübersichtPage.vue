<template>
  <div>
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

    <v-row no-gutters style="text-align: center">
      <v-col>
        <h1>35</h1>
        <h3>Bestellungen</h3>
      </v-col>
      <v-col>
        <h1>94%</h1>
        <h3>Bestellungen bestatigt</h3>
      </v-col>
      <v-col>
        <h1>5</h1>
        <h3>Bestellungen/Uhr</h3>
      </v-col>
    </v-row>
      <v-container fluid>
        <v-row align="center">
          <v-col cols="6">
            <v-subheader>
              Bestellungen
            </v-subheader>
          </v-col>

          <v-col cols="6">
            <v-select
                v-model="select"
                :items="amount"
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
        <GChart type="LineChart" :data="chartData" :options="chartOptions" />
      </v-container>
    <v-spacer></v-spacer>
    <v-card-actions>
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
            Pdf Export
          </v-btn>
        </v-col>
      </v-row>
    </v-card-actions>
  </div>
</template>




<script>
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
        ["period", "Status"],
        ["5-10", 5],
        ["10-20", 12],
        ["20-25", 9],
        ["25-30", 6],
        ["30-50", 6],
      ],

      amount: [
        { state: 'Sample1',},
        { state: 'Sample2',},
        { state: 'sample3',},
        { state: 'sample4',},
        { state: 'sample5',},
      ],


      chartOptions: {
        hAxis: {
          title: "Request",
          titleTextStyle: { color: "#FF0000" },
        },
        vAxis: {
          title: "Time",
          titleTextStyle: { color: "#FF0000" }
        },
        chart: {
          title: "order status",
        },
      },
    };
  },
  computed: {
    dateRangeText () {
      return this.dates.join(' - ')
    },
  },
};
</script>

