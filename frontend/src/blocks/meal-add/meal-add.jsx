// src/components/MealAdd.jsx
import React from "react";
import { Button, Col, Form, Input, InputNumber, Row } from "antd";
import { useAddMealMutation } from "./meal-add-query.js";

export function MealAdd() {
    const { mutate, isPending } = useAddMealMutation();
    const onFinish = (values) => mutate(values);
    const initialValues = { genAiToken: localStorage.getItem("genAiToken") || "" };

    return (
        <Form layout="vertical" autoComplete="off" onFinish={onFinish} initialValues={initialValues}>
            {/* 1-я строка: Gen AI токен (всегда на всю ширину) */}
            <Row gutter={16}>
                <Col span={24}>
                    <Form.Item
                        label="Gen AI токен"
                        name="genAiToken"
                        rules={[{ required: true, message: "Введите токен" }]}
                    >
                        <Input placeholder="Введите токен" />
                    </Form.Item>
                </Col>
            </Row>

            {/* 2-я строка: Продукт + Объём
          - на мобильных (xs) каждая колонка = 24 → встанут друг под другом
          - на десктопе (md+) продукт 16, объём 8 */}
            <Row gutter={16}>
                <Col xs={24} md={16}>
                    <Form.Item
                        label="Продукт"
                        name="product"
                        rules={[{ required: true, message: "Введите продукт" }]}
                    >
                        <Input placeholder="Например, яблоко" />
                    </Form.Item>
                </Col>
                <Col xs={24} md={8}>
                    <Form.Item
                        label="Объём (гр/мл)"
                        name="volume"
                        rules={[{ required: true, message: "Укажите объём" }]}
                    >
                        <InputNumber min={0} step={1} />
                    </Form.Item>
                </Col>
            </Row>

            {/* 3-я строка: Кнопка на всю ширину */}
            <Row gutter={16}>
                <Col span={24}>
                    <Form.Item>
                        <Button type="primary" htmlType="submit" loading={isPending} block>
                            Сохранить
                        </Button>
                    </Form.Item>
                </Col>
            </Row>
        </Form>
    );
}