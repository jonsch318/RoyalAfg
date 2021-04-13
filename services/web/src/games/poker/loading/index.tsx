import React, { FC } from "react";
import { usePoker } from "../provider";
import Connecting from "./connecting";
import PreLobby from "./preLobby";
import NextGameStart from "./nextGame";

const GameLoading: FC = () => {
    const { connected, lobbyInfo, loaded } = usePoker();

    return (
        <div className="w-screen h-screen flex justify-center items-center" hidden={!loaded}>
            {!connected && <Connecting />}
            {connected && lobbyInfo.playerCount < lobbyInfo.minPlayersToStart && <PreLobby />}
            {connected && lobbyInfo.playerCount >= lobbyInfo.minPlayersToStart && <NextGameStart />}
        </div>
    );
};

export default GameLoading;
