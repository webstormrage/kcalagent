import { useMutation, useQueryClient } from "@tanstack/react-query";
import { App } from "antd";
import { api } from "../../network/client";
import { getRoute, isDesktop, ROUTE_NAMES } from "../../route.js";
import { useNavigate } from "react-router-dom";

export function useEditMealMutation() {
    const { notification } = App.useApp();
    const qc = useQueryClient();
    const navigate = useNavigate();

    return useMutation({
        mutationFn: async (payload) => {
            const { data } = await api.post("/meals/edit", payload);
            return data;
        },

        onSuccess: (data) => {
            notification.success({
                message: "Приём пищи обновлён",
                description: `${data.name}: ${Number(data.kcal).toFixed(1)} ккал • ${Number(data.proteins).toFixed(1)} б • ${Number(data.fats).toFixed(1)} ж • ${Number(data.carbohydrates).toFixed(1)} у`,
                duration: 2,
            });

            // Обновим связанные запросы
            qc.invalidateQueries({ queryKey: ["daily"] });

            // На мобильном — вернёмся к сводке
            if (!isDesktop()) {
                navigate(getRoute(ROUTE_NAMES.SUMMARY), { replace: true });
            }
        },

        onError: (err) => {
            const msg =
                err?.response?.data?.message ||
                (typeof err?.response?.data === "string" ? err.response.data : "") ||
                err?.message ||
                "Не удалось обновить приём пищи";
            notification.error({ message: "Ошибка", description: msg });
        },
    });
}