import React from "react";
import { SlotGame } from "./models/slot";
import { motion } from "framer-motion";
import SlotWheel from "./wheel";
import PlayButton from "../../widgets/games/slot/playButton";
import { requestSlotSpin } from "./provider";

type PlaySlotProps = {
    game: SlotGame;
    onError: (err: string) => void;
};

const displayNumber = ["1", "2", "3", "4", "5", "6", "7", "J"];

const PlaySlot: React.FC<PlaySlotProps> = ({ game, onError }) => {
    const play = async () => {
        console.log("play");
        /* await requestSlotSpin({ doubleFactor: false }, "test", crypto)
            .then((newGame) => {
                console.log(newGame);
                setGame(game);
            })
            .catch((err) => {
                onError(err.message);
            }); */
    };

    return (
        <div className="py-24">
            <h1>Play {JSON.stringify(game)}</h1>
            <motion.div className="flex flex-row justify-center py-24">
                {/*game.numbers.slice(0, 3).map((num, i) => {
                    return <SlotWheel number={num} key={i} order={i} started={tr} />;
                })*/}
            </motion.div>
            <PlayButton onClick={play}></PlayButton>
        </div>
    );
};

export default PlaySlot;
