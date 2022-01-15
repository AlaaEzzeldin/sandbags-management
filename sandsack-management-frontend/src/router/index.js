import Vue from 'vue'
import VueRouter from 'vue-router'
import auth from '../store/modules/auth'

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
        component: LoginPage,
        meta: { requiresGuest: true }
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
        meta: { requiresAuth: true },
        children: [
            // signin
            {
                path: '/orders-list',
                name: 'BestellungslistePage',
                component: BestellungslistePage,
                meta: { requiresAuth: true }
            },
            // signin
            {
                path: '/account',
                name: 'KontoPage',
                component: KontoPage,
                meta: { requiresAuth: true }

            },
            // unterabschnit only
            {
                path: '/new-order',
                name: 'NeueBestellungPage',
                component: NeueBestellungPage
            },
            // einsatz/ haupt and leiter
            {
                path: '/overview',
                name: 'BestellübersichtPage',
                component: BestellübersichtPage
            },
            //signed in
            {
                path: '/order-details/:orderId',
                name: 'BestelldetailsPage',
                component: BestelldetailsPage,
                meta: { requiresAuth: true }

            },
            //signed in
            {
                path: '/order-edit/:orderId',
                name: 'BestellBearbeitenPage',
                component: BestellBearbeitenPage,
                meta: { requiresAuth: true }
            },
            // mollnhof
            {
                path: '/equipment-edit',
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

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (!auth.state.isLoggedIn) {
            next({
                path: '/',
                query: { redirect: to.fullPath }
            })
        } else {
            next()
        }
    } else if (to.matched.some(record => record.meta.requiresGuest)) {
        if (auth.state.isLoggedIn) { //already singed in ant is trying to relogin-> redirect to homepage
            next({
                path: '/orders-list',
            })
        }
        else
            next()
    }
    else
        next() // make sure to always call next()!
})

export default router
