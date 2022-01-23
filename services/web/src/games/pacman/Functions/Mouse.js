/**
 * @file Stellt Funktionen zum Zugriff auf die Maus zur verfügung
 * @module mouse.js
 */

/**
 * @var mouseX Die x-Position der Maus
 * @type {number}
 */
var mouseX = 0;

/**
 * @var mouseY Die y-Position der Maus
 * @type {number}
 */
var mouseY = 0;

/**
 * @var mouseDown Die Variabel speichert, ob eine der Maustasten gedrückt ist
 * @type {boolean}
 * @example
 * if(mouseDown)
 *      console.log("Eine Maustaste wurde gedrückt");
 * else
 *      console.log("Es wurde keine Taste gedrückt");
 * //Kürzer:
 * console.log(mouseDown ? "Eine Maustaste wurde gedrückt" : "Es wurde keine Taste gedrückt")
 */
var mouseDown = 0;
document.body.onmousedown = function () {
    mouseDown = 1;
};
document.body.onmouseup = function () {
    mouseDown = 0;
};
//Die Funktion wird ausgeführt wenn die Maus bewegt wird.
function updateMouse(e) {
    mouseX = e.pageX;
    mouseY = e.pageY;
}

export { mouseDown, mouseX, mouseY, updateMouse };
