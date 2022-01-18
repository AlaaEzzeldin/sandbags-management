<template>
  <v-card class="elevation-0 rounded-lg" outlined>
    <v-data-table
        :headers="headers"
        :items="orders"
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

      <!---------------------------------------- TIME ----------------------------------->
      <template v-slot:item.priority_id="{ item }">
        {{getPriorityByID(item.priority_id).name}}
      </template>

      <!---------------------------------------- TIME ----------------------------------->
      <template v-slot:item.create_date="{ item }">
        {{format_time(item.create_date)}}
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
            <v-tooltip top v-if="['Einsatzabschnitt','Hauptabschnitt','Einsatzleiter', 'Unterabschnitt'].includes(getCurrentUserRole)">
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                    v-bind="attrs"
                    v-on="on"
                    style="text-transform: capitalize; font-weight: bolder;"
                    @click="editItem(item)"
                    small
                    :disabled="(item.status!=='ANSTEHEND' && ['Einsatzabschnitt', 'Unterabschnitt'].includes(getCurrentUserRole))
                    || (item.status!=='WEITERGELEITET BEI EINSATZABSCHNITT' && getCurrentUserRole=== 'Hauptabschnitt')
                    || (item.status!=='WEITERGELEITET BEI HAUPTABSCHNITT' && getCurrentUserRole=== 'Einsatzleiter') "

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

import {Mixin} from '../mixin/mixin.js'

export default {
  name: 'Bestelltabelle',
  props: ['orders'],
  mixins: [Mixin],
  components: {},
  data: () => ({
    search: '',
    headers: [
      {
        text: 'id',
        align: 'start',
        value: 'id',
      },
      {text: 'Zeit', value: 'create_date'},
      {text: 'Von', value: 'name'},
      {text: 'Priorität', value: 'priority_id', sortable: false},
      {text: 'Status', value: 'status_name', align: 'center'},
      {text: 'Aktionen', value: 'actions', sortable: false, align: 'center'},
    ],
    options: {
      itemsPerPage: 10,
    },
  }),
  computed: {
    getCurrentUserRole(){
      return this.$store.getters.getCurrentUserRole
    },
    getPriorityByID(){
      return this.$store.getters.getPriorityByID
    }
  },
  methods: {

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
