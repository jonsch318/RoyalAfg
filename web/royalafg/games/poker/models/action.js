//FOLD descibes the action of a player quiting this hand
export const FOLD = 1;

//BET descibes the action of a player betting the same amount as the highes bet and therefore go along or callling the hand
export const BET = 2;

//RAISE raises sets the highest bet a certain amount
export const RAISE = 3;

//CHECK action pushes the action requirement to the next player
export const CHECK = 4;

class Action {
    constructor(action, position, amount) {
        this.action = action;
        this.position = position;
        this.amount = amount;
    }
}

export { Action };
