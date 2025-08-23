import React from "react";
import { Outlet } from "react-router-dom";
import { Layout, Row, Col, Typography } from "antd";

const { Header, Content } = Layout;

export function AppLayout() {
    const today = React.useMemo(
        () =>
            new Date().toLocaleDateString("ru-RU", {
                weekday: "long",
                day: "2-digit",
                month: "long",
                year: "numeric",
            }),
        []
    );

    return (
        <Layout>
            <Header>
                <Typography.Title level={3}  type="success">
                    {today}
                </Typography.Title>
            </Header>

            <Content>
                <br />
                <Row justify="center">
                    {/* ограничение ширины и центрирование без стилей */}
                    <Col xs={24} sm={22} md={20} lg={16} xl={14} xxl={12}>
                        <Outlet />
                    </Col>
                </Row>
            </Content>
        </Layout>
    );
}