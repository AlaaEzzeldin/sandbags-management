import Vue from 'vue'
import VueRouter from 'vue-router'

// import pages
import Bestellungliste from '../views/Bestellungliste.vue'
import Konto from '../views/Konto.vue'
import NeueBestellung from '../views/NeueBestellung.vue'
import Bestellübersicht from '../views/Bestellübersicht.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'NeueBestellung',
    component: NeueBestellung
  },
  {
    path: '/konto',
    name: 'Konto',
    component: Konto
  },
  {
    path: '/bestellungliste',
    name: 'Bestellungliste',
    component: Bestellungliste
  },
  {
    path: '/bestellübersicht',
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
