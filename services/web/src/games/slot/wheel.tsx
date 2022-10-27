import React, { useEffect, useState } from "react";
import { animate, AnimatePresence, motion, useMotionValue } from "framer-motion";
import lodash from "lodash";
import { SlotNumberDisplay } from "./display";
import { color } from "@mui/system";

type SlotWheelProps = {
    number: NumberDisplay;
    order: number;
    started: boolean;
    finished?: (order: number) => void;
};

type NumberDisplay = {
    number: number;
    display: string;
};

const initNumbers: NumberDisplay[] = [
    { number: 0, display: SlotNumberDisplay[0] },
    { number: 1, display: SlotNumberDisplay[1] },
    { number: 2, display: SlotNumberDisplay[2] },
    { number: 3, display: SlotNumberDisplay[3] },
    { number: 4, display: SlotNumberDisplay[4] },
    { number: 5, display: SlotNumberDisplay[5] },
    { number: 6, display: SlotNumberDisplay[6] },
    { number: 7, display: SlotNumberDisplay[7] }
];

//Durstenfeld shuffle
const shuffle = <T,>(array: T[]): void => {
    for (let i = array.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [array[i], array[j]] = [array[j], array[i]];
    }
};

const SLOT_WHEEL_HEIGHT = 130;
const SLOT_NUMBER_COUNT = initNumbers.length;
const STARTING_DURATION = 0.025;
const SLOT_DURATION_FACTOR = 1.2;
const SLOT_SPIN_COUNT = 10;

const SlotWheel: React.FC<SlotWheelProps> = ({ number, order, started, finished }) => {
    const y = useMotionValue(2 * -SLOT_WHEEL_HEIGHT);
    const [wheelPos, setWheelPos] = React.useState(1);
    const [phase, setPhase] = useState(0);
    const [displayNumbers, setDisplayNumbers] = React.useState<NumberDisplay[]>([]);

    const [initialized, setInitialized] = useState(false);

    const prepareNumbers = () => {
        const newNumbers = initNumbers.slice();
        shuffle(newNumbers);
        const first = newNumbers.slice(0, 2);
        newNumbers.unshift(newNumbers[newNumbers.length - 1]);
        first.forEach((x) => newNumbers.push(x));
        setDisplayNumbers(newNumbers);
    };

    useEffect(() => {
        if (!initialized) {
            prepareNumbers();
            y.onChange(() => {
                if (y.get() <= SLOT_NUMBER_COUNT * -SLOT_WHEEL_HEIGHT) {
                    y.set(0);
                }
            });
            setInitialized(true);
        }
        return () => {
            y.clearListeners();
        };
    }, [initialized]);

    const moveToCorrectNumber = async (dur: number) => {
        const target = displayNumbers.slice(1, displayNumbers.length - 2).findIndex((x) => x.number === number.number);
        if (target == -1) {
            throw new Error("Could not find target in displayed numbers");
        }
        if (target === wheelPos) {
            return move(8, dur);
        } else {
            return move(target > wheelPos ? target - wheelPos : SLOT_NUMBER_COUNT - wheelPos + target, dur);
        }
    };

    const moveDigit = (dur: number): Promise<void> => {
        if (y.get() <= SLOT_NUMBER_COUNT * -SLOT_WHEEL_HEIGHT) y.set(1);
        return new Promise<void>((resolve) => {
            animate(y, y.get() - SLOT_WHEEL_HEIGHT, {
                duration: dur,
                onComplete: () => {
                    resolve();
                }
            });
        });
    };

    const move = async (num: number, dur: number) => {
        for (let i = 0; i < num; i++) {
            if (!started) return;
            dur *= 1 + 1.15 / num;
            setWheelPos((x) => (x + 1) % SLOT_NUMBER_COUNT);
            await moveDigit(dur);
        }
    };

    const ensureCorrectNumber = () => {
        const target = displayNumbers.slice(1, displayNumbers.length - 2).findIndex((x) => x.number === number.number);
        y.set(target * -SLOT_WHEEL_HEIGHT);
        if (target % SLOT_NUMBER_COUNT !== wheelPos) {
            setWheelPos(target % SLOT_NUMBER_COUNT);
        }
    };

    useEffect(() => {
        if (started) {
            //setDuration(STARTING_DURATION);
            setTimeout(() => {
                setPhase(1);
            }, order * 400);
        } else {
            y.stop();
        }
    }, [started]);

    useEffect(() => {
        if (phase === 1) {
            spin(SLOT_SPIN_COUNT, SLOT_DURATION_FACTOR).then(() => {
                y.set(0);
                setWheelPos(0);
                setPhase(2);
            });
        } else if (phase === 2) {
            moveToCorrectNumber(STARTING_DURATION * 1.25 * 8).then(() => {
                setPhase(3);
                finished(order);
            });
        } else if (phase === 3) {
            setTimeout(() => {
                ensureCorrectNumber();
            }, 100);
        }
    }, [phase]);

    const spin = async (spinCount, durationFactor: number, startingDur = STARTING_DURATION) => {
        let dur = startingDur;
        for (let i = 0; i < spinCount; i++) {
            await spinAnime(dur);
            dur *= durationFactor;
        }
    };

    const spinAnime = async (duration) => {
        return new Promise<void>((resolve) => {
            y.set(0);
            animate(y, SLOT_NUMBER_COUNT * -SLOT_WHEEL_HEIGHT, {
                duration: 8 * duration,
                ease: "linear",
                onComplete: () => {
                    resolve();
                }
            });
        });
    };

    const color = (phase: number, i: number): string => {
        if (phase < 2) return "text-black";
        if (wheelPos === 0 && (i === 0 || i === 8)) {
            return "text-red-700";
        } else if (wheelPos === 1 && (i === 1 || i === 9)) {
            return "text-red-700";
        } else if (wheelPos === i) {
            return "text-red-700";
        }
        return "text-black";
    };

    return (
        <motion.div
            className="px-8"
            initial={{ paddingRight: "8rem", paddingLeft: "8rem" }}
            animate={{ paddingRight: "2rem", paddingLeft: "2rem" }}
            transition={{ duration: 2, ease: "anticipate" }}>
            <motion.div className="h-[380px] truncate">
                <motion.div className="flex flex-col" style={{ y: y }}>
                    <AnimatePresence>
                        {displayNumbers.map((num, i) => {
                            return (
                                <motion.div
                                    initial={{ scale: 0.25 }}
                                    exit={{ scale: 0 }}
                                    animate={{ scale: 1 }}
                                    transition={{ duration: 3, ease: [0, 0.7, 0.2, 1] }}
                                    className={"text-9xl py-[1px] font-bold transition-colors duration-500 px-12 " + color(phase, i - 1)}
                                    key={i}>
                                    {num.display}
                                </motion.div>
                            );
                        })}
                    </AnimatePresence>
                </motion.div>
            </motion.div>
        </motion.div>
    );
};

export default SlotWheel;
