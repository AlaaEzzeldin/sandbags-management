import ordersAPI from '../../api/orders';

const state = {
    orders: [],
    order:''
}

const getters = {
    getOrders(state) {
        return state.orders;
    },
    getOrder(state) {
        return state.order;
    },
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
    },
    loadOrder({ commit }, id) {
        ordersAPI.show(id)
            .then(function (response) {
                commit('SET_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    createOrder({ commit }, payload) {
        ordersAPI.create(payload)
            .then(function (response) {
                commit('ADD_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    updateOrder({ commit }, payload) {
        ordersAPI.update( payload.id,  payload.data)
            .then(function (response) {
                commit('UPDATE_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },

}

const mutations = {
    SET_ORDERS(state, orders) {
        state.orders = orders;
    },
    SET_ORDER(state, order) {
        state.order = order;
    },
    ADD_ORDER(state, order) {
        state.orders.push(order);
    },
    UPDATE_ORDER(state, updatedOrder) {
        const index = state.orders.findIndex(order => {
            return (updatedOrder.id === order.id)
        })
        state.orders.splice(index, 1, updatedOrder)
    },
}

export default {
    state,
    getters,
    actions,
    mutations
}