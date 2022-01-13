import axios from 'axios'
import authHeader from "../services/auth-header";
import {BASE_URL} from "../api/config";

export default {
    index(params) {
        return axios.get(BASE_URL + 'order', {
            headers: authHeader(),
            params: params
        })
    },
    show(id) {
        return axios.get('http://localhost:3001/order/' + id);
    },

    update(id, data) {
        return axios.patch('http://localhost:3001/orders/' + id, data);
    },

    create(data) {
        return axios.post(BASE_URL + 'order', data, {
            headers: authHeader(),
        });
    },

    delete(id) {
        return axios.delete('http://localhost:3001/orders/' + id)
    }
}