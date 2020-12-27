import * as PIXI from "pixi.js";
import { Board } from "./board";
import { isMobile } from "./utils";

class Player extends PIXI.Container {
    constructor(id, options, state, index) {
        super();

        this.state = state;
        this.playerState = this.state.getPlayerState(index);
        this.index = index;

        this.options = {
            width: 250,
            marginX: 10,
            marginY: 8,
            avatarRadius: 20,
            angle: 0,
            fontFamily: "Source Sans Pro",
            fontSize: 20,
            foreground: 0x000000,
            fillAvatar: 0xffffff,
            waitingAnimationVelocity: 5,
            background: {
                value: 0x000000,
                alpha: 0.12
            }
        };

        if (isMobile()) {
            this.options.width = 150;
        }

        this.buyInLabel = new PIXI.Text("", { align: "center" });
        this.betLabel = new PIXI.Text("", { align: "center" });
        this.usernameLabel = new PIXI.Text("", { align: "center" });
        this.avatar = new PIXI.Graphics();
        this.board = new Board(id, {
            paddingX: 8,
            paddingY: 5,
            paddingCardX: 5
        });
        this.background = new PIXI.Graphics();
        this.activeBackground = new PIXI.Graphics();
        this.waiting = new PIXI.Graphics();
        this.waitingVelocity = this.options.waitingAnimationVelocity;
        this.waitingAnimeDir = true;

        this.dealerButton = new PIXI.Graphics();

        this.topRow = new PIXI.Container();
        this.topRow.addChild(
            this.avatar,
            this.usernameLabel,
            this.betLabel,
            this.buyInLabel,
            this.dealerButton,
            this.waiting
        );
        this.addChild(this.background, this.topRow, this.board, this.activeBackground);

        this.update(options);
    }

    gameLoop(delta) {
        if (this.waitingVelocity < 3.25) {
            this.waitingAnimeDir = true;
        } else if (this.waitingVelocity > 12) {
            this.waitingAnimeDir = false;
        }
        if (this.waitingAnimeDir) {
            this.waitingVelocity += 0.075;
        } else {
            this.waitingVelocity -= 0.075;
        }
        this.waiting.angle = (this.waiting.angle + this.waitingVelocity + delta) % 360;
    }

    update(options) {
        this.options = {
            ...this.options,
            ...options
        };

        this.usernameLabel.text = this.playerState.username;
        this.usernameLabel.style = {
            fontFamily: this.options.fontFamily,
            fontSize: this.options.fontSize,
            fill: this.options.foreground
        };
        this.usernameLabel.pivot.set(this.usernameLabel.width, 0);

        this.dealerButton.clear();
        this.dealerButton.visible = false;

        if (this.state.state.dealer === this.index) {
            this.dealerButton.lineStyle(1, 0xee5015);
            this.dealerButton.beginFill(0xeeee00);
            this.dealerButton.drawCircle(0, 0, this.options.avatarRadius / 2);
            this.dealerButton.endFill();
            this.dealerButton.visible = true;
        }

        this.buyInLabel.text = this.playerState.buyIn;
        this.buyInLabel.style = {
            fontFamily: this.options.fontFamily,
            fontSize: this.options.fontSize,
            fill: 0x000000
        };
        this.buyInLabel.pivot.set(this.buyInLabel.width, 0);

        this.betLabel.text = this.playerState.bet;
        this.betLabel.style = {
            fontFamily: this.options.fontFamily,
            fontSize: this.options.fontSize,
            fill: this.state.state.lastIndex === this.index ? 0xff0000 : this.options.foreground
        };
        this.betLabel.pivot.set(this.betLabel.width, 0);

        this.avatar.clear();
        this.avatar.beginFill(this.options.fillAvatar);
        this.avatar.drawCircle(
            this.options.avatarRadius,
            this.options.avatarRadius,
            this.options.avatarRadius
        );
        this.avatar.endFill();

        if (this.playerState.cards.length <= 0) {
            this.board.clear();
        } else {
            // updateCards
            this.board.pushOrUpdate(this.playerState.cards);
        }

        this.calcWidth = this.options.width - 2 * this.options.marginX;

        this.board.position.set(
            this.calcWidth / 2 - this.board.width / 2 + this.options.marginX,
            this.avatar.height + 2 * this.options.marginY
        );
        this.avatar.position.set(this.options.marginX, this.options.marginY);
        this.usernameLabel.position.set(
            this.calcWidth - this.options.avatarRadius,
            this.avatar.height / 2
        );
        this.buyInLabel.position.set(
            this.calcWidth -
                this.options.avatarRadius -
                this.usernameLabel.width -
                this.options.marginX,
            this.avatar.height / 2
        );
        this.betLabel.position.set(
            this.calcWidth -
                this.options.avatarRadius -
                this.usernameLabel.width -
                2 * this.options.marginX -
                this.buyInLabel.width,
            this.avatar.height / 2
        );
        this.dealerButton.position.set(
            this.betLabel.x - this.betLabel.width - this.options.marginX,
            this.options.avatarRadius + this.options.marginX
        );

        if (this.playerState.waiting) {
            this.waiting.lineStyle(4, 0x000000);
            this.waiting.arc(
                this.options.avatarRadius,
                this.options.avatarRadius,
                this.options.avatarRadius,
                0,
                Math.PI
            );
            this.waiting.position.set(
                this.avatar.position.x + this.options.avatarRadius,
                this.avatar.position.y + this.options.avatarRadius
            );
            this.waiting.pivot.set(this.options.avatarRadius);
            this.waiting.visible = true;
        } else {
            this.waiting.visible = false;
            this.waiting.clear();
        }

        this.background.clear();
        this.background.beginFill(this.options.background.value, this.options.background.alpha);
        this.background.drawRoundedRect(
            0,
            0,
            this.options.width,
            this.topRow.height + 3 * this.options.marginY + this.board.height,
            10
        );
        this.background.endFill();

        if (!this.playerState.in) {
            this.background.clear();
            this.background.beginFill(0x000000, 0.25);
            this.background.drawRoundedRect(
                0,
                0,
                this.options.width,
                this.topRow.height + 3 * this.options.marginY + this.board.height,
                10
            );
            this.background.endFill();
            this.activeBackground.visible = true;
        } else {
            this.activeBackground.visible = false;
        }

        this.onResize();
    }

    onResize() {
        this.pivot.set(this.width * 0.5, this.height * 0.5);
    }

    updateFromState() {
        this.playerState = this.state.getPlayerState(this.index);
        this.update({});
        console.log("Player [", this.index, "] has: ", this.playerState.buyIn);
    }
}

export { Player };
