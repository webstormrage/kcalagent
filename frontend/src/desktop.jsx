import React, { useEffect } from "react";
import {Link, Outlet, useLocation, useNavigate} from "react-router-dom";
import {Col, Layout, Menu, Row} from "antd";
import {BarChartOutlined, ForkOutlined} from "@ant-design/icons";
import {DESKTOP_MQ} from "./constants.js";

const {Sider, Content} = Layout;

const items = [
    {key: "/desktop/meal", icon: <ForkOutlined/>, label: <Link to="/desktop/meal">Meal</Link>},
    {key: "/desktop/summary", icon: <BarChartOutlined/>, label: <Link to="/desktop/summary">Daily</Link>},
];


export  function DesktopLayout() {
    const navigate = useNavigate();
    const location = useLocation();

    const selectedKey = items.find(i => location.pathname.startsWith(i.key))?.key;

    useEffect(() => {
        const isDesktop = window.matchMedia(DESKTOP_MQ).matches;
        if (!isDesktop) {
            navigate(`/mobile/`, { replace: true });
        }
        // только на маунте
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return (
        <Layout style={{minHeight: "100vh"}}>
            {/*selectedKey && (<Sider>
                <Menu
                    theme="dark"
                    mode="inline"
                    selectedKeys={[selectedKey]}
                    items={items}
                />
            </Sider>)*/}
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