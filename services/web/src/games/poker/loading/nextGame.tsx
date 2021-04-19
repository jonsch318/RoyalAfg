import React, { FC, useEffect, useState } from "react";
import { useDots } from "./dots";
import { CircularProgress, LinearProgress } from "@material-ui/core";
import { usePoker } from "../provider";
import { useTranslation } from "next-i18next";

const NextGameStartAfterTime: FC = () => {
    const { t } = useTranslation("poker");
    const dots = useDots();

    return (
        <div className="grid justify-center items-center">
            <div className="mx-auto flex mb-12">
                <CircularProgress variant="indeterminate" size="6rem" color="primary" />
            </div>
            <h1 className="font-sans font-semibold text-5xl text-center mb-4">{t("Royalafg Poker")}</h1>
            <h2 className="text-center">{t("Waiting for next game to start") + dots}</h2>
        </div>
    );
};

type NextGameStartTimedProps = {
    time: number;
    timeout: number;
};
const NextGameStartTimed: FC<NextGameStartTimedProps> = ({ time, timeout }) => {
    const { t } = useTranslation("poker");
    const dots = useDots();

    return (
        <div
            className="flex flex-col justify-center items-center w-full"
            style={{
                width: "100%",
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center"
            }}>
            <div className="box mb-12" style={{ width: "65%" }}>
                <LinearProgress variant={"determinate"} value={(time / timeout) * 100} />
            </div>
            <h1 className="font-sans font-semibold text-5xl text-center mb-4">{t("Royalafg Poker")}</h1>
            <h2 className="text-center">{t("Waiting for next game to start") + "[" + time + "]" + dots}</h2>
        </div>
    );
};

const NextGameStart: FC = () => {
    const { lobbyInfo, gameRunning } = usePoker();
    const [tick, setTick] = useState(lobbyInfo.gameStartTimeout);
    useEffect(() => {
        const inter = setInterval(() => {
            if (lobbyInfo.playerCount >= lobbyInfo.minPlayersToStart && tick >= 0) {
                setTick((t) => t - 1);
            }
        }, 1000);
        return () => {
            clearInterval(inter);
        };
    }, [lobbyInfo, tick]);

    return (
        <>
            {(tick < 0 || gameRunning) && <NextGameStartAfterTime />}
            {tick >= 0 && !gameRunning && <NextGameStartTimed time={tick} timeout={lobbyInfo.gameStartTimeout} />}
        </>
    );
};

export default NextGameStart;
