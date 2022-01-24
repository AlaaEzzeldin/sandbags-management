import api from './api';

export default {
    login( data ){
        return api.post('users/login', data );
    }
}
