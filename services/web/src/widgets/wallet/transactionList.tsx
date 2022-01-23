import { useTranslation } from "next-i18next";
import React, { FC } from "react";
import * as sprintf from "sprintf-js";
import { History } from "../../pages/wallet";

const Deposited = "Deposited";

const Transaction: FC<History> = ({ game, amount, lobby, time, type }) => {
    const { t } = useTranslation("wallet");

    let contentString = "";
    if (type == Deposited) {
        if (game !== "" && game !== undefined) {
            contentString = lobby ? sprintf(t("Won %s in %s [%s]"), amount, game, lobby) : sprintf(t("Won %s in %s"), amount, game);
        } else {
            contentString = sprintf(t("Deposited %s"), amount);
        }
    } else {
        if (game !== "" && game !== undefined) {
            contentString = lobby ? sprintf(t("Lost %s in %s [%s]"), amount, game, lobby) : sprintf(t("Lost %s in %s"), amount, game);
        } else {
            contentString = sprintf(t("Withdrawn %s"), amount);
        }
    }

    return (
        <div className="grid grid-cols-2 py-1 px-4 rounded bg-gray-300 my-2 hover:bg-gray-400 cursor-pointer">
            <span className="">{contentString}</span>
            <span className="flex justify-end">{time}</span>
        </div>
    );
};

type TransactionListProps = {
    transactions: History[];
};

const TransactionList: FC<TransactionListProps> = ({ transactions }) => {
    return (
        <div className="bg-gray-200 px-16 py-10 rounded-xl">
            {transactions.map((x, i) => {
                console.log("Event: ", x);
                return (
                    <div key={i}>
                        <Transaction game={x.game} lobby={x.lobby} amount={x.amount} time={x.time} type={x.type} />
                    </div>
                );
            })}
        </div>
    );
};

export default TransactionList;
