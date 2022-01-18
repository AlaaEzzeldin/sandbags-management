<template>
  <v-card elevation="0" class="pt-10">
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
              v-model="getOrder.priority"
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
              v-model="getOrder.deliveryAddress"
              :rules="[v => !!v || 'Die Adresse ist erforderlich']"
              outlined
          ></v-text-field>
        </v-col>
      </v-row>
      <!--      <v-row no-gutters v-if="getOrder.notesByUnterabschnitt">
              <v-col cols="12" sm="12">
                <h3 style="font-weight: bolder; color: black">Anmerkungen des Anforderers:</h3>
              </v-col>
              <v-col cols="12" sm="12">
                <v-textarea
                    class="mt-3"
                    v-model="getOrder.notesByUnterabschnitt"
                    outlined
                    :disabled="getCurrentUserRole!=='Unterabschnitt'"
                ></v-textarea>
              </v-col>
            </v-row>
            <v-row v-if="getOrder.notesByEinsatzabschnitt">
              <v-col cols="12" sm="12">
                <h3 style="font-weight: bolder; color: black">Notizen aus dem Einsatzabschnitt</h3>
              </v-col>
              <v-col cols="12" sm="12">
                <v-textarea
                    outlined
                    v-model="getOrder.notesByEinsatzabschnitt"
                ></v-textarea>
              </v-col>
            </v-row>
            <v-row v-if="getOrder.notesByHauptabschnitt">
              <v-col cols="12" sm="12">
                <h3 style="font-weight: bolder; color: black">Notizen aus dem hauptabschnitt</h3>
              </v-col>
              <v-col cols="12" sm="12">
                <v-textarea
                    outlined
                    v-model="getOrder.notesByHauptabschnitt"
                ></v-textarea>
              </v-col>
            </v-row>-->

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
               getOrder.quantity <= 0 ||
              !getOrder.deliveryAddress"
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
  created() {
    this.$store.dispatch("loadOrder", this.$route.params.orderId);
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
    goBack() {
      this.$router.go(-1)
    },
    submitUpdatedOrder() {
      let data = {
        "equipments": [
          {
            "id": 1,
            "quantity": this.getOrder.equipments[0].quantity
          }
        ],
        "order_id": this.getOrder.id,
        "priority": this.priorities.findIndex(x => x === this.selectedPriority)
      }
      this.$store.dispatch("editOrder", {data})
      this.gotToOrderDetails()
    },
    gotToOrderDetails(){
      const orderId = this.getOrder.id;
      this.$router.push({name: 'BestelldetailsPage', params: {orderId}})
    }

  }
}
</script>