import { toCell } from "./math";
import { xyToI } from "./Array";
import { COLUMN_COUNT } from "../constants";

const env = process.env.NODE_ENV;
const dev = "development";
/**
 * @file Die Datei stellt einen Logger zur verfügung.
 */
class objLog {
    out;
    objects;
    lvl: Level;
    /**
     *
     * @param {string} id Die id ist die Id des Elements, dem der Debugger den Text hinzufügt
     * @param {obj[] | obj} objList Das {@link obj|Objekt} oder die Liste der {@link obj|Objekte} die ausgelesen werden sollen.
     * @example
     * var log = new objLog("debug", [object1, object2]);
     * log.showOutput();
     */
    constructor(id, objList, lvl) {
        this.lvl = lvl;
        /**
         * @member {element} objLog~out Das Element, das den Text ausgibt
         */
        this.out = document.createElement("pre");
        this.out.setAttribute("id", "logOut");
        this.out.style.fontSize = "10px";
        document.getElementById(id).appendChild(this.out);
        /**
         * @member {element} objLog~objects Die Liste, der {@link obj|Objekte} die der geloggt werden sollen.
         */
        this.objects = objList;
        if (env == dev) this.showOutput();
    }

    /**
     * Die Methode fügt ein {@link obj|Objekt} zur {@link objLog~objects|Liste} hinzu.
     * @param {obj} pObject das Object, das zur {@link objLog~objects|Liste} hinzugefügt wird.
     */
    addObject(pObject) {
        this.objects.push(pObject);
    }

    /**
     * Die Methode aktualisiert den Output
     */
    showOutput() {
        try {
            if (env == dev) {
                this.out.innerHTML = "";
                this.objects.forEach((element) => {
                    this.out.innerHTML += "\n";
                    this.out.innerHTML += element.name + ": \n";
                    this.out.innerHTML += "\tPosition: \n";
                    this.out.innerHTML +=
                        "\t\tx: " +
                        (element.loc.x < 10 ? "00" + element.loc.x : element.loc.x < 100 ? "0" + element.loc.x : element.loc.x) +
                        (Number.isInteger(element.loc.x) ? ".0, Zelle: " : ", Zelle: ") +
                        toCell(element.loc).x +
                        "\n";
                    this.out.innerHTML +=
                        "\t\ty: " +
                        (element.loc.y < 10 ? "00" + element.loc.y : element.loc.y < 100 ? "0" + element.loc.y : element.loc.y) +
                        (Number.isInteger(element.loc.y) ? ".0, Zelle: " : ", Zelle: ") +
                        toCell(element.loc).y +
                        "\n";
                    this.out.innerHTML += "\tGeschwindigkeit: \n";
                    this.out.innerHTML += "\t\tx: " + element.vel.x + "\n";
                    this.out.innerHTML += "\t\ty: " + element.vel.y + "\n";
                    this.out.innerHTML += "\t\tv: " + element.vel.mag + "\n";
                    this.out.innerHTML += "\tCoin in Zelle: " + lvl.coins[xyToI(element.cell.x, element.cell.y, COLUMN_COUNT)] + "\n";
                });
            } else {
                this.out.innerHTML = "";
            }
            return 0;
        } catch {
            return 1;
        }
    }
}

export { objLog };
