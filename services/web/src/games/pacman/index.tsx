import React, { FC, useEffect, useRef, useState } from "react";
import { PacmanMain } from "./main";

const pcMain = new PacmanMain();

const PacmanVisual: FC = () => {
    const ref = useRef<HTMLDivElement>();
    const [started, setStarted] = useState(false);
    useEffect(() => {
        if (ref.current != null && !started && typeof window != "undefined") {
            pcMain.mainMenu(ref.current);
            setStarted(true);
        }
    }, [ref]);
    //<Script src="/lib/PathFinder/easystart-0.4.2.js" strategy="beforeInteractive" />
    return (
        <>
            <div id="canvasID">
                <div id="left" className="split" ref={ref}></div>
            </div>
        </>
    );
};

export default PacmanVisual;
