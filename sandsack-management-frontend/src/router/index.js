import Vue from 'vue'
import VueRouter from 'vue-router'

// import pages
import Bestellungsliste from '../views/Bestellungsliste.vue'
import Konto from '../views/Konto.vue'
import NeueBestellung from '../views/NeueBestellung.vue'
import Bestellübersicht from '../views/Bestellübersicht.vue'

Vue.use(VueRouter)

const routes = [
    {
        path: '/:userRole',
        name: 'NeueBestellung',
        component: NeueBestellung
    },
    {
        path: '/account/:userRole',
        name: 'Konto',
        component: Konto
    },
    {
        path: '/requests-list/:userRole',
        name: 'Bestellungsliste',
        component: Bestellungsliste
    },
    {
        path: '/overview/:userRole',
        name: 'Bestellübersicht',
        component: Bestellübersicht
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
