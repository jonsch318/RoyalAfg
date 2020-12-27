import React from "react";

import dynamic from "next/dynamic";
const Actions = dynamic(import("../../../games/poker/actions.js"), { ssr: false });
const Join = dynamic(import("../../../games/poker/join.js"), { ssr: false });
const Lobbies = dynamic(import("../../../games/poker/lobbies.js"), { ssr: false });
const View = dynamic(import("../../../games/poker/view"), { ssr: false });
import { GameState } from "../../../games/poker/game/state.js";
import { Game } from "../../../games/poker/connection/socket.js";

import "../poker.module.css";
import Layout from "../../../components/layout";

class Play extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            id: "",
            lobbyClass: { i: 0, c: [0, 0, 0] },
            game: {},
            credentials: {},
            joined: false,
            possibleActions: 0
        };
    }

    possibleActionsChange(actions) {
        this.setState({ possibleActions: actions });
    }

    start(cred) {
        let gameState = new GameState();
        gameState.setOnPossibleActions(this.possibleActionsChange.bind(this));
        let game = new Game(gameState, cred, () => {
            this.setState({ joined: false });
        });
        game.start();
        this.setState({ game: game, joined: true });
    }

    onJoin(values) {
        this.start(values);
    }

    onLobbySelect(id, lobbyClassIndex, lobbyClass) {
        this.setState({ id: id, lobbyClass: { i: lobbyClassIndex, c: lobbyClass } });
    }

    render() {
        return (
            <Layout footerAbsolute="true">
                <div className="App">
                    {this.state.joined ? (
                        <div>
                            <Actions game={this.state.game} actions={this.state.possibleActions} />
                            <View game={this.state.game} />
                        </div>
                    ) : (
                        <div>
                            <Join
                                onJoin={this.onJoin.bind(this)}
                                lobbyId={this.state.id}
                                buyInClass={this.state.lobbyClass.i}
                                minBuyIn={this.state.lobbyClass.c[0]}
                                maxBuyIn={this.state.lobbyClass.c[1]}
                            />
                            <Lobbies onLobbySelect={this.onLobbySelect.bind(this)} />
                        </div>
                    )}
                </div>
            </Layout>
        );
    }
}
export default Play;
