import { COLUMN_COUNT, ROW_COUNT } from "../constants";
import { vector2D } from "../Functions/vector";

export interface ILevelObj {
    grid: number[][];
    startPos: Record<string, vector2D>;
    identification: number;
    color: string;
    coins: boolean[];
    coinCount: number;
}

/**
 * @file Die Klasse speichert die Daten für verschiedene Level
 */
class Level {
    grid: number[][];
    private startPos: Record<string, vector2D>;
    identification: number;
    color: string;
    coins: boolean[];
    coinCount: number;
    name: string;
    mode: number;

    /**
     *
     * @param {number[][]} grid Die tilemap
     * @param {Record<string, vector2D>} startPositions Die Map der Startpositionen mit Name als key und position als value.
     * @param {string} [color="blue"] Die Farbe des Levels
     * @param {number} [id=-1] Die Id
     */
    constructor(grid: number[][], startPositions: Record<string, vector2D>, color = "blue", id = -1) {
        /**
         * @member {number[][]} level~grid Die Tilemap ist eine vereinfachte Darstellung des Levels, in der Wände mit einer 1 dargestellt werden
         */
        this.grid = grid;
        this.startPos = startPositions;

        /**
         * @member {number} level~identification Die Identification Nummer ist die Nummer, die Angibt an welcher Stelle das Level ist.
         */
        this.identification = id;

        /**
         * @member {string} level~color Die Farbe die das Level haben soll
         */
        this.color = color;

        /**
         * @member {boolean[]} level~coins Ein Array in dem für jede Position gespeichert wird ob sich dort eine Münze befindet
         */
        this.coins = [];
        for (let i = 0; i < ROW_COUNT * COLUMN_COUNT; i++) {
            this.coins.push(false);
        }

        /**
         * @member {number} level~coinAnzahl Die Anzahl an Münzen im Level
         */
        this.coinCount = 0;

        /**
         * @member {string} [level~name = "Level"] Der Name dieser Klasse
         * @private
         */
        this.name = "Level";

        /**
         * @member {number} level~mode Die Art der Zielbestimmung der Geister
         */
        this.mode = -1;
    }

    addNewCharakter(name: string, startPos: vector2D): void {
        this.startPos[name] = startPos;
    }

    getStartPos(charName: string): vector2D {
        return this.startPos[charName] ? this.startPos[charName] : new vector2D(0, 0);
    }

    /**
     * Erstellt aus einem JavaScript eine Instanz der Klasse
     * @param {object} json Das Objekt aus dem die Klasse erstellt werden soll
     * @example
     * var level = level.from(JSON.parse(Jsonfile));
     */
    static from(json: ILevelObj): Level {
        const lvl = new Level(json.grid, json.startPos, json.color, json.identification);
        lvl.coins = json.coins;
        lvl.coinCount = json.coinCount;
        return lvl;
    }

    /**
     * Der getter id gibt die Levelnummer zurück
     * @return {number}
     */
    get id(): number {
        return this.identification;
    }

    /**
     * Der getter gibt das Level als JSON-Objekt zurück.
     * @return {string} Das Level
     */
    get JSON(): string {
        return JSON.stringify(this);
    }
}

export { Level };
