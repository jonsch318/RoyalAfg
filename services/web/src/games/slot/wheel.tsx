import React, { useEffect, useRef } from "react";
import { AnimatePresence, motion, useMotionValue, useSpring, useTime, useTransform } from "framer-motion";
import { VALID_LOADERS } from "next/dist/shared/lib/image-config";

type SlotWheelProps = {
    number: number;
};

const SlotWheel: React.FC<SlotWheelProps> = ({ number }) => {
    const SLOT_WHEEL_HEIGHT = 128;
    const SLOT_NUMBER_COUNT = 8;

    const time = useTime();
    const y = useMotionValue(2 * -SLOT_WHEEL_HEIGHT);
    const ySpring = useSpring(y, { stiffness: 100, damping: 10 });

    //const rotation = useTransform(time, [0, 1000], [0, 0], { clamp: false });

    const [rotationCount, setRotationCount] = React.useState(50);
    const [wheelPos, setWheelPos] = React.useState(0);
    const ref = useRef<HTMLDivElement>(null);

    const [displayNumbers, setDisplayNumbers] = React.useState(["J", "1", "2", "3", "4", "5", "6", "7", "J", "1", "2"]);

    const ROTATION_DURATION = 0.25;

    const numPositionInWheel = (idx) => -idx * (360 / displayNumbers.length);

    const nextDisplayNumber = () => {
        setDisplayNumbers((x) => {
            const r = x.shift();
            x.push(r);
            return x;
        });
    };

    const nextNumber = () => {
        //setSpinAmount((x) => x - 1);
        setWheelPos((x) => (x + 1) % SLOT_NUMBER_COUNT);
        //setDisplayNumbers((x) => x.slice(1).concat(x[0]));
        if (wheelPos !== 0) {
            y.set((wheelPos + 1) * -SLOT_WHEEL_HEIGHT);
        } else {
            y.set(0, false);
        }
        //setDisplayNumbers((numbers) => numbers.slice(1, numbers.length).concat(tmp));
    };

    useEffect(() => {
        y.set(1 * -SLOT_WHEEL_HEIGHT);
        ySpring.onChange((v) => {
            console.log("On Change Animating: ", ySpring.isAnimating());
            if (ySpring.isAnimating()) {
                return;
            } else {
                //nextDisplayNumber();
            }
        });
    }, []);

    return (
        <motion.div className="px-8">
            <button onClick={() => nextNumber()}>
                Reload POS: {wheelPos} Numbers: {displayNumbers}
            </button>
            <motion.div className="h-[384px] truncate">
                <motion.div
                    ref={ref}
                    className="flex flex-col"
                    style={{ y: wheelPos == 0 ? 0 : ySpring }}
                    initial={{ y: 0 }}
                    transition={{ duration: 0.5 }}>
                    <AnimatePresence>
                        {displayNumbers.map((num, i) => {
                            return (
                                <motion.div
                                    initial={{ height: 0 }}
                                    exit={{ height: 0 }}
                                    animate={{ height: 128 }}
                                    transition={{ duration: 0.5 }}
                                    className="text-9xl font-bold px-12 "
                                    key={i}>
                                    {num}
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
