<template>
  <v-card elevation="0" class="pt-10" v-if="getOrder && getPriorityByID">
    <v-card-title class="pt-10">
      <v-btn icon @click="goBack">
        <v-icon large color="black" class="pr-5">mdi-keyboard-backspace</v-icon>
      </v-btn>
      <h1 style="font-weight: bolder; ">Bestellung # {{ getOrder.id }}</h1>
      <v-chip
          class="ml-5"
          :color="getColor(getOrder.status_name)" outlined
          dark>
        {{ getOrder.status_name }}
      </v-chip>
    </v-card-title>

    <v-card-text class="pt-16 " >

      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Von:</h3>
        </v-col>
        <v-col cols="12" sm="4">
          <h3 style="font-weight: bolder; color: black">{{ getOrder.name }}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Typ:</h3>
        </v-col>
        <v-col cols="12" sm="4">
          <h3 style="font-weight: bolder; color: black">{{ getOrder.equipments[0].name }}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Anzahl:</h3>
        </v-col>
        <v-col cols="12" sm="4">
          <h3 style="font-weight: bolder; color: black">{{ getOrder.equipments[0].quantity }}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Priorität:</h3>
        </v-col>
        <v-col cols="12" sm="4">
          <h3 style="font-weight: bolder; color: black">{{ getPriorityByID(getOrder.priority_id).name }}</h3>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="2">
          <h3 style="font-weight: bolder; color: black">Lieferadresse:</h3>
        </v-col>
        <v-col cols="12" sm="4">
          <h3 style="font-weight: bolder; color: black">{{ getOrder.address_to }}</h3>
        </v-col>
      </v-row>


      <!------------------------------------------------ comment -------------------------------->
      <v-row v-if="getOrder.comments.find(comment=> comment.role === 'Unterabschnitt')">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Anmerkungen des Anforderers:</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              readonly
              :value="getOrder.comments.find(comment=> comment.role === 'Unterabschnitt').comment_text"
              outlined
          ></v-textarea>
        </v-col>
      </v-row>
      <v-row v-if="getOrder.comments.find(comment=> comment.role === 'Einsatzabschnitt')">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Notizen aus dem Einsatzabschnitt</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              readonly
              outlined
              :value="getOrder.comments.find(comment=> comment.role === 'Einsatzabschnitt').comment_text"
          ></v-textarea>
        </v-col>
      </v-row>
      <v-row v-if="getOrder.comments.find(comment=> comment.role === 'Hauptabschnitt')">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Notizen aus dem hauptabschnitt</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              readonly
              outlined
              :value="getOrder.comments.find(comment=> comment.role === 'Hauptabschnitt').comment_text"
          ></v-textarea>
        </v-col>
      </v-row>
      <v-row v-if="getOrder.comments.find(comment=> comment.role === 'Einsatzleiter')">
        <v-col cols="12" sm="12">
          <h3 style="font-weight: bolder; color: black">Notizen aus dem Einsatzleiter</h3>
        </v-col>
        <v-col cols="12" sm="12">
          <v-textarea
              readonly
              outlined
              :value="getOrder.comments.find(comment=> comment.role === 'Einsatzleiter').comment_text"
          ></v-textarea>
        </v-col>
      </v-row>


      <!------------------------------------------------ logs -------------------------------->

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
        <v-col cols="4">
          <b>{{ format_time(item.create_date) }}</b>
        </v-col>
        <v-col>
          <div class="float-right">
            {{ item.description }}
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
              :disabled="getOrder.status_name!=='ANSTEHEND'"
              @click="editOrder"
          >
            Bestellung bearbeiten
          </v-btn>
        </v-col>

        <v-col cols="12" sm="6" offset="3" v-if="getOrder.status_name==='AUF DEM WEG'">
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
              :disabled="getOrder.status_name!=='ANSTEHEND'"
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
              :disabled="getOrder.status_name!=='ANSTEHEND'"
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
              :disabled="getOrder.status_name!=='ANSTEHEND'"
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
              :disabled="getOrder.status_name!=='ANSTEHEND'"
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
              :disabled="getOrder.status_name!=='WEITERGELEITET BEI EINSATZABSCHNITT'"
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
              :disabled="getOrder.status_name!=='WEITERGELEITET BEI EINSATZABSCHNITT'"
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
              :disabled="getOrder.status_name!=='WEITERGELEITET BEI EINSATZABSCHNITT'"
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
              :disabled="getOrder.status_name!=='WEITERGELEITET BEI HAUPTABSCHNITT'"
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
              :disabled="getOrder.status_name!=='WEITERGELEITET BEI HAUPTABSCHNITT'"
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
              :disabled="getOrder.status_name!=='WEITERGELEITET BEI HAUPTABSCHNITT'"
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
              v-bind:disabled="getOrder.status_name!=='AKZEPTIERT'"
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
    action: '',
    confirmationDialog: false,
  }),

  created() {
    this.$store.dispatch("loadPriorities");
    this.$store.dispatch("loadEquipment");
    this.$store.dispatch("loadOrder", this.$route.params.orderId)
  },
  computed: {
    getOrder() {
      return this.$store.getters.getOrder
    },
    getPriorityByID() {
      return this.$store.getters.getPriorityByID
    },
    getCurrentUserRole() {
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

    changeStatus(action) {
      this.action = action;
      this.confirmationDialog = true;
    }
  }
}
</script>
