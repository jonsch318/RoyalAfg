/* eslint-disable jsx-a11y/anchor-is-valid */
import Link from "next/link";
import React, { FC } from "react";

type FooterCardItemProps = {
    href?: string;
};

const FooterCardItem: FC<FooterCardItemProps> = ({ href, children }) => {
    if (href) {
        return (
            <Link href={href}>
                <a className="font-sans font-thin text-sm hover:opacity-75 transition-opacity duration-100 ease-out">{children}</a>
            </Link>
        );
    }
    return <span className="font-sans font-thin text-sm">{children}</span>;
};

export default FooterCardItem;
