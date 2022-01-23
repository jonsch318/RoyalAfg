import { CANVAS_WIDTH, COLUMN_COUNT } from "../constants";
import { xyToI } from "../Functions/Array";
import { Level } from "../Klassen/level";

/**
 * canvasClass ist eine Klasse, die das zugreifen auf das Canvas Objekt erleichtert.
 * @file Stellt Methoden zum Zugriff auf das HTML5 Canvas zur verfügung
 * @example
 * let c = new canvasClass("canvas", w, h, "green");
 */
export class Canvas {
    ctx: CanvasRenderingContext2D;
    canvas: HTMLCanvasElement;
    lvl: Level;
    fpsStats: {
        FPS: number;
        minFPS: number;
        maxFPS: number;
        frameTime: number;
    };
    showFPS = true;

    private time = 0;
    private times: number[] = [];

    /**
     * Der Construktor der Funktion
     * @param {string} canvas Der Name für für die Id des Canvas, welches der Constructor erzeugt.
     * @param {number} width Die Breite des Canvas.
     * @param {number} height Die Höhe des Canvas.
     * @param {string} [color= ""] Die Hintergrundfarbe des Canvas.
     * @example
     * let c = new canvasClass("canvas", w, h, "green");
     */
    constructor(lvl: Level, canvas: string, width: number, height: number, color = "", el: HTMLDivElement) {
        this.lvl = lvl;
        const node = document.createElement("canvas");
        node.setAttribute("id", canvas);
        el.appendChild(node);
        /**
         * @member {canvas} canvasClass~canvas
         */
        this.canvas = document.getElementById(canvas) as HTMLCanvasElement;
        this.canvas.setAttribute("width", width.toString());
        this.canvas.setAttribute("height", height.toString());
        //this.canvas.setAttribute("onmousemove", "updateMouse(event)");
        //this.canvas.setAttribute("onclick", "console.log(toCell(new vector2D(mouseX+constXOffset, mouseY+constYOffset)))");
        /**
         * @member {getContext} canvasClass~ctx
         */
        this.ctx = this.canvas.getContext("2d");
        this.canvas.style.backgroundColor = color;
        this.fpsStats = {
            FPS: 0,
            maxFPS: 0,
            minFPS: 0,
            frameTime: 0
        };
        this.time = performance.now();
    }

    /**
     * Gibt die hintergrundfarbe zurück
     * @type {string}
     * @returns {string}    Die Farbe des Hintergrunds
     * @example
     * console.log(c.background);
     */
    get background(): string {
        return this.canvas.style.backgroundColor;
    }

    /**
     * Setzt die Hintergrundfarbe
     * @param {string} c Die Farbe die als neuer Hintergrund gesetzt werden soll
     * @example
     * c.background = "Blue";
     */
    set background(c: string) {
        this.canvas.style.backgroundColor = c;
    }

    /**
     * Gibt die Breite des Canvas zurück
     * @type {number}
     * @returns {number} Die Breite des Canvas
     * @example
     * console.log(c.width);
     */
    get width(): number {
        return this.canvas.width;
    }
    /**
     * Setzt eine neue Breite
     * @param {number} w Die neue breite des Canvas
     * @example
     * c.width = 800;
     */
    set width(w: number) {
        this.canvas.width = w;
    }

    /**
     * Gibt die höhe des canvas zurück
     * @type {number}
     * @returns {number} Die Höhe des Canvas
     * @example
     * console.log(c.height);
     */
    get height(): number {
        return this.canvas.height;
    }
    /**
     * Setzt eine neue Höhe für das Canvas
     * @param h {number} Die neue höhe des Canvas
     * @example
     * c.height = 600;
     */
    set height(h: number) {
        this.canvas.height = h;
    }

    setLevel(lvl: Level): void {
        this.lvl = lvl;
    }

