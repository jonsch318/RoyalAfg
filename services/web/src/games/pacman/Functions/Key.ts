/**
 * @var key Die Variabel, in welcher die letzte Richtungstaste gespeichert ist.
 * @type {string}
 */
export let key;

export const keyCodes = [];
const KC = [38, 38, 40, 40, 37, 39, 37, 39, 66, 65, "\0"];
let i_KC = 0;

const NL = [78, 88, 84, "\0"];
let i_NL = 0;

const B = [66, 85, 71, "\0"];
let i_B = 0;

document.onkeydown = function (event) {
    keyCodes[event.key] = true;
    if (KC[i_KC] === event.keyCode) {
        i_KC++;
        if (KC[i_KC] == "\0") {
            i_KC = 0;
            points += 100;
        }
    } else {
        i_KC = 0;
    }

    if (NL[i_NL] === event.keyCode) {
        i_NL++;
        if (NL[i_NL] == "\0") {
            i_NL = 0;
            nextLevel();
        }
    } else {
        i_NL = 0;
    }

    if (B[i_B] === event.keyCode) {
        i_B++;
        if (B[i_B] == "\0") {
            i_B = 0;
            bbug = !bbug;
        }
    } else {
        i_B = 0;
    }
    return false;
};

document.onkeyup = function (evt) {
    keyCodes[evt.key] = false;
    return false;
};

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
function detectSwipe(id: string, func: (dir: string) => void, deltaMin = 90): void {
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
        function (e) {
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

function setKey(direction: string | "down" | "up"): void {
    //el ignoriere ich, da ich die Funktion eh nur für ein Element nutze
    if (direction == "down" || direction == "up") key = direction == "down" ? "up" : "down";
    else key = direction;
}

export { setKey, detectSwipe };
