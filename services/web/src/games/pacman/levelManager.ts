import { ILevelObj, Level } from "./Klassen/level";

export class LevelManager {
    levelIndex: number;
    levelList = [
        "/static/games/pacman/levels/lvl1.json",
        "/static/games/pacman/levels/lvl2.json",
        "/static/games/pacman/levels/lvl3.json",
        "/static/games/pacman/levels/lvl4.json"
    ];
    current: Level;
    bbug = false;
    stop: () => void;

    constructor(start = 0, stop: () => void) {
        this.levelIndex = start;
        this.stop = stop;
    }

    async load(): Promise<Level> {
        //Load json file
        const res = await fetch(this.levelList[this.levelIndex % this.levelList.length], {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                Accepts: "application/json"
            }
        });

        if (!res.ok) {
            console.error(`Could not load level ${res.statusText}`);
            //TODO ERROR
            return;
        }

        try {
            const json = (await res.json()) as ILevelObj;

            //Parse object
            this.current = Level.from(json);
            console.log(this.current);
            return this.current;
        } catch (error) {
            console.error(`Could not load level: ${error}`);
            //TODO: ERROR
        }
    }

    async next(): Promise<Level> {
        //stop gameLoop -- How is irrelevant in this class
        this.stop();

        if (this.levelIndex == 256) this.bbug = true;
        else this.bbug = false;
        this.levelIndex++;
        return this.load();
    }
}
