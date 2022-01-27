<template>
  <v-card elevation="0" v-if="getOrder && getPriorityByID">
    <v-card-title>
          <v-btn icon @click="goBack" v-if="$vuetify.breakpoint.mdAndUp">
            <v-icon large color="black" class="pr-5">mdi-keyboard-backspace</v-icon>
          </v-btn>
          <h1 style="font-weight: bolder; ">Bestellung # {{ getOrder.id }}</h1>
          <v-chip
              class="ml-5"
              :color="getColor(getOrder.status_name)" outlined
              dark
              v-if="$vuetify.breakpoint.mdAndUp"
              >
            {{ getOrder.status_name }}
          </v-chip>
    </v-card-title>

    <v-card-text class="pt-10" >
      <v-row v-if="$vuetify.breakpoint.mdAndDown">
        <v-col cols="12" sm="4">
          <h2 style="font-weight: bolder; color: black">Status:</h2>
        </v-col>
        <v-col cols="12" sm="7">
          <h2 style="font-weight: normal; color: black">{{ getOrder.status_name }}</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="4">
          <h2 style="font-weight: bolder; color: black">Von:</h2>
        </v-col>
        <v-col cols="12" sm="7">
          <h2 style="font-weight: normal; color: black">{{ getOrder.name }}</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="4">
          <h2 style="font-weight: bolder; color: black">Typ:</h2>
        </v-col>
        <v-col cols="12" sm="7">
          <h2 style="font-weight: normal; color: black" v-if="getOrder.equipments">{{ getOrder.equipments[0].name }}</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="4">
          <h2 style="font-weight: bolder; color: black">Anzahl:</h2>
        </v-col>
        <v-col cols="12" sm="7">
          <h2 style="font-weight: normal; color: black" v-if="getOrder.equipments">{{ getOrder.equipments[0].quantity }}</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="4">
          <h2 style="font-weight: bolder; color: black">Priorität:</h2>
        </v-col>
        <v-col cols="12" sm="7">
          <h2 style="font-weight: normal; color: black" v-if="getOrder.priority_id">{{ getPriorityByID(getOrder.priority_id).name }}</h2>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" sm="4">
          <h2 style="font-weight: bolder; color: black">Lieferadresse:</h2>
        </v-col>
        <v-col cols="12" sm="7">
          <h2 style="font-weight: normal; color: black" v-if="getOrder.address_to">{{ getOrder.address_to }}</h2>
        </v-col>
      </v-row>


      <!------------------------------------------------ comment -------------------------------->
      <div v-if="getOrder.comments" class="mt-10">
        <v-row v-if="getOrder.comments.find(comment=> comment.role === 'Unterabschnitt')">
          <v-col cols="12" sm="12">
            <h2 style="font-weight: bolder; color: black">Anmerkungen des Anforderers:</h2>
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
            <h2 style="font-weight: bolder; color: black">Notizen aus dem Einsatzabschnitt</h2>
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
            <h2 style="font-weight: bolder; color: black">Notizen aus dem hauptabschnitt</h2>
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
            <h2 style="font-weight: bolder; color: black">Notizen aus dem Einsatzleiter</h2>
          </v-col>
          <v-col cols="12" sm="12">
            <v-textarea
                readonly
                outlined
                :value="getOrder.comments.find(comment=> comment.role === 'Einsatzleiter').comment_text"
            ></v-textarea>
          </v-col>
        </v-row>
      </div>


      <!------------------------------------------------- Unterabschnitt ------------------------------------------->
      <v-card-actions v-if="getCurrentUserRole === 'Unterabschnitt'" class="mt-10">
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
                @click="editOrder"
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
              Bestellung absenden
            </v-btn>
          </v-col>
        </v-row>
      </v-card-actions>

      <!------------------------------------------------ logs -------------------------------->

      <v-row v-if="getOrder.logs" class="mt-10">
        <v-col cols="12" sm="12">
          <h2 style="font-weight: bolder; color: black">Bestellverlauf</h2>
        </v-col>
      </v-row>
      <v-row
          v-for="item in getOrder.logs"
          :key="item.id"
          style="color: black;"
      >
        <v-col cols="4">
          <b>{{ format_time(item.create_date) }}</b>
        </v-col>
        <v-col cols="8">
          <div class="float-right">
            <h3 style="color: black; font-weight: normal;">{{ item.description }}</h3>
          </div>
        </v-col>
      </v-row>
    </v-card-text>

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
