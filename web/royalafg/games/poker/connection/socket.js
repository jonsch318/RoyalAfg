import { JOIN } from "../events/constants";
import { SendCreateEvent, SendEvent } from "../events/event";

class Game {
    constructor(state, entry, onClose) {
        this.state = state;
        this.credentials = entry;
        this.onClose = onClose;
        this.started = new Promise((resolve, reject) => {
            this.state.setOnGameStart(() => {
                console.log("Resolving game start");
                resolve();
            });
        });
    }

    start() {
        this.ws = new WebSocket("ws://" + window.location.hostname + ":5000/join");
        this.ws.onclose = (e) => {
            console.log("close", e);
            this.onClose();
        };

        this.ws.onmessage = (e) => {
            if (e.data) {
                console.log("Event: ", e);
                this.state.decodeChange(JSON.parse(e.data));
            }
        };

        this.ws.onopen = (e) => {
            if (e.type === "error") {
                this.onClose();
                return;
            }

            // Join event
            this.ws.send(SendCreateEvent(JOIN, this.credentials));
        };

        this.ws.onerror = (e) => {
            this.onClose();
        };
    }

    send(event) {
        this.ws.send(SendEvent(event));
    }
}

export { Game };
