<template>
  <v-card elevation="0" class="pt-10" v-if="editedOrder && getPriorities && getEquipment">
    <v-card-title class="pt-10">
      <v-btn icon @click="goBack">
        <v-icon large color="black" class="pr-5">mdi-keyboard-backspace</v-icon>
      </v-btn>
      <h1 style="font-weight: bolder; ">Bestellung bearbeiten # {{ editedOrder.id }} </h1>
      <v-chip
          class="ml-5"
          :color="getColor(editedOrder.status_name)" outlined
          dark>
        {{ editedOrder.status_name }}
      </v-chip>
    </v-card-title>

    <v-card-text class="pt-16 ">
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h2 style="font-weight: bolder; color: black">Von:</h2>
        </v-col>
        <v-col cols="12" sm="8">
          <v-text-field
              disabled
              :value="editedOrder.name"
              outlined
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h2 style="font-weight: bolder; color: black">Typ:</h2>
        </v-col>
        <v-col cols="12" sm="8">
          <v-select
              disabled
              :value="editedOrder.equipments[0].name"
              :items="getEquipment.map(a => a.name)"
              outlined
          ></v-select>
        </v-col>
      </v-row>
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h2 style="font-weight: bolder; color: black">Anzahl:</h2>
        </v-col>
        <v-col cols="12" sm="8">
          <v-text-field
              v-model="editedOrder.equipments[0].quantity"
              outlined
              :rules="[v => (!!v && v <= getCurrentEquipment.quantity && v > 0)|| 'Die Menge ist nicht correct']"
              :hint="'Die Restmenge ist '+getCurrentEquipment.quantity.toString() + ' ' + getCurrentEquipment.measure"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h2 style="font-weight: bolder; color: black">Priorität:</h2>
        </v-col>
        <v-col cols="12" sm="8">
          <v-select
              v-model="selectedPriority"
              :items="getPriorities.map(x => x.name)"
              outlined
          ></v-select>
        </v-col>
      </v-row>
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h2 style="font-weight: bolder; color: black">Lieferadresse:</h2>
        </v-col>
        <v-col cols="12" sm="8">
          <v-text-field
              v-model="editedOrder.address_to"
              outlined
              disabled
          ></v-text-field>
        </v-col>
      </v-row>
    </v-card-text>

    <!------------------------------------------------- Actions ------------------------------------------->
    <v-card-actions
        v-if="['Einsatzabschnitt','Hauptabschnitt','Einsatzleiter', 'Unterabschnitt'].includes(getCurrentUserRole)">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              block
              outlined
              @click="submitUpdatedOrder"
              :disabled="
              (editedOrder.equipments[0].quantity > this.getCurrentEquipment.quantity) ||
               editedOrder.equipments[0].quantity <= 0 "
          >
            Speichern
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              @click="gotToOrderDetails"
          >
            Abbrechen
          </v-btn>
        </v-col>
      </v-row>

    </v-card-actions>

  </v-card>
</template>

<script>
import {Mixin} from '../mixin/mixin.js'

export default {
  name: 'EditRequestCard',
  mixins: [Mixin],
  props:['editedOrder'],
  data: () => ({
    selectedPriority: ''
  }),

  mounted() {
    this.initializePriority()
  },

  computed: {

    getCurrentUserRole() {
      return this.$store.getters.getCurrentUserRole
    },
    getPriorities() {
      return this.$store.getters.getPriorities
    },
    getEquipment() {
      return this.$store.getters.getEquipment
    },
    getCurrentEquipment() {
      if (this.editedOrder.equipments[0].name) {
        return this.$store.getters.getEquipmentByType(this.editedOrder.equipments[0].name);
      }
      return {
        "id": 0,
        "measure": "",
        "quantity": 0,
        "name": ""
      };
    },
  },
  methods: {
    initializePriority(){
      this.selectedPriority= this.getPriorities.find(item => item.level === this.editedOrder.priority_id).name
    },
    goBack() {
      this.$router.go(-1)
    },
    submitUpdatedOrder() {
      let data = {
        "equipments": [
          {
            "id": this.editedOrder.equipments[0].id,
            "quantity": parseInt(this.editedOrder.equipments[0].quantity)
          }
        ],
        "order_id": this.editedOrder.id,
        "priority": this.getPriorities.find(item => item.name === this.selectedPriority).id
      }
      this.$store.dispatch("editOrder", data)
      this.gotToOrderDetails()
    },
    gotToOrderDetails(){
      const orderId = this.editedOrder.id;
      this.$router.push({name: 'BestelldetailsPage', params: {orderId}})
    }

  }
}
</script>
