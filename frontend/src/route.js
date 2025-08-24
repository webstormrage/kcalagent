import { DESKTOP_MQ } from "./constants.js";

export function isDesktop() {
    return window.matchMedia(DESKTOP_MQ).matches;
}

export const ROUTE_NAMES = {
    LOGIN: 'LOGIN',
    MEAL: 'MEAL',
    SUMMARY: 'SUMMARY'
}

export const ROUTES = {
    MOBILE: {
        LOGIN: '/mobile/login',
        MEAL: '/mobile/meal',
        SUMMARY: '/mobile/summary',
    },
    DESKTOP:{
        LOGIN: '/desktop/login',
        MEAL: '/desktop/meal',
    }
};

export const getRoute = (name) => {
    if(isDesktop()){
        return ROUTES.DESKTOP[name];
    }
    return ROUTES.MOBILE[name];
}

export const getFullRoute = (name) => {
    return '/web'+getRoute(name);
}