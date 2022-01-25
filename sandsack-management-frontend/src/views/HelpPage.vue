<template>
  <v-container>
    <v-row>
      <v-btn icon @click="goBack">
        <v-icon large color="black" class="pr-5">mdi-keyboard-backspace</v-icon>
      </v-btn>
      <h2>Hilfe</h2>
    </v-row>
    <v-row class="pt-4">
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
            <v-img v-if="item.image" :src="item.image"/>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-row>
  </v-container>
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
        question: "Wie hoch ist die Mindestbestellmenge?",
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
        question: "Was its der Ablauf einer Bestellung? Welchen Status kann eine Bestellung haben?",
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
        question: "Erläutern Sie die Bearbeitung von Aufträgen?",
        answer: ["Die Bearbeitung der Bestellung hängt von der Priorität ab, die Hohe, Niedrige, Mittlere ist. Unter Berücksichtigung der Priorität wird jeder Betrag der Bestellung erfolgreich verarbeitet."]
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
        answer: ["Die zusätzliche Menge wird im System angezeigt, am Ende oder in verschiedenen Zeitintervallen werden zusätzliche Säcke an verschiedenen Orten abgeholt."]
      },
      {
        question: "Wie funktionieren die PDF-Exportiern auf Basis von Rollen?",
        answer: ["Pdf exportieren funktioniert für die verschiedenen im System genannten Rollen wie Hauptabschnitt, Einsatzabschnitt etc. Sie können die komplette Bestellliste einfach herunterladen."]
      },
      {
        question: "Wie funktioniert der Lieferschein?",
        answer: ["Lieferschein arbeitet für die verschiedenen Rollen, wie im System erwähnt, wie z.B. Mollhof, Sie können ein pdf erstellen, das mit id, zeit, von, prioritat, status gedruckt werden kann."]
      },
      {
        question: "Im Falle eines Systemausfalls oder wenn Sie Hilfe benötigen, rufen Sie den Systemadministrator an?",
        answer: ["Im Falle eines Systemfehlers wird eine leere Seite mit zusätzlichen Dialogen eingeblendet, wie z. B. 1) Wen soll ich kontaktieren? 2) Welche Nummern müssen verwendet werden? 3) Wie ist die Situation zu behandeln?"]
      },
      {
        question: "Im Falle eines 404-Fehlers?",
        answer: ["Im Falle eines Systemfehlers wird eine leere Seite mit zusätzlichen Dialogen eingeblendet, wie z. B. 1) Wen soll ich kontaktieren? 2) Welche Nummern müssen verwendet werden? 3) Wie ist die Situation zu behandeln?"]
      },
      {
        question: "Wie funktioniert der Systemfluss in der Anwendung?",
        answer: ["Der Systemfluss hängt von den Benutzerrollen ab, wie sie das System nutzen. Hängt hauptsächlich von der, 1) Registrierung 2) Anmeldung 3) Bestellung Liste 4) Bestellungen & Status 5) Erstellter Bericht"]
      },
    ]
  })
}
</script>

<style scoped>

</style>
