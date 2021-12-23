import axios from 'axios'
export default {
    index( params ){
        return axios.get( 'http://localhost:3001/equipment', {
            params: params
        })
    },

    update( id, data ){
        return axios.patch( 'http://localhost:3001/equipment/'+id, data);
    },

}
