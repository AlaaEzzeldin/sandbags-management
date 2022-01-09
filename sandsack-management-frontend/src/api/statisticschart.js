import axios from 'axios'

export default {
    index( params ){
        return axios.get( 'http://localhost:3001/statisticschart', {
            params: params
        })
    }
}
