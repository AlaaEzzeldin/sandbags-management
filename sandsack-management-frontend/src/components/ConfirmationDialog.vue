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
            @click="this.$emit('close')"
        >
          Nein
        </v-btn>
        <v-spacer/>
        <v-btn
            rounded
            outlined
            color="green"
            style="min-width: 120px;"
            @click="this.$emit('close')"
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
      "cardText"
  ],
}
</script>

<style scoped>

</style>
