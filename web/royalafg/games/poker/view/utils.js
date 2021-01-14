import * as PIXI from "pixi.js";

let app;

function registerApp(r) {
    app = r;
}

function isMobile() {
    return PIXI.utils.isMobile.any;
}

// Converts a relativ width to the respective width. relative weight is for a normal 1080p browser window
function rW(r) {
    const f = r / 1920;
    return f * app.renderer.width;
}

// Converts a relativ Height to the respective Height relative Height is for a normal 1080p browser window
function rH(r) {
    const f = r / 937;
    return f * app.renderer.height;
}

export { rW, rH, registerApp, isMobile };
