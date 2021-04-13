import React, { FC, useContext, useEffect, useRef, useState } from "react";
import { IPoker, PokerInitState } from "./models/poker";
import { SendEvent } from "./events";
import { JOIN } from "./events/constants";
import { OnMessage } from "./update";
import { IEvent } from "./models/event";
import { useRouter } from "next/router";

export function usePoker(): IPoker {
    const context = useContext(PokerContext);
    if (context === undefined) {
        return PokerInitState;
    }
    return context;
}

export function usePokerConn(): PokerConn | undefined {
    const context = useContext(SendMessageContext);
    if (context === undefined) {
        return undefined;
    }
    return context;
}

export type Send = (e: IEvent) => void;

export type PokerConn = {
    send: Send;
    close: (code?: number, reason?: string) => void;
};

export const PokerContext = React.createContext<IPoker | undefined>(undefined);
export const SendMessageContext = React.createContext<PokerConn | undefined>(undefined);

type PokerTicket = {
    address: string;
    token: string;
};

type PokerConnectionProps = {
    ticket: PokerTicket;
    csrf: string;
};

const PokerProvider: FC<PokerConnectionProps> = ({ children, ticket }) => {
    const [poker, setPoker] = useState(PokerInitState);
    const ws = useRef<WebSocket>();
    const router = useRouter();

    useEffect(() => {
        ws.current = new WebSocket(`ws://${ticket.address}/api/poker/join`);
        ws.current.onopen = (e) => {
            console.log("poker websocket session opened");
            if (e.type === "error") {
                ws.current?.close();
                return;
            }
            ws.current?.send(SendEvent({ event: JOIN, data: { token: ticket.token } }));
            setPoker((p) => {
                return { ...p, connected: true };
            });
            console.log("Connected to poker");
        };
        ws.current.onclose = () => console.log("poker websocket session closed");
        ws.current.onerror = (e) => {
            console.log("Websocket error: ", e);
            ws.current?.close();
            router.push("/games/poker").then();
        };

        return () => {
            console.log("Closing poker session on deconstruction");
            ws.current?.close();
        };
    }, [ticket]);

    useEffect(() => {
        if (!ws.current) return;

        ws.current.onmessage = (e) => {
            if (e.data && setPoker !== undefined) {
                const message = JSON.parse(e.data);
                console.log("Message:", message);
                OnMessage(setPoker, message);
            }
        };
    }, []);

    useEffect(() => {
        console.log("Poker Changed: ", poker);
    }, [poker]);

    return (
        <PokerContext.Provider value={poker}>
            <SendMessageContext.Provider
                value={{
                    send: (e) => {
                        ws.current?.send(SendEvent(e));
                    },
                    close: (code = 1000, reason = "") => {
                        ws.current?.close(code, reason);
                    }
                }}>
                {children}
            </SendMessageContext.Provider>
        </PokerContext.Provider>
    );
};
export default PokerProvider;