    /**
     * Die Funktion zeichnet einen Ausschnitt eines Bildes auf das Canvas
     * @param {Image} img Das Bild aus dem der Ausschnitt kommen soll
     * @param {number} ix Die x-Position des Bildabschnitts
     * @param {number} iy Die y-Position des Bildabschnitts
     * @param {number} iw Die Breite des bildabschnitts
     * @param {number} ih Die Höhe des Bildabschnitts
     * @param {number} x Die x-Position an der es gezeichnet werden soll
     * @param {number} y Die y-position an der es gezeichnet werden soll
     * @param {number} w Die Breite auf die es gestreckt werden soll
     * @param {number} h Die Höhe auf die es gestreckt werden soll.
     */
    drawSprite(img: CanvasImageSource, ix: number, iy: number, iw: number, ih: number, x: number, y: number, w: number, h: number): void {
        this.ctx.drawImage(img, ix, iy, ih, iw, x, y, w, h);
    }

    /**
     * Zeichnet ein Bild aufs canvas
     * @param {Image} image Das zu zeichnende Bild
     * @param {number} x Die x-Position des Bilds
     * @param {number} y Die y-Position des Bilds
     * @param {number} w Die Breite des Bilds
     * @param {number} h Die Höhe des Bilds
     */
    drawImage(image: CanvasImageSource, x: number, y: number, w: number, h: number): void {
        this.ctx.drawImage(image, x, y, w, h);
    }

    /**
     * Zeichnet die Zellen eines 2D Arrays auf das canvas
     * @param {number[][]} arr Das 2D Array
     * @param {string[]} colorArray Die Zahlen in den Zellen des 2D Arrays geben den Index für die Zellen mit der entsprechenden Farbe an
     * @param {number} w Die Breite jeder Zelle
     * @param {number} h Die Höhe jeder Zelle
     */
    fillArray(
        arr: number[][],
        colorArray: string[],
        w = Math.floor(this.canvas.width / arr[0].length),
        h = Math.floor(this.canvas.height / arr.length),
        xOffset = 0,
        yOffset = 0
    ): void {
        const abstand = 0;
        for (let y = 0; y < arr.length; y++) {
            for (let x = 0; x < arr[0].length; x++) {
                this.fillRect(
                    x * w + abstand + xOffset,
                    y * h + abstand + yOffset,
                    w - abstand,
                    h - abstand,
                    colorArray[arr[y][x]],
                    abstand == 0 ? false : true
                );
                if (this.lvl.coins[xyToI(x, y, COLUMN_COUNT)]) {
                    this.fillCircle(
                        Math.round(x * w + abstand + xOffset + w / 2),
                        Math.round(y * h + abstand + yOffset + h / 2),
                        Math.round(h / 5),
                        "yellow"
                    );
                }
                if (this.lvl.coins[xyToI(x, y, COLUMN_COUNT)] && this.lvl.grid[y][x] == 1) {
                    this.lvl.coinCount -= 1;
                    this.lvl.coins[xyToI(x, y, COLUMN_COUNT)] = false;
                }
            }
        }
    }

    /**
     * Zeichnet die die Bilder in den Zellen eines 2D Arrays auf das canvas
     * @param {Image[][]} arr Das 2D Array, dass die Bilder enthält
     * @param {number} w Die Breite jeder Zelle
     * @param {number} h Die Höhe jeder Zelle
     */
    imageArray(
        arr: CanvasImageSource[][],
        w = Math.floor(this.canvas.width / arr[0].length),
        h = Math.floor(this.canvas.height / arr.length),
        xOffset = 0,
        yOffset = 0
    ): void {
        for (let y = 0; y < arr.length; y++) {
            for (let x = 0; x < arr[0].length; x++) {
                this.drawImage(arr[y][x], x * w + xOffset, y * h + yOffset, w, h);
            }
        }
    }

    /**
     * Schreibt einen Text aufs Canvas
     * @param {number} x Die x-Position des Texts
     * @param {number} y Die y-Position des Texts
     * @param {string} text Der zu zeichnende Text
     * @param {string} [font = "30px Arial"] Die zu benutzende Schriftart
     * @param {string} [color = "black"] Die zu benutzende Farbe
     */
    fillText(x: number, y: number, text: string, font = "30px Arial", color = "black"): void {
        this.ctx.font = font;
        this.ctx.fillStyle = color;
        this.ctx.fillText(text, x, y);
    }

