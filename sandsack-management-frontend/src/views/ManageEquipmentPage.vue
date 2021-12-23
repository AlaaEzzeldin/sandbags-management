<template>
  <div>
    <v-row no-gutters>
      <v-col class="pt-13 justify-center align-center">
        <h1 style="font-weight: bolder;">Restliche Ausrüstung Bearbeiten</h1>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-select
            :items="getEquipment.map(item => item.type)"
            @change="setCurrentType"
            label="Ausrüstungtyp"
            outlined
        ></v-select>
      </v-col>
      <v-col>
        <v-text-field
            style="max-width: 50px;"
            :value="getCurrentEquipment.amount"
            v-model="newAmount"
            outlined
        ></v-text-field>
      </v-col>
    </v-row>
    <v-btn
        rounded
        color="green"
        style="min-width: 120px;"
        @click="updateEquipment"
        :disabled="getBtnDisabled"
    >
      Speichern
    </v-btn>
  </div>
</template>

<script>
export default {
  name: "ManageEquipmentPage",

  data: () => ({
    currentType: "",
    newAmount: "0",
  }),

  created() {
    this.$store.dispatch("loadEquipment");
  },

  computed: {
    getEquipment() {
      return this.$store.getters.getEquipment;
    },
    getCurrentEquipment() {
      if (this.currentType) {
        return this.$store.getters.getEquipmentByType(this.currentType);
      }
      return 0;
    },
    getBtnDisabled() {
      return !this.currentType || (parseInt(this.newAmount) === parseInt(this.getCurrentEquipment.amount))
    }
  },

  methods: {
    updateEquipment() {
      let id = this.getCurrentEquipment.id;
      let type = this.currentType;
      let amount = this.newAmount;
      let data = {
        "id": id,
        "amount": amount,
        "type": type
      }
      this.$store.dispatch("updateEquipment", {id, data} );
    },
    setCurrentType(e) {
      this.currentType = e;
      this.newAmount = this.getCurrentEquipment.amount;
    }
  }
}
</script>

<style scoped>

</style>
