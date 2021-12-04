<template>
  <v-dialog
      v-model="dialog"
      width="500"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-btn
          style="text-transform: capitalize; font-weight: bolder;"
          rounded
          dark
          block
          v-bind:color=buttonColor
          v-bind:outlined=buttonOutlined
          v-bind:disabled=buttonDisabled
          v-bind="attrs"
          v-on="on"
      >
        {{buttonText}}
      </v-btn>
    </template>

    <v-card>
      <v-card-title>
        {{cardText}}
      </v-card-title>
      <v-card-actions>
        <v-btn
            rounded
            outlined
            @click="dialog = false"
        >
          Nein
        </v-btn>
        <v-btn
            rounded
            outlined
            @click="dialog = false"
        >
          Ja
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "ConfirmationDialog",
  props: [
      "buttonText",
      "buttonColor",
      "buttonOutlined",
      "buttonDisabled",
      "cardText"
  ],
  data: () => ({
    dialog: false
  })
}
</script>

<style scoped>

</style>
