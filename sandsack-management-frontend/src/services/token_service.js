class Token_service {
    getLocalRefreshToken() {
        const user = JSON.parse(localStorage.getItem("user"));
        return user?.refresh_token;
    }

    getLocalAccessToken() {
        const user = JSON.parse(localStorage.getItem("user"));
        return user?.access_token;
    }

    updateLocalAccessToken(response) {
        let user = JSON.parse(localStorage.getItem("user"));
        console.log("old user",user)
        user.access_token = response.data.access_token;
        user.refresh_token = response.data.refresh_token;
        console.log("new user",user)
        localStorage.setItem("user", JSON.stringify(user));
    }

    getUser() {
        return JSON.parse(localStorage.getItem("user"));
    }

    setUser(user) {
        localStorage.setItem("user", JSON.stringify(user));
    }

    removeUser() {
        localStorage.removeItem("user");
    }
}

export default new Token_service();