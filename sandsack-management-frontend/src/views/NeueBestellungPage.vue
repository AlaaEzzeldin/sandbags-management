<template>
  <v-card elevation="0" class="pt-10">
    <v-card-title class="pt-10">
      <h1 style="font-weight: bolder; ">Neue Bestellung</h1>
    </v-card-title>

    <v-card-text class="pt-16 ">
      <v-row >
        <v-col cols="12">
          <v-text-field
              :value="getBranchName"
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
              v-model="selectedEquipmentIndex"
              :items="equipments"
              :item-text="equipments.name"
              filled
              outlined
              :menu-props="{ top: true, offsetY: true }"
              prepend-icon="mdi-home"
              label="was möchten Sie bestellen?"
              :rules="[v => !!v || 'Bitte geben Sie ein, was genau Sie bestellen möchten?']"
          ></v-select>
        </v-col>
        <v-col sm="6">
          <v-text-field
              v-model="newOrder.equipments[0].quantity"
              filled
              outlined
              :menu-props="{ top: true, offsetY: true }"
              :rules="[v => !!v || 'Die Menge ist erforderlich']"
              prepend-icon="mdi-pound"
              label="Anzahl"
          ></v-text-field>
        </v-col>
      </v-row>

      <v-row >
        <v-col cols="12">
          <v-text-field
              v-model="newOrder.address_to"
              filled
              outlined
              prepend-icon="mdi-map-marker"
              :rules="[v => !!v || 'Die Adresse ist erforderlich']"
              :menu-props="{ top: true, offsetY: true }"
              label="Wohim Liefern wir?"
          ></v-text-field>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="12">
          <v-select
              v-model="selectedPriorityIndex"
              :items="priorities"
              :rules="[v => !!v || 'Die Priorität ist erforderlich']"
              label="Priorität"
              prepend-icon="mdi-alert"
              required
              filled
              outlined
          ></v-select>
        </v-col>
      </v-row>

      <v-row >
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
              style="text-transform: capitalize; font-weight: bolder;"
              block
              rounded
              color="red"
              dark
              @click="createOrder"
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
    selectedEquipmentIndex:'',
    selectedPriorityIndex: '',
    loggedIn: '',
    priority: '',
    abschnitt: '',
    equipments: [
      {
        id: 0,
        name: "Sandsäcke",
        quantity: 0
      }
    ],
    priorities: [
      'Niedrig',
      'Mittle',
      'Hohe',
    ],

    newOrder:{
      address_to: "",
      comment: "",
      equipments: [
        {
          id: "1",
          quantity: "10"
        }
      ],
      priority: '1'
    }

  }),

  computed:{
    getCurrentUserRole(){
      return this.$store.getters.getCurrentUserRole
    },
    getBranchName() {
      return this.$store.getters.getCurrentUserRole
    },
  },

  methods: {
    createOrder() {
      console.log("new order", this.newOrder)
      // Priority and equipment needs to be fetched from the drop down select input
      //console.log("type",this.equipments.findIndex(x => x.name === this.selectedEquipmentIndex))
     // console.log("prirotity index", this.priorities.findIndex(x => x === this.selectedPriorityIndex))
     // this.newOrder.priority = this.priorities.findIndex(x => x === this.selectedPriorityIndex)
      this.$store.dispatch("createOrder", this.newOrder)
      this.$router.push({name: 'BestellungslistePage'})
    }
  },


}
</script>