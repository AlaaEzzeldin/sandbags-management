import AuthService from '../../services/auth.service';


const state = {
    isLoggedIn: !!JSON.parse(localStorage.getItem('user')),
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
    updatePassword({commit}, payload) {
        return AuthService.updatePassword(payload).then(
          data => {
              commit('UPDATE_PASSWORD_SUCCESS', data);
              return Promise.resolve(data);
          },
          error => {
              commit('UPDATE_PASSWORD_FAILURE');
              return Promise.reject(error);
          }
        );
    }
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
}

export default {
    state,
    actions,
    mutations,
    getters
}
