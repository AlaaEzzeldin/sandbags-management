<template>
  <v-card elevation="0" class="pt-10">
    <v-card-title class="pt-10">
      <v-btn icon @click="goBack">
        <v-icon large color="black" class="pr-5">mdi-keyboard-backspace</v-icon>
      </v-btn>
      <h1 style="font-weight: bolder; ">Bestellung # {{ getOrder.id }}</h1>
      <v-chip
          class="ml-5"
          :color="getColor(getOrder.status)" outlined
          dark>
        {{ getOrder.status }}
      </v-chip>
    </v-card-title>

    <v-card-text class="pt-16 ">

      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Von:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <h3 style="font-weight: bolder; color: black">{{ getOrder.from }}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Typ:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <h3 style="font-weight: bolder; color: black">{{ getOrder.equipments[0].name }}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Anzahl:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <h3 style="font-weight: bolder; color: black">{{ getOrder.quantity }}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Priorität:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <h3 style="font-weight: bolder; color: black">{{ getOrder.priority }}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Lieferadresse:</h3>
        </v-col>
        <v-col cols="12" sm="3">
          <h3 style="font-weight: bolder; color: black">{{ getOrder.deliveryAddress }}</h3>
        </v-col>
      </v-row>
      <v-row v-if="getOrder.notesByUnterabschnitt">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Anmerkungen des Anforderers:</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              readonly
              :value="getOrder.notesByUnterabschnitt"
              outlined
          ></v-textarea>
        </v-col>
      </v-row>
      <v-row v-if="getOrder.notesByEinsatzabschnitt">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Notizen aus dem Einsatzabschnitt</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              readonly
              outlined
              :value="getOrder.notesByEinsatzabschnitt"
          ></v-textarea>
        </v-col>
      </v-row>
      <v-row v-if="getOrder.notesByHauptabschnitt">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Notizen aus dem hauptabschnitt</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              readonly
              outlined
              :value="getOrder.notesByHauptabschnitt"
          ></v-textarea>
        </v-col>
      </v-row>
      <v-row v-if="getOrder.logs">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Bestellverlauf</h3>
        </v-col>
      </v-row>
      <v-row
          v-for="item in getOrder.logs"
          :key="item.id"
          style="color: black"
      >
        <v-col cols="3">
          <b>{{item.create_date}}</b>
        </v-col>
        <v-col>
          <div class="float-right">
            {{item.description}}
          </div>
        </v-col>
      </v-row>
    </v-card-text>

    <!------------------------------------------------- Unterabschnitt ------------------------------------------->
    <v-card-actions v-if="getCurrentUserRole === 'Unterabschnitt'">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="orange"
              dark
              block
              outlined
              :disabled="getOrder.status!=='ANSTEHEND'"
              @click="editOrder"
          >
            Bestellung bearbeiten
          </v-btn>
        </v-col>

        <v-col cols="12" sm="6" offset="3" v-if="getOrder.status==='AUF DEM WEG'">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="green"
              dark
              block
              outlined
              @click="changeStatus('confirm_delivery')"
          >
            Lieferung bestätigen
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              :disabled="getOrder.status!=='ANSTEHEND'"
              @click="changeStatus('cancel')"
          >
            Bestellung stornieren
          </v-btn>
        </v-col>
      </v-row>

    </v-card-actions>

    <!---------------------------------- Einsatzabschnitt  -------------------------------->
    <v-card-actions v-if="this.getCurrentUserRole === 'Einsatzabschnitt'">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="green"
              dark
              block
              outlined
              :disabled="getOrder.status!=='ANSTEHEND'"
              @click="changeStatus('accept')"
          >
            Bestellung weiterleiten an Hauptabschnitt
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              outlined
              :disabled="getOrder.status!=='ANSTEHEND'"
              @click="editOrder"
          >
            Bestellung bearbeiten
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              :disabled="getOrder.status!=='ANSTEHEND'"
              @click="changeStatus('cancel')"
          >
            Bestellung ablehnen
          </v-btn>
        </v-col>
      </v-row>
    </v-card-actions>

    <!----------------------------------  Hauptabschnitt -------------------------------->
    <v-card-actions v-if="getCurrentUserRole === 'Hauptabschnitt'">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="green"
              dark
              block
              outlined
              :disabled="getOrder.status!=='WEITERGELEITET BEI EINSATZABSCHNITT'"
              @click="changeStatus('accept')"
          >
            Bestellung annehmen
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              outlined
              :disabled="getOrder.status!=='WEITERGELEITET BEI EINSATZABSCHNITT'"
          >
            Bestellung bearbeiten
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              :disabled="getOrder.status!=='WEITERGELEITET BEI EINSATZABSCHNITT'"
              @click="changeStatus('cancel')"
          >
            Bestellung ablehnen
          </v-btn>
        </v-col>
      </v-row>
    </v-card-actions>

    <!----------------------------------  Einsatzleiter -------------------------------->
    <v-card-actions v-if="getCurrentUserRole === 'Einsatzleiter'">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="green"
              dark
              block
              outlined
              :disabled="getOrder.status!=='WEITERGELEITET BEI HAUPTABSCHNITT'"
              @click="changeStatus('accept')"
          >
            Bestellung annehmen
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              outlined
              :disabled="getOrder.status!=='WEITERGELEITET BEI HAUPTABSCHNITT'"
              @click="editOrder"
          >
            Bestellung bearbeiten
          </v-btn>
        </v-col>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="red"
              dark
              block
              :disabled="getOrder.status!=='WEITERGELEITET BEI HAUPTABSCHNITT'"
              @click="changeStatus('cancel')"
          >
            Bestellung ablehnen
          </v-btn>
        </v-col>
      </v-row>
    </v-card-actions>

    <!------------------------------------------------- Mollhof ------------------------------------------->
    <v-card-actions v-if="getCurrentUserRole === 'Mollnhof'">
      <v-row>
        <v-col cols="12" sm="6" offset="3">
          <v-btn
              style="text-transform: capitalize; font-weight: bolder;"
              rounded
              color="green"
              dark
              block
              outlined
              v-bind:disabled="getOrder.status!=='AKZEPTIERT'"
              @click="changeStatus('dispatch')"
          >
            Bestellung abgesendet
          </v-btn>
        </v-col>
      </v-row>
    </v-card-actions>

    <ConfirmationDialog
        :action="action"
        :orderID="getOrder.id"
        :dialog="confirmationDialog"
        @close="confirmationDialog = false"
    />
  </v-card>
</template>

<script>
import ConfirmationDialog from "./ConfirmationDialog";
import {Mixin} from '../mixin/mixin.js'

export default {
  name: 'BestelldetailsCard',
  components: {ConfirmationDialog},
  mixins: [Mixin],

  data: () => ({
    action:'',
    confirmationDialog: false,
  }),

  created() {
    this.$store.dispatch("loadOrder", this.$route.params.orderId)
  },
  computed: {
    getOrder() {
      return this.$store.getters.getOrder
    },
    getCurrentUserRole(){
      return this.$store.getters.getCurrentUserRole
    }
  },
  methods: {
    goBack() {
      this.$router.go(-1)
    },
    editOrder() {
      const orderId = this.getOrder.id;
      this.$router.push({name: 'BestellBearbeitenPage', params: {orderId}})
    },

    changeStatus(action){
      this.action = action;
      this.confirmationDialog = true;
    }
  }
}
</script>
