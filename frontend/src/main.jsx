import React from "react";
import {createRoot} from "react-dom/client";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import {Daily} from "./pages/daily/daily.jsx";
import {Login} from "./pages/login/login.jsx";
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import "antd/dist/reset.css";
import {AppLayout} from "./layout.jsx";
import {App as AntdApp, ConfigProvider} from "antd";
import {MealAdd} from "./pages/meal-add/meal-add.jsx";

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
        element: <AppLayout/>,
        children: [
            {path: "login", element: <Login/>},
            {path: "summary", element: <Daily/>},
            {path: "meal", element: <MealAdd />}
        ],
    },
], {basename: "/web"});

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