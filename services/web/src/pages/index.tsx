import React, { FC } from "react";
import Layout from "../components/layout";
import { useIntl } from "react-intl";
import { useSession } from "../hooks/auth";
import Head from "next/head";
import Link from "next/link";
import Image from "next/image";
import { motion } from "framer-motion";

const Index: FC = () => {
    const [session] = useSession();
    const { formatMessage } = useIntl();
    const f = (id: string) => formatMessage({ id });

    return (
        <>
            <Head>
                <title>Royalafg Online Casino</title>
                <meta name="description" content="A online casino from a special learning achievement" />
            </Head>
            <Layout>
                <div>
                    <div className="header font-bold font-sans">
                        <div>
                            <motion.div
                                initial="in"
                                animate={{
                                    y: 100,
                                    transition: {
                                        type: "spring",
                                        bounce: 0.45,
                                        delay: 0.1
                                    }
                                }}>
                                <h1 className="md:text-8xl text-7xl text-center z-50 md:my-44 mt-16 mb-24">Royalafg</h1>
                            </motion.div>
                        </div>

                        <div className="flex justify-center items-center mt-52 md:mb-64 mb-44">
                            <Link href={session ? "/game" : "/auth/register"}>
                                <motion.button
                                    whileHover={{ scale: 1.15, transition: { duration: 0.2, ease: "easeInOut" } }}
                                    initial={{ y: 50 }}
                                    animate={{ y: 0, transition: { duration: 0.4, ease: "backIn" } }}
                                    className="my-auto px-24 py-2 bg-black text-white hover:bg-blue-600 transition-colors  rounded-md md:text-2xl text-lg">
                                    <span>{session ? f("PlayAuthenticated") : f("PlayUnauthenticated")}</span>
                                </motion.button>
                            </Link>
                        </div>
                    </div>
                    <article>
                        <h1 className="font-sans text-4xl font-semibold text-center">{f("WelcomeHeader")}</h1>
                        <div className="grid md:grid-cols-2 grid-cols-1 md:px-32 px-8 my-12">
                            <Image src="/index/play.jpg" alt="Play" className="rounded-l-lg" width="500px" height="400px" />
                            <div className="p-12 bg-gray-200 rounded-r-lg">
                                <p className="text-center">{f("WelcomeArticle")}</p>
                            </div>
                        </div>
                    </article>
                </div>
            </Layout>
        </>
    );
};

export default Index;
