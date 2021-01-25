import { JOIN } from "../events/constants";
import { SendCreateEvent, SendEvent } from "../events/event";

class Game {
    constructor(state, ticket, onClose) {
        this.state = state;
        this.ticket = ticket;
        this.onClose = onClose;
        this.started = new Promise((resolve) => {
            this.state.setOnGameStart(() => {
                console.log("Resolving game start");
                resolve();
            });
        });
    }

    start() {
        console.log("Connect to webserver: ", `ws://${this.ticket.address}/api/poker/join`);
        this.ws = new WebSocket(`ws://${this.ticket.address}/api/poker/join`);
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
            this.ws.send(SendCreateEvent(JOIN, { token: this.ticket.token }));
        };

        this.ws.onerror = (e) => {
            console.log(e);
            this.onClose();
        };
    }

    send(event) {
        this.ws.send(SendEvent(event));
    }
}

export { Game };
