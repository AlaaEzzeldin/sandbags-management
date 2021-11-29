<template>
  <v-card class="elevation-0 rounded-lg" outlined>
    <v-data-table
        :headers="headers"
        :items="getOrders"
        class="elevation-2 rounded-lg"
        :search="search"
        :options="options"
    >

      <!----------------------------------------SEARCH----------------------------------->
      <template v-slot:top>
        <v-toolbar
            flat
        >
          <v-spacer></v-spacer>
          <v-text-field
              v-model="search"
              append-icon="mdi-magnify"
              label="Suche nach Status, Abschnitt, Absender.."
              filled
              rounded
              dense
              single-line
              class="mt-6"
          ></v-text-field>
        </v-toolbar>
      </template>

      <!----------------------------------------STATUS CHIP----------------------------------->
      <template v-slot:item.status="{ item }">
        <v-chip
            small
            :color="getColor(item.status)"
            dark
            outlined
        >
          {{ item.status }}
        </v-chip>
      </template>

      <!----------------------------------------ACTIONS----------------------------------->
      <template v-slot:item.actions="{ item }">
        <v-row>
          <v-col cols="12">
            <v-tooltip top>
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                    v-bind="attrs"
                    v-on="on"
                    style="text-transform: capitalize; font-weight: bolder;"
                    @click="editItem(item)"
                    small
                    :disabled="item.status!=='anstehend'"
                    class="elevation-0"
                    color="primary"
                    rounded
                    icon
                >
                  <v-icon> mdi-pencil</v-icon>
                </v-btn>
              </template>
              <h4 class="font-weight-light"> Edit</h4>
            </v-tooltip>
            <v-tooltip top>
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                    @click="inspect(item)"
                    v-bind="attrs"
                    v-on="on"
                    small
                    class="elevation-0"
                    color="primary"
                    icon
                >
                  <v-icon> mdi-information-outline</v-icon>
                </v-btn>
              </template>
              <h4 class="font-weight-light">Inspect</h4>
            </v-tooltip>
          </v-col>
        </v-row>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>


export default {
  name: 'Bestelltabelle',
  components: {},
  data: () => ({
    search: '',
    headers: [
      {
        text: 'id',
        align: 'start',
        value: 'id',
      },
      {text: 'Zeit', value: 'created_at'},
      {text: 'Von', value: 'from'},
      {text: 'Priorität', value: 'priority', sortable: false},
      {text: 'Status', value: 'status', align: 'center'},
      {text: 'Aktionen', value: 'actions', sortable: false, align: 'center'},
    ],
    options: {
      itemsPerPage: 10,
    },
    orders: [
      {
        'id': '0',
        'created_at': '10.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Ost',
        'priority': 'hohe',
        'status': 'anstehend',
        'quantity': '12',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '1',
        'created_at': '11.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Ost',
        'priority': 'hohe',
        'status': 'akzeptiert',
        'quantity': '5',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '2',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
        'quantity': '54',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '3',
        'created_at': '13.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Mitte',
        'priority': 'mittel',
        'status': 'geliefert',
        'quantity': '7',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '4',
        'created_at': '14.11.2021 12:01',
        'from': 'EA 2.1 Nuemarkt- Nord',
        'priority': 'hohe',
        'status': 'Auf dem Weg',
        'quantity': '3',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'

      },
      {
        'id': '5',
        'created_at': '15.11.2021 12:01',
        'from': 'EA 3.2 Universität-West',
        'priority': 'niedrige',
        'status': 'abgelehnt',
        'quantity': '9',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '6',
        'created_at': '11.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Ost',
        'priority': 'hohe',
        'status': 'akzeptiert',
        'quantity': '12',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '7',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
        'quantity': '2',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '8',
        'created_at': '13.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Mitte',
        'priority': 'mittel',
        'status': 'geliefert',
        'quantity': '26',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '9',
        'created_at': '14.11.2021 12:01',
        'from': 'EA 2.1 Nuemarkt- Nord',
        'priority': 'hohe',
        'status': 'Auf dem Weg',
        'quantity': '9',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '10',
        'created_at': '15.11.2021 12:01',
        'from': 'EA 3.2 Universität-West',
        'priority': 'niedrige',
        'status': 'abgelehnt',
        'quantity': '17',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '11',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
        'quantity': '14',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '12',
        'created_at': '13.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Mitte',
        'priority': 'mittel',
        'status': 'geliefert',
        'quantity': '20',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '13',
        'created_at': '14.11.2021 12:01',
        'from': 'EA 2.1 Nuemarkt- Nord',
        'priority': 'hohe',
        'status': 'Auf dem Weg',
        'quantity': '17',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '14',
        'created_at': '15.11.2021 12:01',
        'from': 'EA 3.2 Universität-West',
        'priority': 'niedrige',
        'status': 'abgelehnt',
        'quantity': '12',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },
      {
        'id': '15',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
        'quantity': '12',
        'deliveryAddress': 'Nikolastraße 4 494032 Passau'
      },

    ],

  }),

  computed: {
    getOrders() {
      return this.orders
    }
  },
  methods: {
    getColor(status) {
      if (status === 'akzeptiert') return 'blue'
      if (status === 'geliefert') return 'green'
      else if (status === 'abgelehnt') return 'red'
      else if (status === 'Auf dem Weg') return 'orange'
      else if (status === 'anstehend') return 'grey'
    },
    editItem(Item) {
      const orderId = Item.id;
      this.$router.push({name: 'BestellBearbeitenPage', params: {orderId}})
    },
    inspect(Item) {
      const orderId = Item.id;
      this.$router.push({name: 'BestelldetailsPage', params: {orderId}})
    },
  },
}
</script>
