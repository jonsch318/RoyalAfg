import React, { FC } from "react";
import { CircularProgress } from "@mui/material";
import { useDots } from "../../../hooks/dots";

const NextGameStart: FC = () => {
    const dots = useDots();

    return (
        <div className="grid justify-center items-center">
            <div className="mx-auto flex mb-12">
                <CircularProgress variant="indeterminate" size="6rem" color="primary" />
            </div>
            <h1 className="font-sans font-semibold text-5xl text-center mb-4">Royalafg Poker</h1>
            <h2 className="text-center">Waiting for next game to start{dots}</h2>
        </div>
    );
};

export default NextGameStart;
