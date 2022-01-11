import axios from 'axios'
import {BASE_URL} from "./config";

export default {
    login( data ){
        return axios.post( BASE_URL + 'users/login', data );
    }
}
