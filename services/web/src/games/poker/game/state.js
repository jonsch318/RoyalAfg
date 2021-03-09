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
    PLAYER_JOIN,
    LOBBY_INFO,
    PLAYER_LEAVE
} = require("../events/constants");
const { BET, RAISE, FOLD, Action } = require("../models/action");
const { Player } = require("../models/player");

class GameState {
    constructor(setLobbyInfo) {
        this.state = {
            players: [],
            player: -1,
            dealer: -1,
            board: [],
            bet: 0,
            roundState: 0,
            waitingFor: -1,
            gameState: -2,
            bigBlind: 0,
            lastAction: -1,
            toStart: 0,
            wallet: "",

            notifications: []
        };
        this.lobby = {
            lobbyId: "",
            blind: 0,
            minBuyIn: 0,
            maxBuyIn: 0,
            gameStarted: false,
            minPlayersToStart: 0,
            playerCount: 0,
            gameStartTimeout: 0
        };
        this.pot = "";
        this.name = "GameState test";
        this.updateQueue = [];
        this.setLobbyInfo = setLobbyInfo;
    }

    setOnPossibleActions(onPossibleActions) {
        this.onPossibleAction = onPossibleActions.bind(this);
    }

    setOnGameStart(handler) {
        this.onGameStart = handler;
    }

    getPlayerCount() {
        return this.state.players.length;
    }

    decodeChange(e) {
        switch (e.event) {
            case LOBBY_INFO:
                this.lobby = e.data;
                break;

            case JOIN_SUCCESS:
                this.state.roundState = -2;
                this.state.player = e.data.position;
                this.state.players = [];
                this._addPlayers(e.data.players);
                this.state.wallet = e.data.wallet;
                this.update({ event: UpdateEvents.lobbyJoin });
                break;

            case PLAYER_JOIN:
                console.log("Player [", e.data.player.username, "] joined: ", e.data.player.buyIn);
                this.state.players.push(new Player(e.data.player.username, e.data.player.id, e.data.player.buyIn));
                if (this.updateQueue[this.updateQueue.length - 1]?.event !== UpdateEvents.playerList) {
                    this.update({ event: UpdateEvents.playerList });
                }
                this.lobby.playerCount = e.data.playerCount;
                console.log({ count: this.lobby.playerCount, toStart: this.lobby.minPlayersToStart });
                this.setLobbyInfo({
                    count: this.lobby.playerCount,
                    toStart: this.lobby.minPlayersToStart,
                    timeout: this.lobby.gameStartTimeout,
                    gameStarted: e.data.gameStarted
                });
                break;
            case PLAYER_LEAVE:
                if (this.state.players[e.data.index].id === e.data.player.id) {
                    this.state.players.splice(e.data.index, 1);
                }
                this.lobby.playerCount = e.data.playerCount;
                this.lobby.gameStarted = e.data.gameStarted;
                this.setLobbyInfo({
                    count: this.lobby.playerCount,
                    toStart: this.lobby.minPlayersToStart,
                    timeout: this.lobby.gameStartTimeout,
                    gameStarted: this.lobby.gameStarted
                });
                this.update({ event: UpdateEvents.playerList });
                break;
            case GAME_START:
                this.resetState();
                this.state.roundState = -1;
                this.state.player = e.data.position;
                this.state.notifications.push({ text: "Game starts...", static: false });
                this.update({ event: UpdateEvents.notification });
                this.onGameStart();
                this.state.players = [];
                this._addPlayers(e.data.players);
                this.pot = e.data.pot;
                this.update({ event: UpdateEvents.gameStart });
                console.log("Players: ", this.state.players);
                break;
            case DEALER_SET:
                this.state.dealer = e.data;
                this.update({ event: UpdateEvents.dealer, data: e.data });
                break;

            case HOLE_CARDS:
                for (let i = 0; i < this.state.players.length; i++) {
                    this.state.players[i].cards = [
                        { color: -1, value: -1 },
                        { color: -1, value: -1 }
                    ];
                }
                this.state.players[this.state.player].cards = e.data.cards;
                this.update({ event: UpdateEvents.updateAllPlayers });
                break;

            case WAIT_FOR_PLAYER_ACTION:
                this.state.waitingFor = e.data.position;
                this.state.players[e.data.position].waiting = true;
                this.update({ event: UpdateEvents.player, data: e.data.position });
                if (e.data.position === this.state.player) {
                    this.onPossibleAction(e.data.possibleActions);
                } else {
                    this.onPossibleAction(0);
                }
                break;

            case ACTION_PROCESSED:
                this.handleActionProcessed(e);
                break;

            case FLOP:
                this.state.roundState = 1;
                this.state.board = e.data.cards;
                this.update({ event: UpdateEvents.board });

                break;
            case TURN:
                this.state.roundState = 2;
                this.state.board = this.state.board.concat(e.data.cards);
                this.update({ event: UpdateEvents.board });

                break;
            case RIVER:
                this.state.roundState = 3;
                this.state.board = this.state.board.concat(e.data.cards);
                this.update({ event: UpdateEvents.board });

                break;

            case GAME_END:
                this._handleGameEnd(e);
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
        this.state.bigBlind = 0;
        this.state.lastAction = -1;
        this.state.notifications = [];
        this.state.gameState = -2;
        this.onPossibleAction(0);

        this.update({ event: UpdateEvents.updateAllPlayers });
        this.update({ event: UpdateEvents.boardReset });
    }

    getPlayerState(position = 0) {
        return this.state.players[position];
    }

    update(update = {}) {
        this.updateQueue.push(update);
    }

    handleActionProcessed(e = {}) {
        if (this.state.lastAction > -1) {
            const lastIndex = this.state.lastAction;
            this.state.players[lastIndex].isLastAction = false;
            this.update({ event: UpdateEvents.player, data: lastIndex });
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
        this.pot = e.data.pot;
        this.update({ event: UpdateEvents.player, data: e.data.position });
    }

    _handleGameEnd(e) {
        this.state.roundState = 4;

        for (let i = 0; i < this.state.players.length; i++) {
            this.state.players[i].bet = 0;
        }

        let w = "";
        for (let i = 0; i < e.data.winners.length; i++) {
            const v = e.data.winners[i];
            const player = this._findPlayer(v.id);
            if (player !== undefined) {
                player.buyIn = v.buyIn;
                w += v.username + (i === e.data.winners.length - 1 ? ", " : "");
            }
        }

        console.log("Winners", e.data.winners);

        this.state.notifications.push({
            text: "Game ended. Winner is " + w,
            static: false
        });
        this.update({ event: UpdateEvents.updateAllPlayers });
        this.update({ event: UpdateEvents.notification });
        this.update({ event: UpdateEvents.gameEnd });
    }

    _findPlayer(id = "") {
        return this.state.players.find((obj) => {
            return obj.id === id;
        });
    }

    _addPlayers(players) {
        for (let i = 0; i < players.length; i++) {
            this.state.players.push(new Player(players[i].username, players[i].id, players[i].buyIn));
            console.log("In Lobby it [", players[i].username, "] joined: ", players[i].buyIn);
        }
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
