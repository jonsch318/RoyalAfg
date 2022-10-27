import React, { FC, useEffect } from "react";
import { motion, useAnimation } from "framer-motion";

type PlayButtonProps = {
    onClick: () => void;
    disabled: boolean;
};

const PlayButton: FC<PlayButtonProps> = ({ onClick, disabled }) => {
    const press = () => {
        onClick();
    };

    return (
        <motion.button disabled={disabled} onClick={press} className="px-32 py-3 rounded-lg bg-red-700 text-white font-bold text-xl transition-all">
            Play
        </motion.button>
    );
};

export default motion(PlayButton);
