export const Mixin = {
    methods: {
        getColor(status) {
            if (status === 'akzeptiert') return 'blue'
            if (status === 'geliefert') return 'green'
            else if (status === 'abgelehnt') return 'red'
            else if (status === 'storniert') return 'red'
            else if (status === 'Auf dem Weg') return 'orange'
            else if (status === 'anstehend') return 'grey'
            else if (status === 'weitergeleitet') return 'black'
        }
    }
};