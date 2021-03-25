import React, { FC } from "react";
import View from "./view";
import PokerProvider, { usePoker, usePokerConn } from "./provider";
import Notification from "./notification";
import Actions from "./actions";
import GameLoading from "./loading";
import { useRouter } from "next/router";

const PokerVisual: FC = () => {
    const poker = usePoker();
    const { close, send } = usePokerConn();
    const router = useRouter();

    return (
        <>
            <button
                onClick={() => {
                    console.log("Closing Game from leave button");
                    close();
                    router.push("/games/poker").then();
                }}
                className="absolute cursor-pointer font-sans font-semibold text-sm ml-6 mt-4 py-1 px-3 bg-gray-300  rounded-full hover:bg-gray-800 hover:text-white transition-colors duration-200 ease-out">
                Leave
            </button>
            {poker.gameRunning && (
                <div>
                    <Actions possibleActions={poker.possibleActions} conn={send} />
                    <View />
                </div>
            )}
            <Notification />
            {!poker.gameRunning && <GameLoading />}
        </>
    );
};

type PokerProps = {
    ticket: {
        address: string;
        token: string;
    };
    csrf: string;
};

const Poker: FC<PokerProps> = ({ ticket, csrf }) => {
    return (
        <div className="poker">
            <PokerProvider ticket={ticket} csrf={csrf}>
                <PokerVisual />
            </PokerProvider>
        </div>
    );
};
export default Poker;
