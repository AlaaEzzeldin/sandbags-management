import axios from 'axios'
export default {
    index( params ){
        return axios.get( 'http://localhost:3001/orders', {
            params: params
        })
    },

    show( id ){
        return axios.get( 'http://localhost:3001/orders/'+id );
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