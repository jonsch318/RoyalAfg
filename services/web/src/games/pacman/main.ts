import { Canvas } from "./Draw/canvas";
import { Clyde } from "./Klassen/Ghosts/clyde";
import { Inky } from "./Klassen/Ghosts/inky";
import { Pinky } from "./Klassen/Ghosts/pinky";
import { Blinky } from "./Klassen/Ghosts/blinky";
import { Pacman } from "./Klassen/Pacman";
import { CHASE, SCATTER, CANVAS_HEIGHT, CANVAS_WIDTH, OFFSET_X, OFFSET_Y, PAUSE, TILESIZE } from "./constants";
import { Level } from "./Klassen/level";
import { random } from "./Functions/random";

import { LevelManager } from "./levelManager";
import { GameState } from "./gameState";
import { InputManager } from "./inputManager";
import { Ghost } from "./Klassen/Ghosts/ghost";

/**
 * @file Bereitet die variabeln vor und startet das Spiel
 * @module main
 */

let canvas: Canvas;
let levelManager: LevelManager;
let inputManager: InputManager;
let gameState: GameState;
let gameLoop: number;

export class PacmanMain {
    running = false;
    ghostMode: NodeJS.Timeout;

    bCanvas: HTMLCanvasElement;
    bImg: HTMLImageElement;
    /**
     * mainMenu is the start of the pacman game.
     * @returns
     */
    mainMenu(el: HTMLDivElement): void {
        if (el) {
            el.style.width = "100%";
        }
        canvas = new Canvas(null, "canvas", CANVAS_WIDTH, CANVAS_HEIGHT, "black", el);
        canvas.cls();
        this.init();
    }
    reset(): void {
        stop();
        this.init();
    }

