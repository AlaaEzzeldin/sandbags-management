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

Vue.use(VueRouter)

const routes = [
    {
        path: '/orders-list/:userRole',
        name: 'BestellungslistePage',
        component: BestellungslistePage
    },
    {
        path: '/account/:userRole',
        name: 'KontoPage',
        component: KontoPage
    },
    {
        path: '/new-order/:userRole',
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
    },
    {
        path: '/auth/login',
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
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
