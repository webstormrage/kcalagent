import React, { useEffect } from 'react';
import { Outlet, useLocation, useNavigate } from "react-router-dom";
import { Layout, Menu } from "antd";
import { ForkOutlined, BarChartOutlined } from "@ant-design/icons";
import {MOBILE_MQ} from "./constants.js";

const { Content } = Layout;

const items = [
    { key: "/mobile/meal", icon: <ForkOutlined />, label: null },
    { key: "/mobile/summary", icon: <BarChartOutlined />, label: null },
]

export  function MobileLayout() {
    const navigate = useNavigate();
    const location = useLocation();
    const selectedKey = items.find(i => location.pathname.startsWith(i.key))?.key;

    useEffect(() => {

        const isMobile = window.matchMedia(MOBILE_MQ).matches;
        if (!isMobile) {
            navigate(`/desktop/`, { replace: true });
        }
    }, []);

    return (
        <Layout style={{ minHeight: "100dvh", overflowX: "hidden", overflowY: "auto" }}>
            <Content style={{padding: '12px 12px 56px'}}>
                <Outlet />
            </Content>

            {/* Фиксированная навигация внизу */}
            <div
                style={{
                    position: "fixed",
                    left: 0,
                    right: 0,
                    bottom: 0,
                    background: "#fff",
                    borderTop: "1px solid #f0f0f0",
                    // небольшой учёт безопасной зоны на iOS
                    paddingBottom: "env(safe-area-inset-bottom)",
                    maxWidth: '100vw',
                    zIndex: 1000,
                }}
            >
                <Menu
                    mode="horizontal"
                    selectable
                    onClick={({key}) => navigate(key)}
                    selectedKeys={[selectedKey]}
                    items={items}
                />
            </div>
        </Layout>
    );
}