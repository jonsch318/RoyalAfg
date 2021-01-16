import React from "react";
import PropTypes from "prop-types";
import Link from "next/link";

const ActionMenuLink = ({ children, href, locale }) => {
    return (
        <Link href={href} locale={locale}>
            <span className="bg-gray-200 inline-flex w-auto px-16 py-16 rounded-xl cursor-pointer hover:bg-gray-300 outline-none hover:outline-none">
                {children}
            </span>
        </Link>
    );
};

ActionMenuLink.propTypes = {
    children: PropTypes.string,
    href: PropTypes.string,
    locale: PropTypes.string
};

export default ActionMenuLink;
