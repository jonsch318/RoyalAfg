/* eslint-disable jsx-a11y/anchor-is-valid */
import React, { FC } from "react";
import Link from "next/link";

const Avatar: FC = () => {
    return (
        <Link href="/account">
            <a className="flex cursor-pointer align-middle justify-center items-center break-normal px-2 w-32 flex-no-wrap hover:opacity-75 transition-opacity duration-200">
                My Account
            </a>
        </Link>
    );
};

export default Avatar;
