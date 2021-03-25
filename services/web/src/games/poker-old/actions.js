import React from "react";
import PropTypes from "prop-types";

import { PLAYER_ACTION } from "./events/constants";
import Raise from "./actions/raise";

const Actions = ({ game, actions }) => {
    const send = ({ action, payload }) => {
        game.send({
            event: PLAYER_ACTION,
            data: {
                action: action,
                payload: payload
            }
        });
    };

    return (
        <div className="flex justify-center items-center">
            <div className="flex mx-auto mt-4 bg-blue-600 rounded-md py-2">
                {(actions & 1) === 1 && (
                    <button
                        className="bg-white px-3 text-black mx-4 rounded hover:bg-yellow-500 transition-colors ease-in-out duration-150"
                        onClick={() => send({ action: 1 })}>
                        FOLD
                    </button>
                )}
                {((actions >> 3) & 1) === 1 && (
                    <button
                        className="bg-white px-3 text-black mx-4 rounded hover:bg-yellow-500 transition-colors ease-in-out duration-150"
                        onClick={() => send({ action: 4 })}>
                        CHECK
                    </button>
                )}
                {((actions >> 1) & 1) === 1 && (
                    <button
                        className="bg-white px-3 text-black mx-4 rounded hover:bg-yellow-500 transition-colors ease-in-out duration-150"
                        onClick={() => send({ action: 2 })}>
                        CALL
                    </button>
                )}
                {((actions >> 2) & 1) === 1 && <Raise onRaise={(e) => send(e)} />}
                {((actions >> 4) & 1) === 1 && (
                    <button
                        className="bg-white px-3 text-black mx-4 rounded hover:bg-yellow-500 transition-colors ease-in-out duration-150"
                        onClick={() => send({ action: 5 })}>
                        ALL IN
                    </button>
                )}
            </div>
        </div>
    );
};

Actions.propTypes = {
    game: PropTypes.object,
    actions: PropTypes.oneOfType([PropTypes.number, PropTypes.string, PropTypes.object])
};

export default Actions;
