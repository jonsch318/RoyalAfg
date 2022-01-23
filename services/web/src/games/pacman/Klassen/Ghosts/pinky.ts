import { CHASE, SCATTER } from "../../constants";
import { Canvas } from "../../Draw/canvas";
import { toCell } from "../../Functions/math";
import { vector2D } from "../../Functions/vector";
import { Pacman } from "../Pacman";
import { Ghost } from "./ghost";

const PinkyImageURL = "/static/games/pacman/sprites/pinky.png";

/**
 * @file Die Pinky Klasse wird als Unterklasse zu {@link ghost} zur verfügung gestellt
 * @example
 * var vPacman = new pacman(x, y);
 * var blinky = new blinky(x,y,0,1,vPacman);
 */
export class Pinky extends Ghost {
    /**
     * @param {number} xPos Die x-Startposition
     * @param {number} yPos Die y-Startposition
     * @param {number} xVel Die Startgeschwindigkeit in x-Richtung
     * @param {number} yVel Die Startgeschwindigkeit in y-Richtung
     * @param {pacman} pac Eine Referenz zu pacman
     * @extends ghost
     */
    constructor(canvas: Canvas, pac: Pacman, pos?: vector2D, vel?: vector2D) {
        super(canvas, PinkyImageURL, pac, "pink", pos, vel);
        this.setVelocity(0, 1);
        this.name = "pinky";
    }

    /**
     * Diese Funktion setzt das Ziel zwei Zellen vor Pacman
     */
    setZiel(): void {
        if (this.lvl.mode == CHASE) {
            const pacPos = toCell(this.pacman.loc);
            const pacVel = this.pacman.vel.get();
            let mul = 3;
            pacVel.mul(mul);
            pacPos.add(pacVel);
            while (
                this.lvl.grid[pacPos.y][pacPos.x] === null ||
                this.lvl.grid[pacPos.y][pacPos.x] === undefined ||
                this.lvl.grid[pacPos.y][pacPos.x] == 1
            ) {
                pacPos.sub(pacVel);
                pacVel.div(mul);
                mul--;
                pacVel.mul(mul);
                pacPos.add(pacVel);
            }
            this.ziel = pacPos;
        } else if (this.lvl.mode == SCATTER) {
            this.ziel = this.lvl.getStartPos(this.name);
        }
    }
}
