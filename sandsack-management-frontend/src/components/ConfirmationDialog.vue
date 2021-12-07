<template>
  <v-dialog
      v-model="dialog"
      width="500"
  >
    <v-card>
      <v-card-title>
        {{cardText}}
      </v-card-title>
      <v-card-actions>
        <v-spacer/>
        <v-btn
            rounded
            outlined
            color="red"
            style="min-width: 120px;"
            @click="closeDialog"
        >
          Nein
        </v-btn>
        <v-spacer/>
        <v-btn
            rounded
            outlined
            color="green"
            style="min-width: 120px;"
            @click="submitNewStatus"
        >
          Ja
        </v-btn>
        <v-spacer/>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "ConfirmationDialog",
  props: [
      "dialog",
      "cardText",
      'newStatus',
      'orderID'
  ],
  methods:{
    closeDialog(){
      this.$emit('close')
  },

    submitNewStatus(){
      let data={
        "status": this.newStatus
      }
      let id= this.orderID
      this.$store.dispatch("updateOrder",  {id, data} )
      this.closeDialog()
    }

}
}
</script>

<style scoped>

</style>
