import React from "react";
import {
    ACTION_PROCESSED,
    DEALER_SET,
    FLOP,
    GAME_END,
    GAME_START,
    HOLE_CARDS,
    JOIN_SUCCESS,
    LOBBY_INFO,
    PLAYER_JOIN,
    PLAYER_LEAVE,
    RIVER,
    TURN,
    WAIT_FOR_PLAYER_ACTION
} from "../events/constants";
import { IPlayer, Player } from "../models/player";
import { IPoker, PokerInitState } from "../models/poker";
import { IEvent } from "../models/event";
import { TFunction } from "next-i18next";

//FOLD describes the action of a player quiting this hand
export const FOLD = 1;

//BET describes the action of a player betting the same amount as the highest bet and therefore go along or calling the hand
export const CALL = 2;

//RAISE raises sets the highest bet a certain amount
export const RAISE = 3;

//CHECK action pushes the action requirement to the next player
export const CHECK = 4;

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const _addPlayers = (data: any, setIn = false): IPlayer[] => {
    const players = [];
    for (let i = 0; i < data.length; i++) {
        players.push(new Player(data[i].id, data[i].username, data[i].buyIn, data[i].buyInNum, setIn));
    }
    return players;
};

const _findPlayer = (id: string, arr: IPlayer[]): number => {
    return arr.findIndex((x) => x.id === id);
};

export const OnMessage = (setPoker: React.Dispatch<React.SetStateAction<IPoker>>, e: IEvent, t: TFunction): void => {
    switch (e.event) {
        case LOBBY_INFO:
            setPoker((p) => {
                return { ...p, lobbyInfo: e.data, gameRunning: e.data.gameStarted };
            });
            break;
        case JOIN_SUCCESS:
            console.log("Self: BuyIn(", e.data.players[e.data.position].buyIn, ")");
            setPoker((p) => {
                const players = _addPlayers(e.data.players);

                players[e.data.position].buyIn = e.data.buyIn;
                players[e.data.position].buyInNum = e.data.buyInNum;

                return { ...p, players: players, index: e.data.position, gameRunning: e.data.gameStarted };
            });
            break;
        case PLAYER_JOIN:
            //console.log("Player [ " + e.data.player.username + " ] joined: " + e.data.player.buyIn);
            setPoker((p) => {
                return {
                    ...p,
                    players: [...p.players, new Player(e.data.id, e.data.username, e.data.player.buyIn, e.data.player.buyInNum)],
                    lobbyInfo: {
                        ...p.lobbyInfo,
                        playerCount: e.data.playerCount
                    }
                };
            });
            break;
        case PLAYER_LEAVE:
            setPoker((p) => {
                const players = [...p.players];
                if (p.players[e.data.index].id === e.data.player.id) {
                    players.splice(e.data.index, 1);
                }

                return {
                    ...p,
                    players: players,
                    lobbyInfo: {
                        ...p.lobbyInfo,
                        playerCount: e.data.playerCount
                    }
                };
            });
            break;
        case GAME_START:
            console.log("GAME START");
            setPoker((p) => {
                return {
                    ...p,
                    dealer: -1,
                    board: [],
                    notification: t("Game started"),
                    players: _addPlayers(e.data.players, true),
                    pot: e.data.pot,
                    potNum: e.data.potNum,
                    index: e.data.position,
                    loaded: true,
                    gameRunning: true
                };
            });
            break;
        case DEALER_SET:
            setPoker((p) => {
                if (p.players[e.data.index].id !== e.data.player) {
                    console.log("Dealer does not match the player array");
                }
                return { ...p, dealer: e.data.index };
            });
            break;
        case HOLE_CARDS:
            setPoker((p) => {
                const players = [...p.players];
                for (let i = 0; i < players.length; i++) {
                    players[i].cards = [
                        { value: -1, color: -1 },
                        { value: -1, color: -1 }
                    ];
                }
                players[p.index].cards = e.data.cards;
                return {
                    ...p,
                    players: players
                };
            });
            break;

        case WAIT_FOR_PLAYER_ACTION:
            setPoker((p) => {
                const players = [...p.players];

                if (e.data.player.id !== players[e.data.position].id) {
                    console.log("Waiting for player position does not match the players array");
                }

                players[e.data.position].waiting = true;
                return {
                    ...p,
                    waitingFor: e.data.position,
                    players: players,
                    possibleActions: e.data.position === p.index ? e.data.possibleActions : 0
                };
            });
            break;
        case ACTION_PROCESSED:
            setPoker((p) => {
                const players = [...p.players];
                const i = _findPlayer(e.data.player.id, players);
                if (i !== e.data.position) {
                    console.log("Action processed position does not match players array");
                }
                const player = players[e.data.position];
                player.waiting = false;
                if (e.data.action === FOLD) {
                    player.in = false;
                }

                player.bet = e.data.totalAmount;
                player.betNum = e.data.totalAmountNum;
                player.buyIn = e.data.wallet;
                player.buyInNum = e.data.walletNum;

                console.log("Player [", e.data.position, "] betting ", e.data.totalAmount);

                return {
                    ...p,
                    pot: e.data.pot,
                    potNum: e.data.potNum,
                    bet: e.data.totalAmount,
                    betNum: e.data.totalAmountNum,
                    players: players
                };
            });
            break;
        case FLOP:
            setPoker((p) => {
                return { ...p, board: e.data.cards };
            });
            break;
        case TURN:
            setPoker((p) => {
                return {
                    ...p,
                    board: [...p.board, ...e.data.cards]
                };
            });
            break;
        case RIVER:
            setPoker((p) => {
                return {
                    ...p,
                    board: [...p.board, ...e.data.cards]
                };
            });
            break;

        case GAME_END:
            setPoker((p) => {
                const players = [...p.players];
                for (let i = 0; i < players.length; i++) {
                    players[i].reset();
                }

                let w = "";
                for (let i = 0; i < e.data.winners.length; i++) {
                    const v = e.data.winners[i];
                    const j = _findPlayer(v.id, players);
                    if (j >= 0) {
                        players[j].buyIn = v.buyIn;
                        w += v.username + (i === e.data.winners.length - 1 ? ", " : "");
                    }
                }
                if (e.data.winners.length < 1) {
                    w = "No one. All players folded!";
                }
                return {
                    ...p,
                    notification: t("Game ended. Winner is: ") + w,
                    pot: PokerInitState.pot,
                    bet: PokerInitState.bet,
                    possibleActions: PokerInitState.possibleActions,
                    board: [],
                    gameRunning: false,
                    loaded: false,
                    lobbyInfo: { ...p.lobbyInfo }
                };
            });
            break;
    }
};
