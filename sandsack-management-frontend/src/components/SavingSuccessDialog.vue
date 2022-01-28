<template>
  <v-dialog
      v-model="dialog"
      width="400"
  >
    <v-card>
      <v-card-title>
        Ihre Änderungen wurden gespeichert
      </v-card-title>
      <v-card-actions>
        <v-spacer/>
        <v-btn
            rounded
            outlined
            color="green"
            @click="closeDialog"
        >
          Ok
        </v-btn>
        <v-spacer/>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "SavingSuccessDialog",
  data: () => ({}),
  props: [
    "dialog",
  ],
  methods: {
    closeDialog() {
      this.$emit('close')
    },
  }
}
</script>

<style scoped>

</style>
