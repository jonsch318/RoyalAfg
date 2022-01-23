/**
 * Die Klasse stellt funktionen zum besseren speichern der Position und Geschwindigkeit eines Objects zur verfügung
 * @file Stellt die Klasse {@link vector2D} zur verfügung
 * @example
 * let location = new vector2D(x, y);
 */
class vector2D {
    x: number;
    y: number;

    /**
     * Der Konstruktor setzt die Attribute x und y.
     * @param {number} x X ist die X Position des Vektors.
     * @param {number} y Y ist die Y Position des Vektors.
     */
    constructor(x: number, y: number) {
        /**
         * @type {number}
         */
        this.x = x;
        /**
         * @type {number}
         */
        this.y = y;
    }

    /**
     * @return {vector2D} Erstellt einen zufälligen 2D Vektor
     */
    static random2D(): vector2D {
        const v = new vector2D(Math.random() * 10, Math.random() * 10);
        v.normalize();
        return v;
    }

    /**
     * @param {vector2D} p1 Der Eingabevektor
     * @param {vector2D} p2 Der Vektor der dem Eingabevektor dazu addiert wird
     * @return {vector2D} Erstellt einen Vektor aus der Summe von p1 und p2.
     */
    static add(p1: vector2D, p2: vector2D): vector2D {
        return new vector2D(p1.x + p2.x, p1.y + p2.y);
    }

    /**
     * @param {vector2D} p1 Der Eingabevektor
     * @param {vector2D} p2 Der Vektor der dem Eingabevektor abgezogen wird
     * @return {vector2D} Erstellt einen Vektor aus der differenz von p1 und p2.
     */
    static sub(p1: vector2D, p2: vector2D): vector2D {
        return new vector2D(p1.x - p2.x, p1.y - p2.y);
    }

    /**
     * @param {vector2D} p1 Der Eingabevektor
     * @param {number} v Die Zahl mit der p1 multipliziert wird
     * @return {vector2D} Erstellt einen Vektor aus dem Produkt von p1 und v.
     */
    static mul(p1: vector2D, v: number): vector2D {
        return new vector2D(p1.x * v, p1.y * v);
    }

    /**
     * @param {vector2D} p1 Der Eingabevektor
     * @param {number} v Die Zahl durch die p1 dividiert wird
     * @return {vector2D} Erstellt einen Vektor aus dem quotienten von p1 und v.
     */
    static div(p1: vector2D, v: number): vector2D {
        return new vector2D(p1.x / v, p1.y / v);
    }

    /**
     * @return {vector2D} Gibt eine Kopie des Vektors zurück
     */
    get(): vector2D {
        const tmp = new vector2D(this.x, this.y);
        return tmp;
    }

    /**
     * Rundet die Werte des Vektors auf eine Grade Zahl
     */
    round(): void {
        this.x = Math.round(this.x);
        this.y = Math.round(this.y);
    }

    addScalar(scalar: number): vector2D {
        this.x += scalar;
        this.y += scalar;
        return this;
    }

    /**
     * Addiert einen Vektor auf den aktuellen
     * @param {vector2D} pVector Der Vektor, mit dem der aktuelle addiert wird
     */
    add(pVector: vector2D): vector2D {
        this.x += pVector.x;
        this.y += pVector.y;
        return this;
    }

    /**
     * Subtrahiert einen Vektor mit dem aktuellen
     * @param {vector2D} pVector Der Vektor, mit dem der aktuelle addiert wird
     */
    sub(pVector: vector2D): vector2D {
        this.x -= pVector.x;
        this.y -= pVector.y;
        return this;
    }

    /**
     * Multipliziert den Vektor mit einem Wert
     * @param {number} pFloat Der Wert mit dem der Vektor multipliziert wird
     */
    mul(pFloat: number): vector2D {
        this.x *= pFloat;
        this.y *= pFloat;
        return this;
    }

    /**
     * Dividiert den Vektor mit einer bestimmten zahl
     * @param {number} pFloat Der Wert, mit dem dividiert wird
     */
    div(pFloat: number): vector2D {
        this.x /= pFloat;
        this.y /= pFloat;
        return this;
    }

    /**
     * Die methode vergleicht den Vektor mit einem anderen Vektor
     * @param {vector2D} pVector Der Vektor mit dem verglichen wird.
     * @return {boolean} Gibt true zurück wenn die x- und y-Werte der Vektoren gleich sind, sonst wird false zurückgegeben
     */
    cmp(pVector: vector2D): boolean {
        if (this.x == pVector.x && this.y == pVector.y) return true;
        return false;
    }

    /**
     * Normalisiert den Vektor; setzt die Länge auf 1
     */
    normalize(): void {
        if (this.mag == 0) return;
        this.div(this.mag);
    }

    /**
     * Die Methode verhindert, dass der Vektor länger als ein bestimmter Wert wird
     * @param {number} v Der Maximalwert
     */
    limit(v: number): void {
        if (this.mag > v) {
            this.normalize();
            this.mul(v);
        }
    }

    /**
     * @return {string} Gibt die Werte des Vektors zurück
     */
    toString(): string {
        return "{ x: " + this.x + ", y: " + this.y + ", mag: " + this.mag + " }";
    }

