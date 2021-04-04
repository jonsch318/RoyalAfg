import React, { FC } from "react";
import Avatar from "../../components/header/id/avatar";
import Link from "next/link";
import { signOut, useSession } from "../../hooks/auth";

type NavButtonProps = {
    onClick: React.MouseEventHandler<HTMLButtonElement>;
};
const NavButton: FC<NavButtonProps> = ({ children, onClick }) => {
    return (
        <button
            className="id-nav-item w-fit px-6 py-1 text break-normal flex mr-0 ml-auto my-0 bg-blue-800 rounded hover:bg-blue-900 md:mx-2 text-white transition-colors duration-150 "
            onClick={onClick}>
            {children}
        </button>
    );
};

const IdNav: FC = () => {
    const [session] = useSession();
    if (!session) {
        return (
            <nav className="flex items-center h-full w-full">
                <div className="flex items-center h-full w-full px-4">
                    <Link href="/auth/register" passHref>
                        <a
                            href="/auth/register"
                            className="id-nav-item md:bg-transparent px-4 py-1 rounded bg-gray-300 md:hover:bg-blue-700 md:mx-2 transition-colors duration-150 flex">
                            Register
                        </a>
                    </Link>
                    <Link href="/auth/login" passHref>
                        <a
                            href="/auth/login"
                            className="id-nav-item bg-blue-800 px-6 py-1 rounded hover:bg-blue-900 md:mx-2 text-white transition-colors duration-150 flex mr-0 ml-auto">
                            Login
                        </a>
                    </Link>
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
};

export default IdNav;