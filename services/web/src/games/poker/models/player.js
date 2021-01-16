class Player {
    constructor(username, id, buyIn) {
        this.username = username;
        this.id = id;
        this.buyIn = buyIn;
        this.bet = 0;
        this.cards = [];
        this.in = true;
        this.isLastAction = false;
        this.waiting = false;
    }

    setCards(cards) {
        this.cards = [];
        this.cards.concat(cards);
    }

    setBet(bet) {
        this.bet = bet;
    }

    setIn(val) {
        this.in = val;
    }

    reset() {
        this.isLastAction = false;
        this.waiting = false;
        this.cards = [];
        this.bet = 0;
        this.in = true;
    }
}

export { Player };
