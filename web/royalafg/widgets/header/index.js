import { faBars, faTimes } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useState } from "react";
import IdNav from "./idnav";
import NavItems from "./nav";


export default function Header() {
    const [isOpen, setIsOpen] = useState(false)

    const toggle = () => {
        setIsOpen(!isOpen);
    }
    const [width, setWidth] = React.useState(0);
    React.useEffect(() => {
        setWidth(window.innerWidth);
    });

    return (
        <div>
            <header className="md:h-10 h-16 bg-blue-600 text-white">
                <div className="flex h-full">
                    <button className="hamburger md:hidden flex h-full items-center ml-6 z-10 cursor-pointer focus:outline-none w-6" onClick={() => toggle()}>
                        <FontAwesomeIcon icon={isOpen ? faTimes : faBars} size="lg" />
                    </button>
                    <div className="logo md:ml-16 flex items-center md:h-full h-16 z-0 w-full md:w-auto justify-center md:relative absolute mt-0 mb-auto">
                        <a className="font-medium font-sans text-xl text-center cursor-pointer" href="/">Royalafg</a>
                    </div>
                    {
                        width <= 768 && !isOpen ? <></>
                            : <div className="nav md:ml-4 md:flex absolute md:relative w-full md:h-full z-50 md:z-10 bg-gray-200 md:bg-transparent text-black md:text-white mt-16 md:mt-0">
                                <NavItems></NavItems>
                            </div>
                    }

                </div>
            </header>
        </div>
    )
}