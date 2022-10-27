import React, { FC, useEffect, useState } from "react";
import Layout from "../../../components/layout";
import { CryptoInfoDTO } from "../../../games/slot/dtos/crypto";
import { AnimatePresence, motion } from "framer-motion";
import { SlotGame } from "../../../games/slot/models/slot";
import SlotVisual from "../../../games/slot/visual";
import PlayButton from "../../../widgets/games/slot/playButton";
import { requestSlotSpin, getCryptoInfo } from "../../../games/slot/provider";
import PlayButton3D from "../../../widgets/games/slot/playButton3d";
import { GetStaticPropsResult } from "next";

const render3D = true;

export async function getStaticProps(context): Promise<GetStaticPropsResult<SlotProps>> {
    const CryptoInfo = await getCryptoInfo();

    return {
        props: {
            crypto: CryptoInfo
        },
        revalidate: 100
    };
}

type SlotProps = {
    crypto: CryptoInfoDTO;
};

const BeachAndSun: FC<SlotProps> = ({ crypto }) => {
    const [game, setGame] = useState<SlotGame>();
    const [started, setStarted] = useState<number>(0);

    const play = async () => {
        requestSlotSpin({ doubleFactor: false }, "test", crypto)
            .then((game) => {
                if (game === null) {
                    throw new Error("Game fetch error");
                }
                setGame(game);
                setStarted(2);
            })
            .catch((err) => {
                console.log(err);
            });
        setStarted(1);
    };

    useEffect(() => {
        console.info(game);
    }, [game]);

    return (
        <Layout footerAbsolute headerAbsolute disableFooter>
            <div className="pt-[56px]">
                <motion.h1 initial={{ y: 600 }} animate={{ y: 0 }} transition={{ delay: 0.2 }} className="font-semibold text-3xl text-center my-4">
                    Beach & Sun {started}{" "}
                    {game?.numbers.map((x, i) => {
                        return (
                            <span className="inline" key={i}>
                                {x}
                            </span>
                        );
                    })}
                </motion.h1>
                <motion.div
                    initial={{ y: 100, scale: 0.5 }}
                    animate={{ y: 0, scale: 1 }}
                    transition={{ delay: 0.3 }}
                    className="flex flex-row justify-center py-16">
                    <SlotVisual
                        game={game}
                        started={started}
                        finished={() => {
                            console.log("set started");
                            setStarted(0);
                        }}></SlotVisual>
                </motion.div>

                <AnimatePresence>
                    <motion.div
                        initial={{ y: 100, scale: 0.5 }}
                        animate={{ y: 0, scale: 1 }}
                        transition={{ delay: 0.3 }}
                        exit={{ scale: 0, transition: { duration: 0.2, delay: 0 } }}
                        className="grid justify-center"
                        key={"playButton"}>
                        {render3D ? <PlayButton3D onClick={play}></PlayButton3D> : <PlayButton disabled={started !== 0} onClick={play} />}
                    </motion.div>
                </AnimatePresence>
            </div>
        </Layout>
    );
};

export default BeachAndSun;
