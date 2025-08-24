// src/AppLayout.jsx
import React from "react";
import {Link, Outlet, useLocation} from "react-router-dom";
import {Col, Layout, Menu, Row, Space} from "antd";
import {BarChartOutlined, ForkOutlined} from "@ant-design/icons";

const {Sider, Content} = Layout;

const items = [
    {key: "/meal", icon: <ForkOutlined/>, label: <Link to="/meal">Meal</Link>},
    {key: "/summary", icon: <BarChartOutlined/>, label: <Link to="/summary">Daily</Link>},
];

export function AppLayout() {
    const location = useLocation();

    const selectedKey = items.find(i => location.pathname.startsWith(i.key))?.key;

    return (
        <Layout style={{minHeight: "100vh"}}>
            {selectedKey && (<Sider>
                <Menu
                    theme="dark"
                    mode="inline"
                    selectedKeys={[selectedKey]}
                    items={items}
                />
            </Sider>)}
            <Layout style={{margin: "20px"}}>
                    <Content>
                        <Row justify="center">
                            <Col xs={24} sm={22} md={20} lg={16} xl={14} xxl={12}>
                                <Outlet/>
                            </Col>
                        </Row>
                    </Content>
            </Layout>
        </Layout>
    );
}