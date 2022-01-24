import api from '../services/api';

export default {
    index( params ){
        return api.get('core/priority', {
            params: params
        })
    },

}
