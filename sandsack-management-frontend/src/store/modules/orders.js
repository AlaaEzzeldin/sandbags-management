import ordersAPI from '../../api/orders';

const state = {
    orders: [],
    order: '',
    statistics:[]
}

const getters = {
    getOrders(state) {
        return state.orders;
    },
    getOrder(state) {
        return state.order;
    },
    getStatistics(state) {
        return state.statistics;
    },
}

const actions = {
    loadOrders({commit}) {
        ordersAPI.index()
            .then(function (response) {
                commit('SET_ORDERS', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    loadOrder({commit}, id) {
        ordersAPI.show(id)
            .then(function (response) {
                commit('SET_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    createOrder({commit}, payload) {
        ordersAPI.create(payload)
            .then(function (response) {
                commit('ADD_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    editOrder({commit}, payload) {
        ordersAPI.editOrder(payload)
            .then(function (response) {
                commit('UPDATE_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    commentOrder({commit}, payload) {
        ordersAPI.commentOrder(payload)
            .then(function (response) {
                commit('UPDATE_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },

    acceptOrder({commit}, id) {
        ordersAPI.acceptOrder(id)
            .then(function (response) {
                commit('UPDATE_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },

    cancelOrder({commit}, id) {
        ordersAPI.cancelOrder(id)
            .then(function (response) {
                commit('UPDATE_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },

    confirmOrderDelivery({commit}, id) {
        ordersAPI.confirmOrderDelivery(id)
            .then(function (response) {
                commit('UPDATE_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    dispatchOrder({commit}, id) {
        ordersAPI.dispatchOrder(id)
            .then(function (response) {
                commit('UPDATE_ORDER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    loadStatistics({commit}, data) {
        ordersAPI.showStatistics(data)
            .then(function (response) {
                commit('SET_STATISTICS', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    }

}

const mutations = {
    SET_ORDERS(state, orders) {
        if(orders)
            state.orders = orders;
        else state.orders= []
    },
    SET_ORDER(state, order) {
        state.order = order;
    },
    ADD_ORDER(state, order) {
        state.orders.unshift(order);
    },
    UPDATE_ORDER(state, updatedOrder) {
        const index = state.orders.findIndex(order => {
            return (updatedOrder.id === order.id)
        })
        state.orders.splice(index, 1, updatedOrder)
        state.order = updatedOrder
    },
    SET_STATISTICS(state, statistics) {
        state.statistics = statistics;
    }
}

export default {
    state,
    getters,
    actions,
    mutations
}