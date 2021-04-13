import { ICard } from "./card";

//a poker player
export interface IPlayer {
    username: string;
    readonly id: string;
    buyIn: string;
    bet: string;
    cards: ICard[];
    in: boolean;
    waiting: boolean;

    reset(): void;
}

export class Player implements IPlayer {
    bet: string;
    buyIn: string;
    cards: ICard[];
    readonly id: string;
    in: boolean;
    username: string;
    waiting: boolean;

    constructor(id: string, username: string, buyIn = "€0.0", active = false) {
        this.id = id;
        this.buyIn = buyIn;
        this.bet = "€0.0";
        this.cards = [];
        this.in = active;
        this.username = username;
        this.waiting = false;
    }

    reset(): void {
        this.bet = "€0.0";
        this.in = true;
        this.waiting = false;
        this.cards = [];
    }
}
