/* eslint-disable jsx-a11y/anchor-is-valid */
import { faBars, faTimes } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import React, { FC, useState } from "react";
import NavItems from "./nav";
import useWindowDimensions from "../../hooks/windowSize";

type HeaderProps = {
    absolute: boolean;
};

const Header: FC<HeaderProps> = ({ absolute }) => {
    const [isOpen, setIsOpen] = useState(false);
    const { width } = useWindowDimensions();

    const toggleMenu = () => {
        setIsOpen((val) => !val);
    };

    return (
        <header className="md:h-14 h-16 bg-blue-600 text-white w-full z-50" style={{ position: absolute ? "absolute" : "relative" }}>
            <div className="grid h-full z-50" style={{ gridTemplateColumns: "auto 1fr" }}>
                <button
                    className="no_highlights hamburger md:hidden grid h-full items-center ml-6 z-10 cursor-pointer focus:outline-none outline-none select-none w-6"
                    onClick={toggleMenu}>
                    <FontAwesomeIcon icon={isOpen ? faTimes : faBars} size="lg" />
                </button>
                <div className="logo md:ml-16 grid md:h-full h-16 w-full md:w-auto justify-center items-center md:relative absolute mt-0 mb-auto md:z-10 z-0">
                    <Link href="/">
                        <a className="font-medium font-sans text-xl text-center cursor-pointer">RoyalAfg</a>
                    </Link>
                </div>
                {width <= 768 && !isOpen ? (
                    <></>
                ) : (
                    <div className="nav md:ml-4 md:grid absolute md:relative w-full md:h-full z-50 md:z-10 bg-gray-200 md:bg-transparent text-black md:text-white mt-16 md:mt-0">
                        <NavItems />
                    </div>
                )}
            </div>
        </header>
    );
};

export default Header;
