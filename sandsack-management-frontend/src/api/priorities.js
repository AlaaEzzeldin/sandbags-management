import axios from 'axios'
import authHeader from "../services/auth-header";
import {BASE_URL} from "./config";

export default {
    index( params ){
        return axios.get( BASE_URL + 'core/priority', {
            headers: authHeader(),
            params: params
        })
    },

}
