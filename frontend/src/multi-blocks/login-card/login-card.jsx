import React from 'react';
import {Card} from "antd";
import {Login} from "../../blocks/login/login.jsx";

export function LoginCard(){
    return (
        <Card title='авторизация'>
            <Login />
        </Card>
    )
}