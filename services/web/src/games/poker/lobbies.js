import React, { useContext } from "react";
import Lobby from "./join/lobby";
import Chip from "./join/chip";
import PropTypes from "prop-types";
import { PokerInfoContext } from "../../pages/games/poker";

const Lobbies = ({ info }) => {
    const { lobby, setLobby } = useContext(PokerInfoContext);

    return (
        <div>
            <h1 className="font-sans text-xl text-center my-4 font-medium">Buy In Classes</h1>
            <div className="flex justify-center items-center">
                {info?.classes?.length &&
                    info?.classes.map((c, i) => (
                        <Chip
                            key={c.min}
                            lobbyClass={c}
                            selected={i === lobby.classIndex}
                            onSelect={() => {
                                setLobby({ id: "", i: -1, classIndex: i, class: c });
                            }}
                        />
                    ))}
            </div>
            <h1 className="font-sans text-xl text-center my-4 font-medium">Lobbies</h1>
            <div className="flex justify-center">
                {info?.lobbies?.length &&
                    info?.lobbies.map((c) => {
                        return (
                            c &&
                            c.map((l, j) => {
                                return (
                                    l.id && (
                                        <div key={l.id}>
                                            <Lobby
                                                lobby={l}
                                                selected={l.id === lobby.id}
                                                onLobbySelect={() => {
                                                    console.log("Selected lobby", j);
                                                    setLobby({ ...l, i: j });
                                                }}
                                            />
                                        </div>
                                    )
                                );
                            })
                        );
                    })}
            </div>
        </div>
    );
};

Lobbies.propTypes = {
    onLobbySelect: PropTypes.func,
    info: PropTypes.shape({
        lobbies: PropTypes.array,
        classes: PropTypes.array
    })
};

export default Lobbies;
