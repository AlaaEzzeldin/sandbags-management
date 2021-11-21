import Vue from 'vue'
import VueRouter from 'vue-router'

// import pages
import BestellungslistePage from '../views/BestellungslistePage.vue'
import KontoPage from '../views/KontoPage.vue'
import NeueBestellungPage from '../views/NeueBestellungPage.vue'
import BestellübersichtPage from '../views/BestellübersichtPage'
import BestelldetailsPage from "../views/BestelldetailsPage";

Vue.use(VueRouter)

const routes = [
    {
        path: '/:userRole',
        name: 'BestellungslistePage',
        component: BestellungslistePage
    },
    {
        path: '/account/:userRole',
        name: 'KontoPage',
        component: KontoPage
    },
    {
        path: '/orders-list/:userRole',
        name: 'NeueBestellungPage',
        component: NeueBestellungPage
    },
    {
        path: '/overview/:userRole',
        name: 'BestellübersichtPage',
        component: BestellübersichtPage
    },
    {
        path: '/:userRole/order-details/:orderId',
        name: 'BestelldetailsPage',
        component: BestelldetailsPage
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
