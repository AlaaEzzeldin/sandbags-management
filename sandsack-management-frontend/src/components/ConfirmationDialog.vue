<template>
  <v-dialog
      v-model="dialog"
      width="500"
  >
    <v-card>
      <v-card-title>
        {{cardText}}
      </v-card-title>
      <v-container>
        <v-text-field
            label="Notizen"
            v-if="hasTextField"
            :error="textFieldError"
            :error-messages="textFieldErrorMessages"
            :value="textFieldValue"
        ></v-text-field>
      </v-container>
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
  data: () => ({
    textFieldError: false,
    textFieldErrorMessages: [],
    textFieldValue: '',
  }),
  props: [
      "dialog",
      "cardText",
      'newStatus',
      'orderID',
      'hasTextField'
  ],
  methods:{
    closeDialog(){
      this.$emit('close')
  },

    submitNewStatus(){
      if (this.newStatus === 'abgelehnt' && !this.textFieldValue) {
        this.textFieldError = true;
        this.textFieldErrorMessages = ['Notizen sind verpflichtend!'];
      }
      else {
        this.textFieldError = false;
        this.textFieldErrorMessages = [];
        let data={
          "status": this.newStatus
        }
        let id= this.orderID
        this.$store.dispatch("updateOrder",  {id, data} )
        this.closeDialog()
      }
    },

}
}
</script>

<style scoped>

</style>
