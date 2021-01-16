import React from "react";
import Link from "next/link";

const Avatar = () => {
    return (
        <div className="flex align-middle justify-center items-center break-normal px-2 w-32 flex-no-wrap hover:opacity-75 transition-opacity duration-200">
            <Link href="/account">My Account</Link>
        </div>
    );
};

export default Avatar;
