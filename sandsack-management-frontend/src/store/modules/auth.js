//import authAPI from '../../api/auth';
import AuthService from '../../services/auth.service';


const state= {
    isLoggedIn: !!JSON.parse(localStorage.getItem('access_token')),
    user :JSON.parse(localStorage.getItem('user'))
}

const getters = {
    getCurrentUserRole(state) {
        return state.user.role;
    },
}
   const actions= {
       login({commit}, user) {
           return AuthService.login(user).then(
               user => {
                   commit('loginSuccess', user);
                   return Promise.resolve(user);
               },
               error => {
                   commit('loginFailure');
                   return Promise.reject(error);
               }
           );
       },
       logout({commit}) {
           AuthService.logout();
           commit('logout');
       },
       register({commit}, user) {
           return AuthService.register(user).then(
               response => {
                   commit('registerSuccess');
                   return Promise.resolve(response.data);
               },
               error => {
                   commit('registerFailure');
                   return Promise.reject(error);
               }
           );
       }
   }
        const mutations= {
        loginSuccess(state) {
            state.loggedIn = true;
        },
        loginFailure(state) {
            state.loggedIn = false;
        },
        logout(state) {
            state.loggedIn = false;
        },
        registerSuccess(state) {
            state.loggedIn = false;
        },
        registerFailure(state) {
            state.loggedIn = false;
        }
    }

export default {
    state,
    actions,
    mutations,
    getters
}