import axios from 'axios'

export default {
    index( params ){
        return axios.get( 'http://localhost:3001/users', {
            params: params
        })
    },

    show( id ){
        return axios.get( 'http://localhost:3001/users/'+id );
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