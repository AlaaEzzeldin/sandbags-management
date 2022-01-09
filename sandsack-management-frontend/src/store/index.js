import Vue from 'vue'
import Vuex from 'vuex'
import users from './modules/users'
import orders from './modules/orders'
import equipment from './modules/equipment'
import statisticschart from './modules/statisticschart'


Vue.use(Vuex)

export default  new Vuex.Store({
  modules: {
    orders,
    users,
    equipment,
    statisticschart,


  }
})
