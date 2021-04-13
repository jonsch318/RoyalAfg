import React, { FC, useContext } from "react";
import Lobby from "./join/lobby";
import Chip from "./join/chip";
import { PokerInfoContext } from "../../pages/games/poker";
import { IClass, ILobby } from "./models/class";

type LobbiesProps = {
    info: {
        lobbies: ILobby[][];
        classes: IClass[];
    };
};

const Lobbies: FC<LobbiesProps> = ({ info }) => {
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
                                console.log("Selected ClassIndex: ", i, " with min: ", c.min);
                                setLobby({ id: "", i: -1, classIndex: i, class: c, changeClass: false });
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
                {info?.lobbies?.length == 0 && <p>No lobbies found.</p>}
            </div>
        </div>
    );
};

export default Lobbies;
