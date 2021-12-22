import axios from 'axios'

const API_URL = 'http://46.101.254.157:8000/';

export default {
    login( data ){
        return axios.post( API_URL + 'users/login', data );
    }
}