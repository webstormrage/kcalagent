import React from "react";
import { Table } from "antd";
import {useDailyQuery} from "./daily-query.js";

const fmt = (v) => Number(v).toFixed(0);

const columns = [
    { title: "Продукт",     dataIndex: "name",           key: "kcal"},
    { title: "Ккал",        dataIndex: "kcal",           key: "kcal",           render: fmt },
    { title: "Белки",       dataIndex: "proteins",       key: "proteins",       render: fmt },
    { title: "Жиры",        dataIndex: "fats",           key: "fats",           render: fmt },
    { title: "Углеводы",    dataIndex: "carbohydrates",  key: "carbohydrates",  render: fmt },
];

export function Daily({onRowClick}) {
    const { data: summary, isLoading } = useDailyQuery();
    const rows = summary
        ? summary.map(((d, i) => ({...d, key: i})))
        : [];

    return (
            <Table
                columns={columns}
                dataSource={rows}
                pagination={false}
                loading={isLoading}
                onRow={(record) => {
                    if(record.id === -1){
                        return null
                    }
                    return { onClick: () => onRowClick(record)};
                }}
                locale={{ emptyText: isLoading ? "Загрузка…" : "Нет данных" }}
            />
    );
}