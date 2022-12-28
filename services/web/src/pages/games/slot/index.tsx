import { GetStaticPropsResult } from "next";
import Image from "next/image";
import { useRouter } from "next/router";
import React, { FC, useState } from "react";
import Layout from "../../../components/layout";
import PlaySlot from "../../../games/slot";
import { CryptoInfoDTO } from "../../../games/slot/dtos/crypto";
import { getCryptoInfo } from "../../../games/slot/provider";
import { AnimatePresence, motion } from "framer-motion";
import Link from "next/link";

const Slot: FC = () => {
    return (
        <Layout disableFooter>
            <motion.h1 initial={{ scale: 0.5 }} animate={{ scale: 1 }} exit={{ scale: 0 }} className="font-bold text-5xl text-center my-8 mx-4">
                Slot Machine (Not yet ready)
            </motion.h1>
            <motion.div
                initial={{ scale: 0.35 }}
                animate={{ scale: 1 }}
                exit={{ scale: 0.35 }}
                className="grid justify-center my-4"
                style={{ gridTemplateRows: "1fr auto" }}>
                <Image src="/static/games/slot/beachsun/pb.png" alt="Profilepicture of the beach & sun slot machine" width={400} height={400} />
                <h2 className="font-bold text-4xl text-center my-8">Beach & Sun</h2>
            </motion.div>
            <motion.div className="grid justify-center w-full">
                <Link href={"/games/slot/beachandsun"}>
                    <motion.a
                        initial={{ scale: 0, borderRadius: "0" }}
                        animate={{ scale: 1, borderRadius: "8px" }}
                        whileHover={{ scale: 1.1, borderRadius: "16px", transition: { duration: 0.2, delay: 0 } }}
                        transition={{ duration: 0.3 }}
                        whileTap={{ scale: 0.8 }}
                        className="px-32 py-3 border-0 bg-red-700 text-white font-bold text-xl"
                        href="/games/slot/beachandsun">
                        Select
                    </motion.a>
                </Link>
            </motion.div>
        </Layout>
    );
};

export default Slot;
