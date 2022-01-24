import api from './api';

export default {
    index(params) {
        return api.get('users', {
            params: params
        })
    },

    show() {
        return api.get('users/me', {
        })
    },

    update(data) {
        return api.patch('users/me', data, {
        })
    }

}