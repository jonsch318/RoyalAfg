/* eslint-disable @typescript-eslint/no-empty-function */
// Auth hooks are inspired by the solution by next-auth https://next-auth.js.org/.
// But a custom solution is required to accommodate the custom requirements

import { GetServerSidePropsContext } from "next";
import { useRouter } from "next/router";
import { createContext, createElement, FC, useContext, useEffect, useState } from "react";
import { LoginDto } from "../../dtos/auth";
import { Session } from "../../dtos/session";

const __AUTH = {
    baseUrl: process.env.NEXT_PUBLIC_AUTH_HOST,
    basePath: "/api/auth",
    session: undefined,
    clientMaxAge: 0, // 0 == disabled (only use cache); 60 == sync if last checked > 60 seconds ago
    _clientLastSync: 0, // used for timestamp since last synced (in seconds)
    _clientSyncTimer: null, // stores timer for poll interval
    _clientSession: undefined, // stores last session response from hook,
    // Used to store to function export by getSession() hook
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    _getSession: (): any => {}
};

type getSessionOptions = {
    ctx?: GetServerSidePropsContext;
    req?: GetServerSidePropsContext["req"];
};

export const getSession = async ({ ctx, req }: getSessionOptions = {}): Promise<Session> => {
    if (!req && ctx && ctx.req) {
        req = ctx.req;
    }
    console.log("AUTH URL: ", _apiBaseUrl());
    try {
        const res = await fetch(`${_apiBaseUrl()}/session`, {
            credentials: "include",
            headers: {
                cookie: req?.headers.cookie
            },
            method: "get",
            mode: "cors"
        });
        if (!res.ok) throw res.status;
        return await res.json();
    } catch (error) {
        console.log("error during fetch", `${_apiBaseUrl()}/session`, error);
        return null;
    }
};

const SessionContext = createContext(undefined);

export const useSession = (session?: Session): [Session, boolean] => {
    //check cache (context)
    const ctx = useContext(SessionContext);
    if (ctx) {
        return ctx;
    }
    return _useSessionHook(session);
};

export const useSecure = (route = "/auth/register"): [Session, boolean] => {
    const [session, loading] = useSession();
    const router = useRouter();
    useEffect(() => {
        if (!session && !loading) router.push(route);
    }, [session, loading]);
    return [session, loading];
};

// Internal hook for getting session from the api.
const _useSessionHook = (session?: Session): [Session, boolean] => {
    const [data, setData] = useState<Session>(session);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        const _getSession = async () => {
            try {
                const clientMaxAge = __AUTH.clientMaxAge;
                const clientLastSync = __AUTH._clientLastSync;
                const currentTime = Math.floor(new Date().getTime() / 1000);
                const clientSession = __AUTH.session;

                if (clientSession !== undefined && clientSession !== null && clientMaxAge > 0 && currentTime < clientLastSync + clientMaxAge) {
                    return;
                }

                if (clientSession !== undefined && clientSession !== null && clientMaxAge === 0) {
                    return;
                }

                if (clientSession === undefined) {
                    __AUTH.session = null;
                }

                __AUTH._clientLastSync = Math.floor(new Date().getTime() / 1000);

                const data = await getSession();

                __AUTH.session = data;

                setData(data);
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

export const signIn = async (args: LoginDto, csrfToken: string): Promise<Response> => {
    const res = await _fetch(`${_apiBaseUrl()}/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "X-CSRF-Token": csrfToken
        },
        body: JSON.stringify(args)
    });
    await getSession();
    return res;
};

interface Register {
    username: string;
    password: string;
    birthdate: string;
    email: string;
    fullName: string;
    acceptTerms: boolean;
}

export const register = async (args: Register, csrfToken: string): Promise<Response> => {
    console.log("REGISTER: ", `${_apiBaseUrl()}/register`);
    console.log("regiser args: ", args, " CSRF: ", csrfToken);
    const res = await _fetch(`${_apiBaseUrl()}/register`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "X-CSRF-Token": csrfToken
        },
        body: JSON.stringify(args)
    });
    await getSession();
    return res;
};

export const verifyUser = async (args: number, csrfToken: string): Promise<Response> => {
    console.log("VERIFING with code: ", args, "CSRF: ", csrfToken);

    return await _fetch(`${_apiBaseUrl()}/verification`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "X-CSRF-Token": csrfToken
        },
        body: JSON.stringify({ verificationCode: args })
    });
};

export const refreshSession = async (): Promise<[Session, boolean]> => {
    console.log("Refreshing Session");
    return __AUTH._getSession();
};

export const signOut = async (): Promise<Response> => {
    return _fetch(`${_apiBaseUrl()}/logout`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        }
    });
};

type ProviderProps = {
    session?: Session;
};
// eslint-disable-next-line react/prop-types
export const Provider: FC<ProviderProps> = ({ children, session }) => {
    return createElement(SessionContext.Provider, { value: useSession(session) }, children);
};

const _apiBaseUrl = (): string => {
    if (process.env.AUTH_HOST == undefined) {
        if (typeof window !== undefined) {
            return window.location.origin + "/api/auth";
        }
    }
    return `${process.env.AUTH_HOST}/api/auth`;
};

const _fetch = async (url: RequestInfo, options?: RequestInit): Promise<Response> => {
    try {
        return await fetch(url, options);
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
