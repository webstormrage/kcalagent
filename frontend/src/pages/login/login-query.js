import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { message } from "antd";
import { api } from "../../network/client";

export function useLoginMutation() {
    const navigate = useNavigate();

    return useMutation({
        // Важно: пропускаем добавление X-Auth у логина (skipAuth: true)
        mutationFn: async ({ login, password }) => {
            const { data } = await api.post(
                "/auth/login",
                { login, password },
                { skipAuth: true }
            );

            if (data?.token) {
                localStorage.setItem("authToken", data.token);
            }
            return data;
        },

        onSuccess: () => {
            navigate("/", { replace: true }); // редирект в индексный роут
        },

        onError: (err) => {
            // покрасивее достанем сообщение с сервера, если есть
            const serverMsg =
                err?.response?.data?.message ||
                (typeof err?.response?.data === "string" ? err.response.data : "") ||
                err?.message;
            message.error(serverMsg || "Не удалось войти");
        },
    });
}