const {
    JOIN_SUCCESS,
    GAME_START,
    DEALER_SET,
    WAIT_FOR_PLAYER_ACTION,
    ACTION_PROCESSED,
    FLOP,
    TURN,
    RIVER,
    GAME_END,
    HOLE_CARDS,
    PLAYER_JOIN
} = require("../events/constants");
const { BET, RAISE, FOLD, Action } = require("../models/action");
const { Player } = require("../models/player");

class GameState {
    constructor() {
        this.state = {
            players: [],
            player: -1,
            dealer: -1,
            board: [],
            bet: 0,
            roundState: 0,
            waitingFor: -1,
            gameState: -2,
            smallBlind: 0,
            bigBlind: 0,
            lastAction: -1,
            notifications: []
        };
        this.name = "GameState test";

        this.updateQueue = [];
    }

    setOnPossibleActions(onPossibleActions) {
        this.onPossibleAction = onPossibleActions.bind(this);
    }

    setOnGameEnd(onGameEnd) {
        this.onGameEnd = onGameEnd.bind(this);
    }

    setOnGameStart(handler) {
        this.onGameStart = handler;
    }

    setOnUpdate(onUpdate) {
        this.onUpdate = onUpdate.bind(this);
    }

    setOnNotification(onNotification) {
        this.onNotification = onNotification;
    }

    decodeChange(e) {
        switch (e.event) {
            case JOIN_SUCCESS:
                this.state.roundState = -2;
                this.state.player = e.data.position;
                for (let i = 0; i < e.data.players.length; i++) {
                    if (
                        this.state.players.find((obj) => {
                            return obj.id === e.data.players[i].id;
                        })
                    ) {
                        continue;
                    }
                    this.state.players.push(new Player(e.data.players[i].username, e.data.players[i].id, e.data.players[i].buyIn));
                    console.log("In Lobby it [", e.data.players[i].username, "] joined: ", e.data.players[i].buyIn);
                }
                this.stateBuild = true;
                this.updateQueue.push({ event: UpdateEvents.lobbyJoin });
                break;

            case PLAYER_JOIN:
                console.log("Player [", e.data.player.username, "] joined: ", e.data.player.buyIn);
                this.state.players.push(new Player(e.data.player.username, e.data.player.id, e.data.player.buyIn));
                if (this.updateQueue[this.updateQueue.length - 1]?.event !== UpdateEvents.playerList) {
                    this.updateQueue.push({ event: UpdateEvents.playerList });
                }
                break;

            case GAME_START:
                this.resetState();
                this.state.roundState = -1;
                this.state.player = e.data.position;
                this.state.notifications.push({ text: "Game starts.", static: false });
                this.updateQueue.push({ event: UpdateEvents.notification });
                this.onGameStart();
                this.updateQueue.push({ event: UpdateEvents.gameStart });
                console.log("Players: ", this.state.players);
                break;
            case DEALER_SET:
                this.state.dealer = e.data;
                this.updateQueue.push({ event: UpdateEvents.dealer, data: e.data });
                break;

            case HOLE_CARDS:
                for (let i = 0; i < this.state.players.length; i++) {
                    this.state.players[i].cards = [
                        { color: -1, value: -1 },
                        { color: -1, value: -1 }
                    ];
                }
                this.state.players[this.state.player].cards = e.data.cards;
                this.updateQueue.push({ event: UpdateEvents.updateAllPlayers });
                break;

            case WAIT_FOR_PLAYER_ACTION:
                this.state.waitingFor = e.data.position;
                this.state.players[e.data.position].waiting = true;
                this.updateQueue.push({ event: UpdateEvents.player, data: e.data.position });
                if (e.data.position === this.state.player) {
                    this.onPossibleAction(e.data.possibleActions);
                } else {
                    this.onPossibleAction(0);
                }
                break;

            case ACTION_PROCESSED:
                if (this.state.lastAction > -1) {
                    const lastIndex = this.state.lastAction;
                    this.state.players[lastIndex].isLastAction = false;
                    this.updateQueue.push({ event: UpdateEvents.player, data: lastIndex });
                }
                const player = this.state.players[e.data.position];
                player.waiting = false;
                const action = new Action(e.data.action, e.data.position, e.data.amount);
                player.lastAction = action;
                player.isLastAction = true;
                if (action.action === FOLD) {
                    player.in = false;
                }
                if (action.action === BET || action.action === RAISE) {
                    this.state.bet = e.data.totalAmount;
                }

                console.log("Player [", e.data.position, "] betting ", e.data.totalAmount);
                this.state.players[e.data.position].bet = e.data.totalAmount;
                this.state.players[e.data.position].buyIn = e.data.wallet;
                this.state.lastAction = e.data.position;
                this.state.waitingFor = -1;
                this.updateQueue.push({ event: UpdateEvents.player, data: e.data.position });
                break;

            case FLOP:
                this.state.roundState = 1;
                this.state.board = e.data.cards;
                this.updateQueue.push({ event: UpdateEvents.board });

                break;
            case TURN:
                this.state.roundState = 2;
                this.state.board = this.state.board.concat(e.data.cards);
                this.updateQueue.push({ event: UpdateEvents.board });

                break;
            case RIVER:
                this.state.roundState = 3;
                this.state.board = this.state.board.concat(e.data.cards);
                this.updateQueue.push({ event: UpdateEvents.board });

                break;

            case GAME_END:
                this.state.roundState = 4;

                for (let i = 0; i < this.state.players.length; i++) {
                    this.state.players[i].bet = 0;
                }

                let w = "";
                for (let i = 0; i < e.data.winners.length; i++) {
                    const v = e.data.winners[i];
                    this.state.players.find((el) => el.id === v.id).buyIn += e.data.share;
                    w += v.username + ", ";
                }

                this.state.notifications.push({
                    text: "Game ended. Winner is " + w,
                    static: false
                });
                this.updateQueue.push({ event: UpdateEvents.updateAllPlayers });
                this.updateQueue.push({ event: UpdateEvents.notification });
                this.updateQueue.push({ event: UpdateEvents.gameEnd });
                break;
            default:
                break;
        }
    }

    resetState() {
        this.state.dealer = -1;
        this.state.board = [];
        this.state.bet = 0;
        for (let i = 0; i < this.state.players.length; i++) {
            this.state.players[i].reset();
        }
        this.state.waitingFor = -1;
        this.state.smallBlind = 0;
        this.state.bigBlind = 0;
        this.state.lastAction = -1;
        this.state.notifications = [];
        this.state.gameState = -2;
        this.onPossibleAction(0);

        this.updateQueue.push({ event: UpdateEvents.updateAllPlayers });
        this.updateQueue.push({ event: UpdateEvents.boardReset });
    }

    getPlayerState(position) {
        return this.state.players[position];
    }
}

export { GameState };

export const UpdateEvents = {
    gameStart: 1,
    playerList: 2,
    player: 3,
    board: 4,
    gameEnd: 5,
    dealer: 6,
    updateAllPlayers: 7,
    lobbyJoin: 8,
    notification: 9,
    boardReset: 10
};
