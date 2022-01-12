import axios from 'axios'
import authHeader from "../services/auth-header";

export default {
    index( params ){
        return axios.get( 'http://18.193.112.34:8001/core/equipment', {
            headers: authHeader(),
            params: params
        })
    },

    update( id, data ){
        return axios.patch( 'https://localhost:3001/equipment/'+id, data);
    },

}
