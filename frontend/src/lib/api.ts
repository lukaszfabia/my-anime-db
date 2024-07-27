import axios from "axios";

const api = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL,
})

api.interceptors.request.use(
    (config) => {
        const token: string | null = ACCESS_TOKEN;
        if (token) {
            config.headers['Authorization'] = `Bearer ${token}`;
        }
        return config;
    }, (error) => Promise.reject(error)
);

export default api;