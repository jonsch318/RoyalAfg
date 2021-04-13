import React, { FC } from "react";
import Link from "next/link";

const Avatar: FC = () => {
    return (
        <Link href="/account">
            <span className="flex align-middle justify-center items-center break-normal px-2 w-32 flex-no-wrap hover:opacity-75 transition-opacity duration-200">
                My Account
            </span>
        </Link>
    );
};

export default Avatar;
