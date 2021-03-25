import React, { FC } from "react";
import { usePoker } from "../provider";
import Connecting from "./connecting";
import PreLobby from "./preLobby";
import NextGameStart from "./nextGame";

const GameLoading: FC = () => {
    const poker = usePoker();

    return (
        <div className="w-screen h-screen flex justify-center items-center">
            {!poker.connected && <Connecting />}
            {poker.connected && poker.lobbyInfo.playerCount < poker.lobbyInfo.minPlayersToStart && <PreLobby />}
            {poker.connected && poker.lobbyInfo.playerCount >= poker.lobbyInfo.minPlayersToStart && <NextGameStart />}
        </div>
    );
};

export default GameLoading;
