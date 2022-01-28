<template>
  <div ref="content" :class="$vuetify.breakpoint.mdAndUp ? 'pt-10 pl-3 pr-3' : ''">
    <BestelldetailsCard/>
  </div>
</template>

<script>

import BestelldetailsCard from "@/components/BestelldetailsCard";

export default {
  name: 'BestelldetailsPage',

  components: {BestelldetailsCard},

  data: () => ({

  }),

}
</script>
