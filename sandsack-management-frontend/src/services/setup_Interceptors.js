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
        (error) => {return Promise.reject(error);
        }
    );

    axiosInstance.interceptors.response.use(
        (res) => {
            return res;
        },
        async (err) => {
            const error = err.toJSON()
            console.log("error", error)
            const originalConfig = error.config;
            if (originalConfig.url !== "users/login" && error.message) {
                // Access Token was expired
                console.log("Access Token was expired, check if refresh should be done", error.status == 401 && !originalConfig._retry)
                if (error.status === 401 && !originalConfig._retry) {
                    originalConfig._retry = true;

                    try {
                        console.log("do a refresh token request")
                        const rs = await axiosInstance.post("users/refresh", {
                            refresh_token: TokenService.getLocalRefreshToken(),
                        });
                        console.log("response of refresh request", rs)
                        const  access_token  = rs.data.access_token;
                        const  refresh_token     = rs.data.refresh_token;
                        console.log("new accessToken", access_token)
                        store.dispatch('refreshToken', refresh_token);
                        TokenService.updateLocalAccessToken(rs);

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