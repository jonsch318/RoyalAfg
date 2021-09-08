import React, { FC } from "react";
import { Send } from "./provider";
import { PLAYER_ACTION } from "./events/constants";
import Raise from "./raise";
import { useTranslation } from "next-i18next";

const send = (s: Send | undefined, action: number, payload?: number) => {
    if (s === undefined) {
        console.log("Tried action with conn being undefined");
        return;
    }
    s({
        event: PLAYER_ACTION,
        data: {
            action: action,
            payload: payload
        }
    });
};

type ActionsProps = {
    conn: Send;
    possibleActions: number;
    bet: number;
};

const Actions: FC<ActionsProps> = ({ possibleActions, conn, bet }) => {
    const { t } = useTranslation("poker");

    return (
        <div style={{ height: 60 }} className="flex justify-center items-center">
            <div
                style={{
                    visibility: possibleActions > 0 ? "visible" : "hidden",
                    height: 50
                }}
                className="text-white flex justify-center items-center bg-blue-600 py-2 px-6 rounded-md">
                {(possibleActions & 1) === 1 && (
                    <button
                        className="bg-white px-3 flex justify-center items-center h-full text-black mx-4 rounded hover:bg-yellow-500 transition-colors ease-in-out duration-150"
                        onClick={() => send(conn, 1)}>
                        {t("FOLD")}
                    </button>
                )}
                {((possibleActions >> 3) & 1) === 1 && (
                    <button
                        className="bg-white px-3 flex justify-center items-center h-full text-black mx-4 rounded hover:bg-yellow-500 transition-colors ease-in-out duration-150"
                        onClick={() => send(conn, 4)}>
                        {t("CHECK")}
                    </button>
                )}
                {/*((possibleActions >> 1) & 1) === 1 && (
                    <button
                        className="bg-white px-3 flex justify-center items-center h-full text-black mx-4 rounded hover:bg-yellow-500 transition-colors ease-in-out duration-150"
                        onClick={() => send(conn, 2)}>
                        {t("CALL")}
                    </button>
                )*/}
                {((possibleActions >> 2) & 1) === 1 && <Raise bet={bet} onRaise={(amount) => send(conn, 3, amount)} />}
                {((possibleActions >> 4) & 1) === 1 && (
                    <button
                        className="bg-white px-3 flex justify-center items-center h-full text-black mx-4 rounded hover:bg-yellow-500 transition-colors ease-in-out duration-150"
                        onClick={() => send(conn, 5)}>
                        {t("ALL IN")}
                    </button>
                )}
            </div>
        </div>
    );
};

export default Actions;
