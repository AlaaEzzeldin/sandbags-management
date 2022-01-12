import Vue from 'vue'
import Vuex from 'vuex'
import users from './modules/users'
import orders from './modules/orders'
import equipment from './modules/equipment'
import statisticschart from './modules/statisticschart'


Vue.use(Vuex)

export default  new Vuex.Store({
<<<<<<< HEAD
    modules: {
        orders,
        users,
        equipment,
        statisticschart,


    }
=======
  modules: {
    orders,
    users,
    equipment,
    statisticschart,


  }
>>>>>>> 285d454f5d7d37370b314a62a759391e4489de7c
})
