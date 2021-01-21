import React, { useState } from "react";
import Layout from "../../../components/layout";
import Join from "../../../games/poker/join";
import Lobbies from "../../../games/poker/lobbies";
import { useRouter } from "next/router";
const Poker = () => {
    const router = useRouter();

    const [selectedLobby, setSelectedLobby] = useState({ id: "", lobbyClass: { i: -1, c: [] } });

    const join = (params) => {
        router
            .push({
                pathname: "/games/poker/play",
                query: {
                    lobbyId: params.lobbyId,
                    buyInClass: params.buyInClass,
                    buyIn: params.buyin
                }
            })
            .then();
    };

    return (
        <Layout footerAbsolute>
            <h1 className="text-center font-sans font-bold text-3xl my-10">Join A Poker Match</h1>
            <Join
                onJoin={join}
                lobbyId={selectedLobby.id}
                buyInClass={selectedLobby.lobbyClass.i}
                minBuyIn={selectedLobby.lobbyClass.c[0]}
                maxBuyIn={selectedLobby.lobbyClass.c[1]}
            />
            <Lobbies
                onLobbySelect={(id, lobbyClassIndex, lobbyCass) => {
                    setSelectedLobby({ id: id, lobbyClass: { i: lobbyClassIndex, c: lobbyCass } });
                }}
            />
        </Layout>
    );
};

Poker.propTypes = {};

export default Poker;
