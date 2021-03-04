import React from "react";
import dynamic from "next/dynamic";
import Link from "next/link";

import Actions from "../../../games/poker/actions.js";
import "../poker.module.css";
import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { GameState } from "../../../games/poker/game/state.js";
import { Game } from "../../../games/poker/connection/socket.js";
import Back from "../../../components/layout/back";
import { formatTitle } from "../../../utils/title";
import Head from "next/head";

const View = dynamic(import("../../../games/poker/view"), { ssr: false });

const _getUrl = (id) => {
    let url = "";
    if (process.env.NEXT_PUBLIC_POKER_TICKET_HOST != undefined) {
        url = process.env.NEXT_PUBLIC_POKER_TICKET_HOST;
    }
    if (id) {
        console.log("Requesting ticket with ID");
        return `${url}/api/poker/ticket/${id}`;
    }
    console.log("Requesting ticket without ID");
    return `${url}/api/poker/ticket`;
};

const _fetch = async (url, params) => {
    return fetch(`${url}?${params.toString()}`, {
        mode: "cors",
        credentials: "include",
        method: "GET"
    });
};

const Play = () => {
    const [game, setGame] = useState({});
    const [joined, setJoined] = useState(false);
    const [actions, setActions] = useState({});
    const router = useRouter();
    const { lobbyId, buyInClass, buyIn } = router.query;

    useEffect(async () => {
        const params = new URLSearchParams({ buyIn: buyIn, class: buyInClass });
        const res = await _fetch(_getUrl(lobbyId), params);

        if (!res.ok) {
            await router.push("/games/poker");
            return;
        }
        const ticket = await res.json();
        console.log("Ticket: ", ticket);
        if (!ticket.address || !ticket.token) {
            await router.push("/games/poker");
            return;
        }
        let gameState = new GameState();
        gameState.setOnPossibleActions((actions) => {
            setActions(actions);
        });
        setGame(
            new Game(gameState, ticket, () => {
                router.push("/games/poker").then();
            })
        );
        game.start();
        console.log("Starting");
        setJoined(true);

        return () => {
            console.log("Closing Websocket connection");
            game.close(1001, "");
        };
    }, []);

    return (
        <div className="App">
            {joined ? (
                <div>
                    <Actions game={game} actions={actions} />
                    <View game={game} />
                </div>
            ) : (
                <div>
                    <h1>Loading</h1>
                    <a href="/games/poker">Back to lobby search</a>
                </div>
            )}
        </div>
    );
};

export default Play;
