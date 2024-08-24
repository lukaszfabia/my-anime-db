import axios from "axios";
import { ACCESS_TOKEN } from "./constants";

const api = axios.create({
    baseURL: `${process.env.NEXT_PUBLIC_API_URL}/api`,
})

if (typeof window !== 'undefined') {
    api.interceptors.request.use(
        (config) => {
            const token: string | null = localStorage.getItem(ACCESS_TOKEN) || sessionStorage.getItem(ACCESS_TOKEN);
            if (token) {
                config.headers["Authorization"] = `Bearer ${token}`;
            }
            config.headers["Content-Type"] = "multipart/form-data";
            return config;
        },
        (error) => Promise.reject(error)
    );
}

export default api;