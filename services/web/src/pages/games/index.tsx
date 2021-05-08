import React, { FC } from "react";
import Layout from "../../components/layout";
import Link from "next/link";
import Head from "next/head";
import { GetStaticProps } from "next";
import { formatTitle } from "../../utils/title";
import { serverSideTranslations } from "next-i18next/serverSideTranslations";
import { useTranslation } from "next-i18next";

type GameProps = {
    href: string;
};

// eslint-disable-next-line react/prop-types
const Game: FC<GameProps> = ({ href, children }) => {
    return (
        <Link href={href}>
            <span className="bg-gray-200 grid justify-center items-center md:m-12 mx-8 my-2 w-auto md:p-16 p-8 rounded-xl cursor-pointer hover:bg-gray-300 outline-none hover:outline-none">
                {children}
            </span>
        </Link>
    );
};

export interface IGame {
    name: string;
    id: string;
}

type GamesProps = {
    games: IGame[];
};

const Games: FC<GamesProps> = ({ games }) => {
    const { t } = useTranslation("games");

    return (
        <Layout>
            <Head>
                <title>{formatTitle(t("Title"))}</title>
            </Head>
            <div className="font-sans bg-gray-200 grid justify-center items-center">
                <div className="px-auto py-20 flex">
                    <h1 className="text-6xl mx-auto font-bold font-sans text-center bg-white px-12 py-8 rounded inline w-auto">
                        {t("Game Selection")}
                    </h1>
                </div>
                <div className="md:p-10 p-5 ">
                    <div className="md:grid-cols-4 sm:grid-cols-3 grid-cols-1 grid md:gap-1 gap-0 gap-y-0 bg-white lg:px-10 py-10 rounded-lg">
                        {games.map((game) => (
                            <Game href={"/games/" + game.id} key={game.id}>
                                {game.name}
                            </Game>
                        ))}
                    </div>
                </div>
            </div>
        </Layout>
    );
};

export const getStaticProps: GetStaticProps = async (ctx) => {
    return {
        props: {
            games: [
                { name: "Poker Texas Holdem", id: "poker" },
                { name: "Pacman", id: "pacman" },
                { name: "Slot Machine", id: "slot" }
            ],
            ...(await serverSideTranslations(ctx.locale, ["common", "games"]))
        }
    };
};
export default Games;
