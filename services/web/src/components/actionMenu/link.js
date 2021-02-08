import React from "react";
import PropTypes from "prop-types";
import Link from "next/link";

const ActionMenuLink = ({ children, href, locale, key }) => {
    return (
        <Link href={href} locale={locale}>
            <span
                key={key}
                className="bg-gray-200 mx-12 inline-flex w-auto px-16 py-16 rounded-xl cursor-pointer hover:bg-gray-300 outline-none hover:outline-none">
                {children}
            </span>
        </Link>
    );
};

ActionMenuLink.propTypes = {
    children: PropTypes.string,
    href: PropTypes.string,
    locale: PropTypes.string,
    key: PropTypes.object
};

export default ActionMenuLink;
