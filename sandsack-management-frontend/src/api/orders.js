import api from '../services/api';

export default {

    index(params) {
        return api.get( 'order/', {
            params: params
        })
    },

    show(id) {
        return api.get( 'order/?id=' + id)
    },


    create(data) {
        return api.post( 'order/', data);
    },

    acceptOrder(id) {
        return api.post( 'order/accept?id=' + id, {});
    },

    cancelOrder(id) {
        return api.post( 'order/cancel?id=' + id, {})
    },
    confirmOrderDelivery(id) {
        return api.post( 'order/delivery/confirm?id=' + id, {});
    },
    dispatchOrder(id) {
        return api.post( 'order/dispatch?id=' + id, {});
    },
    commentOrder(data) {
        return api.post( 'order/comment', data);
    },
    editOrder(data) {
        return api.patch( 'order/edit', data)
    }
}
