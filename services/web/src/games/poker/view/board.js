import * as PIXI from "pixi.js";
import { Card, CARDHEIGHT, CARDWIDTH } from "./card";
import { rW, rH, isMobile } from "./utils";

class Board extends PIXI.Container {
    constructor(state) {
        super();
        this.state = state;
        this.options = {
            paddingX: 20,
            paddingY: 20,
            paddingCardX: 10,
            minWidth: 100,
            updatedWidth: () => {}
        };

        this.background = new PIXI.Graphics();
        this.cards = new PIXI.Container();

        this.addChild(this.background, this.cards);
    }

    setup(id) {
        this.id = id;
        this.update({});
    }

    push(card) {
        let itemX = this.options.paddingX + this.cards.children.length * CARDWIDTH + this.cards.children.length * 2 * this.options.paddingCardX;
        const registered = this.insertCard(itemX);
        if (card) {
            registered.update(card);
        }
        this.onResize();
    }

    addCards(count) {
        for (let i = 0; i < count; i++) {
            this.push({});
        }
    }

    updateCard(i, card) {
        if (i >= 0 && i < this.cards.children.length) {
            this.cards.children[i].update(card);
        }
    }

    updateCards(startIndex, cards = []) {
        if (startIndex >= 0 && startIndex + cards.length <= this.cards.children.length) {
            for (let i = startIndex; i < startIndex + cards.length; i++) {
                this.updateCard(i, cards[i - startIndex]);
            }
        }
    }

    pop() {
        const card = this.cards[this.cards.length - 1];
        this.cards.pop();
        this.removeChild(card);
        this.onResize();
    }

    clear() {
        this.removeChildren();
        this.cards.removeChildren();
        this.addChild(this.background, this.cards);
        this.onResize();
    }

    update(options) {
        this.options = {
            ...this.options,
            ...options
        };
        this.onResize();
    }

    pushOrUpdate(cards = []) {
        if (this.cards.children.length > 0) {
            if (cards.length > this.cards.children.length) {
                let update = cards.slice(0, this.cards.children.length);
                this.updateCards(0, update);
                for (let i = this.cards.children.length; i < cards.length; i++) {
                    this.push(cards[i]);
                }
            } else {
                this.updateCards(0, cards);
            }
        } else {
            for (let i = 0; i < cards.length; i++) {
                this.push(cards[i]);
            }
        }
    }

    onResize() {
        this.w = CARDWIDTH * this.cards.children.length + this.options.paddingX * 2 + this.options.paddingCardX * this.cards.children.length * 2;
        if (isMobile()) {
            this.h = 15 + rH(2 * this.options.paddingY);
        } else {
            this.h = rH(CARDHEIGHT + 2 * this.options.paddingY);
        }
        this.background.clear();
        this.background.beginFill(0x000000, 0.12);
        this.background.drawRoundedRect(0, 0, Math.max(rW(this.w), rW(this.options.minWidth)), this.h, rH(10));
        this.background.endFill();
        this.options.updatedWidth();
    }

    insertCard(itemX) {
        const card = new Card(this.id, { value: -1, color: -1 });
        itemX += this.options.paddingCardX;
        card.position.set(rW(itemX), rW(this.options.paddingY));
        itemX += this.options.paddingCardX + CARDWIDTH;
        this.cards.addChild(card);
        return card;
    }

    updateFromState() {
        this.pushOrUpdate(this.state.board);
    }
}
export { Board };
