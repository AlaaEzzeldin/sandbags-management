<template>
  <v-dialog
      v-model="dialog"
      width="800"
  >
    <v-card outlined >
      <v-card-title>
        <v-col>
          <h2>Kennwort ändern</h2>
        </v-col>
      </v-card-title>
      <v-card-text>
        <v-alert
            type="error"
            outlined
            :v-if="showAlert"
        >
          {{error}}
        </v-alert>
        <v-form
            ref="form"
            v-model="valid"
            lazy-validation
        >
          <v-text-field
              v-model="oldPassword"
              :rules="passwordRules"
              label="Altes Kennwort"
              required
              filled
              outlined
              prepend-icon="mdi-key"
          ></v-text-field>
          <v-text-field
              v-model="newPassword"
              :rules="passwordRules"
              label="Neues Kennwort"
              required
              filled
              outlined
              prepend-icon="mdi-key"
          ></v-text-field>
        </v-form>
      </v-card-text>
        <v-card-actions>
          <v-row>
            <v-col class="align-center justify-center" cols="3" offset="3">
              <v-btn
                  style="text-transform: capitalize; font-weight: bolder;"
                  block
                  rounded
                  color="red"
                  outlined
                  dark
                  @click="closeDialog"
              >
                Abbrechen
              </v-btn>
            </v-col>
            <v-col cols="3">
              <v-btn
                  style="text-transform: capitalize; font-weight: bolder;"
                  block
                  rounded
                  color="red"
                  @click="submitUpdatedInfo"
                  :disabled="!(newPassword && oldPassword)"
              >
                Speichern
              </v-btn>
            </v-col>

          </v-row>

        </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>

export default {
  name: 'PasswordEditDialog',
  props: ['dialog'],

  data() {
    return {
      valid: true,
      oldPassword: "",
      newPassword: "",
      error: "",
      passwordRules: [
        v => !!v || 'Password is required',
      ],
    }
  },

  computed: {
    getLoggedInUser() {
      return this.$store.getters.getLoggedInUser
    },
    showAlert() {
      return !this.valid && this.error;
    }
  },
  methods: {
    closeDialog() {
      this.$emit('close')
    },
    submitUpdatedInfo(){
      let data={
        "new_password": this.newPassword,
        "old_password": this.oldPassword,
      }
      this.$store.dispatch("updatePassword",  data).then(
          () => {
            this.newPassword = '';
            this.oldPassword = '';
            this.error = '';
            this.$emit("close")
          },
          error => {
            this.valid = false;
            this.error = error.response.data.err_message;
          }
      );
    }
  }
}
</script>
