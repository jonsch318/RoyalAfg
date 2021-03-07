import React from "react";
import Layout from "../components/layout";
import CardListItem from "../widgets/cardList/cardItem";
import CardList from "../widgets/cardList/cardList";
import { useIntl } from "react-intl";
import Head from "next/head";
import { formatTitle } from "../utils/title.ts";
import Link from "next/link";

const About = () => {
    const { formatMessage } = useIntl();
    const f = (id) => formatMessage({ id });

    return (
        <>
            <Head>
                <title>{formatTitle("About Royalafg")}</title>
                <meta name="description" content="About the Royalafg Online Casino" />
            </Head>
            <Layout>
                <div className="">
                    <div className="bg-gray-200 pb-24">
                        <h1 className="md:px-10 py-24 font-sans text-5xl font-semibold text-center">{f("header")}</h1>

                        <CardList>
                            <CardListItem header={f("contactHeader")}>
                                <span className="block">
                                    {f("emailLabel")}:{" "}
                                    <a className="text-blue-700 hover:text-blue-800" href="jonas.max.schneider@gmail.com">
                                        jonas.max.schneider@gmail.com
                                    </a>
                                </span>
                                <span className="block">{f("nameLabel")}: Jonas Schneider</span>
                                <span className="block">
                                    {f("githubLabel")}: <a href="github.com/JohnnyS318/RoyalAfgInGo">JohnnyS318/RoyalAfgInGo</a>
                                </span>
                            </CardListItem>
                            <CardListItem header={f("privacyHeader")}>
                                <span className="block">
                                    {f("privacyLabel")}: <Link href="/legal/privacy">To the Privacy terms</Link>
                                </span>
                                <span className="block">
                                    {f("termsLabel")}: <Link href="/legal/terms">Found Here</Link>
                                </span>
                            </CardListItem>
                        </CardList>
                    </div>

                    <div className="my-10">
                        <h1 className="text-center md:text-4xl text-3xl md:p-12 p-4 pt-8 font-sans font-semibold">{f("origin")}</h1>
                        <h2 className="text-center md:text-2xl text-2xl p-10 font-sans">
                            {f("testWarning")} <span className="font-black ">{f("productionWarning")}</span>
                        </h2>
                    </div>
                </div>
            </Layout>
        </>
    );
};

export default About;
