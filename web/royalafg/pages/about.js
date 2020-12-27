import React from "react";
import Layout from "../components/layout";

const About = () => {
    return (
        <Layout>
            <div>
                <div className="md:px-10 md:py-8 p-4">
                    <h1 className="font-sans text-4xl font-bold mb-4 text-center md:text-left">
                        About Royalafg
                    </h1>
                    <div className="cardList grid md:gap-20 gap-5 md:grid-cols-2 md:grid-rows-1">
                        <div className="card bg-gray-300 rounded-lg p-12 flex flex-col hover:opacity-75 transition-opacity duration-150">
                            <div className="content text-center md:text-xl py-8 flex-auto font-sans">
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
                            </div>
                            <h2 className="md:text-3xl text-xl font-medium text-center font-sans">
                                Contact
                            </h2>
                        </div>
                        <div className="card bg-gray-300 rounded-lg p-12 flex flex-col hover:opacity-75 transition-opacity duration-150">
                            <div className="content text-center text-xl py-8 flex-auto font-sans">
                                <span className="block">
                                    Privacy:
                                    <a
                                        href="/privacy"
                                        className="text-blue-700 hover:text-blue-800">
                                        To the Privacy terms
                                    </a>
                                </span>
                                <span className="block">
                                    Terms of Use:{" "}
                                    <a href="/terms" className="text-blue-700 hover:text-blue-800">
                                        Found here
                                    </a>
                                </span>
                            </div>
                            <h2 className="text-3xl font-medium text-center font-sans">Privacy</h2>
                        </div>
                    </div>
                    <h1 className="text-center md:text-5xl text-3xl md:p-12 p-4 pt-8 font-sans font-semibold">
                        This website and it&apos;s serverside environment was created out of a
                        special learning achievement
                    </h1>
                    <h2 className="text-center md:text-3xl text-2xl p-10 font-sans">
                        It has not been subjected to stability and security testing!{" "}
                        <span className="font-black ">Do Not Use In Production!</span>
                    </h2>
                </div>
            </div>
        </Layout>
    );
};

export default About;
