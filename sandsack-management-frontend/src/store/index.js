import Vue from 'vue'
import Vuex from 'vuex'
import users from './modules/users'
import orders from './modules/orders'
import equipment from './modules/equipment'
import auth from './modules/auth'
<<<<<<< HEAD
import statisticschart from './modules/statisticschart'
=======
import priorities from './modules/priorities'
>>>>>>> development

Vue.use(Vuex)

export default  new Vuex.Store({
<<<<<<< HEAD
    modules: {
        orders,
        users,
        auth,
        equipment,
        statisticschart
    }
=======
  modules: {
    orders,
    users,
    auth,
    equipment,
    priorities
  }
>>>>>>> development
})
