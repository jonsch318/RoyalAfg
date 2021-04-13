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
        <div>
            <header className="md:h-10 h-16 bg-blue-600 text-white w-full" style={{ position: absolute ? "absolute" : "relative" }}>
                <div className="flex h-full">
                    <button
                        className="hamburger md:hidden flex h-full items-center ml-6 z-10 cursor-pointer focus:outline-none w-6"
                        onClick={toggleMenu}>
                        <FontAwesomeIcon icon={isOpen ? faTimes : faBars} size="lg" />
                    </button>
                    <div className="logo md:ml-16 flex items-center md:h-full h-16 w-full md:w-auto justify-center md:relative absolute mt-0 mb-auto md:z-10 z-0">
                        <Link href="/">
                            <span className="font-medium font-sans text-xl text-center cursor-pointer">RoyalAfg</span>
                        </Link>
                    </div>
                    {width <= 768 && !isOpen ? (
                        <></>
                    ) : (
                        <div className="nav md:ml-4 md:flex absolute md:relative w-full md:h-full z-50 md:z-10 bg-gray-200 md:bg-transparent text-black md:text-white mt-16 md:mt-0">
                            <NavItems />
                        </div>
                    )}
                </div>
            </header>
        </div>
    );
};

export default Header;
