<template>
  <v-card class="elevation-0 rounded-lg">
    <v-list rounded>
      <h3 class="ml-4 mb-4">Restliche Ausrüstung</h3>
      <v-list-item-group>
        <v-list-item
            v-for="(item, i) in getEquipment"
            :key="i"
        >
          {{item.type}} ({{item.measure}})
          <v-spacer/>
          <b>{{item.amount}}</b>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </v-card>
</template>

<script>
export default {
  name: "EquipmentQuantityTable",

  created() {
    this.$store.dispatch("loadEquipment")
  },

  computed: {
    getEquipment() {
      return this.$store.getters.getEquipment
    }
  },

}
</script>

<style scoped>

</style>
