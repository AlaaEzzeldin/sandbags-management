<template>
  <v-card class="elevation-0 rounded-lg" outlined v-if="orders">
    <v-data-table
        :headers="headers"
        :items="orders"
        class="elevation-2 rounded-lg"
        :search="search"
        :options="options"
        :custom-filter="customFilter"
        @click:row="inspect"
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
      <template v-slot:item.status_name="{ item }">
        <v-chip
            small
            :color="getColor(item.status_name)"
            dark
            outlined
        >
          {{ item.status_name }}
        </v-chip>
      </template>

      <!----------------------------------------ACTIONS----------------------------------->
      <template v-slot:item.actions="{ item }">
        <v-row>
          <v-col cols="12">
            <v-btn
                v-for="action in getActions(item)"
                :key="action.name"
                style="text-transform: capitalize; font-weight: bolder;"
                @click.stop="doAction(action.actionType, item)"
                small
                class="elevation-0"
                :color="action.color"
                rounded
                icon
            >
              <v-icon>{{action.icon}}</v-icon>
            </v-btn>
          </v-col>
        </v-row>
      </template>
    </v-data-table>
    <ConfirmationDialog
        :action="action"
        :orderID="currentOrder ? currentOrder.id : null"
        :dialog="confirmationDialog"
        @close="confirmationDialog = false"
    />
  </v-card>
</template>

<script>
import ConfirmationDialog from "../components/ConfirmationDialog";
import {Mixin} from '../mixin/mixin.js'

export default {
  name: 'Bestelltabelle',
  props: ['orders'],
  mixins: [Mixin],
  components: {ConfirmationDialog},
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
    action: null,
    confirmationDialog: false,
    currentOrder: null,
  }),
  computed: {
    getCurrentUserRole(){
      return this.$store.getters.getCurrentUserRole
    },
    getPriorityByID(){
      return this.$store.getters.getPriorityByID
    },
    getPriorities(){
      return this.$store.getters.getPriorities
    }
  },
  methods: {
    customFilter (value, search, item) {
      console.log(value, search, item)
      if (this.getPriorities.find(p => p.name.toString().toLowerCase().indexOf(search.toString().toLowerCase())!== -1))
       return value != null &&
          search != null &&
          typeof value === 'string' &&
          this.getPriorityByID(item.priority_id).name.toLowerCase().indexOf(search.toString().toLowerCase())!== -1
      else return value != null &&
          search != null &&
          typeof value === 'string' &&
          value.toString().toLowerCase().indexOf(search.toString().toLowerCase()) !== -1
    },
    doAction(action, order) {
      if (action === 'edit') {
        const orderId = order.id;
        this.$router.push({name: 'BestellBearbeitenPage', params: {orderId}})
      }
      else if (['cancel', 'confirm_delivery', 'accept'].includes(action)) {
        this.action = action;
        this.currentOrder = order;
        this.confirmationDialog = true;
      }
    },
    inspect(Item) {
      const orderId = Item.id;
      this.$router.push({name: 'BestelldetailsPage', params: {orderId}})
    },
    getActions(Item) {
      const status = Item.status_name;
      if (this.getCurrentUserRole === 'Unterabschnitt') {
        if (status === 'ANSTEHEND') {
          return [
              {
                name: 'Bearbeiten',
                actionType: 'edit',
                icon: 'mdi-pencil',
                color: 'primary'
              },
              {
                name: 'Ablehnen',
                actionType: 'cancel', // todo: change
                icon: 'mdi-cancel',
                color: 'red'
              }
          ]
        }
        else if (status === 'AUF DEM WEG') {
          return [
            {
              name: 'Lieferung bestätigen',
              actionType: 'confirm_delivery',
              icon: 'mdi-check',
              color: 'green'
            },
          ]
        }
      }
      else if (this.getCurrentUserRole === 'Einsatzabschnitt') {
        if (status === 'ANSTEHEND') {
          return [
            {
              name: 'Bearbeiten',
              actionType: 'edit',
              icon: 'mdi-pencil',
              color: 'primary'
            },
            {
              name: 'Weiterleiten',
              actionType: 'accept',
              icon: 'mdi-check',
              color: 'green'
            },
            {
              name: 'Ablehnen',
              actionType: 'cancel',
              icon: 'mdi-cancel',
              color: 'red'
            }
          ]
        }
      }
    }
  },
}
</script>
