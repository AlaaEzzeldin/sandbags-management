<template>
  <v-container>
    <v-expansion-panels>
      <v-expansion-panel
          v-for="(item,i) in questionsAnswers"
          :key="i"
      >
        <v-expansion-panel-header>
          {{item.question}}
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          {{item.answer}}
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
  </v-container>
</template>

<script>
export default {
  name: "HelpPage",
  data: () => ({
    questionsAnswers: [
      {
        question: "Wie funktioniert das login formular?",
        answer: "Durch die Eingabe der genauen Anmeldedaten wie Benutzername und Passwort während der Registrierung kann für die Anmeldung der Anwendung verwendet werden."
      },
      {
        question: "Wie funktioniert das Anmeldeng formular?",
        answer: "Durch Hinzufügen der entsprechenden E-Mail-Adresse und des Passworts können sich die Mitarbeiter im System registrieren. Anschließend wird dieser Benutzername als Login für die erfolgreiche Authentifizierung in der System-App verwendet."
      },
      {
        question: "Wenn Passwort vergessen?",
        answer: "Verwenden Sie immer den Link 'Passwort vergessen', um das Passwort mit den angegebenen Anmeldedaten erfolgreich zurückzusetzen."
      },
      {
        question: "Wie kann man das Passwort wiederherstellen?",
        answer: "Verwenden Sie immer die Seite zur Wiederherstellung des Kennworts. Fügen Sie dort die erforderlichen Anmeldedaten hinzu, bevor Sie sich am System anmelden. Auf diese Weise wird eine E-Mail zum Zurücksetzen des Passworts an die angegebene Adresse gesendet. Setzen Sie von dort aus das Passwort zurück und melden Sie sich erneut erfolgreich an."
      },
      {
        question: "Wie hoch ist die Mindestbestellmenge?",
        answer: "Es gibt keine Mindestbestellmenge, jede Bestellung kann je nach Anfrage aufgegeben werden, die dann priorisiert und bearbeitet wird."
      },
      {
        question: "Wie hoch ist die Höchstzahl der Bestellungen?",
        answer: " Wie bereits erwähnt, gibt es keine Maximalaufträge, es kommt immer auf den Bedarf in der Notsituation an."
      },
      {
        question: " Wie bearbeite ich einen Auftrag basierend auf einer Rolle?",
        answer: "Die Bearbeitung einer Bestellung hängt von der Benutzerrolle ab, die im System arbeitet. Nicht jeder im System hat das Recht, die eingegangenen Bestellungen zu bearbeiten."
      },
      {
        question: "Erläutern Sie den Status von Aufträgen?",
        answer: "Status of the oder depends on akzeptiert, Auf dem weg, Geliefert, abgelehnt etc. Dies kann auf der Seite Bestellliste zu sehen, die je nach Benutzerrolle und -berechtigung angezeigt wird."
      },
      {
        question: "Erläutern Sie die Bearbeitung von Aufträgen?",
        answer: "Die Bearbeitung der Bestellung hängt von der Priorität ab, die Hohe, Niedrige, Mittlere ist. Unter Berücksichtigung der Priorität wird jeder Betrag der Bestellung erfolgreich verarbeitet."
      },
      {
        question: "Welche grundlegenden Anforderungen sind bei der Erstellung eines Auftrags zu beachten?",
        answer: "Alle Bestellungen müssen mit genauer Straße, Adresse und Postleitzahl aufgegeben werden, damit sie rechtzeitig ankommen. Dadurch werden Verzögerungen bei der Bestellung und Nachbearbeitungszeiten vermieden."
      },
      {
        question: "Wie viel Zeit muss für die Genehmigung mindestens aufgewendet werden?",
        answer: "Die Mindestdauer der Genehmigung hängt davon ab, wie viele Aufträge noch ausstehen?wie viele bearbeitet werden? Je nach Bedarf wird es in der Notsituation schnell bearbeitet."
      },
      {
        question: "Wie kann ich eine zusätzliche Menge an Taschen zurückgeben?",
        answer: "Die zusätzliche Menge wird im System angezeigt, am Ende oder in verschiedenen Zeitintervallen werden zusätzliche Säcke an verschiedenen Orten abgeholt."
      },
      {
        question: "Wie funktionieren die PDF-Exportiern auf Basis von Rollen?",
        answer: "Pdf exportieren funktioniert für die verschiedenen im System genannten Rollen wie Hauptabschnitt, Einsatzabschnitt etc. Sie können die komplette Bestellliste einfach herunterladen."
      },
      {
        question: "Wie funktioniert der Lieferschein?",
        answer: "Lieferschein arbeitet für die verschiedenen Rollen, wie im System erwähnt, wie z.B. Mollhof, Sie können ein pdf erstellen, das mit id, zeit, von, prioritat, status gedruckt werden kann."
      },
      {
        question: "Im Falle eines Systemausfalls oder wenn Sie Hilfe benötigen, rufen Sie den Systemadministrator an?",
        answer: "Im Falle eines Systemfehlers wird eine leere Seite mit zusätzlichen Dialogen eingeblendet, wie z. B. 1) Wen soll ich kontaktieren? 2) Welche Nummern müssen verwendet werden? 3) Wie ist die Situation zu behandeln?"
      },
      {
        question: "Im Falle eines 404-Fehlers?",
        answer: "Im Falle eines Systemfehlers wird eine leere Seite mit zusätzlichen Dialogen eingeblendet, wie z. B. 1) Wen soll ich kontaktieren? 2) Welche Nummern müssen verwendet werden? 3) Wie ist die Situation zu behandeln?"
      },
      {
        question: "Wie funktioniert der Systemfluss in der Anwendung?",
        answer: "Der Systemfluss hängt von den Benutzerrollen ab, wie sie das System nutzen. Hängt hauptsächlich von der, 1) Registrierung 2) Anmeldung 3) Bestellung Liste 4) Bestellungen & Status 5) Erstellter Bericht"
      },
    ]
  })
}
</script>

<style scoped>

</style>
