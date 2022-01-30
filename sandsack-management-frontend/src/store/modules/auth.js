import AuthService from '../../services/auth_service';


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
            response => {
                commit('LOGIN_SUCCESS', response.data);
                return Promise.resolve(response.data);
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
    refreshToken({ commit }, accessToken) {
        commit('REFRESH_TOKEN', accessToken);
    },
    updatePassword({commit}, payload) {
        return AuthService.updatePassword(payload).then(
          data => {
              console.log(commit);
              return Promise.resolve(data);
          },
          error => {
              console.log(commit);
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
    REFRESH_TOKEN(state, accessToken) {
        state.isLoggedIn = true;
        state.user = { ...state.user, accessToken: accessToken };
    }
}

export default {
    state,
    actions,
    mutations,
    getters
}
