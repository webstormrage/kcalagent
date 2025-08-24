import { useMutation, useQueryClient } from "@tanstack/react-query";
import { App } from "antd";
import { api } from "../../network/client";

export function useAddMealMutation() {
    const { notification } = App.useApp();
    const qc = useQueryClient();

    return useMutation({
        mutationFn: async ({ product, volume, genAiToken }) => {
            localStorage.setItem("genAiToken", genAiToken);
            // Если бекенд ждёт токен В ТЕЛЕ запроса:
            const { data } = await api.post("/meals/add", {
                product,
                volume,
                genAiToken,
            });
            console.log(data)

            return data;
        },

        onSuccess: (data) => {
            notification.success({
                message: "Прием пищи добавлен",
                description: `${data.product}: ${data.kcal.toFixed(1)} ккал ${data.proteins.toFixed(1)} б. ${data.fats.toFixed(1)} ж. ${data.carbohydrates.toFixed(1)} у.`,
            });
        },
        onError: (err) => {
            const serverMsg = err?.response?.data ?? "Неизвестная ошибка";
            notification.error({ message: "Ошибка", description: serverMsg });
        },
    });
}