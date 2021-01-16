import React from "react";
import PropTypes from "prop-types";
import Dinero from "dinero.js";
import * as sprintf from "sprintf";

const DepositSprintf = "Deposited %s";
const WinSprintf = "Won %s in %s";
const LostSprintf = "Lost %s in %s";
const LobbyWonSprintf = "Won %s in %s lobby %s";
const LobbyLostSprintf = "Lost %s in %s lobby %s";

const Transaction = ({ game, amount, lobby, time }) => {
    let contentString = "";
    if (amount.isPositive()) {
        if (game) {
            contentString = lobby ? sprintf(LobbyWonSprintf, amount.toFormat(), game, lobby) : sprintf(WinSprintf, amount.toFormat(), game);
        } else {
            contentString = sprintf(DepositSprintf, amount.toFormat());
        }
    } else {
        const val = amount.multiply(-1);
        contentString = lobby ? sprintf(LobbyLostSprintf, val.toFormat(), game, lobby) : sprintf(LostSprintf, val.toFormat(), game);
    }

    return (
        <div className="grid grid-cols-2 py-1 px-4 rounded bg-gray-300 my-2 hover:bg-gray-400 cursor-pointer">
            <span className="">{contentString}</span>
            <span className="flex justify-end">{time.format("MMMM Do YYYY, h:mm:ss a")}</span>
        </div>
    );
};

Transaction.propTypes = {
    game: PropTypes.string,
    amount: PropTypes.object,
    lobby: PropTypes.string,
    time: PropTypes.object
};

const TransactionList = ({ transactions }) => {
    return (
        <div className="bg-gray-200 px-16 py-10 rounded-xl">
            {transactions.map((x, i) => {
                return <Transaction key={i} game={x.game} lobby={x.lobby} amount={Dinero({ amount: x.amount, currency: "EUR" })} time={x.time} />;
            })}
        </div>
    );
};

TransactionList.propTypes = {
    transactions: PropTypes.array
};

export default TransactionList;
