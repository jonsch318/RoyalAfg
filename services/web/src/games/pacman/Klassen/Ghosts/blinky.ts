import { CHASE, SCATTER } from "../../constants";
import { Canvas } from "../../Draw/canvas";
import { toCell } from "../../Functions/math";
import { vector2D } from "../../Functions/vector";
import { Pacman } from "../Pacman";
import { Ghost } from "./ghost";

const BlinkyImageURL = "/static/games/pacman/sprites/blinky2.png";

/**
 * @file Die Blinky Klasse wird als Unterklasse zu {@link Ghost} zur verfügung gestellt
 * @example
 * var vPacman = new Pacman(x, y);
 * var Blinky = new Blinky(x,y,0,1,vPacman);
 */
export class Blinky extends Ghost {
    /**
     * @param {number} xPos Die x-Startposition
     * @param {number} yPos Die y-Startposition
     * @param {number} xVel Die Startgeschwindigkeit in x-Richtung
     * @param {number} yVel Die Startgeschwindigkeit in y-Richtung
     * @param {pacman} pac Eine Referenz zu pacman
     * @extends ghost
     */
    constructor(canvas: Canvas, pac: Pacman, pos?: vector2D, vel?: vector2D) {
        super(canvas, BlinkyImageURL, pac, "red", pos, vel);
        this.setVelocity(0, 1);
        this.name = "blinky";
    }

    /**
     * Diese Funktion setzt das Ziel auf die aktuelle Position von Pacman
     */
    setZiel(): void {
        if (this.lvl.mode == CHASE) {
            const pacPos = toCell(this.pacman.loc);
            this.ziel = pacPos;
        } else if (this.lvl.mode == SCATTER) {
            this.ziel = this.lvl.getStartPos(this.name);
        }
    }
}
