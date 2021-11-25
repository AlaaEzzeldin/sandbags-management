<template>
  <v-form v-model="valid">
    <v-text-field
        v-model="email"
        :rules="emailRules"
        label="E-mail"
        required
    ></v-text-field>
    <v-text-field
        v-model="password"
        :rules="passwordRules"
        label="Password"
        required
    ></v-text-field>
    <v-text-field
        v-model="password"
        :rules="passwordRules"
        label="Repeat password"
        required
    ></v-text-field>
    <v-btn
        :disabled="!valid"
        color="success"
        class="mr-4"
        @click="validate"
    >
      Anmelden
    </v-btn>
  </v-form>
</template>

<script>

export default {
  name: 'SignupPage',

  data: () => ({

  }),

}
</script>