    /**
     * Gibt die Länge des Vektors zurück
     * @type {number} */
    get mag() {
        return Math.sqrt(this.x * this.x + this.y * this.y);
    }
}

class vector3D {
    /**
     * The x coord. of the vector
     */
    x: number;
    /**
     * The y coord. of the vector.
     */
    y: number;
    /**
     * The z coord. of the vector.
     */
    z: number;

    /**
     * Der Konstruktor setzt die Attribute x und y.
     * @param {number} x X ist die X Position des Vektors.
     * @param {number} y Y ist die Y Position des Vektors.
     * @param {number} z Y ist die Z Position des Vektors.
     */
    constructor(x = 0, y = 0, z = 0) {
        this.x = x;
        this.y = y;
        this.z = z;
    }

    /**
     * @return {vector3D} Erstellt einen zufälligen 3D Vektor
     */
    static random3D(): vector3D {
        const v = new vector3D(Math.random() * 10, Math.random() * 10);
        v.normalize();
        return v;
    }

    /**
     * @param {vector3D} p1 Der Eingabevektor
     * @param {vector3D} p2 Der Vektor der dem Eingabevektor dazu addiert wird
     * @return {vector3D} Erstellt einen Vektor aus der Summe von p1 und p2.
     */
    static add(p1: vector3D, p2: vector3D): vector3D {
        return new vector3D(p1.x + p2.x, p1.y + p2.y, p1.z + p2.z);
    }

    /**
     * @param {vector3D} p1 Der Eingabevektor
     * @param {vector3D} p2 Der Vektor der dem Eingabevektor abgezogen wird
     * @return {vector3D} Erstellt einen Vektor aus der differenz von p1 und p2.
     */
    static sub(p1: vector3D, p2: vector3D): vector3D {
        return new vector3D(p1.x - p2.x, p1.y - p2.y, p1.z - p2.z);
    }

    /**
     * @param {vector3D} p1 Der Eingabevektor
     * @param {number} v Die Zahl mit der p1 multipliziert wird
     * @return {vector3D} Erstellt einen Vektor aus dem Produkt von p1 und v.
     */
    static mul(p1: vector3D, v: number): vector3D {
        return new vector3D(p1.x * v, p1.y * v, p1.z * v);
    }

    /**
     * @param {vector3D} p1 Der Eingabevektor
     * @param {number} v Die Zahl durch die p1 dividiert wird
     * @return {vector3D} Erstellt einen Vektor aus dem quotienten von p1 und v.
     */
    static div(p1: vector3D, v: number): vector3D {
        return new vector3D(p1.x / v, p1.y / v, p1.z / v);
    }

    /**
     * @return {vector3D} Gibt eine Kopie des Vektors zurück
     */
    get(): vector3D {
        const tmp = new vector3D(this.x, this.y, this.z);
        return tmp;
    }

    /**
     * Addiert einen Vektor auf den aktuellen
     * @param {vector3D} pVector Der Vektor, mit dem der aktuelle addiert wird
     */
    add(pVector: vector3D): void {
        this.x += pVector.x;
        this.y += pVector.y;
        this.z += pVector.z;
    }

    /**
     * Subtrahiert einen Vektor mit dem aktuellen
     * @param {vector3D} pVector Der Vektor, mit dem der aktuelle addiert wird
     */
    sub(pVector: vector3D): void {
        this.x -= pVector.x;
        this.y -= pVector.y;
        this.z -= pVector.z;
    }

    /**
     * Multipliziert den Vektor mit einem Wert
     * @param {number} pFloat Der Wert mit dem der Vektor multipliziert wird
     */
    mul(pFloat: number): void {
        this.x *= pFloat;
        this.y *= pFloat;
        this.z *= pFloat;
    }

    /**
     * Dividiert den Vektor mit einer bestimmten zahl
     * @param {number} pFloat Der Wert, mit dem dividiert wird
     */
    div(pFloat: number): void {
        this.x /= pFloat;
        this.y /= pFloat;
        this.z /= pFloat;
    }

    /**
     * Die methode vergleicht den Vektor mit einem anderen Vektor
     * @param {vector3D} pVector Der Vektor mit dem verglichen wird.
     * @return {boolean} Gibt true zurück wenn die x- und y- und z-Werte der Vektoren gleich sind, sonst wird false zurückgegeben
     */
    cmp(pVector: vector3D): boolean {
        if (this.x == pVector.x && this.y == pVector.y && this.z == pVector.z) return true;
        return false;
    }

    /**
     * Normalisiert den Vektor; setzt die Länge auf 1
     */
    normalize(): void {
        if (this.mag == 0) return;
        this.div(this.mag);
    }

    /**
     * Die Methode verhindert, dass der Vektor länger als ein bestimmter Wert wird
     * @param {number} v Der Maximalwert
     */
    limit(v: number): void {
        if (this.mag > v) {
            this.normalize();
            this.mul(v);
        }
    }

    /**
     * @return {string} Gibt die Werte des Vektors zurück
     */
    toString(): string {
        return "{ x: " + this.x + ", y: " + this.y + ", z: " + this.z + ", mag: " + this.mag + " }";
    }

    /**
     * Gibt die Länge des Vektors zurück
     * @type {number} */
    get mag(): number {
        return Math.sqrt(this.x * this.x + this.y * this.y + this.z * this.z);
    }
}

export { vector2D, vector3D };