    init() {
        //loadLevel("{\"grid\":[[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1],[1,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,1],[1,0,1,1,1,1,0,1,1,0,1,1,1,1,1,1,1,1,0,1,1,0,1,1,1,1,0,1],[1,0,1,1,1,1,0,1,1,0,1,1,1,1,1,1,1,1,0,1,1,0,1,1,1,1,0,1],[1,0,1,1,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,1,1,0,1],[1,0,1,1,0,1,1,1,1,0,1,1,0,1,1,0,1,1,0,1,1,1,1,0,1,1,0,1],[1,0,1,1,0,1,1,1,1,0,1,1,0,1,1,0,1,1,0,1,1,1,1,0,1,1,0,1],[1,0,1,1,0,1,1,1,1,0,1,1,0,1,1,0,1,1,0,1,1,1,1,0,1,1,0,1],[1,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,1],[1,0,1,1,1,1,0,1,1,0,1,1,1,1,1,1,1,1,0,1,1,0,1,1,1,1,0,1],[1,0,1,1,1,1,0,1,1,0,1,1,1,1,1,1,1,1,0,1,1,0,1,1,1,1,0,1],[1,0,1,1,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,1,1,0,1],[1,0,1,1,0,1,1,1,1,0,1,1,1,0,0,1,1,1,0,1,1,1,1,0,1,1,0,1],[1,0,1,1,0,1,1,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,1,0,1,1,0,1],[1,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,1],[1,1,1,1,0,1,1,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,1,0,1,1,1,1],[1,1,1,1,0,1,1,1,1,0,1,1,1,1,1,1,1,1,0,1,1,1,1,0,1,1,1,1],[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1],[1,0,1,1,1,1,1,1,1,1,1,1,0,1,1,0,1,1,1,1,1,1,1,1,1,1,0,1],[1,0,1,1,1,1,1,1,1,1,1,1,0,1,1,0,1,1,1,1,1,1,1,1,1,1,0,1],[1,0,0,0,0,0,0,0,1,1,0,0,0,1,1,0,0,0,1,1,0,0,0,0,0,0,0,1],[1,0,1,1,1,1,1,0,1,1,0,1,1,1,1,1,1,0,1,1,0,1,1,1,1,1,0,1],[1,0,1,1,1,1,1,0,1,1,0,1,1,1,1,1,1,0,1,1,0,1,1,1,1,1,0,1],[1,0,0,0,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0,0,0,1],[1,0,1,1,0,1,1,0,1,1,1,1,0,1,1,0,1,1,1,1,0,1,1,0,1,1,0,1],[1,0,1,1,0,1,1,0,1,1,1,1,0,1,1,0,1,1,1,1,0,1,1,0,1,1,0,1],[1,0,1,1,0,0,0,0,1,1,0,0,0,1,1,0,0,0,1,1,0,0,0,0,1,1,0,1],[1,0,1,1,1,1,1,0,1,1,0,1,1,1,1,1,1,0,1,1,0,1,1,1,1,1,0,1],[1,0,1,1,1,1,1,0,1,1,0,1,1,1,1,1,1,0,1,1,0,1,1,1,1,1,0,1],[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1],[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]],\"pacmanStart\":{\"x\":1,\"y\":1},\"blinkyStart\":{\"x\":1,\"y\":5},\"clydeStart\":{\"x\":10,\"y\":5},\"speedyStart\":{\"x\":10,\"y\":10},\"inkyStart\":{\"x\":10,\"y\":10},\"identification\":1}");
        levelManager = new LevelManager(0, stop);
        const pac = new Pacman(canvas);
        gameState = new GameState(pac, [new Blinky(canvas, pac), new Clyde(canvas, pac), new Pinky(canvas, pac), new Inky(canvas, pac)]);
        inputManager = new InputManager(levelManager, gameState);
        levelManager
            .load()
            .then((lvl) => this.onLevelLoad(lvl))
            .catch((e) => console.error("Catched: ", e));
    }
    registerKeys(): void {
        inputManager.registerKeyHandler("w", () => gameState.changePacmanDir("up"));
        inputManager.registerKeyHandler("ArrowUp", () => gameState.changePacmanDir("up"));
        inputManager.registerKeyHandler("s", () => gameState.changePacmanDir("down"));
        inputManager.registerKeyHandler("ArrowDown", () => gameState.changePacmanDir("down"));
        inputManager.registerKeyHandler("a", () => gameState.changePacmanDir("left"));
        inputManager.registerKeyHandler("ArrowLeft", () => gameState.changePacmanDir("left"));
        inputManager.registerKeyHandler("d", () => gameState.changePacmanDir("right"));
        inputManager.registerKeyHandler("ArrowRight", () => gameState.changePacmanDir("right"));

        inputManager.registerKeyHandler("r", (e, s) => {
            if (s == 0b101) this.reset();
        });
        inputManager.registerKeyHandler("f", (e, s) => {
            if (s == 0b1) canvas.toggleShowFPS();
        });
        inputManager.detectSwipe("left", inputManager.setKey);
    }

    onLevelLoad(lvl: Level): void {
        console.log("called level load: " + canvas);
        canvas.lvl = lvl;
        this.registerKeys();
        inputManager.register();
        console.log("GameState: ", gameState);
        gameState.newLevel(lvl);
        setTimeout(() => {
            this.running = true;
            //_logger_ = new objLog("debugText", objects);
            lvl.mode = CHASE;
            tick();
            lvl.mode = SCATTER;
            tick();
            this.ghostMode = setTimeout(
                (t) => {
                    console.log("CHASE");
                    lvl.mode = CHASE;
                    this.ghostMode = setTimeout(
                        (t) => {
                            if (t != levelManager.levelIndex) return;
                            console.log("SCATTER");
                            gameState.ghosts.forEach((element) => {
                                element.velocity.mul(-1);
                            });
                            lvl.mode = SCATTER;
                            this.ghostMode = setTimeout(
                                (t) => {
                                    if (t != levelManager.levelIndex) return;
                                    console.log("CHASE");
                                    lvl.mode = CHASE;
                                    this.ghostMode = setTimeout(
                                        (t) => {
                                            if (t != levelManager.levelIndex) return;
                                            console.log("SCATTER");
                                            gameState.ghosts.forEach((element) => {
                                                element.velocity.mul(-1);
                                            });
                                            lvl.mode = SCATTER;
                                            this.ghostMode = setTimeout(
                                                (t) => {
                                                    if (t != levelManager.levelIndex) return;
                                                    console.log("CHASE");
                                                    lvl.mode = CHASE;
                                                    this.ghostMode = setTimeout(
                                                        (t) => {
                                                            if (t != levelManager.levelIndex) return;
                                                            console.log("SCATTER");
                                                            gameState.ghosts.forEach((element) => {
                                                                element.velocity.mul(-1);
                                                            });
                                                            lvl.mode = SCATTER;
                                                            this.ghostMode = setTimeout(
                                                                (t) => {
                                                                    if (t != levelManager.levelIndex) return;
                                                                    console.log("CHASE");
                                                                    lvl.mode = CHASE;
                                                                },
                                                                5000,
                                                                t
                                                            );
                                                        },
                                                        20000,
                                                        t
                                                    );
                                                },
                                                5000,
                                                t
                                            );
                                        },
                                        20000,
                                        t
                                    );
                                },
                                7000,
                                t
                            );
                        },
                        20000,
                        t
                    );
                },
                7000,
                levelManager.levelIndex
            );
            gameLoop = requestAnimationFrame(tick);
        }, 50);
        this.running = false;
    }

