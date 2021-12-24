import axios from 'axios';

const API_URL = 'http://46.101.254.157:8000/';

class AuthService {
    login(user) {
        return axios
            .post(API_URL + 'users/login', {
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

    /*register(user) {
        return axios.post(API_URL + 'admin/create_user', {
            username: user.username,
            email: user.email,
            password: user.password

        });
    }*/
}

export default new AuthService();
