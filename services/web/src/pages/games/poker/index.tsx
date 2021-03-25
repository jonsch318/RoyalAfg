import React, { createContext, FC, useEffect, useState } from "react";
import Layout from "../../../components/layout";
import Join, { JoinOptions } from "../../../games/poker/join";
import Lobbies from "../../../games/poker/lobbies";
import { useRouter } from "next/router";
import Head from "next/head";
import { formatTitle } from "../../../utils/title";
import { useSnackbar } from "notistack";
import { IClass, ILobby, LobbyInit } from "../../../games/poker/models/class";
import { GetServerSideProps } from "next";

type PokerInfoContext = {
    lobby: ILobby;
    setLobby: React.Dispatch<React.SetStateAction<ILobby>>;
};
export const PokerInfoContext = createContext<PokerInfoContext | undefined>(undefined);

const PokerConnectError = ({ onRefresh, onBack }) => {
    return (
        <div className="bg-gray-200">
            <div className="bg-gray-200 p-20 pt-32 grid justify-center items-center">
                <h1 className="font-sans text-5xl font-bold text-center inlines bg-white p-8 rounded-xl">
                    Unable to connect to the Poker matchmaking server
                </h1>
            </div>
            <div className="pb-72 pt-32">
                <div className="flex items-center justify-center bg-gray-200">
                    <button
                        className="mr-16 ml-auto bg-black text-white px-8 py-4 rounded hover:scale-105 transition-transform transform"
                        onClick={() => onBack()}>
                        Go Back
                    </button>
                    <button
                        className="ml-16 mr-auto bg-black text-white px-8 py-4 rounded hover:scale-105 transition-transform transform"
                        onClick={() => onRefresh}>
                        Refresh
                    </button>
                </div>
            </div>
        </div>
    );
};

type PokerProps = {
    error: string;
    info: {
        lobbies: ILobby[][];
        classes: IClass[];
    };
};

const Poker: FC<PokerProps> = ({ info, error }) => {
    const router = useRouter();
    const [lobby, setLobby] = useState<ILobby>(LobbyInit);
    const { enqueueSnackbar } = useSnackbar();

    useEffect(() => {
        if (error !== "") {
            enqueueSnackbar("Unable to reach the poker server", { variant: "error" });
        }
    }, [error]);

    const join = (params: JoinOptions) => {
        router
            .push({
                pathname: "/games/poker/play",
                query: {
                    lobbyId: params.lobbyId,
                    buyIn: params.buyIn,
                    buyInClass: params.class
                }
            })
            .then();
    };

    return (
        <>
            <Head>
                <title>{formatTitle("Play Poker")}</title>
            </Head>
            <Layout footerAbsolute>
                {info && Array.isArray(info?.classes) && Array.isArray(info?.lobbies) ? (
                    <PokerInfoContext.Provider value={{ lobby, setLobby }}>
                        <h1 className="text-center font-sans font-bold text-3xl my-10">Join A Poker Match</h1>
                        <Join onJoin={join} classes={info?.classes} />
                        <Lobbies info={info} />
                    </PokerInfoContext.Provider>
                ) : (
                    <PokerConnectError onRefresh={() => router.reload()} onBack={() => router.push("/games")} />
                )}
            </Layout>
        </>
    );
};

export const getServerSideProps: GetServerSideProps = async (ctx) => {
    try {
        console.log("Calling: ", process.env.POKER_INFO_HOST ? `${process.env.POKER_INFO_HOST}/api/poker/pokerinfo` : "/api/pokerinfo");
        const res = await fetch(process.env.POKER_INFO_HOST ? `${process.env.POKER_INFO_HOST}/api/poker/pokerinfo` : "/api/pokerinfo", {
            method: "GET",
            mode: "cors",
            credentials: "include"
        });

        if (!res.ok) {
            console.log("not ok: ", res.status);
            return {
                props: {
                    error: "Status invalid: " + res.status
                }
            };
        }

        const info = await res.json();

        if (!info || !info?.classes || !info?.classes.length || !info?.lobbies || !info?.lobbies.length) {
            return {
                props: {
                    error: "Status invalid: " + res.status
                }
            };
        }
        return {
            props: {
                info: info,
                error: ""
            }
        };
    } catch (e) {
        console.log("e", e);
        return {
            props: {
                error: "Error: " + e
            }
        };
    }
};

export default Poker;
