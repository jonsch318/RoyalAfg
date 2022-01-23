import { Ghost } from "./Klassen/Ghosts/ghost";
import { Level } from "./Klassen/level";
import { Pacman } from "./Klassen/Pacman";

export class GameState {
    private _points: number;
    private _bbug: boolean;
    private lvl: Level;

    readonly pacman: Pacman;
    readonly ghosts: Ghost[];

    constructor(pacman: Pacman, ghosts: Ghost[]) {
        this.pacman = pacman;
        this.ghosts = ghosts;
    }

    public get points(): number {
        return this._points;
    }

    public get bbug(): boolean {
        return this._bbug;
    }

    public set bbug(val: boolean) {
        this._bbug = val;
    }

    addPoints(): void {
        this._points += 100;
    }

    removeLife(): void {
        this.pacman.lives--;
    }

    move(): void {
        if (this.pacman) this.pacman.move();
        this.ghosts.forEach((val) => {
            val.move();
        });
    }

    changePacmanDir(key: "up" | "down" | "left" | "right") {
        if (this.pacman) this.pacman.changeDir(key);
    }

    /**
     * @returns the flag, whether the game has ended or not.
     */
    newLevel(lvl: Level): boolean {
        this.lvl = lvl;
        if (this.pacman.lives <= 0) return true;

        this.pacman.loadLevel(this.lvl);
        this.pacman.getToNearestFree();

        for (const cha of this.ghosts) cha.loadLevel(this.lvl);
        return false;
    }
}
