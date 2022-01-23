import { CHASE, SCATTER } from "../../constants";
import { Canvas } from "../../Draw/canvas";
import { toCell } from "../../Functions/math";
import { vector2D } from "../../Functions/vector";
import { Pacman } from "../Pacman";
import { Ghost } from "./ghost";

const InkyImageURL = "/static/games/pacman/sprites/inky.png";

/**
 * @file Die Inky Klasse wird als Unterklasse zu {@link ghost} zur verfügung gestellt
 * @example
 * var vPacman = new pacman(x, y);
 * var inky = new inky(x,y,0,1,vPacman);
 */
export class Inky extends Ghost {
    //blinky: Blinky;
    /**
     * @param {number} xPos Die x-Startposition
     * @param {number} yPos Die y-Startposition
     * @param {number} xVel Die Startgeschwindigkeit in x-Richtung
     * @param {number} yVel Die Startgeschwindigkeit in y-Richtung
     * @param {pacman} pac Eine Referenz zu pacman
     * @extends ghost
     */
    constructor(canvas: Canvas, pac: Pacman, pos?: vector2D, vel?: vector2D) {
        super(canvas, InkyImageURL, pac, "blue", pos, vel);
        this.setVelocity(1, 1);
        this.name = "inky";
        //this.blinky = blinky;
    }

    /**
     * Diese Funktion setzt das Ziel auf drei Felder hinter die aktuelle Position von Pacman
     */
    setZiel(): void {
        if (this.lvl.mode == CHASE) {
            const pacPos = toCell(this.pacman.loc);
            const pacVel = this.pacman.vel.get();
            let mul = -3;
            pacVel.mul(mul);
            pacPos.add(pacVel);
            while (
                this.lvl.grid[pacPos.y][pacPos.x] === null ||
                this.lvl.grid[pacPos.y][pacPos.x] === undefined ||
                this.lvl.grid[pacPos.y][pacPos.x] == 1
            ) {
                pacPos.sub(pacVel);
                pacVel.div(mul);
                mul++;
                pacVel.mul(mul);
                pacPos.add(pacVel);
            }
            this.ziel = pacPos;
        } else if (this.lvl.mode == SCATTER) {
            this.ziel = this.lvl.getStartPos(this.name);
        }
    }
    /*
    setZiel() {
        if (lvl.mode == CHASE) {
            let pacPos = toCell(this.pacman.loc);
            let pacVel = this.pacman.vel.get();
            pacVel.mul(2);
            pacPos.add(pacVel);
            let ziel = vector2D.sub(this.blinky.cell, pacPos);
            ziel.mul(2);
            ziel = vector2D.add(this.blinky.cell, ziel);
            if (ziel.x > spalten || ziel.x < 0 || ziel.y > zeilen || ziel.y < 0) {
                let suchbereich = 15;
                let wege = [];
                for (let _x = 0 - suchbereich; _x < suchbereich; _x++) {
                    for (let _y = 0 - suchbereich; _y < suchbereich; _y++) {
                        let nx = ziel.x + _x;
                        let ny = ziel.y + _y;
                        if (lvl.grid[ny][nx] && lvl.grid[ny][nx] != 1) {
                            wege.push({
                                x: nx,
                                y: ny,
                                dist: euclideanDistance(nx, ny, ziel.x, ziel.y)
                            });
                        }
                    }
                }
                wege.sort(function (a, b) {
                    return b.dist - a.dist;
                });
                ziel = new vector2D(wege[0].x, wege[0].y);
            }
            this.ziel = [ziel.x, ziel.y];
        } else if (lvl.mode == SCATTER) {
            this.ziel = [lvl.inkyStart.x, lvl.inkyStart.y];
        }
    }*/
}
