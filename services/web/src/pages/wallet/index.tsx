import React, { FC } from "react";
import Layout from "../../components/layout";
import Front from "../../components/layout/front";
import Dinero from "dinero.js";
import ActionMenu from "../../components/actionMenu";
import TransactionList from "../../widgets/account/wallet/transactionList";
import BackToAccount from "../../widgets/account/back";
import ActionMenuLink from "../../components/actionMenu/link";
import { getSession } from "../../hooks/auth";
import moment from "moment";
import { GetServerSideProps } from "next";
import { serverSideTranslations } from "next-i18next/serverSideTranslations";
import { useTranslation } from "next-i18next";

type WalletHeaderProps = {
    value: string;
};

const WalletHeader: FC<WalletHeaderProps> = ({ value }) => {
    const { t } = useTranslation("wallet");

    if (!value) {
        return;
    }
    return (
        <Front>
            <div className="md:px-10 font-sans text-5xl font-semibold text-center grid grid-cols-2 justify-center items-center">
                <h1 className="text-6xl h-auto align-middle">{value}</h1>
                <h1>{t("Your wallet")}</h1>
            </div>
        </Front>
    );
};

export interface History {
    type: string;
    game: string;
    lobby: string;
    time: string;
    amount: string;
}

type WalletProps = {
    balance: string;
    history: History[];
};

const Wallet: FC<WalletProps> = ({ balance, history }) => {
    const { t } = useTranslation("wallet");

    return (
        <Layout>
            <BackToAccount />
            {balance && <WalletHeader value={balance} />}
            <div className="px-10 pb-10 bg-gray-200">
                <ActionMenu>
                    <ActionMenuLink href="/wallet/deposit">{t("Deposit")}</ActionMenuLink>
                </ActionMenu>
            </div>
            <div className="p-10 bg-white">{history && <TransactionList transactions={history} />}</div>
        </Layout>
    );
};

export const getServerSideProps: GetServerSideProps = async (ctx) => {
    const { req } = ctx;
    const session = getSession();

    if (!session) {
        console.log("No session");
        return {
            redirect: {
                destination: "/",
                permanent: true
            },
            props: {
                ...(await serverSideTranslations(ctx.locale, ["common", "wallet"]))
            }
        };
    }

    console.log("getting history");
    const history = await getHistory(req);

    console.log("getting balance");
    const balance = await getBalance(req);

    return {
        props: {
            history: history,
            balance: balance,
            ...(await serverSideTranslations(ctx.locale, ["common", "wallet"]))
        }
    };
};

async function getBalance(req): Promise<string> {
    console.log("API ADDRESS", process.env.API_ADRESS);
    const res = await fetch(`${process.env.API_ADRESS}/api/bank/balance`, {
        headers: {
            cookie: req.headers.cookie ?? ""
        }
    });

    const value = await res.json();
    return Dinero({ amount: value?.balance?.value, currency: value?.balance?.currency }).toFormat();
}

async function getHistory(req) {
    const res = await fetch(`${process.env.API_ADRESS}/api/bank/history`, {
        headers: {
            cookie: req.headers.cookie ?? ""
        }
    });

    const value = await res.json();

    const history = [];
    for (let i = 0; i < value.history?.length; i++) {
        if (!value.history[i].amount) {
            continue;
        }

        console.log("Game: ", value.history[i].amount.gameId, " Lobby: ", value.history[i].amount.roundId);
        history.push({
            amount: Dinero({ amount: value.history[i].amount.value, currency: value.history[i].amount.currency }).toFormat(),
            time: moment(value.history[i].time).format("MMMM Do YYYY, h:mm:ss a"),
            game: value.history[i].gameId,
            lobby: value.history[i].roundId,
            type: value.history[i].type
        });
    }

    return history;
}

export default Wallet;
