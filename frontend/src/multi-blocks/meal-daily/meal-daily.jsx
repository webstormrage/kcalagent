import React, {useState} from 'react';
import {Card} from "antd";
import {Daily} from "../../blocks/daily/daily.jsx";
import {MealAdd} from "../../blocks/meal-add/meal-add.jsx";
import {MealEdit} from "../../blocks/meal-edit/meal-edit.jsx";
import {Drawer} from "antd";

export function MealDaily(){

    const [open, setOpen] = useState(false);
    const [selectedRow, setSelectedRow] = useState(null);

    const handleRowClick = (record) => {
        setSelectedRow(record);
        setOpen(true);
    };


    const handleClose = () => {
        setOpen(false);
        setSelectedRow(null);
    };

    return (
        <>
            <Card title='прием пищи'>
                <MealAdd />
            </Card>
            <Card title='сводка за сегодня'>
                <Daily onRowClick={handleRowClick}/>
            </Card>
            <Drawer
                title={selectedRow ? `Редактирование: ${selectedRow.name}` : "Редактирование"}
                placement="right"
                width={480}
                open={open}
                onClose={handleClose}
            >
                {selectedRow && (
// ключ нужен, чтобы форма пере-монтировалась при выборе другой строки,
// так как initialValues в antd Form задаются только при первом маунте.
                    <MealEdit
                        key={selectedRow.id}
                        initialValues={{
                            id: selectedRow.id,
                            product: selectedRow.name,
                            volume: selectedRow.weight,
                            kcal: selectedRow.kcal,
                            proteins: selectedRow.proteins,
                            fats: selectedRow.fats,
                            carbohydrates: selectedRow.carbohydrates,
                        }}
                        onSubmit={handleClose}
                    />
                )}
            </Drawer>
        </>
    )
}