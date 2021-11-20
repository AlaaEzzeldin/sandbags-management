<template>
  <v-card class="elevation-0 rounded-lg" outlined>
    <v-data-table
        :headers="headers"
        :items="getRequests"
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
                  <v-icon> mdi-eye-outline</v-icon>
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
  name: 'RequestsTable',
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
    requests: [
      {
        'id': '#1',
        'created_at': '11.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Ost',
        'priority': 'hohe',
        'status': 'akzeptiert',
      },
      {
        'id': '#2',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
      },
      {
        'id': '#3',
        'created_at': '13.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Mitte',
        'priority': 'mittel',
        'status': 'geliefert',
      },
      {
        'id': '#4',
        'created_at': '14.11.2021 12:01',
        'from': 'EA 2.1 Nuemarkt- Nord',
        'priority': 'hohe',
        'status': 'Auf dem Weg',
      },
      {
        'id': '#5',
        'created_at': '15.11.2021 12:01',
        'from': 'EA 3.2 Universität-West',
        'priority': 'niedrige',
        'status': 'abgelehnt',
      },
      {
        'id': '#6',
        'created_at': '11.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Ost',
        'priority': 'hohe',
        'status': 'akzeptiert',
      },
      {
        'id': '#7',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
      },
      {
        'id': '#8',
        'created_at': '13.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Mitte',
        'priority': 'mittel',
        'status': 'geliefert',
      },
      {
        'id': '#9',
        'created_at': '14.11.2021 12:01',
        'from': 'EA 2.1 Nuemarkt- Nord',
        'priority': 'hohe',
        'status': 'Auf dem Weg',
      },
      {
        'id': '#10',
        'created_at': '15.11.2021 12:01',
        'from': 'EA 3.2 Universität-West',
        'priority': 'niedrige',
        'status': 'abgelehnt',
      },
      {
        'id': '#11',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
      },
      {
        'id': '#12',
        'created_at': '13.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- Mitte',
        'priority': 'mittel',
        'status': 'geliefert',
      },
      {
        'id': '#13',
        'created_at': '14.11.2021 12:01',
        'from': 'EA 2.1 Nuemarkt- Nord',
        'priority': 'hohe',
        'status': 'Auf dem Weg',
      },
      {
        'id': '#14',
        'created_at': '15.11.2021 12:01',
        'from': 'EA 3.2 Universität-West',
        'priority': 'niedrige',
        'status': 'abgelehnt',
      },
      {
        'id': '#15',
        'created_at': '12.11.2021 12:01',
        'from': 'EA 1.1 Altstadt- West',
        'priority': 'niedrige',
        'status': 'anstehend',
      },

    ],

  }),

  computed: {
    getRequests() {
      return this.requests
    }
  },
  methods: {
    getColor(status) {
      if (status === 'akzeptiert') return 'blue'
      if (status === 'geliefert') return 'green'
      else if (status === 'abgelehnt') return 'red'
      else if (status === 'Auf dem Weg') return 'yellow'
      else if (status === 'anstehend') return 'orange'
    },
    editItem() {
      //pass
    },
    inspect() {
      //pass
    },
  },
}
</script>
