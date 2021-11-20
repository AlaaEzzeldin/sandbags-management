<template>
  <div>
    <h1 style="font-weight: bolder; " class="pt-15">Bestellungsliste</h1>
    <RequestsTable class="mt-10"></RequestsTable>
  </div>
</template>

<script>

import RequestsTable from "@/components/RequestsTable";

export default {
  name: 'Bestellungsliste',
  components: {
    RequestsTable
  },
}
</script>
