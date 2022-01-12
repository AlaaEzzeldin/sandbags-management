import axios from 'axios'
<<<<<<< HEAD
=======

>>>>>>> 285d454f5d7d37370b314a62a759391e4489de7c
export default {
    index( params ){
        return axios.get( 'http://localhost:3001/statisticschart', {
            params: params
        })
    }
}
