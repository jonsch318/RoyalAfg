import React, { FC, useEffect, useState } from "react";
import { LinearProgress } from "@material-ui/core";

type WaitingForPlayerProps = {
    joined: number;
    minNumber: number;
    timeout: number;
};

const getDots = (dots: number): string => {
    return ".".repeat(dots);
};

const WaitingForPlayers: FC<WaitingForPlayerProps> = ({ joined, minNumber, timeout }) => {
    const [dots, setDots] = useState(1);
    const [tick, setTick] = useState(timeout);
    useEffect(() => {
        const inter = setInterval(() => {
            if (dots >= 3) {
                setDots(1);
            } else {
                setDots(dots + 1);
            }
        }, 500);
        return () => {
            clearInterval(inter);
        };
    }, [dots]);

    useEffect(() => {
        const inter = setInterval(() => {
            if (joined >= minNumber && tick > 0) {
                setTick(tick - 1);
            } else {
                console.log(joined, " > ", minNumber);
            }
        }, 1000);
        return () => {
            console.log("Clear interval");
            clearInterval(inter);
        };
    }, [joined, minNumber, tick]);

    return (
        <div className="flex flex-col justify-center items-center w-full">
            <div className="box mb-12" style={{ width: "65%" }}>
                <LinearProgress variant={"determinate"} value={joined < minNumber ? (joined / minNumber) * 100 : (tick / timeout) * 100} />
            </div>
            <h1 className="font-sans font-semibold text-5xl text-center mb-4">Royalafg Poker</h1>
            <h2 className="text-center">
                {joined < minNumber
                    ? "Waiting for more players [" + joined + " out of " + minNumber + "]"
                    : "Waiting " + tick + " seconds to game start"}
                {getDots(dots)}
            </h2>
        </div>
    );
};

export default WaitingForPlayers;
