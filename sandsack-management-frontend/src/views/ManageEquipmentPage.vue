<template>
  <div v-if="getEquipment.length!==0">
    <v-row no-gutters>
      <v-col class="pt-13 justify-center align-center">
        <h1 style="font-weight: bolder;">Ausrüstung Zurückgeben</h1>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-select
            :label="getEquipment.find(item=>item.name==='Sandsack').name"
            :items="getEquipment.map(item => item.name)"
            @change="setCurrentType"
            outlined
        ></v-select>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <h3>Aktueller Betrag:</h3>
      </v-col>
      <v-col>
        {{getCurrentEquipment.quantity}}
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <h3>Zurückgeben:</h3>
      </v-col>
      <v-col>
        <v-text-field
            :value="0"
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
    this.$store.dispatch("loadEquipment")
    if(this.getEquipment.length!==0)
      this.currentType= this.getEquipment.find(item=>item.name==='Sandsack')
  },
  computed: {
    getEquipment() {
      return this.$store.getters.getEquipment;
    },
    getCurrentEquipment() {
      if (this.currentType) {
        return this.$store.getters.getEquipmentByType(this.currentType);
      }
      return this.$store.getters.getEquipmentByType('Sandsack');
    },
    getBtnDisabled() {
      return (parseInt(this.newAmount) === 0)
    }
  },

  methods: {
    updateEquipment() {
      let id = this.getCurrentEquipment.id;
      let type = this.currentType? this.currentType: 'Sandsack' ;
      let amount = this.newAmount;
      let data = {
        "id": id,
        "quantity": parseInt(amount),
        "name": type
      }
      this.$store.dispatch("updateEquipment", {id, data} );
    },
    setCurrentType(e) {
      this.currentType = e;
    }
  }
}
</script>

<style scoped>

</style>
