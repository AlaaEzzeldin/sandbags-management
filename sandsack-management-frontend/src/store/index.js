import Vue from 'vue'
import Vuex from 'vuex'
import users from './modules/users'
import orders from './modules/orders'
import equipment from './modules/equipment'
import auth from './modules/auth'

Vue.use(Vuex)

export default  new Vuex.Store({
    modules: {
        orders,
        users,
        auth,
        equipment
    }
})
