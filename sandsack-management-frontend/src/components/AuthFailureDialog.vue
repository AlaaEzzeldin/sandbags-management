<template>
  <v-dialog
      v-model="dialog"
      width="500"
  >
    <v-card class="rounded-lg" outlined>
      <v-card-title>
        <v-icon color="black" class="mr-2">mdi-alert-circle-outline</v-icon>
        Ein Fehler ist aufgetreten
      </v-card-title>
      <v-divider></v-divider>

      <v-card-text class="mt-5 black--text">
          {{message}}
      </v-card-text>
      <v-card-actions>
        <v-spacer/>
        <v-btn
            text
            style="min-width: 120px;"
            @click="closeDialog"
        >
          schließen
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "AuthFailureDialog",
  data: () => ({

  }),
  props: [
    "dialog",
    "message",
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
