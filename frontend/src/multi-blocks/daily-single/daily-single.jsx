import React from 'react';
import {useNavigate} from "react-router-dom";
import {Daily} from "../../blocks/daily/daily.jsx";
import {getRoute, ROUTE_NAMES} from "../../route.js";

export const DailySingle = () => {
    const navigate = useNavigate();
    const toMealEdit = ({id}) => {
        const route = `${getRoute(ROUTE_NAMES.MEAL)}/${id}`;
        navigate(route);
    }
    return (<Daily onRowClick={toMealEdit} /> );
}