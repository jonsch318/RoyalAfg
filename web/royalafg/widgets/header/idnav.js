import React from "react";
//import { useSelector } from "react-redux";
import Avatar from "../../components/header/id/avatar";
import { signIn, signOut, useSession } from "next-auth/client";
import PropTypes from "prop-types";

const NavButton = ({ children, onClick }) => {
    return (
        <button
            className="id-nav-item w-fit px-6 py-1 text break-normal flex mr-0 ml-auto my-0 bg-blue-800 rounded hover:bg-blue-900 md:mx-2 text-white transition-colors duration-150 "
            onClick={onClick}>
            {children}
        </button>
    );
};

NavButton.propTypes = {
    children: PropTypes.string,
    onClick: PropTypes.func
};

export default function IdNav() {
    //const isLoggedIn = useSelector((state) => state.auth.isLoggedIn);
    const [session] = useSession();

    if (!session) {
        return (
            <nav className="flex items-center h-full w-full">
                <div className="flex items-center h-full w-full px-4">
                    {/* <a
                        className="id-nav-item md:bg-transparent px-4 py-1 rounded bg-gray-300 md:hover:bg-blue-700 md:mx-2 transition-colors duration-150 flex"
                        href="/register">
                        Register
                    </a>
                    <a
                        className="id-nav-item bg-blue-800 px-6 py-1 rounded hover:bg-blue-900 md:mx-2 text-white transition-colors duration-150 flex mr-0 ml-auto"
                        href="/login">
                        Login
                    </a> */}
                    <NavButton onClick={signIn}>Login</NavButton>
                </div>
            </nav>
        );
    }

    return (
        <nav className="flex items-center h-full">
            <Avatar />
            <NavButton onClick={signOut}>Logout</NavButton>
        </nav>
    );
}
