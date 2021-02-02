import React, { useEffect, useState } from "react";

import dynamic from "next/dynamic";

import Actions from "../../../games/poker/actions.js";
import { GameState } from "../../../games/poker/game/state.js";
import { Game } from "../../../games/poker/connection/socket.js";

import "../poker.module.css";
import { useRouter } from "next/router";
import { usePokerTicketRequest } from "../../../hooks/games/poker/connect";

const View = dynamic(import("../../../games/poker/view"), { ssr: false });

const Play = () => {
    const [game, setGame] = useState({});
    const [joined, setJoined] = useState(false);
    const [actions, setActions] = useState({});
    const router = useRouter();
    const { lobbyId, buyInClass, buyIn } = router.query;

    useEffect(() => {
        usePokerTicketRequest({ id: lobbyId, class: buyInClass, buyIn: buyIn }).then((ticket) => {
            console.log("Ticket: ", ticket);
            if (!ticket.address || !ticket.token) {
                router.push("/games/poker").then();
            } else {
                let gameState = new GameState();
                gameState.setOnPossibleActions((actions) => {
                    setActions(actions);
                });
                let game = new Game(gameState, ticket, () => {
                    router.push("/games/poker").then();
                });
                game.start();
                console.log("Starting");
                setGame(game);
                setJoined(true);
            }
        });
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
