import axios from 'axios';
import {BASE_URL} from "../api/config";
import authHeader from "./auth-header";

class AuthService {
    login(user) {
        return axios
            .post(BASE_URL + 'users/login', {
                email: user.email,
                password: user.password
            })
            .then(response => {
                if (response.data.access_token) {
                    localStorage.setItem('user', JSON.stringify(response.data));
                }

                return response;
            });
    }

    logout() {
        localStorage.removeItem('user');
    }

    updatePassword(data) {
        return axios
          .patch(BASE_URL + 'users/change_password', data, {
              headers: authHeader(),
          })
    }

    /*register(user) {
        return axios.post(API_URL + 'admin/create_user', {
            username: user.username,
            email: user.email,
            password: user.password

        });
    }*/
}

export default new AuthService();
