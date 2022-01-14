<template>
  <v-dialog
      v-model="dialog"
      width="800"
  >
    <v-card>
      <v-card-title>
        {{ getMessageText() }}
      </v-card-title>
      <v-container>
        <v-textarea
            label="Notizen"
            outlined
            v-if="action==='cancel'"
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
    'action',
    'orderID',
  ],
  computed: {
    getCurrentUserRole() {
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
      if (this.action === 'cancel') { // in case of cancel send a comment
        if (!this.textFieldValue) {
          this.textFieldError = true;
          this.textFieldErrorMessages = ['Notizen sind verpflichtend!']
        } else this.commentOrder()
      }
      // else dispatch directly the new status
      this.changeOrderStatus()
      this.closeDialog()

    },
    commentOrder() {
      let data = {
        "comment": this.textFieldValue,
        "order_id": this.orderID
      }
      this.$store.dispatch("commentOrder", data)
      this.closeDialog()
    },

    changeOrderStatus() {
      let id = this.orderID
      if (this.action === 'accept')
        this.$store.dispatch("acceptOrder", id)
      else if (this.action === 'cancel')
        this.$store.dispatch("cancelOrder", id)
      else if (this.action === 'confirm_delivery')
        this.$store.dispatch("confirmOrderDelivery", id)
      else if (this.action === 'dispatch')
        this.$store.dispatch("dispatchOrder", id)

      this.closeDialog()
    },
    getMessageText() {
      if (this.getCurrentUserRole === 'Unterabschnitt') {
        if (this.action === 'confirm_delivery')
          return 'Lieferung bestätigen?'
        else if (this.action === 'cancel')
          return 'Bestellung stornieren?'
      } else if (this.getCurrentUserRole === 'Einsatzabschnitt') {
        if (this.action === 'accept')
          return 'Bestellung weiterleiten an Hauptabschnitt?'
        else if (this.action === 'cancel')
          return 'Bestellung ablehnen?'
      } else if (this.getCurrentUserRole === 'Hauptabschnitt') {
        if (this.action === 'accept')
          return 'Bestellung weiterleiten an Einsatzleiter?'
        else if (this.action === 'cancel')
          return 'Bestellung ablehnen?'
      } else if (this.getCurrentUserRole === 'Einsatzleiter') {
        if (this.action === 'accept')
          return 'akzeptiert'
        else if (this.action === 'cancel')
          return 'Bestellung ablehnen?'
      } else if (this.getCurrentUserRole === 'Mollnhof') {
        if (this.action === 'dispatch')
          return 'Bestellung senden?'
      }

    }

  }
}
</script>

<style scoped>

</style>
