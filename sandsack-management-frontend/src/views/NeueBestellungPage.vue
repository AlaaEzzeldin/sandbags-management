<template>
  <v-card elevation="0" class="pt-10">
    <v-card-title class="pt-10">
      <h1 style="font-weight: bolder; ">Neue Bestellung</h1>
    </v-card-title>

    <v-card-text class="pt-16 ">
      <v-row>
        <v-col cols="12">
          <v-text-field
              :value="getLoggedInUserName"
              readonly
              prepend-icon="mdi-account"
              filled
              outlined
              disabled
          ></v-text-field>
        </v-col>
      </v-row>

      <v-row>
        <v-col sm="6">
          <v-select
              v-model="chosenEquipmentType"
              :items="getEquipment.map(a => a.name)"
              filled
              outlined
              :menu-props="{ top: true, offsetY: true }"
              prepend-icon="mdi-home"
              label="Was möchten Sie bestellen?"
              :rules="[v => !!v || 'Bitte geben Sie ein, was genau Sie bestellen möchten?']"
          ></v-select>
        </v-col>
        <v-col sm="6">
          <v-text-field
              v-model="orderQuantity"
              filled
              outlined
              :menu-props="{ top: true, offsetY: true }"
              :rules="[v => (!!v && v <= getCurrentEquipment.quantity && v > 0)|| 'Die Menge ist nicht correct']"
              prepend-icon="mdi-pound"
              label="Anzahl"
              :hint="'Die Restmenge ist '+getCurrentEquipment.quantity.toString() + ' ' + getCurrentEquipment.measure"
          ></v-text-field>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="12">
          <v-text-field
              v-model="newOrder.address_to"
              filled
              outlined
              prepend-icon="mdi-map-marker"
              :rules="[v => !!v || 'Die Adresse ist erforderlich']"
              :menu-props="{ top: true, offsetY: true }"
              label="wohin liefern wir?"
          ></v-text-field>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="12">
          <v-select
              v-model="selectedPriority"
              :items="getPriorities.map(x => x.name)"
              :rules="[v => !!v || 'Die Priorität ist erforderlich']"
              label="Priorität"
              prepend-icon="mdi-alert"
              required
              filled
              outlined
          ></v-select>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="12">
          <v-textarea
              v-model="newOrder.comment"
              outlined
              filled
              prepend-icon="mdi-message-bulleted"
              name="input-7-4"
              label="Irgendwelche zusätzlichen Anmerkungen?"

          ></v-textarea>
        </v-col>
      </v-row>

    </v-card-text>

    <v-card-actions>
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;color: white"
              block
              rounded
              color="red"
              @click="createOrder"
              :disabled="!(
                  chosenEquipmentType &&
                  selectedPriority &&
                  newOrder.address_to &&
                  orderQuantity > 0 &&
                  orderQuantity <= getCurrentEquipment.quantity
              )"
          >
            Bestellen
          </v-btn>
        </v-col>
      </v-row>

    </v-card-actions>

  </v-card>
</template>

<script>

export default {
  name: 'BestelldetailsPage',

  data: () => ({
    selectedPriority: '',
    chosenEquipmentType: '',
    orderQuantity: '',
    newOrder:
        {
          "address_to": "",
          "comment": "",
          "equipments": [
            {
              "id": 0,
              "measure": "",
              "name": "",
              "quantity": 0
            }
          ],
          "priority": 0
        }
  }),
  created() {
    this.$store.dispatch("loadEquipment");
    this.$store.dispatch("loadPriorities");
  },

  computed: {
    getCurrentUserRole() {
      return this.$store.getters.getCurrentUserRole
    },
    getLoggedInUserName() {
      return this.$store.getters.getLoggedInUser.name
    },
    getEquipment() {
      return this.$store.getters.getEquipment
    },
    getCurrentEquipment() {
      if (this.chosenEquipmentType) {
        return this.$store.getters.getEquipmentByType(this.chosenEquipmentType);
      }
      return {
        "id": 0,
        "measure": "",
        "quantity": 0,
        "name": ""
      };
    },
    getPriorities() {
      return this.$store.getters.getPriorities
    },
  },

  methods: {
    createOrder() {
      this.newOrder.equipments =
          [
            {
              "id": this.getEquipment.find(a => a.name === this.chosenEquipmentType).id,
              "quantity": parseInt(this.orderQuantity)
            }
          ]
      this.newOrder.priority = this.getPriorities.find(item => item.name === this.selectedPriority).id

      console.log("new order", this.newOrder)
      this.$store.dispatch("createOrder", this.newOrder)
      this.$router.push({name: 'BestellungslistePage'})
    }
  },


}
</script>
