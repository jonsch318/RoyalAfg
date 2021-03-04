import * as PIXI from "pixi.js-legacy";

class Notification extends PIXI.Container {
    constructor(state, appWidth, appHeight) {
        super();
        this.state = state;
        this.notification = {};
        this.appWidth = appWidth;
        this.appHeight = appHeight;

        this.label = new PIXI.Text("Bye", { fill: 0xffffff });
        this.label.resolution = 2;
        this.bg = new PIXI.Graphics();
        this.textBg = new PIXI.Graphics();

        this.visible = false;
        this.addChild(this.bg, this.textBg, this.label);
    }

    update() {
        if (this.notification?.text) {
            this.label.text = this.notification.text;
        }

        this.bg.clear();
        this.bg.beginFill(0x000000, 0.5);
        this.bg.drawRect(0, 0, this.appWidth, this.appHeight);
        this.bg.endFill();

        const paddingX = 50;
        const paddingY = 30;
        this.textBg.clear();
        this.textBg.beginFill(0x000000, 0.75);
        this.textBg.drawRect(
            this.appWidth / 2 - this.label.width / 2 - paddingX / 2,
            this.appHeight / 2 - this.label.height / 2 - paddingY / 2,
            this.label.width + paddingX,
            this.label.height + paddingY
        );
        this.textBg.endFill();

        this.label.position.set(this.appWidth / 2 - this.label.width / 2, this.appHeight / 2 - this.label.height / 2);
    }

    updateFromState() {
        this.update();
    }

    onNotification() {
        if (this.state.notifications.length > 0) {
            this.notification = this.state.notifications[0];
            this.state.notifications.shift();
            this.visible = true;
            this.update();
            if (!this.notification.static) {
                setTimeout(() => {
                    this.visible = false;
                    this.onNotification();
                }, 2500);
            }
        }
    }

    reset() {
        this.text = "";
        this.visible = false;
    }
}

export { Notification };
