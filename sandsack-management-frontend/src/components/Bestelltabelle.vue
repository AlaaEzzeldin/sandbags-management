<template>
  <v-card class="elevation-0 rounded-lg" outlined v-if="orders && getPriorities">
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

      <!---------------------------------------- PRIORITY ----------------------------------->
      <template v-slot:item.priority_id="{ item }">
        <v-icon
            :color="getIcon(item.priority_id).color"
        >
          {{getIcon(item.priority_id).icon}}
        </v-icon>
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
            <v-tooltip top
                   v-for="action in getActions(item)"
                   :key="action.name">
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                    style="text-transform: capitalize; font-weight: bolder;"
                    @click.stop="doAction(action.actionType, item)"
                    small
                    class="elevation-0"
                    :color="action.color"
                    rounded
                    icon
                    v-bind="attrs"
                    v-on="on"
                >
                  <v-icon>{{action.icon}}</v-icon>
                </v-btn>
              </template>
              <span>{{action.name}}</span>
            </v-tooltip>
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
      {text: 'Menge', value: 'equipments[0].quantity', align: 'right'},
      {text: 'Ausrüstung', value: 'equipments[0].name'},
      {text: 'Status', value: 'status_name', align: 'center'},
      {text: 'Aktionen', value: 'actions', sortable: false, align: 'right'},
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
      else if (action === 'inspect') {
        this.inspect(order);
      }
      else if (['cancel', 'confirm_delivery', 'accept', 'dispatch'].includes(action)) {
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
      let actions = []
      if (this.getCurrentUserRole === 'Unterabschnitt') {
        if (status === 'ANSTEHEND') {
          actions = [
              {
                name: 'Bearbeiten',
                actionType: 'edit',
                icon: 'mdi-pencil',
                color: 'primary'
              },
              {
                name: 'Ablehnen',
                actionType: 'cancel', // todo: change
                icon: 'mdi-close',
                color: 'red'
              }
          ]
        }
        else if (status === 'AUF DEM WEG') {
          actions = [
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
          actions = [
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
              icon: 'mdi-close',
              color: 'red'
            }
          ]
        }
      }
      else if (this.getCurrentUserRole === 'Hauptabschnitt') {
        if (status === 'WEITERGELEITET BEI EINSATZABSCHNITT') {
          actions = [
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
              icon: 'mdi-close',
              color: 'red'
            }
          ]
        }
      }
      else if (this.getCurrentUserRole === 'Einsatzleiter') {
        if (status === 'WEITERGELEITET BEI HAUPTABSCHNITT') {
          actions = [
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
              icon: 'mdi-close',
              color: 'red'
            }
          ]
        }
      }
      else if (this.getCurrentUserRole === 'Mollnhof') {
        if (status === 'AKZEPTIERT') {
          actions = [
            {
              name: 'Absenden',
              actionType: 'dispatch',
              icon: 'mdi-cube-send',
              color: 'green'
            }
          ]
        }
      }
      actions.push({
        name: 'Inspektieren',
        actionType: 'inspect',
        icon: 'mdi-chevron-right',
        color: 'gray'
      });
      return actions;
    }
  },
}
</script>
