import { useQuery } from "@tanstack/react-query";
import { api } from "../../network/client.js";
import {App} from "antd";

/**
 * Запрашивает суточную сводку.
 * Возвращает нормализованные поля: kcal, proteins, fats, carbohydrates (числа).
 */
export function useDailyQuery() {
    const { notification } = App.useApp();
    return useQuery({
        queryFn: async () => {
            const { data } = await api.get("/get-daily-summary");
            return data
        },
        staleTime: 30_000,
        refetchOnWindowFocus: false,
        onError: (err) => {
            // покрасивее достанем сообщение с сервера, если есть
            const serverMsg = err?.response?.data ?? "Неизвестная ошибка";
            notification.error({ message: "Ошибка", description: serverMsg});
        },
    });
}