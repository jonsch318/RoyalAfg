import React, { FC } from "react";
import { Suspense, useState } from "react";
import { motion, MotionConfig, useMotionValue } from "framer-motion";
import { Shapes, transition } from "./playButtonShapes";
import useMeasure from "react-use-measure";
import { css } from "@emotion/react";

type PlayButton3DProps = {
    onClick: () => void;
};

const PlayButton3D: FC<PlayButton3DProps> = ({ onClick }) => {
    const [ref, bounds] = useMeasure({ scroll: false });
    const [isHover, setIsHover] = useState(false);
    const [isPress, setIsPress] = useState(false);
    const mouseX = useMotionValue(0);
    const mouseY = useMotionValue(0);

    const resetMousePosition = () => {
        mouseX.set(0);
        mouseY.set(0);
    };

    return (
        <MotionConfig transition={transition}>
            <motion.button
                className="bg-[#acc7ed] text-white font-bold px-3 py-3 relative text-center flex items-center rounded-full text-5xl"
                ref={ref}
                initial={false}
                animate={isHover ? "hover" : "rest"}
                whileTap="press"
                onClick={onClick}
                variants={{
                    rest: { scale: 1 },
                    hover: { scale: 2 },
                    press: { scale: 1.4 }
                }}
                onHoverStart={() => {
                    resetMousePosition();
                    setIsHover(true);
                }}
                onHoverEnd={() => {
                    resetMousePosition();
                    setIsHover(false);
                }}
                onTapStart={() => setIsPress(true)}
                onTap={() => setIsPress(false)}
                onTapCancel={() => setIsPress(false)}
                onPointerMove={(e) => {
                    mouseX.set(e.clientX - bounds.x - bounds.width / 2);
                    mouseY.set(e.clientY - bounds.y - bounds.height / 2);
                }}>
                <motion.div
                    className="shapes absolute w-full h-full right-0 rounded-full"
                    css={css`
                        background: linear-gradient(60deg, #61dafb 0%, #d6cbf6 30%, #f2056f 70%);
                    `}
                    variants={{
                        rest: { opacity: 0 },
                        hover: { opacity: 1 }
                    }}>
                    <div className="pink blush absolute bottom-[-15px] w-24 blur-md h-8" />
                    <div className="blue blush" />
                    <div className="container absolute w-[calc(100% + 200px)]">
                        <Suspense fallback={null}>
                            <Shapes isHover={isHover} isPress={isPress} mouseX={mouseX} mouseY={mouseY} />
                        </Suspense>
                    </div>
                </motion.div>
                <motion.div variants={{ hover: { scale: 0.85 }, press: { scale: 1.1 } }} className="label w-[180px] px-5 ">
                    Play
                </motion.div>
            </motion.button>
        </MotionConfig>
    );
};

export default PlayButton3D;
