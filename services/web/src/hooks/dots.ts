import { useEffect, useState } from "react";

//small react hook for a dotted loading animation...
export const useDots = (time = 500): string => {
    const [dots, setDots] = useState(".");
    useEffect(() => {
        const inter = setInterval(() => {
            if (dots.length >= 3) {
                setDots(".");
            } else {
                setDots(dots + ".");
            }
        }, time);
        return () => {
            clearInterval(inter);
        };
    }, [dots]);

    return dots;
};
