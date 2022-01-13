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
              v-model="getLoggedInUser.phone"
              :rules="phoneRules"
              label="Phone"
              required
              filled
              outlined
              prepend-icon="mdi-phone"
          ></v-text-field>

<!--          <v-text-field
              v-model="getLoggedInUser.password"
              :rules="passwordRules"
              label="Altes Passwort"
              prepend-icon="mdi-lock"
              required
              filled
              outlined
          ></v-text-field>
          <v-text-field
              v-model="getLoggedInUser.password"
              :rules="passwordRules"
              label="Neues Passwort"
              prepend-icon="mdi-lock"
              required
              filled
              outlined
          ></v-text-field>-->
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
                  dark
                  @click="closeDialog"
              >
                Abrechen
              </v-btn>

            </v-col>
            <v-col cols="3">

              <v-btn
                  style="text-transform: capitalize; font-weight: bolder;"
                  block
                  rounded
                  color="red"
                  dark
                  @click="submitUpdatedInfo"
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
      passwordRules: [
        v => !!v || 'Password is required',
      ],
    }
  },

  computed: {
    getLoggedInUser() {
      return this.$store.getters.getLoggedInUser
    }
    },
  methods: {
    closeDialog() {
      this.$emit('close')
    },
    submitUpdatedInfo(){
      let data={
        "name": this.getLoggedInUser.name,
        "phone": this.getLoggedInUser.phone,
      }
      this.$store.dispatch("updateUserInfo",  data)
      this.$emit("close")
    }
  }
}
</script>
