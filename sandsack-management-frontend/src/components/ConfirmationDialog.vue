<template>
  <v-dialog
      v-model="dialog"
      width="800"
  >
    <v-card>
      <v-card-title>
        {{ cardText }}
      </v-card-title>
      <v-container>
        <v-textarea
            label="Notizen"
            outlined
            v-if="hasTextField"
            :error="textFieldError"
            :error-messages="textFieldErrorMessages"
            v-model="textFieldValue"
        ></v-textarea>
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
  computed: {
    getCurrentUserRole(){
      return this.$store.getters.getCurrentUserRole
    },
    getLoggedInUser() {
      return this.$store.getters.getLoggedInUser
    },
  },
  methods: {
    closeDialog() {
      this.$emit('close')
    },

    submitNewStatus() {
      if (this.newStatus === 'abgelehnt') {
        if(!this.textFieldValue){
          this.textFieldError = true;
          this.textFieldErrorMessages = ['Notizen sind verpflichtend!']
        }
        else this.submitUpdatedOrder()
      } else {
        let data = {
          "status": this.newStatus
        }
        let id = this.orderID
        this.$store.dispatch("updateOrder", {id, data})
        this.closeDialog()
      }
    },
    submitUpdatedOrder() {
      let TextBy = this.getLoggedInUser.name
      let data = {
        "status": this.newStatus,
        [TextBy] : this.textFieldValue
      }
      let id = this.orderID
      this.$store.dispatch("updateOrder", {id, data})
      this.closeDialog()
    }
  }
}
</script>

<style scoped>

</style>
