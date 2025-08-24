import React from "react";
import { Form, Input, Button, Card } from "antd";
import { useLoginMutation } from "./login-query.js";

export  function Login() {
    const { mutate, isPending } = useLoginMutation();
    const onFinish = (values) => mutate(values);

    return (
        <Card title="Авторизация">
        <Form layout="vertical" onFinish={onFinish}>
            <Form.Item name="login" label="Логин" rules={[{ required: true }]}>
                <Input autoComplete="username" />
            </Form.Item>
            <Form.Item name="password" label="Пароль" rules={[{ required: true }]}>
                <Input.Password autoComplete="current-password" />
            </Form.Item>
            <Button type="primary" htmlType="submit" loading={isPending} block>
                Войти
            </Button>
        </Form>
        </Card>
    );
}