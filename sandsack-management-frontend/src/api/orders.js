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
        return axios.post(BASE_URL + 'order/accept' + id, {
            headers: authHeader(),
        });
    },
    cancelOrder(id) {
        return axios.post(BASE_URL + 'order/cancel' + id, {
            headers: authHeader(),
        });
    },
    confirmOrderDelivery(id) {
        return axios.post(BASE_URL + 'order/delivery/confirm' + id, {
            headers: authHeader(),
        });
    },
    dispatchOrder(id) {
        return axios.post(BASE_URL + 'order/dispatch' + id, {
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
    }
}
