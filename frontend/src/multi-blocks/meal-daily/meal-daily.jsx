import React from 'react';
import {Card} from "antd";
import {Daily} from "../../blocks/daily/daily.jsx";
import {MealAdd} from "../../blocks/meal-add/meal-add.jsx";


export function MealDaily(){
    return (
        <>
            <Card title='прием пищи'>
                <MealAdd />
            </Card>
            <Card title='сводка за сегодня'>
                <Daily />
            </Card>
        </>
    )
}