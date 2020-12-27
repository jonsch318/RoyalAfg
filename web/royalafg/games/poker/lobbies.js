import React, { useEffect, useState } from "react";
import Lobby from "./join/lobby";
import Chip from "./join/chip";
import PropTypes from "prop-types";

const Lobbies = ({ onLobbySelect }) => {
    const [classes, setClasses] = useState([]);
    const [lobbies, setLobbies] = useState([]);

    const update = () => {
        fetch(`http://${window.location.hostname}:5000/options`, {
            mode: "cors"
        })
            .then((res) => {
                if (!res.ok) {
                    throw res;
                }
                return res.json();
            })
            .then((res) => {
                setClasses(res.buyInClasses);
                setLobbies(res.lobbies);
            })
            .catch((err) => {
                console.log(err);
            });
    };

    useEffect(() => {
        update();
        const interval = setInterval(() => update(), 5000);
        return () => {
            clearInterval(interval);
        };
    }, []);

    const f = [];
    if (lobbies.length > 0) {
        for (let i = 0; i < lobbies.length; i++) {
            if (lobbies[i]) {
                const c = lobbies[i];
                if (c.length > 0) {
                    f.push(
                        c.map((l) => (
                            <Lobby
                                key={l.id}
                                id={l.id}
                                min={l.minBuyIn}
                                max={l.maxBuyIn}
                                lobbyClass={l.lobbyClass}
                                lobbyClasses={classes}
                                smallBlind={l.smallBlind}
                                onLobbySelect={onLobbySelect}
                            />
                        ))
                    );
                }
            }
        }
    }

    return (
        <div>
            <h1 className="font-sans text-xl text-center my-4 font-medium">Buy In Classes</h1>
            <div className="flex justify-center items-center">
                {classes.map((c, i) => (
                    <Chip
                        key={c[0]}
                        min={c[0]}
                        max={c[1]}
                        smallBlind={c[2]}
                        onSelect={() => onLobbySelect("", i, c)}
                    />
                ))}
            </div>
            <h1 className="font-sans text-xl text-center my-4 font-medium">Lobbies</h1>
            <div className="flex justify-center">{f}</div>
        </div>
    );
};

Lobbies.propTypes = {
    onLobbySelect: PropTypes.func
};

export default Lobbies;
