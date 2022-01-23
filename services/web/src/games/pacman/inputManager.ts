import { GameState } from "./gameState";
import { LevelManager } from "./levelManager";

let keys: string[] = [];
let key = "";
let specialChars = 0;
let handlers: { key: string; handler: (e: KeyboardEvent, specialChars: number) => void }[];
export class InputManager {
    levelManager: LevelManager;
    gameTracker: GameState;

    // 1. bit = shift, 2. bit = ctrl, 3. bit = alt, 4. bit = alt_gr

    constructor(levelManager: LevelManager, gameTracker: GameState) {
        this.levelManager = levelManager;
        this.gameTracker = gameTracker;
        keys = [];
        handlers = [];
    }

    registerKeyHandler(key: string, handler: (e: KeyboardEvent, specialChars: number) => void): void {
        handlers.push({ key, handler });
    }

    register(): void {
        document.addEventListener("keydown", this.onKeyDown);
        document.addEventListener("keyup", this.onKeyUp);
    }

    onKeyUp(event: KeyboardEvent): void {
        // search handlers with keys
        switch (event.key) {
            case "shift":
                specialChars &= ~(1 << 0);
                break;
            case "ctrl":
                specialChars &= ~(1 << 1);
                break;
            case "alt":
                specialChars &= ~(1 << 2);
                break;
            case "alt_gr":
                specialChars &= ~(1 << 3);
                break;
            default:
                keys = keys.filter((val) => val !== event.key);
                key = "";
                break;
        }
    }

    onKeyDown(event: KeyboardEvent): void {
        switch (event.key) {
            case "shift":
                specialChars |= 1 << 0;
                break;
            case "ctrl":
                specialChars |= 1 << 1;
                break;
            case "alt":
                specialChars |= 1 << 2;
                break;
            case "alt_gr":
                specialChars |= 1 << 3;
                break;
            default:
                keys.push(event.key);
                key = event.key;
                //console.log("Received key Input: ", event);
                handlers
                    .filter((val) => val.key === event.key)
                    .forEach(({ handler }) => {
                        handler(event, specialChars);
                    });
                break;
        }
    }

    clear(): void {
        removeEventListener("keydown", onkeydown);
    }

    setKey(direction: string | "down" | "up"): void {
        //el ignoriere ich, da ich die Funktion eh nur für ein Element nutze
        if (direction == "down" || direction == "up") key = direction == "down" ? "up" : "down";
        else key = direction;
    }

    /**
     * Die Methode erkennt Swipe-Bewegungen. Quelle: {@link https://stackoverflow.com/a/58719294}
     * @param {string} id Die Id des Elements das für touch-events genutzt werden soll.
     * @param {Function} func Die callback-Funktion
     * @param {number} deltaMin Die minimale distanz bis die Funktion auslöst
     */
    //detectSwipe('swipeme', (el, dir) => alert(`you swiped on element with id ${el.id} to ${dir} direction`));

    // source code

    // Tune deltaMin according to your needs. Near 0 it will almost
    // always trigger, with a big value it can never trigger.
    detectSwipe(id: string, func: (dir: string) => void, deltaMin = 90): void {
        const swipe_det = {
            sX: 0,
            sY: 0,
            eX: 0,
            eY: 0
        };
        // Directions enumeration
        const directions = Object.freeze({
            UP: "up",
            DOWN: "down",
            RIGHT: "right",
            LEFT: "left"
        });

        let direction = null;
        const el = document.getElementById(id);
        el.addEventListener(
            "touchstart",
            function (e) {
                const t = e.touches[0];
                swipe_det.sX = t.screenX;
                swipe_det.sY = t.screenY;
            },
            false
        );
        el.addEventListener(
            "touchmove",
            function (e) {
                // Prevent default will stop user from scrolling, use with care
                // e.preventDefault();
                const t = e.touches[0];
                swipe_det.eX = t.screenX;
                swipe_det.eY = t.screenY;
            },
            false
        );
        el.addEventListener(
            "touchend",
            function () {
                const deltaX = swipe_det.eX - swipe_det.sX;
                const deltaY = swipe_det.eY - swipe_det.sY;
                // Min swipe distance, you could use absolute value rather
                // than square. It just felt better for personnal use
                if (deltaX ** 2 + deltaY ** 2 < deltaMin ** 2) return;
                // horizontal
                if (deltaY === 0 || Math.abs(deltaX / deltaY) > 1) direction = deltaX > 0 ? directions.RIGHT : directions.LEFT;
                // vertical
                else direction = deltaY > 0 ? directions.UP : directions.DOWN;

                if (direction && typeof func === "function") func(direction);

                direction = null;
            },
            false
        );
    }
}
