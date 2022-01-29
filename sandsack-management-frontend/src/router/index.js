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
import HelpPage from "../views/HelpPage";

Vue.use(VueRouter)


const routes = [
    {
        path: '/',
        name: 'LoginPage',
        component: LoginPage,
        meta: {requiresGuest: true}
    },
    {
        path: '/auth/sign-up',
        name: 'SignupPage',
        component: SignupPage,
        meta: {requiresGuest: true}
    },
    {
        path: '/auth/recover-password',
        name: 'RecoverPasswordPage',
        component: RecoverPasswordPage,
        meta: {requiresGuest: true}
    },
    {
        path: '/auth/recover-password-confirm',
        name: 'RecoverPasswordConfirmPage',
        component: RecoverPasswordConfirmPage,
        meta: {requiresGuest: true}
    },
    {
        path: '/homepage',
        name: 'HomePage',
        component: HomePage,
      //  meta: {requiresAuth: true},
        children: [
            {
                path: '/orders-list',
                name: 'BestellungslistePage',
                component: BestellungslistePage,
            },
            {
                path: '/account',
                name: 'KontoPage',
                component: KontoPage,
            },
            // unterabschnit only
            {
                path: '/new-order',
                name: 'NeueBestellungPage',
                component: NeueBestellungPage,

            },
            // einsatz/ haupt and leiter
            {
                path: '/overview',
                name: 'BestellübersichtPage',
                component: BestellübersichtPage
            },
            {
                path: '/order-details/:orderId',
                name: 'BestelldetailsPage',
                component: BestelldetailsPage,

            },
            // all exept for mollnhof
            {
                path: '/order-edit/:orderId',
                name: 'BestellBearbeitenPage',
                component: BestellBearbeitenPage,
            },
            // mollnhof
            {
                path: '/equipment-edit',
                name: 'ManageEquipmentPage',
                component: ManageEquipmentPage
            },
            // all
            {
                path: '/help',
                name: 'HelpPage',
                component: HelpPage,
            },
        ]
    },


]


const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

function hasPermissionsNeeded(to) {
    let userRole = auth.state.user.role
    let toPage = to.name

    if (['Mollnhof'].includes(userRole) &&
        ['ManageEquipmentPage'].includes(toPage)) // only Mollnhof
        return true
    else if (['Unterabschnitt'].includes(userRole) &&
        ['NeueBestellungPage'].includes(toPage)) // only Unterabschnitt
        return true
    else if (['Einsatzabschnitt', 'Hauptabschnitt', 'Einsatzleiter'].includes(userRole) &&
        ['BestellübersichtPage'].includes(toPage)) // only Einsatzabschnitt, Hauptabschnitt, Einsatzleiter
        return true
    else if (['Unterabschnitt', 'Einsatzabschnitt', 'Hauptabschnitt', 'Einsatzleiter'].includes(userRole) &&
        ['BestellBearbeitenPage'].includes(toPage)) // all except for Mollnhof
        return true
    else if (['BestellungslistePage', 'KontoPage', 'BestelldetailsPage', 'HelpPage'].includes(toPage)) // all users
        return true
    else return false
}

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (!auth.state.isLoggedIn) { // unsigned-in user tries to access auhtenticated route
            next({
                path: '/',
                query: {redirect: to.fullPath}
            })
        } else {
            if (!hasPermissionsNeeded(to)) { //the user has no access to this page so will be redirected to homepage
                alert("you don't have access to this page")
                next('/orders-list');
            } else {
                next();
            }
        }
    } else if (to.matched.some(record => record.meta.requiresGuest)) {
        if (auth.state.isLoggedIn) { // already singed in and is trying to re-login-> redirect to homepage
            next({
                path: '/orders-list',
            })
        } else
            next()
    } else
        next()
})

export default router
