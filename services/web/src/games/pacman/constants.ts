export const SCATTER = 0;
export const CHASE = 1;

export const ORIGPATH = false;

/**
 * pause
 */
export const PAUSE = false;

/**
 * The size of one tile
 */
export const TILESIZE = 20;

/**
 * The (for now fixed) window width
 */
export const WINDOW_WIDTH = window.innerWidth;

/**
 * The (for now fixed) window height
 */
export const WINDOW_HEIGHT = window.innerHeight;

/**
 * The total row count
 */
export const ROW_COUNT = 31;

/**
 * The displayed column count
 */
export const COLUMN_COUNT = 28;

/**
 * Die Variabel beschreibt den Abstand nach links
 * @constant {number}   constXOffset
 */
export const OFFSET_X = (WINDOW_WIDTH - COLUMN_COUNT * TILESIZE) / 2;

/**
 * Die Variabel beschreibt den Abstand nach oben
 * @constant {number}   constYOffset
 */
export const OFFSET_Y = 80;

export const DEBUG = process.env.NODE_ENV == "development";

export const CANVAS_WIDTH = TILESIZE * COLUMN_COUNT + 2 * OFFSET_X;
export const CANVAS_HEIGHT = TILESIZE * ROW_COUNT + 2 * OFFSET_Y;
