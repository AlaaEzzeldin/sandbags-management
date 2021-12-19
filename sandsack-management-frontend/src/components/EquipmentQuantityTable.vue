<template>
  <v-card class="elevation-0 rounded-lg" outlined>
    <v-list rounded>
      <v-list-item-group>
        <v-list-item
            v-for="(item, i) in getEquipment"
            :key="i"
        >
          {{item.type}}
          {{item.amount}}
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

  data() {
    return {
      options: {
        itemsPerPage: -1,
      },
    }
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
