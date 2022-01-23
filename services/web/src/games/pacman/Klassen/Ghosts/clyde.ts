import { CHASE, SCATTER } from "../../constants";
import { Canvas } from "../../Draw/canvas";
import { toCell, manhattenDistance } from "../../Functions/math";
import { vector2D } from "../../Functions/vector";
import { Pacman } from "../Pacman";
import { Ghost } from "./ghost";

const ClydeImageURL = "/static/games/pacman/sprites/clyde.png";

/**
 * @file Die Clyde Klasse wird als Unterklasse zu {@link Ghost} zur verfügung gestellt
 * @example
 * var vPacman = new Pacman(x, y);
 * var clyde = new Clyde(x,y,0,1,vPacman);
 */
export class Clyde extends Ghost {
    /**
     * @param {number} xPos Die x-Startposition
     * @param {number} yPos Die y-Startposition
     * @param {number} xVel Die Startgeschwindigkeit in x-Richtung
     * @param {number} yVel Die Startgeschwindigkeit in y-Richtung
     * @param {pacman} pac Eine Referenz zu pacman
     * @extends ghost
     */
    constructor(canvas: Canvas, pac: Pacman, pos?: vector2D, vel?: vector2D) {
        super(canvas, ClydeImageURL, pac, "orange", pos, vel);
        this.setVelocity(0, 1);
        this.name = "clyde";
    }

    /**
     * Diese Funktion setzt das Ziel auf die aktuelle Position von Pacman, außer die Entfernung ({@link manhattenDistance}) ist kleiner als 8, worauf er in seine Ecke vom {@link level#grid|grid} geht.
     */
    setZiel(): void {
        if (this.lvl.mode == CHASE) {
            const pacPos = toCell(this.pacman.loc);
            const pX = pacPos.x;
            const pY = pacPos.y;
            if (manhattenDistance(this.cell.x, this.cell.y, pX, pY) >= 8) this.ziel = pacPos;
            else this.ziel = this.lvl.getStartPos(this.name);
        } else if (this.lvl.mode == SCATTER) {
            this.ziel = this.lvl.getStartPos(this.name);
        }
    }
}
