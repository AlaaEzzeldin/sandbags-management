import api from "./api";
import TokenService from "./token_service";

class Auth_service {
    login(user) {
        return api
            .post('users/login', {
                email: user.email,
                password: user.password
            })
            .then(response => {
                if (response.data.access_token) {
                    TokenService.setUser(response.data);
                }
                return response;
            });
    }

    logout() {
        TokenService.removeUser();
    }

    updatePassword(data) {
        return api
          .patch('users/change_password', data)
    }
}

export default new Auth_service();