    stop() {
        clearTimeout(this.ghostMode);
        cancelAnimationFrame(gameLoop);
        this.running = false;
        gameState = null;
    }

    bug() {
        if (this.bImg === null) {
            this.bImg = new Image();
            //Quellen: https://www.pinterest.com/pin/462393086712054233/ http://kafumble.blogspot.com/2011/05/pacman.html
            this.bImg.src = ".\\img\\BugSprites.jpg";
            this.bImg.onload = function () {
                createBug();
            };
        }
        if (gameState.bbug) {
            //bCtx.clearRect(0, 0, bCanvas.width, bCanvas.height);
            canvas.ctx.drawImage(this.bCanvas, 0, 0);
        }
    }
}

function tick(): void {
    if (canvas) canvas.tick();

    if (PAUSE) {
        return;
    }
    canvas.cls();
    drawInfoText();
    gameState.move();
    draw();
    const collision = collisionCheck();
    if (collision.length > 0) {
        gameState.removeLife();
        levelManager.next().then((lvl) => {
            if (gameState.newLevel(lvl)) {
                stop();
                alert("Du hast verloren!");
            }
        });
    }

    //check for level complete
    if (levelManager.current.coinCount == 0) levelManager.next().then();

    //draw FPS
    if (canvas.showFPS) canvas.drawFPS();
    canvas.calcFPS();

    gameLoop = requestAnimationFrame(tick);
}

function createBug(): void {
    this.bImg = null;
    this.bCanvas = document.createElement("canvas");
    this.bCanvas.setAttribute("width", CANVAS_WIDTH.toString());
    this.bCanvas.setAttribute("height", CANVAS_HEIGHT.toString());
    const bCtx = this.bCanvas.getContext("2d");
    bCtx.fillRect(canvas.width / 2, 0, canvas.width, canvas.height);
    for (let i = 0; i < 256; i++) {
        const ix = random.getRandomInt(0, this.bImg.width);
        const iy = random.getRandomInt(0, this.bImg.height);
        const x = random.getRandomInt(canvas.width / 2, canvas.width - OFFSET_X);
        const y = random.getRandomInt(OFFSET_Y, canvas.height - OFFSET_Y);
        bCtx.drawImage(this.bImg, ix, iy, TILESIZE, TILESIZE, x, y, TILESIZE, TILESIZE);
    }
}

function drawInfoText(): void {
    canvas.fillArray(levelManager.current.grid, ["black", levelManager.current.color, "red"], TILESIZE, TILESIZE, OFFSET_X, OFFSET_Y);
    canvas.fillText(OFFSET_X, OFFSET_Y, "Score: " + gameState.points, "20px Arial", "white");
    canvas.fillText(OFFSET_X + (CANVAS_WIDTH - OFFSET_X * 1.15) / 2, OFFSET_Y, "Lives: " + gameState.pacman.lives, "20px Arial", "white");
}

function draw(): void {
    gameState.pacman.draw();
    gameState.ghosts.forEach((el) => el.draw());
}
function collisionCheck(): Ghost[] {
    const collisions = [];
    gameState.ghosts.forEach((el) => {
        if (el.cell.cmp(gameState.pacman.cell)) {
            collisions.push(el);
        }
    });
    if (collisions.length > 0) {
        console.log("Collisions: ", collisions);
    }
    return collisions;
}
