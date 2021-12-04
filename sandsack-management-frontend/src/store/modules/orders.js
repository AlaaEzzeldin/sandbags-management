import ordersAPI from '../../api/orders';

const state = {
    orders: [],
    order:''
}

const getters = {
    getOrders(state) {
        return state.orders;
    }
}

const actions = {
    loadOrders({ commit }) {
        ordersAPI.index()
            .then(function (response) {
                commit('SET_ORDERS', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    }
}

const mutations = {
    SET_ORDERS(state, orders) {
        state.orders = orders;
    },
}

export default {
    state,
    getters,
    actions,
    mutations
}