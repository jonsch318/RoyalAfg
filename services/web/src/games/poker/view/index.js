 import * as PIXI from "pixi.js";
import React from "react";
import { Board } from "./board";
import { Players } from "./players";
import { registerApp, rH, rW } from "./utils";
import PropTypes from "prop-types";
import { Game } from "../connection/socket";
import { UpdateEvents } from "../game/state";
import { Notification } from "./notifcation";
import { withRouter } from "next/router";

const textureUrl = window.location.origin + "/static/games/poker/textures/cards.json";

class View extends React.Component {
    constructor(props) {
        super(props);
        this.gameState = props.game.state;
        this.game = props.game.state.state;
        this.loader = props.loader;

        this.state = {
            loaded: false
        };
        PIXI.Loader.shared.reset();
    }

    componentDidMount() {
        this.setState({ loaded: false });

        this.props.game.state.setOnGameStart(() => {
            this.gameStart();
        });
    }

    gameStart() {
        const d = document.getElementById("view");

        this.setState({ loaded: true });
        PIXI.settings.PRECISION_FRAGMENT = PIXI.PRECISION.HIGH;
        this.app = new PIXI.Application({
            antialias: false,
            transparent: false,
            resolution: 1,
            resizeTo: window
        });

        PIXI.Renderer.registerPlugin("interaction", PIXI.InteractionManager);
        this.app.renderer.backgroundColor = 0xffffff;
        d.appendChild(this.app.view);

        registerApp(this.app);

        PIXI.Loader.shared.add(textureUrl).load(this.setup.bind(this));

        this.table = new PIXI.Graphics();
        this.board = new Board(this.game);

        this.players = new Players(this.gameState);
        this.players.position.set(0, 0);

        this.notification = new Notification(this.game, this.app.renderer.width, this.app.renderer.height);
        this.notification.position.set(0, 0);

        this.app.stage.addChild(this.table, this.board, this.players, this.notification);

        this.tableWidth = rW(200);
        this.tableHeight = rH(125);
        // if (isMobile()) {
        //     this.tableWidth = 85;
        //     this.tableHeight = 50;
        // }
        // this.table.beginFill(0x1daf08);
        // this.table.drawEllipse(this.tableWidth, this.tableHeight, this.tableWidth, this.tableHeight)
        // //table.drawRoundedRect(0, 0, this.tableWidth * 2, this.tableHeight * 2, this.tableHeight)
        // this.table.endFill();
        // this.table.position.set(this.app.renderer.width / 2 - this.tableWidth, this.app.renderer.height / 2 - this.tableHeight)

        this.board.update({
            updatedWidth: () => {
                this.board.position.set(this.app.renderer.width / 2 - this.board.width / 2, this.app.renderer.height / 2 - this.board.height / 2);
            }
        });
        this.board.position.set(this.app.renderer.width / 2 - this.board.width / 2, this.app.renderer.height / 2 - this.board.height / 2);
    }

    setup() {
        this.didSetup = true;
        if (!PIXI.Loader.shared.resources[textureUrl]) {
            this.props.router.push("/games/poker");
        }
        this.id = PIXI.Loader.shared.resources[textureUrl]?.textures;

        this.notification.reset();
        this.players.setup(this.id, {
            x: this.app.renderer.width / 2,
            y: this.app.renderer.height / 2,
            width: this.tableWidth,
            height: this.tableHeight
        });
        this.board.setup(this.id);
        this.app.ticker.add((delta) => this.gameLoop(delta));
    }

    gameLoop(delta) {
        this.players.gameLoop(delta);
        this.workUpdateQueue();
    }

    workUpdateQueue() {
        if (this.gameState.updateQueue.length > 0) {
            for (let i = 0; i < this.gameState.updateQueue.length; i++) {
                const work = this.gameState.updateQueue[0];
                this.updateFromState(work.event, work.data);
                this.gameState.updateQueue.shift();
            }
        }
    }

    updateFromState(event, data) {
        if (event === UpdateEvents.lobbyJoin) {
            this.players.updateFromState();
        }
        if (event === UpdateEvents.playerList) {
            this.players.updateFromState();
        }
        if (event === UpdateEvents.updateAllPlayers) {
            this.players.updateAllPlayersFromState();
        }
        if (event === UpdateEvents.player) {
            this.players.updatePlayerFromState(data);
        }
        if (event === UpdateEvents.dealer) {
            this.players.updatePlayerFromState(data);
        }
        if (event === UpdateEvents.board) {
            this.board.updateFromState();
        }
        if (event === UpdateEvents.notification) {
            this.notification.onNotification();
        }
        if (event === UpdateEvents.boardReset) {
            this.board.clear();
        }
    }

    render() {
        return <div id="view">{!this.state.loaded ? <h1>Waiting for next game to start.</h1> : <></>}</div>;
    }
}

View.propTypes = {
    game: PropTypes.instanceOf(Game),
    loader: PropTypes.object,
    router: PropTypes.object
};

export default withRouter(View);
