import api from '../services/api';

export default {
    login( data ){
        return api.post('users/login', data );
    }
}
