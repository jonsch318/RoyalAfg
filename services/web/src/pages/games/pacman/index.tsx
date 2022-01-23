import React, { FC, useEffect, useState } from "react";
import dynamic from "next/dynamic";
import Head from "next/head";
import Layout from "../../../components/layout";

const PacmanVisual = dynamic(() => import("../../../games/pacman/index"), { ssr: false });

const Pacman: FC = () => {
    const [playing, setPlaying] = useState(false);

    return (
        <Layout footerAbsolute={!playing}>
            <Head>
                <title>Pacman</title>
            </Head>
            {!playing ? (
                <div>
                    <div className="grid justify-center items-center mt-64">
                        <button
                            className="px-6 py-2 bg-yellow-500 hover:bg-yellow-600 rounded"
                            onClick={() => {
                                setPlaying(true);
                            }}>
                            Play
                        </button>
                    </div>
                    Bet on how many Levels you pass and for each extra
                </div>
            ) : (
                <PacmanVisual></PacmanVisual>
            )}
        </Layout>
    );
};

export default Pacman;
