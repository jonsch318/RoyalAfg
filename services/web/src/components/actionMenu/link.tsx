import React, { FC } from "react";
import Link from "next/link";

type ActionMenuLinkProps = {
    href: string;
};

const ActionMenuLink: FC<ActionMenuLinkProps> = ({ children, href }) => {
    return (
        <Link href={href}>
            <span className="bg-gray-200 mx-12 inline-flex w-auto px-16 py-16 rounded-xl cursor-pointer hover:bg-gray-300 outline-none hover:outline-none">
                {children}
            </span>
        </Link>
    );
};

export default ActionMenuLink;
