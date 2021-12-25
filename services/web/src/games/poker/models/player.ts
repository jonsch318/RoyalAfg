import { ICard } from "./card";

//a poker player
export interface IPlayer {
    username: string;
    readonly id: string;
    buyIn: string;
    buyInNum: number;
    bet: string;
    betNum: number;
    cards: ICard[];
    in: boolean;
    waiting: boolean;

    reset(): void;
}

export class Player implements IPlayer {
    bet: string;
    betNum: number;
    buyIn: string;
    buyInNum: number;
    cards: ICard[];
    readonly id: string;
    in: boolean;
    username: string;
    waiting: boolean;

    constructor(id: string, username: string, buyIn = "€0.0", buyInNum = 0.0, active = false) {
        this.id = id;
        this.buyIn = buyIn;
        this.buyInNum = buyInNum;
        this.bet = "€0.0";
        this.betNum = 0.0;
        this.cards = [];
        this.in = active;
        this.username = username;
        this.waiting = false;
    }

    reset(): void {
        this.bet = "€0.0";
        this.betNum = 0.0;
        this.in = true;
        this.waiting = false;
        this.cards = [];
    }
}
