export const Mixin = {
        methods: {
            getColor(status) {
                if (status === 'ANSTEHEND') return 'grey'
                else if (status === 'STORNIERT') return 'red'

                else if (status === 'WEITERGELEITET BEI EINSATZABSCHNITT') return 'black'
                else if (status === 'WEITERGELEITET BEI HAUPTABSCHNITT') return 'black'

                else if (status === 'ABGELEHNT BEI EINSATZABSCHNITT') return 'red'
                else if (status === 'ABGELEHNT BEI HAUPTABSCHNITT') return 'red'
                else if (status === 'ABGELEHNT BEI EINSATZLEITER') return 'red'

                else if (status === 'AKZEPTIERT') return 'blue'
                else if (status === 'AUF DEM WEG') return 'orange'
                else if (status === 'GELIEFERT') return 'green'

            },
            getIcon (priority) {
                if (priority === 1) return {icon: 'mdi-chevron-up', color: 'red'}
                else if (priority === 2) return {icon: 'mdi-equal', color: 'orange'}
                else if (priority === 3) return {icon: 'mdi-chevron-down', color: 'blue'}
            },
            format_time(s) {
                return new Date(s).toLocaleString('en-GB', { timeZone: 'UTC' })
            }
        }
    }
;
