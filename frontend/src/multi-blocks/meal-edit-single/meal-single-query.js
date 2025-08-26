import { useQuery } from "@tanstack/react-query";
import { api } from "../../network/client.js";
import { App } from "antd";

/**
 * Запрашивает один meal по id.
 * Возвращает нормализованные поля: id, product, volume, kcal, proteins, fats, carbohydrates.
 */
export function useMealQuery(mealId) {
    const { notification } = App.useApp();

    return useQuery({
        queryKey: ["meal", mealId],
        enabled: mealId != null, // не дергаем запрос без id
        queryFn: async () => {
            const { data } = await api.get("/meal", { params: { id: mealId } });
            return data;
        },
        select: (data) => ({
            id: Number(data?.id ?? 0),
            product: data?.product ?? "",
            volume: Number(data?.volume ?? 0),
            kcal: Number(data?.kcal ?? 0),
            proteins: Number(data?.proteins ?? 0),
            fats: Number(data?.fats ?? 0),
            carbohydrates: Number(data?.carbohydrates ?? 0),
        }),
        staleTime: 30_000,
        refetchOnWindowFocus: false,
        onError: (err) => {
            const serverMsg = err?.response?.data ?? "Неизвестная ошибка";
            notification.error({ message: "Ошибка", description: serverMsg });
        },
    });
}