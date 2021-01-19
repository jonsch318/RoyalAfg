import React from "react";
import Link from "next/link";

const Avatar = () => {
    return (
        <Link href="/account" passHref>
            <a
                href="/account"
                className="flex align-middle justify-center items-center break-normal px-2 w-32 flex-no-wrap hover:opacity-75 transition-opacity duration-200">
                My Account
            </a>
        </Link>
    );
};

export default Avatar;
