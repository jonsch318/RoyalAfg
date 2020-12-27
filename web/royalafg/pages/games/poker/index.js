import React, { useState } from "react";
import Layout from "../../../components/layout";
import Join from "../../../games/poker/join";
import Lobbies from "../../../games/poker/lobbies";

const Poker = () => {
    const [selectedLobby, setSelectedLobby] = useState({ id: "", lobbyClass: { i: -1, c: [] } });

    const join = () => {
        Router.push()
    };

    return (
        <Layout>
            <div>Play Poker</div>
            <Join
                onJoin={join}
                lobbyId={selectedLobby.id}
                buyInClass={selectedLobby.lobbyClass.i}
                minBuyIn={selectedLobby.lobbyClass.c[0]}
                maxBuyIn={selectedLobby.lobbyClass.c[1]}
            />
            <Lobbies
                onLobbySelect={(id, lobbyClassIndex, lobbyCass) => {
                    setSelectedLobby({ id: id, lobbyCass: { i: lobbyClassIndex, c: lobbyCass } });
                }}
            />
        </Layout>
    );
};

Poker.propTypes = {};

export default Poker;
