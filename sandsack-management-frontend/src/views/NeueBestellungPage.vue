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
              v-model="newOrder.type"
              :items="types"
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
              v-model="newOrder.quantity"
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
              v-model="newOrder.deliveryAddress"
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
              v-model="newOrder.priority"
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
    loggedIn: '',
    priority: '',
    abschnitt: '',
    types: [
      'Sandsäcke',
    ],
    priorities: [
      'Niedrig',
      'Mittle',
      'Hohe',
    ],
    order: {
      type: "",
      quantity: "",
      priority: "",
      deliveryAddress: "",
      notesByUnterabschnitt: ""
    },
    newOrder:{
      address_to: "Johann-bergler Strasse 3",
      comment: "test comment 2",
      equipments: [
        {
          "id": 1,
          "quantity": 6
        }
      ],
      "priority": 2
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
      this.$store.dispatch("createOrder", this.newOrder)
      this.$router.push({name: 'BestellungslistePage'})
    }
  },


}
</script>