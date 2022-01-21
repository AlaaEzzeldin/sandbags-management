import usersAPI from '../../api/users';

const state = {
    loggedInUser: ''
}

const getters = {
    getLoggedInUser(state) {
        return state.loggedInUser;
    }
}

const actions = {
    getUserInfo({commit}) {
        usersAPI.show()
            .then(function (response) {
                commit('SET_LOGGED_IN_USER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },
    updateUserInfo({commit}, payload) {
        usersAPI.update( payload)
            .then(function (response) {
                commit('SET_LOGGED_IN_USER', response.data);
            })
            .catch(function (error) {
                console.log(error);
            });
    },

}

const mutations = {

    SET_LOGGED_IN_USER(state, user) {
        state.loggedInUser = user;
    }
}

export default {
    state,
    getters,
    actions,
    mutations
}