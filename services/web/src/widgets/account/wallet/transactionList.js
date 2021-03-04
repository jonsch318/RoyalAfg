import React from "react";
import PropTypes from "prop-types";
import * as sprintf from "sprintf";

const DepositSprintf = "Deposited %s";
const WinSprintf = "Won %s in %s";
const LostSprintf = "Lost %s in %s";
const LobbyWonSprintf = "Won %s in %s lobby %s";
const LobbyLostSprintf = "Lost %s in %s lobby %s";

const Transaction = ({ game, amount, lobby, time, positive, money }) => {
    console.log("Money: ", money, " amount: ", amount);
    let contentString = "";
    if (positive) {
        if (game) {
            contentString = lobby ? sprintf(LobbyWonSprintf, amount, game, lobby) : sprintf(WinSprintf, amount, game);
        } else {
            contentString = sprintf(DepositSprintf, amount);
        }
    } else {
        contentString = lobby ? sprintf(LobbyLostSprintf, amount, game, lobby) : sprintf(LostSprintf, amount, game);
    }

    return (
        <div className="grid grid-cols-2 py-1 px-4 rounded bg-gray-300 my-2 hover:bg-gray-400 cursor-pointer">
            <span className="">{contentString}</span>
            <span className="flex justify-end">{time}</span>
        </div>
    );
};

Transaction.propTypes = {
    game: PropTypes.string,
    amount: PropTypes.string,
    lobby: PropTypes.string,
    time: PropTypes.object,
    positive: PropTypes.bool,
    money: PropTypes.object
};

const TransactionList = ({ transactions }) => {
    return (
        <div className="bg-gray-200 px-16 py-10 rounded-xl">
            {transactions.map((x, i) => {
                console.log("Event: ", x);
                return (
                    <div key={i}>
                        <Transaction game={x.game} lobby={x.lobby} amount={x.amount} time={x.time} positive={x.positive} money={x.amountDinero} />
                    </div>
                );
            })}
        </div>
    );
};

TransactionList.propTypes = {
    transactions: PropTypes.array
};

export default TransactionList;
