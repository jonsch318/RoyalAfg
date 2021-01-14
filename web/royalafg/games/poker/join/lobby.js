import React from "react";
import PropTypes from "prop-types";

const Lobby = ({ id, lobbyClass, lobbyClasses, selected, onLobbySelect }) => {
    return (
        <button
            style={{ background: selected ? "#ecc94b" : "#3182ce" }}
            className="bg-blue-600 my-2 mx-10 height text-white rounded cursor-pointer px-4 py-2 outline-none justify-end hover:shadow-xl transition-shadow duration-200"
            onMouseUp={() => {
                onLobbySelect(id, lobbyClass, lobbyClasses[lobbyClass]);
            }}>
            <h1 className="lobby">Lobby [{id}]</h1>
            <span className="buyIn">{`Buy in: ${lobbyClasses[lobbyClass][0]} - ${lobbyClasses[lobbyClass][1]}  Blinds: ${
                lobbyClasses[lobbyClass][2]
            } - ${lobbyClasses[lobbyClass][2] * 2}`}</span>
        </button>
    );
};

Lobby.propTypes = {
    id: PropTypes.string,
    min: PropTypes.number,
    max: PropTypes.number,
    smallBlind: PropTypes.number,
    lobbyClass: PropTypes.number,
    lobbyClasses: PropTypes.array,
    onLobbySelect: PropTypes.func,
    selected: PropTypes.bool
};

export default Lobby;
