<template>
  <v-dialog
      v-model="dialog"
      width="800"
      persistent
  >
    <v-card outlined >
      <v-card-title>
        <v-col cols="2" v-if="$vuetify.breakpoint.mdAndUp">
          <v-avatar
              color="white"
              size="60">
            <v-icon color="black" large>mdi-home</v-icon>
          </v-avatar>
        </v-col>
        <v-col :cols="$vuetify.breakpoint.mdAndUp ? 10 : 12">
          <h3>{{getUser.name}}</h3>
          <h4>{{getUser.branch_name}}</h4>
        </v-col>
      </v-card-title>
      <v-card-text>
        <v-alert
            type="info"
            outlined
        >
          Wenn Sie Ihre E-Mail-Adresse ändern, werden Sie aufgefordert, sich erneut anzumelden
        </v-alert>
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
        </v-form>
      </v-card-text>
        <v-card-actions>
          <v-row>
            <v-col class="align-center justify-center">
              <v-btn
                  style="text-transform: capitalize; font-weight: bolder;"
                  block
                  rounded
                  color="red"
                  outlined
                  @click="closeDialog"
              >
                Abbrechen
              </v-btn>

            </v-col>
            <v-col>

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
  mounted: function() {
    this.phone = this.$store.getters.getLoggedInUser.phone
    this.name = this.$store.getters.getLoggedInUser.name
    this.email = this.$store.getters.getLoggedInUser.email

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
        "email": this.email,
      }
      this.$store.dispatch("updateUserInfo",  data).then(
          () => {
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
