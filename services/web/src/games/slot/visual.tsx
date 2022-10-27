import React, { FC, useEffect, useState } from "react";
import { SlotNumberDisplay } from "./display";
import { SlotGame } from "./models/slot";
import SlotWheel from "./wheel";

type SlotVisualProps = {
    game: SlotGame;
    started: number;
    finished: () => void;
};

type NumberType = {
    number: number;
    display: string;
};

const numberToDTO = (number: number): NumberType => {
    return { number: number, display: SlotNumberDisplay[number] };
};

const SlotVisual: FC<SlotVisualProps> = ({ game, started, finished }) => {
    const [progress, setProgress] = React.useState(0);
    const [extended, setExtended] = useState(0);

    useEffect(() => {
        console.log("Finieshed: ", progress);
        if (progress === 3) {
            if (!shouldExtend) {
                setProgress(6);
            } else {
                setProgress(4);
                setTimeout(() => {
                    extend();
                }, 1000);
            }
        } else if (progress === 6) {
            if (shouldExtend && extended) {
                finished();
            }
        }
    }, [progress]);

    useEffect(() => {
        if (started > 0) {
            setProgress(0);
            setExtended(0);
        }
    }, [started]);

    const shouldExtend = (numbers: number[]): boolean => {
        for (let i = 0; i < 3; i++) {
            if (numbers[i] !== 0) {
                return false;
            }
        }
        return true;
    };

    const wheelUpdate = (wheel: number) => {
        console.log("Wheel notifiesw: ", wheel);
        setProgress((x) => x + 1);
    };

    const extend = () => {
        setProgress(5);
        setExtended(1);
    };

    return (
        <>
            {progress !== 5 && !extended ? (
                [1, 1, 1].map((n, i) => {
                    return (
                        <SlotWheel
                            number={numberToDTO(game?.numbers[i])}
                            finished={wheelUpdate}
                            key={i}
                            order={i}
                            started={!extended && started === 2}></SlotWheel>
                    );
                })
            ) : (
                <></>
            )}
            {progress >= 5 && extended ? (
                <SlotWheel number={numberToDTO(game?.numbers[game.numbers.length - 1])} finished={wheelUpdate} started={started === 2} order={0} />
            ) : (
                <></>
            )}
        </>
    );
};

export default SlotVisual;
