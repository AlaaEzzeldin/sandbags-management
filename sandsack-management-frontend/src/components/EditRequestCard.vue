<template>
  <v-card elevation="0" class="pt-10" v-if="getOrder && getPriorities && getEquipment">
    <v-card-title class="pt-10">
      <v-btn icon @click="goBack">
        <v-icon large color="black" class="pr-5">mdi-keyboard-backspace</v-icon>
      </v-btn>
      <h1 style="font-weight: bolder; ">Bestellung bearbeiten # {{ getOrder.id }} </h1>
      <v-chip
          class="ml-5"
          :color="getColor(getOrder.status_name)" outlined
          dark>
        {{ getOrder.status_name }}
      </v-chip>
    </v-card-title>

    <v-card-text class="pt-16 ">
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Von:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <v-text-field
              disabled
              :value="getOrder.name"
              outlined
          ></v-text-field>
        </v-col>
      </v-row>
            <v-row no-gutters>
              <v-col cols="12" sm="2">
                <h3 style="font-weight: bolder; color: black">Typ:</h3>
              </v-col>
              <v-col cols="12" sm="3">
                <v-select
                    disabled
                    :value="getOrder.equipments[0].name"
                    :items="getEquipment.map(a => a.name)"
                    outlined
                ></v-select>
              </v-col>
            </v-row>
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Anzahl:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <v-text-field
              v-model="getOrder.equipments[0].quantity"
              outlined
              :rules="[v => (!!v && v <= getCurrentEquipment.quantity && v > 0)|| 'Die Menge ist nicht correct']"
              :hint="'Die Restmenge ist '+getCurrentEquipment.quantity.toString() + ' ' + getCurrentEquipment.measure"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Priorität:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <v-select
              v-model="selectedPriority"
              :items="getPriorities.map(x => x.name)"
              outlined
          ></v-select>
        </v-col>
      </v-row>
      <v-row no-gutters>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Lieferadresse:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <v-text-field
              v-model="getOrder.address_to"
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
              (getOrder.quantity > this.getCurrentEquipment.quantity) ||
               getOrder.quantity <= 0 "
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
  data: () => ({
    selectedPriority: ''
  }),
  created() {
    this.$store.dispatch("loadOrder", this.$route.params.orderId);
    this.$store.dispatch("loadEquipment");
    this.$store.dispatch("loadPriorities");
  },

  computed: {
    getOrder() {
      return this.$store.getters.getOrder
    },
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
      if (this.getOrder.equipments[0].name) {
        return this.$store.getters.getEquipmentByType(this.getOrder.equipments[0].name);
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
      this.selectedPriority= this.getPriorities.find(item => item.level === this.getOrder.priority_id).name
    },
    goBack() {
      this.$router.go(-1)
    },
    submitUpdatedOrder() {
      console.log("order ", this.getOrder)
      this.initializePriority()
      let data = {
        "equipments": [
          {
            "id": this.getOrder.equipments[0].id,
            "quantity": parseInt(this.getOrder.equipments[0].quantity)
          }
        ],
        "order_id": this.getOrder.id,
        "priority": this.getPriorities.find(item => item.name === this.selectedPriority).id
      }
      console.log("edited order", data)
      this.$store.dispatch("editOrder", data)
      this.gotToOrderDetails()
    },
    gotToOrderDetails(){
      const orderId = this.getOrder.id;
      this.$router.push({name: 'BestelldetailsPage', params: {orderId}})
    }

  }
}
</script>