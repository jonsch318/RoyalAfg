import React, { createContext, useState } from "react";
import PropTypes from "prop-types";
import Layout from "../../../components/layout";
import Join from "../../../games/poker/join";
import Lobbies from "../../../games/poker/lobbies";
import { useRouter } from "next/router";
import Head from "next/head";
import { formatTitle } from "../../../utils/title";

export const PokerInfoContext = createContext({});

// eslint-disable-next-line react/prop-types
const PokerConnectError = ({ onRefresh, onBack }) => {
    return (
        <div className="bg-gray-200">
            <div className="bg-gray-200 p-20 pt-32 grid justify-center items-center">
                <h1 className="font-sans text-5xl font-bold text-center inlines bg-white p-8 rounded-xl">Unable to connect to Poker Matchmaking</h1>
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

const Poker = ({ info }) => {
    const router = useRouter();
    const [lobby, setLobby] = useState({ i: -1, classIndex: -1 });

    const join = (params) => {
        router
            .push({
                pathname: "/games/poker/play",
                query: {
                    lobbyId: params.lobbyId,
                    buyInClass: params.class,
                    buyIn: params.buyIn
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

export async function getServerSideProps() {
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
                redirect: {
                    destination: "/games",
                    permanent: true
                }
            };
        }

        const info = await res.json();
        console.log("Info not expected: ", info);

        if (!info || !info?.classes || !info?.classes.length || !info?.lobbies || !info?.lobbies.length) {
            return {
                redirect: {
                    destination: "/games",
                    permanent: true
                }
            };
        }
        return {
            props: {
                info: info
            }
        };
    } catch (e) {
        console.log("e", e);
        return {
            props: {},
            redirect: {
                destination: "/games",
                permanent: true
            }
        };
    }

    /*  return {
        props: {
            info: {
                classes: [
                    { min: 1000, max: 4999, blind: 100 },
                    { min: 5000, max: 14999, blind: 500 }
                ],
                lobbies: [
                    [{ id: "abc", class: { min: 1000, max: 4999, blind: 100 }, classIndex: 0 }],
                    [{ id: "def", class: { min: 5000, max: 14999, blind: 500 }, classIndex: 1 }]
                ]
            }
        }
    };*/
}

Poker.propTypes = {
    info: PropTypes.shape({
        lobbies: PropTypes.array,
        classes: PropTypes.array
    })
};

export default Poker;
