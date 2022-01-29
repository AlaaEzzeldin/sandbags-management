import axiosInstance from "./api";
import TokenService from "./token_service";

const setup = (store) => {
    axiosInstance.interceptors.request.use(
        (config) => {
            const token = TokenService.getLocalAccessToken();
            if (token) {
                config.headers["Authorization"] = 'Bearer ' + token;
            }
            return config;
        },
        (error) => {
            return Promise.reject(error);
        }
    );

    axiosInstance.interceptors.response.use(
        (res) => {
            return res;
        },
        async (err) => {

            const originalConfig = err.config;

            if (originalConfig.url !== "users/login" && err.response) {
                // Access Token was expired
                if (err.response.status === 401 && !originalConfig._retry) {
                    originalConfig._retry = true;
                    try {
                        const rs = await axiosInstance.post("users/refresh", {
                            refresh_token: TokenService.getLocalRefreshToken(),
                        });
                        const {accessToken} = rs.data;
                        store.dispatch('REFRESH_TOKEN', accessToken);
                        TokenService.updateLocalAccessToken(accessToken);

                        return axiosInstance(originalConfig);
                    } catch (_error) {
                        return Promise.reject(_error);
                    }
                }
            }
            return Promise.reject(err).catch(err => {
                if (err.toJSON().message=="Network Error")
                    alert('Please check your internet connection, for emergency please contact contact:')})
        }
    );
};

export default setup;