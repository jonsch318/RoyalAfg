import React, { FC } from "react";
import Layout from "../components/layout";
import CardListItem from "../widgets/cardList/cardItem";
import CardList from "../widgets/cardList/cardList";
import Head from "next/head";
import Link from "next/link";
import { formatTitle } from "../utils/title";
import { useTranslation } from "next-i18next";
import { GetStaticProps } from "next";
import { serverSideTranslations } from "next-i18next/serverSideTranslations";

export const getStaticProps: GetStaticProps = async ({ locale }) => ({
    props: {
        ...(await serverSideTranslations(locale, ["about", "common"]))
    }
});

const About: FC = () => {
    const { t } = useTranslation("about");

    return (
        <>
            <Head>
                <title>{formatTitle(t("Title"))}</title>
                <meta name="description" content={t("Description")} />
            </Head>
            <Layout>
                <div>
                    <div className="bg-gray-200 pb-24">
                        <h1 className="md:px-10 py-24 font-sans text-5xl font-semibold text-center">{t("About this project")}</h1>

                        <CardList>
                            <CardListItem header={t("Contact")}>
                                <span className="block">
                                    {t("Email")}:{" "}
                                    <a className="text-blue-700 hover:text-blue-800" href="jonas.max.schneider@gmail.com">
                                        jonas.max.schneider@gmail.com
                                    </a>
                                </span>
                                <span className="block">{t("Name")}: Jonas Schneider</span>
                                <span className="block">
                                    {t("Github")}: <a href="github.com/jonsch318/RoyalAfg">JohnnyS318/RoyalAfg</a>
                                </span>
                            </CardListItem>
                            <CardListItem header={t("Privacy")}>
                                <span className="block">
                                    {t("Privacy")}: <Link href="/legal/privacy">{t("To the privacy terms")}</Link>
                                </span>
                                <span className="block">
                                    {t("Terms of use")}: <Link href="/legal/terms">{t("Found here")}</Link>
                                </span>
                            </CardListItem>
                        </CardList>
                    </div>

                    <div className="my-10">
                        <h1 className="text-center md:text-4xl text-3xl md:p-12 p-4 pt-8 font-sans font-semibold">{t("Origin")}</h1>
                        <h2 className="text-center md:text-2xl text-2xl p-10 font-sans">
                            {t("Warning")} <span className="font-black ">{t("ProductionWarning")}</span>
                        </h2>
                    </div>
                </div>
            </Layout>
        </>
    );
};

export default About;
