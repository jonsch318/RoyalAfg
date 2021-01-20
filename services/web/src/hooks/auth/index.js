// Auth hooks are inspired by the solution by next-auth https://next-auth.js.org/.
// But a custom solution is required to accommodate the custom requirements

import { createContext, createElement, useContext, useEffect, useState } from "react";

const __AUTH = {
    baseUrl: "http://localhost:8080",
    basePath: "/api/auth",
    session: undefined,
    clientMaxAge: 0, // 0 == disabled (only use cache); 60 == sync if last checked > 60 seconds ago
    _clientLastSync: 0, // used for timestamp since last synced (in seconds)
    _clientSyncTimer: null, // stores timer for poll interval
    _clientSession: undefined, // stores last session response from hook,
    // Used to store to function export by getSession() hook
    _getSession: () => {}
};

export const getSession = async ({ req, ctx } = {}) => {
    if (!req && ctx && ctx.req) {
        req = ctx.req;
    }

    try {
        const res = await fetch(`${_apiBaseUrl()}/session`, {
            credentials: "include",
            headers: req?.headers.cookie,
            method: "get",
            mode: "cors"
        });
        if (!res.ok) throw res.statusCode;
        return await res.json();
    } catch (error) {
        console.log("error during fetch", `${_apiBaseUrl()}/session`, error);
        return null;
    }
};

const SessionContext = createContext();

export const useSession = (session) => {
    const value = useContext(SessionContext);
    //check cache (context)
    return value === undefined ? _useSessionHook(session) : value;
};

// Internal hook for getting session from the api.
const _useSessionHook = (session) => {
    const [data, setData] = useState(session);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const _getSession = async () => {
            try {
                const clientMaxAge = __AUTH.clientMaxAge;
                const clientLastSync = parseInt(__AUTH._clientLastSync);
                const currentTime = Math.floor(new Date().getTime() / 1000);
                const clientSession = __AUTH.session;

                if (clientSession !== undefined && clientMaxAge > 0 && currentTime < clientLastSync + clientMaxAge) {
                    return;
                }

                if (clientSession !== undefined && clientMaxAge === 0) {
                    return;
                }

                if (clientSession === undefined) {
                    __AUTH.session = null;
                }

                __AUTH._clientLastSync = Math.floor(new Date().getTime() / 1000);

                const newClientSessionData = await getSession();

                __AUTH.session = newClientSessionData;

                setData(newClientSessionData);
                setLoading(false);
            } catch (error) {
                console.log("session error", error);
            }
        };

        __AUTH._getSession = _getSession;

        _getSession();
    });
    return [data, loading];
};

export const signIn = async (args = {}) => {
    console.log("signin: ", args);
    const options = {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        credentials: "include",
        mode: "cors",
        body: JSON.stringify({ ...args })
    };

    return _fetch(`${_apiBaseUrl()}/login`, options);
};

export const signOut = async () => {
    const options = {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        credentials: "include",
        mode: "cors",
        body: {}
    };

    return _fetch(`${_apiBaseUrl()}/logout`, options);
};

// eslint-disable-next-line react/prop-types
export const Provider = ({ children, session }) => {
    return createElement(SessionContext.Provider, { value: useSession(session) }, children);
};

const _apiBaseUrl = () => {
    return "http://localhost:8080/api/auth";
    /*    if (typeof window === "undefined") {
        if (!process.env.API_URL) {
            console.log("API_URL", "API_URL environment variable not set");
        }

        // Return absolute path when called server side
        return `${__AUTH.baseUrl}${__AUTH.basePath}`;
    } else {
        // Return relative path when called client side
        return __AUTH.basePath;
    }*/
};

const _fetch = async (url, options) => {
    try {
        await fetch(url, options);
        window.location = "/";
        return Promise.resolve();
    } catch (error) {
        console.log("error during fetch", url, error);
        return Promise.reject(error);
    }
};

export default {
    getSession,
    useSession,
    signIn,
    signOut,
    Provider
};
