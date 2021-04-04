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

//FOLD describes the action of a player quiting this hand
export const FOLD = 1;

//BET describes the action of a player betting the same amount as the highest bet and therefore go along or calling the hand
export const CALL = 2;

//RAISE raises sets the highest bet a certain amount
export const RAISE = 3;

//CHECK action pushes the action requirement to the next player
export const CHECK = 4;

const _addPlayers = (data: any): IPlayer[] => {
    const players = [];
    for (let i = 0; i < data.length; i++) {
        players.push(new Player(data[i].id, data[i].username, data[i].buyIn));
    }
    return players;
};

const _findPlayer = (id: string, arr: IPlayer[]): number => {
    return arr.findIndex((x) => x.id === id);
};

export const OnMessage = (setPoker: React.Dispatch<React.SetStateAction<IPoker>>, e: IEvent): void => {
    switch (e.event) {
        case LOBBY_INFO:
            setPoker((p) => {
                return { ...p, lobbyInfo: e.data };
            });
            break;
        case JOIN_SUCCESS:
            console.log("Self: BuyIn(", e.data.players[e.data.position].buyIn, ")");
            setPoker((p) => {
                return { ...p, players: _addPlayers(e.data.players), index: e.data.position };
            });
            break;
        case PLAYER_JOIN:
            //console.log("Player [ " + e.data.player.username + " ] joined: " + e.data.player.buyIn);
            setPoker((p) => {
                return {
                    ...p,
                    players: [...p.players, new Player(e.data.id, e.data.username, e.data.player.buyIn)],
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
            setPoker((p) => {
                console.log("Position: ", e.data.position);
                return {
                    ...p,
                    dealer: -1,
                    board: [],
                    notification: "Game Started",
                    gameRunning: true,
                    players: _addPlayers(e.data.players),
                    pot: e.data.pot,
                    index: e.data.position
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
                    if (i === e.data.player) {
                        players[i].cards = e.data.cards;
                    } else {
                        players[i].cards = [
                            { value: -1, color: -1 },
                            { value: -1, color: -1 }
                        ];
                    }
                }
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
                player.buyIn = e.data.wallet;
                console.log("Player [", e.data.position, "] betting ", e.data.totalAmount);

                return {
                    ...p,
                    pot: e.data.pot,
                    bet: e.data.totalAmount,
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
                    notification: "Game Ended. Winner is: " + w,
                    pot: PokerInitState.pot,
                    bet: PokerInitState.bet,
                    possibleActions: PokerInitState.possibleActions,
                    board: [],
                    gameRunning: false,
                    lobbyInfo: { ...p.lobbyInfo }
                };
            });
            break;
    }
};
