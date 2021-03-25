import React from "react";
import { useDots } from "./dots";
import { CircularProgress } from "@material-ui/core";

const Connecting = () => {
    const dots = useDots();

    return (
        <div className="grid justify-center items-center">
            <div className="mx-auto flex mb-12">
                <CircularProgress variant="indeterminate" size="6rem" color="primary" />
            </div>
            <h1 className="font-sans font-semibold text-5xl text-center mb-4">Royalafg Poker</h1>
            <h2 className="text-center">Connecting{dots}</h2>
        </div>
    );
};

export default Connecting;
