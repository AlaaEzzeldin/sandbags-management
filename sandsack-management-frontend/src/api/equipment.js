import api from './api';

export default {
    index( params ){
        return api.get('core/equipment', {
            params: params
        })
    },

    update( id, data ){
        return api.patch( 'core/equipment/return', data, {
        });
    },

}
