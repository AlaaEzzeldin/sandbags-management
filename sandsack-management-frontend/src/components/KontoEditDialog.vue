<template>
  <v-dialog
      v-model="dialog"
      width="800"
  >
    <v-card outlined >
      <v-card-title>
        <v-col cols="2">
          <v-avatar
              color="white"
              size="60">
            <v-icon color="black" large>mdi-home</v-icon>
          </v-avatar>
        </v-col>
        <v-col cols="10">
          <h2>{{getUser.name}}</h2>
          <h3>{{getUser.branch_name}}</h3>
        </v-col>
      </v-card-title>
      <v-card-text>
        <v-form
            ref="form"
            v-model="valid"
            lazy-validation
        >
<!--          <v-text-field
              v-model="getLoggedInUser.email"
              :rules="emailRules"
              label="E-mail"
              required
              filled
              outlined
              prepend-icon="mdi-email"
          ></v-text-field>-->
          <v-text-field
              v-model="name"
              :rules="nameRules"
              label="Name"
              required
              filled
              outlined
              prepend-icon="mdi-account"
          ></v-text-field>
          <v-text-field
              v-model="phone"
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
                  @click="dialog = false"
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
  name: 'KontoEditDialog',
  props: ['dialog'],

  data() {
    return {
      valid: true,
      name: '',
      phone: '',
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
      ],
      phoneRules: [
        v => !!v || 'Phone is required',
        v => /\(?\+\(?49\)?[ ()]?([- ()]?\d[- ()]?){10}/.test(v) || 'Phone must be valid',
      ],
      nameRules: [
        v => !!v || 'Name is required',
      ],
      passwordRules: [
        v => !!v || 'Password is required',
      ],
    }
  },
  created: function() {
    this.phone = this.$store.getters.getLoggedInUser.phone
    this.name = this.$store.getters.getLoggedInUser.name
  },
  computed: {
    getUser(){
      return this.$store.getters.getLoggedInUser
    },

  },
  methods: {
    submitUpdatedInfo(){
      let data={
        "name": this.name,
        "phone": this.phone,
      }
      this.$store.dispatch("updateUserInfo",  data)
      this.$emit("close")
    }
  }
}
</script>
