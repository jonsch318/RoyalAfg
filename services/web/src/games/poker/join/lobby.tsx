import React, { FC } from "react";
import { ToFormat } from "../../../utils/currency";
import Dinero from "dinero.js";

type LobbyProps = {
    lobby: {
        id: string;
        class: {
            min: number;
            max: number;
            blind: number;
        };
        classIndex: number;
    };
    onLobbySelect: () => void;
    selected: boolean;
};

const Lobby: FC<LobbyProps> = ({ lobby, selected, onLobbySelect }) => {
    if (!lobby || !lobby.class) {
        return <></>;
    }
    return (
        <button
            style={{ background: selected ? "#f59e0b" : "#3182ce" }}
            className="bg-blue-600 my-2 mx-10 height text-white rounded cursor-pointer px-4 py-2 outline-none justify-end hover:shadow-xl transition-shadow duration-200"
            onMouseUp={() => {
                onLobbySelect();
            }}>
            <h1 className="lobby">Lobby [{lobby.id}]</h1>
            <span className="buyIn">{`Buy in: ${ToFormat(lobby.class.min)} - ${ToFormat(lobby.class.min)}  Blinds: ${ToFormat(
                lobby.class.blind
            )} - ${Dinero({ amount: lobby.class.blind, currency: "USD" }).multiply(2).toFormat()}`}</span>
        </button>
    );
};

export default Lobby;
