import { OFFSET_X, OFFSET_Y, PAUSE, TILESIZE } from "../constants";
import { Canvas } from "../Draw/canvas";
import { toCell } from "../Functions/math";
import { vector2D } from "../Functions/vector";
import { Character } from "./character";

/**
 * @file Stellt die Pacman Klasse zur verfügung
 * @example
 * var vPacman = new pacman(x, y);
 */
export class Pacman extends Character {
    dir: number;
    pctOpen: number;
    moving: boolean;
    r: number;
    baseVel: number;
    lives: number;

    /**
     *
     * @param {number} x Die x-Position
     * @param {number} y Die y-Position
     * @augments obj
     */
    constructor(canvas: Canvas, pos?: vector2D) {
        super(canvas, null, pos, new vector2D(0, -1));
        /**
         * @member {number} pacman~dir Die Geschwindigkeit, mit der sich der "Mund" von Pacman öffnet und schließt
         */
        this.dir = -10;
        /**
         * @member {number} pacman~pctOpen Mundöffnung in Prozent
         */
        this.pctOpen = 100;
        /**
         * @member {boolean} pacman~moving Gibt an, ob sich der Pacman bewegt und sich der Mund grade bewegen soll
         */
        this.moving = true;
        /**
         * @member {number} pacman~r Der Radius von Pacman
         */
        this.r = TILESIZE / 2;

        /**
         * @member {string} pacman~name Der Name von Pacman
         * @private
         */
        this.name = "pacman";
        this.baseVel = 100;
        this.lives = 3;
    }

    changeDir(key: "up" | "down" | "left" | "right") {
        const geschwindigkeit = this.getPacVel();
        switch (key) {
            case "up":
                this.setVelocity(0, -geschwindigkeit);
                break;
            case "down":
                this.setVelocity(0, geschwindigkeit);
                break;
            case "left":
                this.setVelocity(-geschwindigkeit, 0);
                break;
            case "right":
                this.setVelocity(geschwindigkeit, 0);
                break;
        }
    }

    /**
     * Die Methode bewegt den Pacman
     * @override
     */
    move(): void {
        this.moving = true;
        const next = toCell(this.loc.add(this.vel));
        //check if current vel would hit a wall.
        if (next != this.cell && this.collide(next)) {
            this.setVelocity(0, 0);
            this.toCellMid();
        } else {
            //fine to move on.
            this.setLocationVec(this.loc.add(this.vel));
        }
    }

    getToNearestFree() {
        if (this.lvl.grid.length <= 0) return;
        if (!this.collide(this.cell)) {
            this.toCellMid();
            return;
        }
        const tmp = this.cell;
        tmp.y -= 1;
        if (!this.collide(tmp)) {
            this.cell = tmp;
            this.toCellMid();
            return;
        }

        tmp.y += 2;
        if (!this.collide(tmp)) {
            this.cell = tmp;
            this.toCellMid();
            return;
        }

        tmp.y -= 1;
        tmp.x -= 1;
        if (!this.collide(tmp)) {
            this.cell = tmp;
            this.toCellMid();
            return;
        }

        tmp.x += 2;
        if (!this.collide(tmp)) {
            this.cell = tmp;
            this.toCellMid();
            return;
        }
    }

    toCellMid() {
        const tmp = this.cell;
        tmp.x = tmp.x * TILESIZE + TILESIZE / 2;
        tmp.y = tmp.y * TILESIZE + TILESIZE / 2;
        this.setLocationVec(tmp);
    }

    /**
     * Berechnet die Position der Zelle, die Pacman als nächstes betritt
     */
    nextCell(): vector2D {
        //this.nCell = this.cell.add(this.vel);
        return toCell(this.loc.add(this.vel));
    }

    /**
     * Zeichnet Pacman auf das Canvas
     * @override
     */
    draw(): void {
        this.c.ctx.save();
        this.drawPacman(this.moving && !PAUSE ? (this.pctOpen += this.dir) : this.pctOpen);
        this.c.ctx.restore();
    }

    /**
     * Zeichnet den Mund von pacman
     * @param {number} pctOpen Die Öffnung des Munds in Prozent
     */
    drawPacman(pctOpen: number): void {
        const geschwindigkeit = this.getPacVel();
        this.c.ctx.translate(this.loc.x + OFFSET_X, this.loc.y + OFFSET_Y);
        if (this.velocity.x == 0) this.c.ctx.rotate(((this.velocity.y == geschwindigkeit ? 90 : 270) * Math.PI) / 180);
        else this.c.ctx.rotate(((this.velocity.x == geschwindigkeit ? 0 : 180) * Math.PI) / 180);
        const fltOpen = pctOpen / 100;

        this.c.ctx.beginPath();
        this.c.ctx.arc(0, 0, this.r, fltOpen * 0.2 * Math.PI, (2 - fltOpen * 0.2) * Math.PI);

        this.c.ctx.lineTo(0, 0);
        this.c.ctx.closePath();

        this.c.ctx.fillStyle = "#FFFF00";
        this.c.ctx.fill();
        this.pctOpen = Math.max(0, Math.min(100, pctOpen));

        if (pctOpen % 100 == 0) {
            this.dir *= -1;
        }
    }

    private getPacVel(): number {
        return Math.ceil(this.c.fpsStats.frameTime / 5);
        //this.c.fpsStats.frameTime *
    }
}
