import Vue from 'vue'
import VueRouter from 'vue-router'

// import pages
import BestellungslistePage from '../views/BestellungslistePage.vue'
import KontoPage from '../views/KontoPage.vue'
import NeueBestellungPage from '../views/NeueBestellungPage.vue'
import BestellübersichtPage from '../views/BestellübersichtPage'
import BestelldetailsPage from "../views/BestelldetailsPage";
import LoginPage from "../views/LoginPage";
import SignupPage from "../views/SignupPage";
import RecoverPasswordPage from "../views/RecoverPasswordPage";
import RecoverPasswordConfirmPage from "../views/RecoverPasswordConfirmPage";
import HomePage from "../views/HomePage";
import BestellBearbeitenPage from "../views/BestellBearbeitenPage";
import ManageEquipmentPage from "../views/ManageEquipmentPage";

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'LoginPage',
        component: LoginPage
    },
    {
        path: '/auth/sign-up',
        name: 'SignupPage',
        component: SignupPage
    },
    {
        path: '/auth/recover-password',
        name: 'RecoverPasswordPage',
        component: RecoverPasswordPage
    },
    {
        path: '/auth/recover-password-confirm',
        name: 'RecoverPasswordConfirmPage',
        component: RecoverPasswordConfirmPage
    },
    {
        path: '/homapage',
        name: 'HomePage',
        component: HomePage,
        children: [
            {
                path: '/orders-list',
                name: 'BestellungslistePage',
                component: BestellungslistePage
            },
            {
                path: '/account',
                name: 'KontoPage',
                component: KontoPage
            },
            {
                path: '/new-order',
                name: 'NeueBestellungPage',
                component: NeueBestellungPage
            },
            {
                path: '/overview',
                name: 'BestellübersichtPage',
                component: BestellübersichtPage
            },
            {
                path: '/order-details/:orderId',
                name: 'BestelldetailsPage',
                component: BestelldetailsPage
            },
            {
                path: '/order-edit/:orderId',
                name: 'BestellBearbeitenPage',
                component: BestellBearbeitenPage
            },
            {
                path: '/:userRole/equipment-edit',
                name: 'ManageEquipmentPage',
                component: ManageEquipmentPage
            },
        ]
    },


]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
