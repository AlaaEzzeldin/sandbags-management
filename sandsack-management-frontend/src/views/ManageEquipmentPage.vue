<template>
  <div>
    <v-row no-gutters>
      <v-col class="pt-13 justify-center align-center">
        <h1 style="font-weight: bolder;">Restliche Ausrüstung Bearbeiten</h1>
      </v-col>
    </v-row>
    <v-row
        v-for="(item, i) in getEquipment"
        :key="i">
      <v-col>
        {{item.type}} ({{item.measure}})
        </v-col>
      <v-col>
        <v-text-field
            style="max-width: 50px;"
            v-model="item.amount"
        ></v-text-field>
      </v-col>
    </v-row>
    <v-btn
        rounded
        color="green"
        style="min-width: 120px;"
        @click="updateEquipment"
    >
      Speichern
    </v-btn>
  </div>
</template>

<script>
export default {
  name: "ManageEquipmentPage",

  data: () => ({
    equipment: {},
  }),

  created() {
    this.$store.dispatch("loadEquipment");
  },

  computed: {
    getEquipment() {
      return this.$store.getters.getEquipment;
    }
  },

  methods: {
    updateEquipment() {
      this.$store.dispatch("updateEquipment",  this.getEquipment );
    }
  }
}
</script>

<style scoped>

</style>
