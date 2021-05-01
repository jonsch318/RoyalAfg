import React, { FC } from "react";
import dynamic from "next/dynamic";

import "../poker.module.css";
import { useEffect } from "react";
import { useRouter } from "next/router";
import { useSnackbar } from "notistack";
import { getCSRF } from "../../../hooks/auth/csrf";
import { GetServerSideProps } from "next";
import Head from "next/head";
import { serverSideTranslations } from "next-i18next/serverSideTranslations";
import { useTranslation } from "next-i18next";

const Poker = dynamic(import("../../../games/poker/index"), { ssr: false });

const _getUrl = (id) => {
    let url = "";
    if (process.env.NEXT_PUBLIC_POKER_TICKET_HOST !== undefined) {
        url = process.env.NEXT_PUBLIC_POKER_TICKET_HOST;
    }
    if (id) {
        console.log("Requesting ticket with ID");
        return `${url}/api/poker/ticket/${id}`;
    }
    console.log("Requesting ticket without ID");
    return `${url}/api/poker/ticket`;
};

const _fetch = async (url, params, args = {}, csrf = "", cookie: string | undefined) => {
    return await fetch(`${url}?${params.toString()}`, {
        mode: "cors",
        credentials: "include",
        method: "POST",
        headers: {
            "X-CSRF-Token": csrf,
            cookie: cookie
        },
        body: JSON.stringify({ ...args })
    });
};

export const getServerSideProps: GetServerSideProps = async (context) => {
    const csrf = await getCSRF(context);
    const { lobbyId, buyInClass, buyIn } = context.query;
    const params = new URLSearchParams();
    params.append("buyIn", buyIn.toString());
    params.append("class", buyInClass.toString());
    const res = await _fetch(
        _getUrl(lobbyId),
        params,
        {
            buyIn: buyIn,
            class: buyInClass,
            lobbyId: lobbyId
        },
        csrf,
        context.req.headers.cookie
    );
    if (!res.ok) {
        return {
            props: {
                csrf: csrf,
                ticket: { address: "", token: "" },
                error: "error during ticket. Code: " + res.status,
                ...(await serverSideTranslations(context.locale, ["common", "poker"]))
            }
        };
    }

    try {
        const ticket = await res.json();
        console.log("Ticket: ", ticket);
        return {
            props: {
                csrf: csrf,
                ticket: ticket,
                error: "",
                ...(await serverSideTranslations(context.locale, ["common", "poker"]))
            }
        };
    } catch (e) {
        console.log("Error during ticket fetch: ", e);
        return {
            props: {
                ticket: { address: "", token: "" },
                error: "error during ticket fetch: " + e,
                csrf: csrf,
                ...(await serverSideTranslations(context.locale, ["common", "poker"]))
            }
        };
    }
};

type PlayProps = {
    csrf: string;
    ticket: { address: string; token: string };
    error: string;
};

const Play: FC<PlayProps> = ({ csrf, ticket, error }) => {
    const { t } = useTranslation("poker");
    const router = useRouter();
    const { enqueueSnackbar } = useSnackbar();

    useEffect(() => {
        if (error !== "") {
            console.log("Error: ", error);
            enqueueSnackbar(t("Could not connect to the poker server"), { variant: "error" });
            router.push("/games/poker").then();
        }
    }, [error]);

    return (
        <>
            <Head>
                <title>{t("TitleGame")}</title>
            </Head>
            <Poker ticket={ticket} csrf={csrf} />
        </>
    );
};

export default Play;
