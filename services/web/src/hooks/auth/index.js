// Auth hooks are inspired by the solution by next-auth https://next-auth.js.org/.
// But a custom solution is required to accommodate the custom requirements

import { createContext, createElement, useContext, useEffect, useState } from "react";

const __AUTH = {
    baseUrl: process.env.API_URL || process.env.VERCEL_URL,
    basePath: "/api/auth/",
    session: undefined,
    // Used to store to function export by getSession() hook
    _getSession: () => {}
};

export const getSession = async ({ req, ctx } = {}) => {
    if (!req && ctx && ctx.req) {
        req = ctx.req;
    }

    const options = req ? { headers: { cookie: req.headers.cookie } } : {};
    try {
        const res = await fetch(`${_apiBaseUrl()}/session`, options);
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

const _useSessionHook = (session) => {
    const [data, setData] = useState(session);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const _getSession = async () => {
            try {
                if (__AUTH.session === undefined) {
                    __AUTH.session = null;
                }

                const newClientSessionData = await getSession();

                // Save session state internally, just so we can track that we've checked
                // if a session exists at least once.
                __AUTH.session = newClientSessionData;

                setData(newClientSessionData);
                setLoading(false);
            } catch (error) {
                console.log("error during session", error);
            }
        };

        __AUTH._getSession = _getSession;

        _getSession();
    });
    return [data, loading];
};

export const signIn = async (provider, args = {}) => {
    const options = {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: _encodedForm({
            ...args,
            json: true
        })
    };
    await fetch(`${_apiBaseUrl()}/login`, options);
};

export const signOut = async () => {
    const options = {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },
        body: _encodedForm({
            json: true
        })
    };

    await fetch(`${_apiBaseUrl()}/signout`, options);
};

// eslint-disable-next-line react/prop-types
export const Provider = ({ children, session }) => {
    return createElement(SessionContext.Provider, { value: useSession(session) }, children);
};

const _encodedForm = (formData) => {
    return Object.keys(formData)
        .map((key) => {
            return encodeURIComponent(key) + "=" + encodeURIComponent(formData[key]);
        })
        .join("&");
};

const _apiBaseUrl = () => {
    if (typeof window === "undefined") {
        if (!process.env.API_URL) {
            console.log("API_URL", "API_URL environment variable not set");
        }

        // Return absolute path when called server side
        return `${__AUTH.baseUrl}${__AUTH.basePath}`;
    } else {
        // Return relative path when called client side
        return __AUTH.basePath;
    }
};

export default {
    getSession,
    useSession,
    signIn,
    signOut,
    Provider
};
