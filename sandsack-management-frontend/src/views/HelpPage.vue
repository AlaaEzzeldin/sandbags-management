<template>
  <div ref="content" :class="$vuetify.breakpoint.mdAndUp ? 'pt-10 pl-3 pr-3' : ''">
    <v-row no-gutters>
      <v-col class="justify-center align-center">
        <h1 style="font-weight: bolder;">Hilfe</h1>
      </v-col>
    </v-row>
    <v-row class="mt-5">
      <v-expansion-panels>
        <v-expansion-panel
            v-for="(item,i) in questionsAnswers"
            :key="i"
        >
          <v-expansion-panel-header>
            <b>{{item.question}}</b>
          </v-expansion-panel-header>
          <v-expansion-panel-content>
            <p v-for="answer in item.answer" :key="answer"> {{answer}}</p>
            <v-img v-if="item.image" :src="item.image" max-height="500" contain/>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-row>
  </div>
</template>

<script>

export default {
  name: "HelpPage",
  methods: {
    goBack() {
      this.$router.go(-1)
    },
  },
  data: () => ({
    questionsAnswers: [
      {
        question: "Wie funktioniert das Login-Formular?",
        answer: ["Durch die Eingabe der genauen Anmeldedaten (Email und Passwort) während der Anmeldung kann für die Anmeldung der Anwendung verwendet werden."]
      },
      {
        question: "Wie kann man das Passwort ändern?",
        answer: ["Benutzen Sie die 'Konto' Seite, dann 'Passwort ändern' wählen."]
      },
      {
        question: "Welche Rollen gibt es im System?",
        answer: ["Unterabschnitt, Einsatzabschnitt, Hauptabschnitt, Einsatzleiter, Mollnhof."],
        image: require('@/assets/images/Rollen.png')
      },
      {
        question: "Wie sieht der Bestellablauf aus? Welcher Status kann eine Bestellung haben?",
        answer: ["Der Ablauf einer Bestellung ist: ",
          "1. Unterabschnitt macht eine neue Bestellung -> Bestellung hat Status 'ANSTEHEND'.",
          "2. Unterabschnitt kann die Bestellung bearbeiten -> Bestellung hat Status 'ANSTEHEND'.",
          "3. Unterabschnitt kann die Bestellung stornieren -> Bestellung hat Status 'STORNIERT'.",
          "4. Einsatzabschnitt kann die anstehende Bestellung an Hauptabschnitt weiterleiten oder ablehnen ->  Bestellung hat Status 'WEITERGELEITET BEI EINSATZABSCHNITT' oder 'ABGELEHNT'.",
          "5. Hauptabschnitt kann die weitergeleitete Bestellung an Einsatzleiter  weiterleiten oder ablehnen ->  Bestellung hat Status 'WEITERGELEITET BEI HAUPTABSCHNITT' oder 'ABGELEHNT'.",
          "6. Ensatzleiter kann die weitergeleitete Bestellung akzeptieren oder ablehnen ->  Bestellung hat Status 'AKZEPTIERT' oder 'ABGELEHNT'.",
          "7. Mollnhof kann die akzeptierte Bestellung senden  ->  Bestellung hat Status 'AUF DEM WEG'.",
          "8. Unterabschnitt bestätigt die Lieferung -> Bestellung hat Status 'GELIEFERT'."
        ],
        image: require('@/assets/images/Flow.png')
      },
      {
        question: "Gibt es Mindestbestellmenge?",
        answer: ["Es gibt keine Mindestbestellmenge, jede Bestellung kann je nach Anfrage aufgegeben werden, die dann priorisiert und bearbeitet wird."]
      },
      {
        question: "Wie hoch ist die Höchstzahl der Bestellungen?",
        answer: ["Sie können nicht mehr als Restmenge bestellen."]
      },
      {
        question: "Wie bearbeite ich einen Auftrag basierend auf einer Rolle?",
        answer: ["Die Bearbeitung einer Bestellung hängt von der Benutzerrolle ab, die im System arbeitet. Nicht jeder im System hat das Recht, die eingegangenen Bestellungen zu bearbeiten."]
      },
      {
        question: "Welche grundlegenden Anforderungen sind bei der Erstellung eines Auftrags zu beachten?",
        answer: ["Alle Bestellungen müssen mit genauer Straße, Adresse und Postleitzahl aufgegeben werden, damit sie rechtzeitig ankommen. Dadurch werden Verzögerungen bei der Bestellung und Nachbearbeitungszeiten vermieden."]
      },
      {
        question: "Wie viel Zeit muss für die Genehmigung mindestens aufgewendet werden?",
        answer: ["Die Mindestdauer der Genehmigung hängt davon ab, wie viele Aufträge noch ausstehen?wie viele bearbeitet werden? Je nach Bedarf wird es in der Notsituation schnell bearbeitet."]
      },
      {
        question: "Wie kann ich eine zusätzliche Menge an Taschen zurückgeben?",
        answer: ["Die Ausrüstung wird beim Mollnhof zurückgegeben. Mollnhof kann 'Ausrüstungverwaltung' dafür benutzen."]
      },
      {
        question: "Wie funktionieren das PDF-Exportieren auf Basis von Rollen?",
        answer: ["Hauptabschnitt, Einsatzabschnitt und Einsatzleite können die komplette Bestellliste in Bestellungsliste mit 'Exportieren' Taste herunterladen."]
      },
      {
        question: "Wie funktioniert der Lieferschein?",
        answer: ["Mollnhof kann den Lieferschein in Bestellungsliste mit 'Lieferschein drücken' Taste herunterladen. Dann kann der Lieferschein am Fahrer gegeben werden."]
      },
      {
        question: "Im Fall eines Systemausfalls oder falls Sie Hilfe benötigen, rufen Sie den Systemadministrator an",
        answer: ["Sie können 0851 43318 anrufen."]
      },
    ]
  })
}
</script>

<style scoped>

</style>
