import { IPlayer } from "./player";
import { ICard } from "./card";
import { ILobby } from "./lobby";

//The interface for a poker game. Includes all necessary information.
export interface IPoker {
    players: IPlayer[];
    index: number;
    dealer: number;
    board: ICard[];
    pot: string;
    potNum: number;
    bet: string;
    betNum: number;
    lobbyInfo: ILobby;
    gameRunning: boolean;
    possibleActions: number;
    notification: string;
    connected: boolean;
    loaded: boolean;
}

export const PokerInitState: IPoker = {
    players: [],
    index: -1,
    dealer: -1,
    board: [],
    pot: "€0.0",
    potNum: 0.0,
    bet: "€0.0",
    betNum: 0.0,
    lobbyInfo: {
        lobbyId: "",
        blind: 0,
        minBuyIn: 0,
        maxBuyIn: 0,
        minPlayersToStart: 3,
        playerCount: 0,
        gameStartTimeout: 15
    },
    gameRunning: false,
    possibleActions: 0,
    notification: "",
    connected: false,
    loaded: false
};
