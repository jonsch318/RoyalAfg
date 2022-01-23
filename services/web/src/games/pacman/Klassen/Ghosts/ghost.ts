import { toCell } from "../../Functions/math";
import { vector2D } from "../../Functions/vector";
import { Character } from "../character";
import { Canvas } from "../../Draw/canvas";
import { Level } from "../level";
import { Pacman } from "../Pacman";
import { COLUMN_COUNT, DEBUG, OFFSET_X, OFFSET_Y, ORIGPATH, ROW_COUNT, TILESIZE } from "../../constants";
import EasyStar from "../../../../lib/PathFinder/easystar-0.4.2";

/**
 * @file Stellt die Ghost Oberklasse zur verfügung
 **/

export abstract class Ghost extends Character {
    ziel: vector2D;
    easystar: any;
    pacman: Pacman;
    color: string;
    start: vector2D;
    path: any;
    v: number;

    abstract setZiel(): void;

    /**
     * Der Constructor setzt den Positions- und Geschwindigkeitsvektor und {@link pacman |Pacman}.
     * @param {string} img Der Pfad zur Textur des Geistes
     * @param {number} x Die x-Startposition
     * @param {number} y Die y-Startposition
     * @param {number} xvel Die x-Startgeschwindigkeit
     * @param {number} yvel Die y-Startgeschwindigkeit
     * @param {pacman} pac Eine Referenz zu Pacman
     * @param {string} color Die Farbe des Geistes
     * @augments Character
     */
    constructor(canvas: Canvas, img: string, pac: Pacman, color: string, pos?: vector2D, vel?: vector2D) {
        super(canvas, img, pos, vel);
        this.ziel = new vector2D(0, 0);
        this.pacman = pac;
        this.path = null;
        this.color = color;
        this.start = this.loc ? this.loc.get() : new vector2D(0, 0);
        this.v = 0.5;
    }

    loadLevel(lvl: Level): void {
        //Easystar ist eine API mit der ich die Wege für die geister berechnen lasse, ich habe diese allerdings Modifiziert um besser die einzelnen geister anzusprechen
        super.loadLevel(lvl);
        this.ziel = this.loc.get();
        this.easystar = new EasyStar.js();
        this.easystar.setGrid(lvl.grid);
        this.easystar.setAcceptableTiles([0]);
        this.easystar.enableSync();
    }

    draw(): void {
        super.draw();
        if (ORIGPATH) this.drawPath();
    }

    drawPath(): void {
        if (this.path != null || this.path?.length > 0) {
            console.log("Path: ", this.path);
            const path = this.path;
            this.c.ctx.beginPath();
            this.c.ctx.moveTo(path[0].x * TILESIZE + TILESIZE / 2 + OFFSET_X, path[0].y * TILESIZE + TILESIZE / 2 + OFFSET_Y);
            for (let i = 1; i < this.path.length; i++) {
                this.c.ctx.lineTo(path[i].x * TILESIZE + TILESIZE / 2 + OFFSET_X, path[i].y * TILESIZE + TILESIZE / 2 + OFFSET_Y);
            }
            this.c.ctx.strokeStyle = this.color;
            this.c.ctx.lineWidth = 5;
            this.c.ctx.stroke();
            this.c.ctx.lineWidth = 2;
            this.c.drawCircle(
                path[path.length - 1].x * TILESIZE + TILESIZE / 2 + OFFSET_X,
                path[path.length - 1].y * TILESIZE + TILESIZE / 2 + OFFSET_Y,
                TILESIZE / 1.5,
                this.color
            );
        }
    }

    /**
     * Bewegt den Geist, indem es mit der aktuellen Position von Pacman das Ziel setzt und in die Richtung geht.
     * Desweiteren wird die Zelle hinter dem Geist für die Dauer der berechnung des Pfades zu einer Wand geändert, um zu verhindern, dass sich der Geist umdreht
     * @override
     */
    move(): void {
        this.cell = toCell(this.loc); //Berechnet die Position des Geistes im Grid
        if (this.cell.x > COLUMN_COUNT - 1 || this.cell.x < 0 || this.cell.y > ROW_COUNT - 1 || this.cell.y < 0) {
            this.setLocationVec(this.start.get());
            this.velocity = new vector2D(0, 0);
        }
        const xOffset = this.cell.x - this.velocity.x * 2;
        const yOffset = this.cell.y - this.velocity.y * 2;
        let cFlag = true;
        try {
            if (this.lvl.grid[yOffset][xOffset] == 1 || (xOffset == this.cell.x && yOffset == this.cell.y)) cFlag = false;
            else this.lvl.grid[yOffset][xOffset] = 2;
        } catch {
            cFlag = false;
        }
        if (this.isInMiddleOfCell()) {
            //if (this.path !== null)
            //this.easystar.cancelPath(this.path);
            this.setZiel();

            if (this.ziel.x == this.cell.x && this.ziel[1] == this.cell.y) {
                this.ziel.x += this.velocity.x * 2;
                this.ziel.y += this.velocity.y * 2;
            }

            if (this.lvl.grid[this.ziel.y][this.ziel.y] == 1) {
                this.ziel.x += 1;
            }

            //Erste Möglichkeit zum Berechnen der neuen Richtung.
            if (ORIGPATH) {
                const richtungen = [];
                //Links
                richtungen.push(new vector2D(-1, 0));
                //Rechts
                richtungen.push(new vector2D(1, 0));
                //Oben
                richtungen.push(new vector2D(0, -1));
                //Unten
                richtungen.push(new vector2D(0, 1));
            } else {
                //throw new Error("UnsupportedBranch");
                //Zweite Möglichkeit
                this.path = this.easystar.findPath(this.cell.x, this.cell.y, this.ziel.x, this.ziel.y, this.parsePath, this);
                //this.easystar.calculate();
            }
        }
        this.setLocationVec(this.loc.add(this.vel));
        if (cFlag) this.lvl.grid[yOffset][xOffset] = 0;
    }

    /**
     * Diese Funktion, ist die callbackfunction für die Pfadfindung
     * @param {object[]} path Der Pfad, den der Pfadfinder gefunden hat
     * @param {Ghost} ghost Eine Referenz auf den Geist, der den Pfad angefordert hat.
     */
    parsePath(path: vector2D[], ghost: Ghost): void {
        if (path !== null && path[0]) {
            /*if (path[0] === undefined)
                stop();*/
            const pTmp = new vector2D(path[1].x - path[0].x, path[1].y - path[0].y);
            pTmp.mul(ghost.v);
            ghost.velocity = pTmp.get();
            ghost.path = path;
        }
    }
}
