import axios from 'axios'

const API_URL = 'http://18.193.112.34:8001/';

export default {
    login( data ){
        return axios.post( API_URL + 'users/login', data );
    }
}