import React, {useEffect} from "react";
import { Form, Input, InputNumber, Button } from "antd";
import {useEditMealMutation} from "./meal-edit-query.js";

export function MealEdit({ initialValues, onSubmit }) {
    const {mutate, isPending, isSuccess} = useEditMealMutation();
    const onFinish = (values) => mutate({...values, id: initialValues.id});
    useEffect(() => {
        if(isSuccess) {
            onSubmit?.();
        }
    }, [isSuccess])
    return (
        <Form
            layout="vertical"
            initialValues={initialValues}
            onFinish={onFinish}
            autoComplete="off"
        >
            <Form.Item
                label="Продукт"
                name="product"
                rules={[{ required: true, message: "Введите название продукта" }]}
            >
                <Input placeholder="Например, яблоко" />
            </Form.Item>

            <Form.Item
                label="Объём (гр/мл)"
                name="volume"
                rules={[{ required: true, message: "Укажите объём" }]}
            >
                <InputNumber min={0} step={1} />
            </Form.Item>

            <Form.Item label="Ккал" name="kcal">
                <InputNumber />
            </Form.Item>

            <Form.Item label="Белки" name="proteins">
                <InputNumber />
            </Form.Item>

            <Form.Item label="Жиры" name="fats">
                <InputNumber />
            </Form.Item>

            <Form.Item label="Углеводы" name="carbohydrates">
                <InputNumber />
            </Form.Item>

            <Form.Item>
                <Button type="primary" htmlType="submit" loading={isPending} block>
                    Сохранить
                </Button>
            </Form.Item>
        </Form>
    );
}