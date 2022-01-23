import { vector2D } from "../Functions/vector";
import { toCell } from "../Functions/math";
import { Canvas } from "../Draw/canvas";
import { OFFSET_X, OFFSET_Y, TILESIZE } from "../constants";
import { Level } from "./level";

/**
 * Die klasse obj ist die Oberklasse von den Geistern und Pacman und stellt Grundfunktionen, sowie Variabeln zur verfügung
 * @file Stellt die Klasse {@link Character} zur verfügung
 */
abstract class Character {
    c: Canvas;
    lvl: Level;
    image: HTMLImageElement;
    private _location: vector2D;
    velocity: vector2D;
    cell: vector2D;
    name: string;
    protected onLevelLoad: () => void;

    /**
     * Gibt den Positionsvektor des Characters zurück
     * @type {vector2D}
     * @return {vector2D} Die Position des Characters
     */
    get loc(): vector2D {
        return this._location;
    }

    /**
     * Gibt den Geschwindigkeitsvektor des Characters zurück
     * @type {vector2D}
     * @return {vector2D} Die Geschwindigkeit des Characters
     */
    get vel(): vector2D {
        return this.velocity.get();
    }

    /**
     * @param {string} img Das Bild, welches genutz werden soll
     * @param {number} xPos Die x-Startposition
     * @param {number} yPos Die y-Startposition
     * @param {number} xVel Die Startgeschwindigkeit in x-Richtung
     * @param {number} yVel Die Startgeschwindigkeit in y-Richtung
     */
    constructor(canvas: Canvas, img: string, pos?: vector2D, vel?: vector2D) {
        this.c = canvas;
        /**
         * @member {Image} [obj~image=null] Das Bild, das dieser Character nutzen soll
         */
        this.image = new Image();
        /**
         * @member {vector2D} obj~location Der Positionsvektor ({@link vector2D}) des Characters
         */
        this._location = pos ? pos : new vector2D(0, 0);
        /**
         * @member {vector2D} obj~velocity Der Geschwindigkeitsvektor ({@link vector2D}) des Characters
         */
        this.velocity = vel ? vel : new vector2D(0, 0);
        if (img === null) this.image = null;
        else this.image.src = img;

        /**
         * @member {vector2D} obj~cell Der Positionsvektor ({@link vector2D}) des Characters konvertiert in die Position im {@link level#grid|grid} (Siehe: {@link toCell})
         */
        this.cell = pos ? toCell(this._location) : null;
        // eslint-disable-next-line @typescript-eslint/no-empty-function
        this.onLevelLoad = function () {};
    }

    loadLevel(lvl: Level): void {
        this.lvl = lvl;
        console.log("start pos:", this.lvl.getStartPos(this.name));
        this.setLocationVec(vector2D.add(vector2D.mul(this.lvl.getStartPos(this.name), TILESIZE), new vector2D(TILESIZE / 2, TILESIZE / 2)));
        this.onLevelLoad();
    }

    /**
     * Die Funktion zeichnet, wenn sie nicht überschrieben ist, den Character auf das Canvas.
     */
    draw(): void {
        this.c.drawImage(this.image, this._location.x - TILESIZE / 2 + OFFSET_X, this._location.y - TILESIZE / 2 + OFFSET_Y, TILESIZE, TILESIZE);
    }

    /**
     * Diese Funktion testet, ob der Character in der Zellmitte ist
     * @return {boolean} True, wenn in Zellmitte
     */
    isInMiddleOfCell(): boolean {
        const tmp = this.cell.get();
        tmp.x = tmp.x * TILESIZE + TILESIZE / 2;
        tmp.y = tmp.y * TILESIZE + TILESIZE / 2;
        return this.loc.cmp(tmp);
    }

    /**
     * Die Funktion zentriert den Character in der Zellmitte
     */
    center(): void {
        if (this.velocity.x != 0) this._location.y = this.cell.y * TILESIZE + TILESIZE / 2;
        else this._location.x = this.cell.x * TILESIZE + TILESIZE / 2;
    }

    /**
     * Testet ob eine Zelle eine Wand ist
     * @param {vector2D} pVector Positionsvektor der aktuellen Zelle
     * @return {boolean}    True, wenn die Zelle 1 ist.
     */
    collide(pVector: vector2D): boolean {
        let res = false;
        if (this.lvl.grid.length <= pVector.y || pVector.y < 0 || this.lvl.grid[0].length <= pVector.x || pVector.x < 0) res = true;
        else if (this.lvl.grid[pVector.y][pVector.x] === 1) res = true;
        else res = false;
        /*console.log(
            "Collides: Grid(",
            this.lvl.grid.length,
            ", ",
            this.lvl.grid[0].length,
            ") pVector: ",
            pVector,
            " res: ",
            res,
            " g(x,y): ",
            this.lvl.grid[pVector.y][pVector.x],
            " cell: ",
            this.cell
        );*/
        return res;
    }

    setLocation(x: number, y: number): void {
        this._location.x = x;
        this._location.y = y;
        //if (this.name == "pacman") console.log("PACMAN CELL: ", toCell(this.loc), " PACMAN LOC: ", this.loc);
        this.cell = toCell(this._location);
    }

    setLocationVec(loc: vector2D): void {
        this.setLocation(loc.x, loc.y);
    }

    setVelocity(x: number, y: number): void {
        this.velocity.x = x;
        this.velocity.y = y;
    }
}

export { Character };
