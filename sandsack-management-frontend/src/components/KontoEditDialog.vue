<template>
  <v-dialog
      v-model="dialog"
      width="800"
      persistent
  >
    <v-card outlined>
      <v-card-title>
        <v-col cols="2">
          <v-avatar
              color="white"
              size="60">
            <v-icon color="black" large>mdi-home</v-icon>
          </v-avatar>
        </v-col>
        <v-col cols="10">
          <h2>{{ getUser.name }}</h2>
          <h3>{{ getUser.branch_name }}</h3>
        </v-col>
      </v-card-title>
      <v-card-text>
        <v-form
            ref="form"
            v-model="valid"
            lazy-validation
        >
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
              v-model="email"
              :rules="emailRules"
              label="E-mail"
              required
              filled
              outlined
              prepend-icon="mdi-email"
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

          <h2 class="mt-3 ml-10" style="font-weight: normal">
            * Sie werden aufgefordert, sich erneut anzumelden, wenn Sie Ihre Daten ändern.
          </h2>
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
              Abbrechen
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
      email: '',
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
  mounted: function () {
    this.phone = this.$store.getters.getLoggedInUser.phone
    this.name = this.$store.getters.getLoggedInUser.name
    this.email = this.$store.getters.getLoggedInUser.email

  },
  computed: {
    getUser() {
      return this.$store.getters.getLoggedInUser
    },

  },
  methods: {
    submitUpdatedInfo() {
      let data = {
        "name": this.name,
        "phone": this.phone,
        "email": this.email,
      }
      this.$store.dispatch("updateUserInfo", data).then(
          () => {
            console.log("old", this.$store.getters.getLoggedInUser.email)
            console.log("new", this.email)
            if (this.email !== this.$store.getters.getLoggedInUser.email)
              this.$store.dispatch("logout").then(
                  this.$router.push({name: 'LoginPage'}))
          },
      );
      this.$emit("close")
    },
    closeDialog() {
      this.$emit('close')
    }
  }
}
</script>
