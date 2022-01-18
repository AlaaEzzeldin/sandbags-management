<template>
  <v-dialog
      v-model="dialog"
      width="800"
  >
    <v-card class="rounded-lg" outlined>
      <v-card-title>
        <h2 style="color: red"> Ops!</h2>
        <v-spacer></v-spacer>
        <v-btn icon @click.stop="closeDialog">
          <v-icon color="black ">mdi-close</v-icon>
        </v-btn>
      </v-card-title>
      <v-divider></v-divider>

      <v-card-text class="mt-5">

          <h2>  {{message}}</h2>
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
