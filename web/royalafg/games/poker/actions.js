import React, { useEffect, useState } from "react";
import { BehaviorSubject } from "rxjs";
import { debounceTime } from "rxjs/operators";
import PropTypes from "prop-types";

const { PLAYER_ACTION } = require("./events/constants");

const Subj = new BehaviorSubject(0);

const Actions = ({ game, actions }) => {
    const [raise, setRaise] = useState(0);

    const send = ({ action, payload }) => {
        game.send({
            event: PLAYER_ACTION,
            data: {
                action: action,
                payload: payload
            }
        });
    };

    useEffect(() => {
        const subscription = Subj.pipe(debounceTime(700)).subscribe((val) => {
            setRaise(val);
        });

        return () => {
            subscription.unsubscribe();
        };
    }, []);

    return (
        <div>
            {actions & 1 ? <button onClick={() => send({ action: 1 })}>FOLD</button> : <></>}
            {(actions >> 3) & 1 ? (
                <button onClick={() => send({ action: 4 })}>CHECK</button>
            ) : (
                <></>
            )}
            {(actions >> 1) & 1 ? <button onClick={() => send({ action: 2 })}>CALL</button> : <></>}
            {(actions >> 2) & 1 ? (
                <div style={{ display: "inline" }}>
                    <input
                        type="text"
                        id="raiseInput"
                        name="raiseInput"
                        value={raise}
                        onChange={(e) => {
                            setRaise(e.target.value);
                            Subj.next(parseFloat(e.target.value));
                        }}
                    />
                    <button onClick={() => send({ action: 3, payload: parseInt(raise * 100) })}>
                        RAISE
                    </button>
                </div>
            ) : (
                <></>
            )}
            {(actions >> 4) & 1 ? (
                <button onClick={() => send({ action: 5 })}>All In</button>
            ) : (
                <></>
            )}
        </div>
    );
};

Actions.propTypes = {
    game: PropTypes.object,
    actions: PropTypes.object
};

export default Actions;
