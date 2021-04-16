import React, { FC, useEffect, useRef, useState } from "react";
import { Text } from "@inlet/react-pixi";
import { useResize } from "../../../hooks/dimensions";
import { StatusPadding } from "./constants";
import { useTranslation } from "next-i18next";

type StatusProps = {
    pot: string;
    bet: string;
    lobbyId: string;
    appWidth: number;
    appHeight: number;
};

const Status: FC<StatusProps> = ({ pot, bet, lobbyId, appWidth, appHeight }) => {
    const { t } = useTranslation("poker");
    const [text, setText] = useState("");
    const ref = useRef(null);
    const { width, height } = useResize(ref);

    useEffect(() => {
        let w = "";
        if (pot) {
            w = t("Pot: ") + pot;
        }
        if (bet) {
            w += t(", To bet: ") + bet;
        }
        if (lobbyId) {
            w += t(", Lobby id: ") + lobbyId;
        }
        setText(w);
    }, [pot, bet, lobbyId]);

    return <Text x={appWidth - width - StatusPadding} y={appHeight - height - StatusPadding} ref={ref} text={text} />;
};

export default Status;