    /**
     * Zeichnet ein den Umriss eines Rechtecks auf das Canvas
     * @param {number} x Die x-Position
     * @param {number} y Die y-Position
     * @param {number} w Die Breite
     * @param {number} h Die Höhe
     * @param {string} [color = "black"] Die Farbe der Umrandung
     */
    drawRect(x: number, y: number, w: number, h: number, color = "black"): void {
        this.ctx.strokeStyle = color;
        this.ctx.strokeRect(x, y, w, h);
    }

    /**
     * Zeichnet ein Rechteck auf das Canvas
     * @param {number} x Die x-Position
     * @param {number} y Die y-Position
     * @param {number} w Die Breite
     * @param {number} h Die Höhe
     * @param {string} color Die Farbe der Umrandung
     * @param {boolean} [border = false] Wenn wahr, wird eine Umrandung gezeichnet
     * @param {string} [bColor = "black"] Gibt die Farbe für die umrandung an
     */
    fillRect(x: number, y: number, w: number, h: number, color: string, border = false, bColor = "black"): void {
        this.ctx.fillStyle = color;
        this.ctx.fillRect(x, y, w, h);
        this.ctx.strokeStyle = bColor;
        if (border) this.ctx.strokeRect(x, y, w, h);
    }

    /**
     * Löscht alles was auf das Canvas gezeichnet wurde
     */
    cls(): void {
        this.ctx.clearRect(0, 0, this.canvas.clientWidth, this.canvas.height);
        this.time = performance.now();
    }

    /**
     * Zeichnet eine Linie durch eine Reihe von Punkten
     * @param {number[][]} path Gibt die Punkte bei einem n Punkte langen Pfad in der Form [[x1,y1],[x2,y2],[xn,yn]] an
     * @param {string} [color = "black"] Gibt die farbe für den pfad an
     * @throws {InvalidPathLength} Wird geworfen, wenn der Pfad kürzer als zwei Punkte ist
     */
    drawPath(path: number[][], color = "black"): void {
        if (path.length <= 1) throw "InvalidPathLength";
        path.push([0, 0]);
        const tmp = path.pop();
        this.ctx.beginPath();
        this.ctx.moveTo(tmp[0], tmp[1]);
        for (let i = 1; i < path.length; i++) {
            this.ctx.lineTo(path[i][0], path[i][1]);
        }
        this.ctx.strokeStyle = color;
        this.ctx.stroke();
    }

    /**
     * Zeichnet einen Pfad durch eine Reihe von Punkten und schließt diesen am Ende
     * @param {number[][]} path Gibt die Punkte bei einem n Punkte langen Pfad in der Form [[x1,y1],[x2,y2],[xn,yn]] an
     * @param {string} [color = "black"] Gibt die farbe für den pfad an
     * @throws {InvalidPathLength} Wird geworfen, wenn der Pfad kürzer als drei Punkte ist
     */
    drawShape(path: number[][], color = "black"): void {
        if (path.length <= 2) throw "InvalidPathLength";
        path.push([0, 0]);
        let tmp = path.pop();
        this.ctx.beginPath();
        this.ctx.moveTo(tmp[0], tmp[1]);
        while (path.length != 0) {
            tmp = path.pop();
            this.ctx.lineTo(tmp[0], tmp[1]);
        }
        this.ctx.strokeStyle = color;
        this.ctx.closePath();
        this.ctx.stroke();
    }

