import React, { useEffect, useState } from "react";
import Lobby from "./join/lobby";
import Chip from "./join/chip";
import PropTypes from "prop-types";


export async function getServerSideProps() {
    const resClasses = await fetch("http://localhost:5000/api/poker/classes", {
        mode: "cors"
    })
    const resLobbies = await fetch("http://localhost:5000/api/poker/lobbies", {
        mode: "cors"
    })
    
    const classes = await resClasses.json()


    const lobbies = await resLobbies.json()


    return {
        props: {
            classes,
            lobbies
        }
    }
}

const Lobbies = ({ onLobbySelect }) => {
    const [classes, setClasses] = useState([]);
    const [lobbies, setLobbies] = useState([]);
    const [selectedLobby, setSelectedLobby] = useState(-1);
    const [selectedClass, setSelectedClass] = useState(-1);

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
                    c.map((l, index) => {
                        f.push(
                            <Lobby
                                key={l.id}
                                id={l.id}
                                min={l.minBuyIn}
                                max={l.maxBuyIn}
                                lobbyClass={l.lobbyClass}
                                lobbyClasses={classes}
                                smallBlind={l.smallBlind}
                                selected={(i + 1) * (index + 1) === selectedLobby}
                                onLobbySelect={(id, lobbyClass, classes) => {
                                    console.log("Selected", (i + 1) * (index + 1));
                                    setSelectedLobby((i + 1) * (index + 1));
                                    onLobbySelect(id, lobbyClass, classes);
                                }}
                            />
                        );
                    });
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
                        selected={i === selectedClass}
                        onSelect={() => {
                            setSelectedClass(i);
                            onLobbySelect("", i, c);
                        }}
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
