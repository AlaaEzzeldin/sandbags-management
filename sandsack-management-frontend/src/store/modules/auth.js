import authAPI from '../../api/auth';

const state = {
    userRole: ''
}

const getters = {
    getUserRole(state) {
        return state.userRole;
    },
}

const actions = {
    login({commit}, payload) {
        authAPI.login(payload)
            .then(function (response) {
                    console.log("response", response)
                    commit('SET_USER_ROLE', response.role);
            })
            .catch(function (error) {
                console.log(error);
            });
    },

}

const mutations = {

    SET_USER_ROLE(state, role) {
        state.userRole = role;
    }
}

export default {
    state,
    getters,
    actions,
    mutations
}