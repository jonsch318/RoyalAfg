import React, { FC } from "react";
import Layout from "../components/layout";
import { useIntl } from "react-intl";
import { useSession } from "../hooks/auth";
import Head from "next/head";
import Link from "next/link";
import Image from "next/image";

const Index: FC = () => {
    const [session] = useSession();
    const { formatMessage } = useIntl();
    const f = (id: string) => formatMessage({ id });
    return (
        <>
            <Head>
                <title>Royalafg Online Casino</title>
                <meta name="description" content="A online casino from a special lerning achivement"></meta>
            </Head>
            <Layout>
                <div>
                    <div className="header font-bold font-sans">
                        <h1 className="text-8xl text-center my-44">Royalafg</h1>
                        <div className="flex justify-center items-center mt-52 mb-64">
                            <Link href={session ? "/game" : "/auth/register"}>
                                <button className="my-auto px-24 py-2 bg-black text-white hover:bg-blue-600 transition-colors  rounded-md text-2xl">
                                    {session ? f("PlayAuthenticated") : f("PlayUnauthenticated")}
                                </button>
                            </Link>
                        </div>
                    </div>
                    <article>
                        <h1 className="font-sans text-4xl font-semibold text-center">{f("WelcomeHeader")}</h1>
                        <div className="grid grid-cols-2 px-32 my-12">
                            <Image src="/index/play.jpg" alt="Play" width="500" height="500" className="rounded-l-lg" />
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
