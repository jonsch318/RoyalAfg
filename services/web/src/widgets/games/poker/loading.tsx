import React, { FC, useEffect } from "react";
import Back from "../../../components/layout/back";
import ConnectionLoading from "./connecting";
import WaitingForPlayers from "./waiting";

type LoadingProps = {
    connecting: boolean;
    joined?: number;
    minNumber?: number;
    loaded: boolean;
    timeout: number;
};

const Loading: FC<LoadingProps> = ({ connecting, joined, minNumber, loaded, timeout }) => {
    console.log("Joined : ", joined, " Min Number: ", minNumber, " Connecting: ", connecting, "loading: ", loaded);

    useEffect(() => {
        console.log("Joined Effect : ", joined, " Min Number: ", minNumber, " Connecting: ", connecting, "loading: ", loaded);
    }, [joined]);

    if (loaded) {
        return null;
    }
    return (
        <>
            <div className="flex w-full justify-center items-center h-screen">
                {connecting && <ConnectionLoading />}
                {joined >= 0 && <WaitingForPlayers joined={joined} minNumber={minNumber} timeout={timeout} />}
            </div>
        </>
    );
};

export default Loading;
