import React from "react";
import Layout from "../components/layout";
import CardListItem from "../widgets/cardList/cardItem";
import CardList from "../widgets/cardList/cardList";

const About = () => {
    return (
        <Layout>
            <div className="">
                <div className="bg-gray-300 pb-24">
                    <h1 className="md:px-10 py-24 font-sans text-5xl font-semibold text-center">
                        About this Project
                    </h1>

                    <CardList>
                        <CardListItem header="Contact">
                            <span className="block">
                                Email:{" "}
                                <a
                                    className="text-blue-700 hover:text-blue-800"
                                    href="jonas.max.schneider@gmail.com">
                                    jonas.max.schneider@gmail.com
                                </a>
                            </span>
                            <span className="block">Name: Jonas Schneider</span>
                            <span className="block">
                                Github:{" "}
                                <a href="github.com/JohnnyS318/RoyalAfgInGo">
                                    JohnnyS318/RoyalAfgInGo
                                </a>
                            </span>
                        </CardListItem>
                        <CardListItem header="Privacy">
                            <span className="block">
                                Privacy:
                                <a href="/privacy" className="text-blue-700 hover:text-blue-800">
                                    To the Privacy terms
                                </a>
                            </span>
                            <span className="block">
                                Terms of Use:{" "}
                                <a href="/terms" className="text-blue-700 hover:text-blue-800">
                                    Found here
                                </a>
                            </span>
                        </CardListItem>
                    </CardList>
                </div>

                <div className="my-10">
                    <h1 className="text-center md:text-4xl text-3xl md:p-12 p-4 pt-8 font-sans font-semibold">
                        This website and it&apos;s serverside environment was created out of a
                        special learning achievement
                    </h1>
                    <h2 className="text-center md:text-2xl text-2xl p-10 font-sans">
                        It has not been subjected to stability and security testing!{" "}
                        <span className="font-black ">Do Not Use In Production!</span>
                    </h2>
                </div>
            </div>
        </Layout>
    );
};

export default About;
