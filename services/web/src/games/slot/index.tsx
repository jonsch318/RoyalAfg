import React from "react";
import { SlotGame } from "./models/slot";
import { motion } from "framer-motion";
import SlotWheel from "./wheel";

type PlaySlotProps = {
    game: SlotGame;
};

const displayNumber = ["1", "2", "3", "4", "5", "6", "7", "J"];

const PlaySlot: React.FC<PlaySlotProps> = ({ game }) => {
    return (
        <div className="py-24">
            <h1>Play {JSON.stringify(game)}</h1>
            <motion.div className="flex flex-row justify-center py-24">
                {game.numbers.slice(0, game.numbers.length - 1).map((num, i) => {
                    return <SlotWheel number={num} key={i} />;
                })}
            </motion.div>
        </div>
    );
};

export default PlaySlot;
