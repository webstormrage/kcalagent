import { useQuery } from "@tanstack/react-query";
import { api } from "../../network/client.js";

/**
 * Запрашивает суточную сводку.
 * Возвращает нормализованные поля: kcal, proteins, fats, carbohydrates (числа).
 */
export function useDailyQuery() {
    return useQuery({
        queryFn: async () => {
            const { data } = await api.get("/get-daily-summary");
            const s = data || {};
            return {
                kcal: Number(s.kcal ?? 0),
                proteins: Number(s.proteins ?? 0),
                fats: Number(s.fats ?? 0),
                carbohydrates: Number(s.carbohydrates ??  0),
            };
        },
        staleTime: 30_000,
        refetchOnWindowFocus: false,
    });
}