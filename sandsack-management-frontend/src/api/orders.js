import axios from 'axios'
import authHeader from "../services/auth-header";
import {BASE_URL} from "../api/config";

export default {

    index(params) {
        return axios.get(BASE_URL + 'order/', {
            headers: authHeader(),
            params: params
        })
    },

    show(id) {
        return axios.get(BASE_URL + 'order/?id=' + id, {
            headers: authHeader(),
        })
    },


    create(data) {
        return axios.post(BASE_URL + 'order/', data, {
            headers: authHeader(),
        });
    },

    acceptOrder(id) {
        return axios.post(BASE_URL + 'order/accept?id=' + id, {}, {
            headers: authHeader(),
        });
    },

    cancelOrder(id) {
        return axios.post(BASE_URL + 'order/cancel?id=' + id, {}, {
            headers: authHeader(),
        })
    },
    confirmOrderDelivery(id) {
        return axios.post(BASE_URL + 'order/delivery/confirm?id=' + id, {}, {
            headers: authHeader(),
        });
    },
    dispatchOrder(id) {
        return axios.post(BASE_URL + 'order/dispatch?id=' + id, {}, {
            headers: authHeader(),
        });
    },
    commentOrder(data) {
        return axios.post(BASE_URL + 'order/comment', data, {
            headers: authHeader(),
        });
    },
    editOrder(data) {
        return axios.patch(BASE_URL + 'order/edit', data, {
            headers: authHeader()
        })
    },

    showStatistics(data) {
        return axios.get(BASE_URL + 'order/stats', {
            headers: authHeader(),
            params: data
        })
    },
}
