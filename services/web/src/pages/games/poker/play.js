import React from "react";
import dynamic from "next/dynamic";

import Actions from "../../../games/poker/actions.js";
import "../poker.module.css";
import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { GameState } from "../../../games/poker/game/state.js";
import { Game } from "../../../games/poker/connection/socket.js";
import Loading from "../../../widgets/games/poker/loading";
import { useSnackbar } from "notistack";

const View = dynamic(import("../../../games/poker/view"), { ssr: false });

const _getUrl = (id) => {
    let url = "";
    if (process.env.NEXT_PUBLIC_POKER_TICKET_HOST !== undefined) {
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
    const [game, setGame] = useState({ start: () => console.log("Started before initializing") });
    const [joined, setJoined] = useState(false);
    const [loaded, setLoaded] = useState(false);
    const [actions, setActions] = useState({});
    const [lobbyInfo, setLobbyInfo] = useState({ count: -1, toStart: -1, timeout: 0 });
    const router = useRouter();
    const { lobbyId, buyInClass, buyIn } = router.query;
    const { enqueueSnackbar } = useSnackbar();

    useEffect(() => {
        const params = new URLSearchParams({ buyIn: buyIn, class: buyInClass });
        _fetch(_getUrl(lobbyId), params)
            .then((res) => {
                if (!res.ok) {
                    throw new Error("Connection could not be established");
                }
                return res.json();
            })
            .then((ticket) => {
                console.log("Ticket: ", ticket);
                if (!ticket.address || !ticket.token) {
                    enqueueSnackbar("Could not connect to the Poker server", { variant: "error" });
                    router.push("/games/poker").then();
                    return;
                }
                let gameState = new GameState((info) => {
                    console.log("LobbyInfo SET: ", info);
                    setLobbyInfo(info);
                });
                gameState.setOnPossibleActions((actions) => {
                    setActions(actions);
                });
                const g = new Game(gameState, ticket, () => {
                    router.push("/games/poker").then();
                });
                g.start();
                console.log("Game", g);
                setGame(g);
                console.log("Starting");
                setJoined(true);

                return () => {
                    console.log("Closing Websocket connection");
                    game.close(1000, "");
                };
            })
            .catch((err) => {
                enqueueSnackbar("Could not connect to the Poker server", { variant: "error" });
                console.log("Error during poker connect ", err);
                router.push("/games/poker").then();
            });
    }, []);

    return (
        <div className="App">
            <button
                onClick={() => {
                    console.log("Closing Game: ", game);
                    game.close(1000, "");
                    router.push("/games/poker").then();
                }}
                className="absolute cursor-pointer font-sans font-semibold text-sm ml-6 mt-4 py-1 px-3 bg-gray-300  rounded-full hover:bg-gray-800 hover:text-white transition-colors duration-200 ease-out">
                Leave
            </button>
            {joined && (
                <div>
                    <Actions game={game} actions={actions} />
                    <View game={game} gameStart={() => setLoaded(true)} />
                </div>
            )}
            <Loading connecting={!joined} joined={lobbyInfo.count} minNumber={lobbyInfo.toStart} loaded={loaded} timeout={lobbyInfo.timeout} />
        </div>
    );
};

export default Play;
