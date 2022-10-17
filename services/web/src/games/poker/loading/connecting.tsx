import React, { FC } from "react";
import { useDots } from "./dots";
import { CircularProgress } from "@mui/material";
import { useTranslation } from "next-i18next";

const Connecting: FC = () => {
    const { t } = useTranslation("poker");
    const dots = useDots();

    return (
        <div className="grid justify-center items-center">
            <div className="mx-auto flex mb-12">
                <CircularProgress variant="indeterminate" size="6rem" color="primary" />
            </div>
            <h1 className="font-sans font-semibold text-5xl text-center mb-4">{t("Royalafg Poker")}</h1>
            <h2 className="text-center">{t("Connecting") + dots}</h2>
        </div>
    );
};

export default Connecting;
