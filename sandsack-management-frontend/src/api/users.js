import axios from 'axios'
import authHeader from "../services/auth-header";
import {BASE_URL} from "../api/config";

export default {
    index(params) {
        return axios.get(BASE_URL + 'users', {
            headers: authHeader(),
            params: params
        })
    },

    show() {
        return axios.get(BASE_URL + 'users/me', {
            headers: authHeader(),
        })
    },

    update(data) {
        console.log("userinfo", data)
        return axios.patch(BASE_URL + 'users/me', data, {
            headers: authHeader()
        })
    }

}