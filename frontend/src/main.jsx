import React from "react";
import {createRoot} from "react-dom/client";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import {Daily} from "./blocks/daily/daily.jsx";
import {Login} from "./blocks/login/login.jsx";
import {createBrowserRouter, Navigate, RouterProvider} from "react-router-dom";
import "antd/dist/reset.css";
import {App as AntdApp, ConfigProvider} from "antd";
import {MealAdd} from "./blocks/meal-add/meal-add.jsx";
import {DesktopLayout} from "./desktop.jsx";
import {MobileLayout} from "./mobile.jsx";
import {MealDaily} from "./multi-blocks/meal-daily/meal-daily.jsx";
import {LoginCard} from "./multi-blocks/login-card/login-card.jsx";
import {MealEditSingle} from "./multi-blocks/meal-edit-single/meal-edit-single.jsx";
import {DailySingle} from "./multi-blocks/daily-single/daily-single.jsx";

const queryClient = new QueryClient({
    defaultOptions: {
        queries: {
            retry: 1,          // меньше ретраев по умолчанию
            refetchOnWindowFocus: false,
            staleTime: 30_000, // 30с считаем кэш «свежим»
        },
    },
});

const router = createBrowserRouter([
    {
        path: "/",
        children: [
            { index: true, element: <Navigate to="desktop/meal" replace /> },
            {
                path: "desktop/",
                element: <DesktopLayout />,
                children: [
                    { index: true, element: <Navigate to="meal" replace /> },
                    {path: "login", element: <LoginCard />},
                    {path: "meal", element: <MealDaily />}
                ]
            },
            {
                path: "mobile/",
                element: <MobileLayout />,
                children: [
                    { index: true, element: <Navigate to="meal" replace /> },
                    {path: "login", element: <Login/>},
                    {path: "summary", element: <DailySingle />},
                    {path: "meal/:mealId", element: <MealEditSingle /> },
                    {path: "meal", element: <MealAdd />},
                ]
            }
        ],
    },
],
    {basename: "/web"}
);

const root = createRoot(document.getElementById("root"));
root.render(
    <ConfigProvider>
        <AntdApp>
            <QueryClientProvider client={queryClient}>
                <RouterProvider router={router}/>
            </QueryClientProvider>
        </AntdApp>
    </ConfigProvider>
);