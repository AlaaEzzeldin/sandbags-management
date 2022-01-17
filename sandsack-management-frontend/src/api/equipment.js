import axios from 'axios'
import authHeader from "../services/auth-header";
import {BASE_URL} from "./config";

export default {
    index( params ){
        return axios.get( BASE_URL +'core/equipment', {
            headers: authHeader(),
            params: params
        })
    },

    update( id, data ){
        return axios.patch( BASE_URL + 'core/equipment/return', data, {
            headers: authHeader()
        });
    },

}
