import React, { useContext, useEffect, useState } from "react";
import PropTypes from "prop-types";
import { PokerInfoContext } from "../../pages/games/poker";
import CurrencyInput from "react-currency-input-field";
import { useRouter } from "next/router";

const GetClass = (classes, v, setLobby) => {
    let val = v * 100;
    for (let i = 0; i < classes.length; i++) {
        if (classes[i].min < val && classes[i].max > val) {
            console.log("Set Lobby: ", i);
            setLobby({ class: classes[i], classIndex: i, changeClass: true });
            return i;
        }
    }
    setLobby({ class: {}, classIndex: -1 });
    return -1;
};

const Join = ({ onJoin, classes }) => {
    const { lobby, setLobby } = useContext(PokerInfoContext);
    const [buyIn, setBuyIn] = useState();
    const { locale } = useRouter();

    useEffect(() => {
        if (lobby?.class?.min) {
            if (!lobby.changeClass) {
                setBuyIn(lobby.class.min / 100);
            } else {
                setLobby({ ...lobby, changeClass: false });
            }
        }
    }, [lobby.class]);

    const onSubmit = (e) => {
        e.preventDefault();

        const values = {
            buyIn: Math.floor(parseFloat(buyIn) * 100),
            lobbyId: lobby.id,
            class: lobby.classIndex
        };
        console.log(values);
        onJoin(values);
    };

    if (!classes || !classes.length) {
        return <div>Cant load poker information</div>;
    }
    return (
        <div>
            <form
                onSubmit={onSubmit}
                className="flex justify-center items-center mx-auto my-5 bg-blue-600 w-screen px-1 py-2 rounded shadow-lg"
                style={{ width: "fit-content" }}>
                <CurrencyInput
                    name="buyIn"
                    className="mx-4 p-1 pl-3 rounded outline-none"
                    placeholder={"Buy In Amount"}
                    intlConfig={{ locale: locale, currency: "USD" }}
                    value={buyIn}
                    onValueChange={(val) => {
                        GetClass(classes, parseFloat(val), setLobby);
                        setBuyIn(val);
                    }}
                    allowNegativeValue={false}
                />
                <input
                    className="mx-4 p-1 pl-3 rounded outline-none"
                    name="lobbyId"
                    id="lobbyId"
                    placeholder="Lobby Id"
                    type="text"
                    value={lobby?.id ?? ""}
                    onChange={(e) => setLobby({ ...lobby, id: e.target.value, i: -1 })}
                />
                <button
                    className="bg-yellow-500 text-gray-800 hover:bg-yellow-600 transition-colors duration-150 ease-in-out rounded py-1 px-3 mr-3 disabled:opacity-75"
                    type="submit"
                    disabled={!buyIn || lobby.classIndex < 0}>
                    Join
                </button>
            </form>
        </div>
    );
};

Join.propTypes = {
    onJoin: PropTypes.func,
    classes: PropTypes.arrayOf(PropTypes.shape({ min: PropTypes.number, max: PropTypes.number, blind: PropTypes.number }))
};

export default Join;
