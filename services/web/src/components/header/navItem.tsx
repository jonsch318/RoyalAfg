/* eslint-disable jsx-a11y/anchor-is-valid */
import React, { FC } from "react";
import Link from "next/link";

type HeaderNavItemProps = {
    href: string;
};

const HeaderNavItem: FC<HeaderNavItemProps> = ({ href, children }) => {
    return (
        <Link href={href}>
            <a className="nav-item cursor-pointer w-auto block py-4 px-4 md:p-0 border-gray-300 border-b-2 border-solid md:border-none">{children}</a>
        </Link>
    );
};

export default HeaderNavItem;
