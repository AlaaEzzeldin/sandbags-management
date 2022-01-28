<template>
  <div :class="$vuetify.breakpoint.mdAndUp ? 'pt-10 pl-3 pr-3' : ''" v-if="getEquipment.length!==0">
    <v-row no-gutters>
      <v-col class="justify-center align-center">
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
      <v-col :cols="$vuetify.breakpoint.mdAndUp ? 4 : 8">
        <h3>Aktueller Betrag:</h3>
      </v-col>
      <v-col :cols="$vuetify.breakpoint.mdAndUp ? 2 : 4">
        <h3 style="font-weight: normal">{{getCurrentEquipment.quantity}} </h3>
      </v-col>
    </v-row>
    <v-row>
      <v-col :cols="$vuetify.breakpoint.mdAndUp ? 4 : 8">
        <h3>Bitte geben Sie den Betrag ein, der zurückgegeben wurde:</h3>
      </v-col>
      <v-col :cols="$vuetify.breakpoint.mdAndUp ? 2 : 4">
        <v-text-field
            :value="0"
            v-model="newAmount"
            outlined
        ></v-text-field>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-btn
            style="text-transform: capitalize; font-weight: bolder;"
            rounded
            color="green"
            block
            outlined
            @click="updateEquipment"
            :disabled="getBtnDisabled"
        >
          Speichern
        </v-btn>
      </v-col>
    </v-row>
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
      this.currentType= this.getEquipment.find(item=>item.name==='Sandsack').name
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
      return (parseInt(this.newAmount) === 0) || this.newAmount.length===0
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
