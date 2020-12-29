import React, { useEffect, useState } from "react";

import dynamic from "next/dynamic";

import Actions from "../../../games/poker/actions.js";
import { GameState } from "../../../games/poker/game/state.js";
import { Game } from "../../../games/poker/connection/socket.js";

import "../poker.module.css";
import { useRouter } from "next/router";

const View = dynamic(import("../../../games/poker/view"), { ssr: false });

const Play = () => {
    const [game, setGame] = useState({});
    const [joined, setJoined] = useState(false);
    const [actions, setActions] = useState({});
    const router = useRouter();
    const { lobbyId, id, username, buyInClass, buyIn } = router.query;

    useEffect(() => {
        let gameState = new GameState();
        gameState.setOnPossibleActions((actions) => {
            setActions(actions);
        });
        let game = new Game(
            gameState,
            {
                lobbyId: lobbyId,
                id: id,
                username: username,
                buyin: parseInt(buyIn),
                buyInClass: parseInt(buyInClass)
            },
            () => {
                router.push("/games/poker").then();
            }
        );
        game.start();
        console.log("Starting");
        setGame(game);
        setJoined(true);
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
