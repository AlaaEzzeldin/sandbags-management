import axios from 'axios'
import authHeader from "../services/auth-header";
//import {BASE_URL} from "../api/config";

export default {
    index( params ){
        return axios.get( 'http://18.193.112.34:8001/users', {
            headers: authHeader(),
            params: params
        })
    },

    show() {
        return axios.get('http://18.193.112.34:8001/users/me', {
            headers: authHeader(),
        })
    },

    update( id, data ){
        return axios.put( 'http://localhost:3001/users/'+id, data );
    },

    create( data ){
        return axios.post( 'http://localhost:3001/users/', data );
    },

    delete( id ){
        return axios.delete( 'http://localhost:3001/users/' + id )
    }
}