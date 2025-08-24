import { useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { api } from "../../network/client";
import { App } from "antd";

export function useLoginMutation() {
    const navigate = useNavigate();
    const { notification } = App.useApp();

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
            navigate("/summary", { replace: true }); // редирект в индексный роут
        },

        onError: (err) => {
            // покрасивее достанем сообщение с сервера, если есть
            const serverMsg = err?.response?.data ?? "Неизвестная ошибка";
            notification.error({ message: "Ошибка", description: serverMsg});
        },
    });
}