import axios from 'axios'
import authHeader from "../services/auth-header";

export default {
    index( params ){
        return axios.get( 'http://18.193.112.34:8001/core/priority', {
            headers: authHeader(),
            params: params
        })
    },

}