    /**
     * Zeichnet einen Pfad durch eine Reihe von Punkten und schließt und füllt diesen am Ende
     * @param {number[][]} path Gibt die Punkte bei einem n Punkte langen Pfad in der Form [[x1,y1],[x2,y2],[xn,yn]] an
     * @param {string} [color = "black"] Gibt die farbe für den pfad an
     * @param {boolean} [outline = "false"] Wenn wahr, wird der Umriss mit gezeichnet
     * @param {string} [strokeColor = "black"] Gibt die farbe für die Umrandung an
     * @throws {InvalidPathLength} Wird geworfen, wenn der Pfad kürzer als drei Punkte ist
     */
    fillShape(path: number[][], color = "black", outline = false, strokeColor = "black"): void {
        if (path.length <= 2) throw "InvalidPathLength";
        path.push([0, 0]);
        let tmp = path.pop();
        this.ctx.beginPath();
        this.ctx.moveTo(tmp[0], tmp[1]);
        while (path.length != 0) {
            tmp = path.pop();
            this.ctx.lineTo(tmp[0], tmp[1]);
        }
        this.ctx.fillStyle = color;
        this.ctx.strokeStyle = strokeColor;
        this.ctx.closePath();
        if (outline) this.ctx.stroke();
        this.ctx.fill();
    }

    /**
     * Zeichnet den Umriss eines Kreises
     * @param {number} x Gibt die Position des Kreises an
     * @param {number} y Gibt die y-Position des Kreises an
     * @param {number} r Gibt den radius des Kreises an
     * @param {string} color Gibt die Farbe des Kreises an
     */
    drawCircle(x: number, y: number, r: number, color: string): void {
        this.ctx.strokeStyle = color;
        this.ctx.beginPath();
        this.ctx.arc(x, y, r, 0, 2 * Math.PI);
        this.ctx.stroke();
    }

    /**
     * Zeichnet eine Kreises
     * @param {number} x Gibt die Position des Kreises an
     * @param {number} y Gibt die y-Position des Kreises an
     * @param {number} r Gibt den radius des Kreises an
     * @param {string} color Gibt die Farbe des Kreises an
     * @param {boolean} [outline=false] Wenn wahr, wird der umriss gezeichnet
     * @param {string} [outColor="blac"] Gibt die Farbe der Umrandung an
     */
    fillCircle(x: number, y: number, r: number, color: string, outline = false, outColor = "black"): void {
        this.ctx.strokeStyle = outColor;
        this.ctx.fillStyle = color;
        this.ctx.beginPath();
        this.ctx.arc(x, y, r, 0, 2 * Math.PI);
        this.ctx.fill();
        if (outline) this.ctx.stroke();
        this.ctx.closePath();
    }

    toggleShowFPS(): void {
        this.showFPS = !this.showFPS;
    }

    drawFPS(): void {
        this.fillText(CANVAS_WIDTH - 80, 25, "FPS: " + this.fpsStats.FPS, "13px Times New Roman", "white");
        this.fillText(CANVAS_WIDTH - 80, 45, "Max. FPS: " + this.fpsStats.maxFPS, "13px Times New Roman", "white");
        this.fillText(CANVAS_WIDTH - 80, 65, "Min. FPS: " + this.fpsStats.minFPS, "13px Times New Roman", "white");
        this.fillText(CANVAS_WIDTH - 80, 85, "TPF: " + this.fpsStats.frameTime, "13px Times New Roman", "white");
    }

    calcFPS(): void {
        const now = performance.now();
        while (this.times.length > 0 && this.times[0] <= now - 1000) {
            this.times.shift();
        }
        this.times.push(now);
        this.fpsStats.FPS = this.times.length;
        if (this.fpsStats.FPS < this.fpsStats.minFPS && now >= 2000) this.fpsStats.minFPS = this.fpsStats.FPS;
        if (this.fpsStats.FPS > this.fpsStats.maxFPS) this.fpsStats.maxFPS = this.fpsStats.FPS;
    }

    tick(): void {
        this.fpsStats.frameTime = (performance.now() - this.time) / 1000;
        this.time = performance.now();
    }

    /**
     * Die Funktion konvertiert ein Canvas zu einem Bild.
     * @param {canvas} canvas Das zu konvertierende Canvas
     * @return {Image} Das Bild
     */
    static convertCanvasToImage(canvas: HTMLCanvasElement): HTMLImageElement {
        const image = new Image();
        image.src = canvas.toDataURL("image/png");
        return image;
    }
}
