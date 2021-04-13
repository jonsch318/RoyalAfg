import React, { FC, useEffect, useState } from "react";
import useWindowDimensions from "../../hooks/windowSize";
import { usePoker } from "./provider";

const VisibleTime = 2000;

const Notification: FC = () => {
    const { height } = useWindowDimensions();
    const [visible, setVisible] = useState(false);
    const { notification } = usePoker();

    useEffect(() => {
        if (notification !== "") {
            setVisible(true);
            setTimeout(() => {
                setVisible(false);
            }, VisibleTime);
        }
    }, [notification]);

    return (
        <div style={{ visibility: visible ? "visible" : "hidden", background: "rgba(0,0,0,0.5)" }} className="absolute w-screen h-screen">
            <div
                style={{
                    top: 60,
                    height: height - 60
                }}
                className="text-white absolute grid justify-center items-center w-screen">
                <div className="rounded px-10 py-6 bg-black">
                    <h1>{notification}</h1>
                </div>
            </div>
        </div>
    );
};

export default Notification;
