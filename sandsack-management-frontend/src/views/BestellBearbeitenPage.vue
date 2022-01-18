<template>

    <EditRequestCard v-if="getOrder" :editedOrder="getOrder"></EditRequestCard>


</template>
<script>

import EditRequestCard from "@/components/EditRequestCard";

export default {
  name: 'BestellBearbeitenPage',

  components: {EditRequestCard},

  data: () => ({

  }),

  created() {
    this.$store.dispatch("loadOrder", this.$route.params.orderId);
    this.$store.dispatch("loadEquipment")
    this.$store.dispatch("loadPriorities")
  },
  computed: {
    getOrder() {
      return this.$store.getters.getOrder
    },
  }

}
</script>