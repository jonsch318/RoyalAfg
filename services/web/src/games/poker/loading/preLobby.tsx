import React, { FC } from "react";
import { useDots } from "./dots";
import { usePoker } from "../provider";
import { LinearProgress } from "@mui/material";
import { useTranslation } from "next-i18next";

const PreLobby: FC = () => {
    const { t } = useTranslation("poker");
    const dots = useDots();
    const { lobbyInfo } = usePoker();

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
                <LinearProgress variant={"determinate"} value={(lobbyInfo.playerCount / lobbyInfo.minPlayersToStart) * 100} />
            </div>
            <h1 className="font-sans font-semibold text-5xl text-center mb-4">{t("Royalafg Poker")}</h1>
            <h2 className="text-center">
                {t("Waiting for more players") + "[" + lobbyInfo.playerCount + " " + t("out of") + " " + lobbyInfo.minPlayersToStart + "]" + dots}
            </h2>
        </div>
    );
};

export default PreLobby;
