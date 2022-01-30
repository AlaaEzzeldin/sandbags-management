import axios from "axios";

const instance = axios.create({
    baseURL: "http://18.193.112.34:8001/",
    headers: {
        "Content-Type": "application/json",
    },
});

export default instance;