import axios from 'axios'
import authHeader from "../services/auth-header";

export default {
    index( params ){
        return axios.get( 'http://18.193.112.34:8001/order', {
            headers: authHeader(),
            params: params
        })
    },

    show( id ){
        return axios.get( 'http://localhost:3001/order/'+id );
    },

    update( id, data ){
        return axios.patch( 'http://localhost:3001/orders/'+id, data );
    },

    create( data ){
        return axios.post( 'http://localhost:3001/orders/', data );
    },

    delete( id ){
        return axios.delete( 'http://localhost:3001/orders/' + id )
    }
}