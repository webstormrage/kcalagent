import React from 'react';
import {Skeleton} from "antd";
import {MealEdit} from "../../blocks/meal-edit/meal-edit.jsx";
import {useMealQuery} from "./meal-single-query.js";
import {useParams} from "react-router-dom";

export function MealEditSingle(){
    const { mealId } = useParams();
    const {data, isPending, isError} = useMealQuery(mealId);
    if(isPending || isError) {
        return (
            <>
                <Skeleton active paragraph={false} title={{ width: 240 }} style={{ marginBottom: 16 }} />

                {/* Поля формы */}
                <Skeleton.Input active size="small" style={{ width: 140, marginBottom: 8 }} />
                <Skeleton.Input active style={{ width: "100%", height: 40, marginBottom: 16 }} />

                <Skeleton.Input active size="small" style={{ width: 180, marginBottom: 8 }} />
                <Skeleton.Input active style={{ width: "100%", height: 40, marginBottom: 16 }} />

                <Skeleton.Input active size="small" style={{ width: 80, marginBottom: 8 }} />
                <Skeleton.Input active style={{ width: "100%", height: 40, marginBottom: 16 }} />

                <Skeleton.Input active size="small" style={{ width: 80, marginBottom: 8 }} />
                <Skeleton.Input active style={{ width: "100%", height: 40, marginBottom: 16 }} />

                <Skeleton.Input active size="small" style={{ width: 110, marginBottom: 8 }} />
                <Skeleton.Input active style={{ width: "100%", height: 40, marginBottom: 16 }} />

                {/* Кнопка */}
                <Skeleton.Button active style={{ width: "100%", height: 40 }} />
            </>
        )
    }
    return (
        <MealEdit
            initialValues={{
                id: data.id,
                product: data.product,
                volume: data.volume,
                kcal: data.kcal,
                proteins: data.proteins,
                fats: data.fats,
                carbohydrates: data.carbohydrates,
            }}
        />
    );
}