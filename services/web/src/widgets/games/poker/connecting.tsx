import React, { FC, useEffect, useState } from "react";
import { CircularProgress } from "@material-ui/core";

const getDots = (dots: number): string => {
    return ".".repeat(dots);
};

const ConnectionLoading: FC = () => {
    const [dots, setDots] = useState(1);
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

    return (
        <div className="grid justify-center items-center">
            <div className="mx-auto flex mb-12">
                <CircularProgress variant="indeterminate" size="6rem" color="primary" />
            </div>
            <h1 className="font-sans font-semibold text-5xl text-center mb-4">Royalafg Poker</h1>
            <h2 className="text-center">Connecting{getDots(dots)}</h2>
        </div>
    );
};

export default ConnectionLoading;
