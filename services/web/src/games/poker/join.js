import React, { useEffect, useState } from "react";
import { v4 as uuidv4 } from "uuid";
import { BehaviorSubject } from "rxjs";
import { debounceTime } from "rxjs/operators";
import PropTypes from "prop-types";

const Subj = new BehaviorSubject(0);

const Join = ({ onJoin, lobbyId, buyInClass, minBuyIn, maxBuyIn }) => {
    const [username, setUsername] = useState("");
    const [buyIn, setBuyIn] = useState(0);
    const [lId, setLobbyId] = useState(lobbyId);
    const [lobbyClass, setLobbyClass] = useState(0);


    useEffect(() => {
        if (isNaN(minBuyIn)) {
            minBuyIn = 0;
        }
        if (isNaN(maxBuyIn)) {
            maxBuyIn = 0;
        }
    }, [minBuyIn, maxBuyIn]);

    const onSubmit = (e) => {
        e.preventDefault();
        const vals = {
            username: username,
            buyin: parseInt(buyIn * 100),
            id: uuidv4(),
            lobbyId: lId,
            buyInClass: lobbyClass
        };
        onJoin(vals);
        console.log(vals);
    };

    useEffect(() => {
        const subscription = Subj.pipe(debounceTime(700)).subscribe((v) => {
            let b = v;
            b = b < minBuyIn ? minBuyIn : b;
            b = b > maxBuyIn ? maxBuyIn : b;
            setBuyIn(b);
        });
        return () => {
            subscription.unsubscribe();
        };
    }, [minBuyIn, maxBuyIn]);

    useEffect(() => {
        setLobbyId(lobbyId);
        setLobbyClass(buyInClass);
        setBuyIn((minBuyIn + maxBuyIn) / 2);
        Subj.next((minBuyIn + maxBuyIn) / 2);
    }, [lobbyId, buyInClass, minBuyIn, maxBuyIn]);

    return (
        <div>
            <form
                onSubmit={onSubmit}
                className="flex justify-center items-center mx-auto my-5 bg-blue-600 w-screen px-1 py-2 rounded shadow-lg"
                style={{ width: "fit-content" }}>
                <input
                    className="mx-4 p-1 pl-3 text-white rounded outline-none bg-yellow-500 fill-current stroke-current"
                    name="buyIn"
                    type="range"
                    disabled={!minBuyIn && !maxBuyIn}
                    min={minBuyIn}
                    max={maxBuyIn}
                    id="buyInRange"
                    placeholder="BuyIn"
                    value={buyIn}
                    onChange={(e) => setBuyIn(e.target.value)}
                />
                <input
                    className="mx-4 p-1 rounded outline-none"
                    name="buyIn"
                    disabled={!minBuyIn && !maxBuyIn}
                    min={minBuyIn}
                    max={maxBuyIn}
                    id="range"
                    placeholder="BuyIn"
                    value={buyIn}
                    onChange={(e) => {
                        setBuyIn(e.target.value);
                        Subj.next(parseFloat(e.target.value));
                    }}
                />
                <input
                    className="mx-4 p-1 pl-3 rounded outline-none"
                    name="lobbyId"
                    id="lobbyId"
                    placeholder="Lobby Id"
                    value={lId}
                    onChange={(e) => setLobbyId(e.target.value)}
                />
                <input
                    name="mx-4 pl-3 rounded"
                    type="number"
                    hidden
                    id="buyInClass"
                    placeholder="Lobby Class"
                    value={lobbyClass + 1}
                    onChange={(e) => setLobbyClass(e.target.value - 1)}
                />
                <button
                    className="bg-yellow-500 text-gray-800 hover:bg-yellow-600 transition-colors duration-150 ease-in-out rounded py-1 px-3 mr-3"
                    type="submit"
                    disabled={(!minBuyIn && !maxBuyIn) || !username}>
                    Join
                </button>
            </form>
        </div>
    );
};

Join.propTypes = {
    onJoin: PropTypes.func,
    lobbyId: PropTypes.string,
    buyInClass: PropTypes.number,
    minBuyIn: PropTypes.number,
    maxBuyIn: PropTypes.number
};

export default Join;
