import React from "react";
import { Card, Table } from "antd";
import {useDailyQuery} from "./daily-query.js";

const fmt = (v) => Number(v).toFixed(2);

const columns = [
    { title: "Продукт",     dataIndex: "name",           key: "kcal"},
    { title: "Ккал",        dataIndex: "kcal",           key: "kcal",           render: fmt },
    { title: "Белки",       dataIndex: "proteins",       key: "proteins",       render: fmt },
    { title: "Жиры",        dataIndex: "fats",           key: "fats",           render: fmt },
    { title: "Углеводы",    dataIndex: "carbohydrates",  key: "carbohydrates",  render: fmt },
];

export function Daily() {
    const { data: summary, isLoading } = useDailyQuery();
    const rows = summary
        ? summary.map(((d, i) => ({...d, key: i})))
        : [];

    return (
        <Card title="Сводка за сегодня">
            <Table
                columns={columns}
                dataSource={rows}
                pagination={false}
                loading={isLoading}
                locale={{ emptyText: isLoading ? "Загрузка…" : "Нет данных" }}
            />
        </Card>
    );
}