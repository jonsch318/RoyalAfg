import React from "react";
import Layout from "../../../components/layout";
import Front from "../../../components/layout/front";
import Dinero from "dinero.js";
import PropTypes from "prop-types";
import ActionMenu from "../../../components/actionMenu";
import TransactionList from "../../../widgets/account/wallet/transactionList";
import BackToAccount from "../../../widgets/account/back";
import ActionMenuLink from "../../../components/actionMenu/link";
import moment from "moment";

const WalletHeader = ({ value }) => {
    return (
        <Front>
            <div className="md:px-10 font-sans text-5xl font-semibold text-center grid grid-cols-2 justify-center items-center">
                <h1 className="text-6xl h-auto align-middle">{value.toFormat()}</h1>
                <h1>Your Wallet</h1>
            </div>
        </Front>
    );
};

WalletHeader.propTypes = {
    value: PropTypes.object
};

const Wallet = ({ value, history }) => {
    const account = {
        value: Dinero({ amount: 25000, currency: "EUR" }),
        transactions: [
            { amount: 1000, game: "", lobby: "", time: moment().subtract(7, "days") },
            { amount: -2500, game: "Poker", lobby: "UXMSK", time: moment().subtract(1, "h") },
            { amount: 1000, game: "Poker", lobby: "UXMSK", time: moment() }
        ]
    };

    return (
        <Layout>
            <BackToAccount />
            <WalletHeader value={account.value} />
            <div className="px-10 pb-10 bg-gray-200">
                <ActionMenu>
                    <ActionMenuLink href="/account/wallet/deposit">Deposit</ActionMenuLink>
                </ActionMenu>
            </div>
            <div className="p-10 bg-white">
                <TransactionList transactions={account.transactions} />
            </div>
        </Layout>
    );
};

export async function getStaticProps(ctx) {
    const { req } = ctx;

    const res = await fetch(`${process.env.API_ADRESS}/api/bank/history`, {
        headers: {
            cookie: req.headers.cookie ?? ""
        }
    });

    const valueRes = await fetch(`${process.env.API_ADRESS}/api/bank/history`, {
        headers: {
            cookie: req.headers.cookie ?? ""
        }
    });

    const history = await res.json();
    const value = await valueRes.json();

    return {
        props: {
            history: history,
            value: value
        }
    };
}

export default Wallet;
