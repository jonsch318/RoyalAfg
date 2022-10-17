import { GetStaticPropsResult } from "next";
import Image from "next/image";
import { useRouter } from "next/router";
import React, { FC } from "react";
import Layout from "../../../components/layout";
import PlaySlot from "../../../games/slot";
import { SlotProvider } from "../../../games/slot/context";
import { CryptoInfoDTO } from "../../../games/slot/dtos/crypto";
import { SlotGame, SlotGameSchema } from "../../../games/slot/models/slot";
import { getCryptoInfo, requestSlotSpin } from "../../../games/slot/provider";
import { useDots } from "../../../hooks/dots";
import ProgressBar from "../../../widgets/progressBar/progressBar";

type SlotProps = {
    crypto: CryptoInfoDTO;
};

export async function getStaticProps(context): Promise<GetStaticPropsResult<SlotProps>> {
    const CryptoInfo = await getCryptoInfo();

    return {
        props: {
            crypto: CryptoInfo
        },
        revalidate: 100
    };
}

const Slot: FC<SlotProps> = ({ crypto }) => {
    const router = useRouter();
    const [playState, setPlayState] = React.useState(0);
    const [game, setGame] = React.useState<SlotGame>();
    const [errorText, setErrorText] = React.useState("");
    const [ready, setReady] = React.useState(false);
    const dots = useDots();

    const play = async () => {
        requestSlotSpin({ doubleFactor: false }, "test", crypto)
            .then((game) => {
                console.log(game);
                setGame(game);
                setReady(true);
            })
            .catch((err) => {
                setPlayState(0);
                setErrorText("An error occured... Error: " + err.message);
            });
        setPlayState(1);
        //setReady(true);
    };

    const playCallback = async () => {
        setPlayState(2);

        //router.push("/games/slot/play", undefined, { shallow: true });
    };

    const renderPlayStateSwitch = () => {
        switch (playState) {
            case 0:
                return (
                    <button
                        onClick={play}
                        className="px-32 py-3 rounded-lg bg-red-700 text-white font-bold text-xl hover:scale-105 hover:rounded-xl transition-all">
                        Play
                    </button>
                );
            case 1:
                return (
                    <div className="flex flex-col justify-center items-center align-middle">
                        <ProgressBar completed={ready} callback={playCallback}></ProgressBar>
                        <h2 className="text-center w-10">Loading{dots}</h2>
                    </div>
                );
        }
    };

    return (
        <Layout disableFooter headerAbsolute={playState == 2}>
            {playState == 2 ? (
                <PlaySlot game={game} />
            ) : (
                <>
                    <h1 className="font-semibold text-4xl text-center mx-4">Slot Machine</h1>
                    <div className="grid justify-center my-4" style={{ gridTemplateRows: "1fr auto" }}>
                        <Image
                            src="/static/games/slot/beachsun/pb.png"
                            alt="Profilepicture of the beach & sun slot machine"
                            width={500}
                            height={500}
                        />
                        <h2 className="font-semibold text-3xl text-center my-4">Beach & Sun</h2>
                    </div>
                    {errorText != "" ? <p className="text-red-700 text-center text-lg font-semibold mb-8">{errorText}</p> : <></>}
                    <div className="grid justify-center w-full">{renderPlayStateSwitch()}</div>
                </>
            )}
        </Layout>
    );
};

export default Slot;
