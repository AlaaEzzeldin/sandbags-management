//import authAPI from '../../api/auth';
import AuthService from '../../services/auth.service';


const state = {
    isLoggedIn: !!JSON.parse(localStorage.getItem('access_token')),
    user: JSON.parse(localStorage.getItem('user'))
}

const getters = {
    getCurrentUserRole(state) {
        return state.user.role;
    },
    getCurrentUserName(state) {
        return state.user.name;
    },
    isLoggedIn(state) {
        return state.isLoggedIn;
    },
}
const actions = {
    login({commit}, user) {
        return AuthService.login(user).then(
            user => {
                commit('LOGIN_SUCCESS', user.data);
                return Promise.resolve(user);
            },
            error => {
                commit('LOGIN_FAILURE');
                return Promise.reject(error);
            }
        );
    },
    logout({commit}) {
        AuthService.logout();
        commit('LOGOUT');
    },
    refreshToken({commit}, accessToken) {
        commit('REFRESH_TOKEN', accessToken);
    },

}
const mutations = {
    LOGIN_SUCCESS(state, user) {
        state.isLoggedIn = true;
        state.user = user
    },
    LOGIN_FAILURE(state) {
        state.isLoggedIn = false;
        state.user = null;
    },
    LOGOUT(state) {
        state.isLoggedIn = false;
        state.user = null;
    },
    REFRESH_TOKEN(state, accessToken) {
        state.status.isLoggedIn = true;
        state.user = {...state.user, accessToken: accessToken};
    }
}

export default {
    state,
    actions,
    mutations,
    getters
}