import axios from "axios";

// База берётся из .env (Vite): VITE_API_BASE_URL="http://localhost:8080"
// Если не задано — используем текущий origin.
export const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL || "",
    // можно добавить timeout и т.п.
});

// Добавляем токен в каждый запрос (заголовок X-Auth)
api.interceptors.request.use((config) => {
    // Позволяем точечно отключать авторизацию: api.get(url, { skipAuth: true })
    if (config.skipAuth) return config;

    const token = localStorage.getItem("authToken");
    if (token) {
        config.headers = config.headers || {};
        config.headers["X-Auth"] = token;
    }
    return config;
});

// Глобальная обработка 401 → редирект на /login
api.interceptors.response.use(
    (res) => res,
    (error) => {
        const status = error?.response?.status;
        if (status === 401) {
            // чистим токен и уходим на /login
            localStorage.removeItem("authToken");
            if (typeof window !== "undefined" && window.location.pathname !== "/web/login") {
                window.location.assign("/web/login");
            }
        }
        return Promise.reject(error);
    }
);