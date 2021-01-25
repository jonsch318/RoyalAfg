import React, { createContext, useState } from "react";
import PropTypes from "prop-types";
import Layout from "../../../components/layout";
import Join from "../../../games/poker/join";
import Lobbies from "../../../games/poker/lobbies";
import { useRouter } from "next/router";

export const PokerInfoContext = createContext({});

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
        <Layout footerAbsolute>
            <PokerInfoContext.Provider value={{ lobby, setLobby }}>
                <h1 className="text-center font-sans font-bold text-3xl my-10">Join A Poker Match</h1>
                <Join onJoin={join} classes={info?.classes} />
                <Lobbies info={info} />
            </PokerInfoContext.Provider>
        </Layout>
    );
};

export async function getServerSideProps() {
    try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_POKER_TICKET_HOST}/api/poker/pokerinfo`, {
            method: "GET",
            mode: "cors",
            credentials: "include"
        });

        if (res.ok) {
            const info = await res.json();
            return {
                props: {
                    info: info
                }
            };
        }
    } catch (e) {
        console.log("e", e);
        return { props: {} };
    }

    return { props: {} };

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
