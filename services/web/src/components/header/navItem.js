import React from "react";
import Link from "next/link";
import PropTypes from "prop-types";

const HeaderNavItem = (props) => {
    return (
        <Link href={props.href}>
            <a className="nav-item block py-4 px-4 md:p-0 border-gray-300 border-b-2 border-solid md:border-none">{props.children}</a>
        </Link>
    );
};

HeaderNavItem.propTypes = {
    href: PropTypes.string,
    children: PropTypes.oneOfType([PropTypes.string, PropTypes.element, PropTypes.array])
};

export default HeaderNavItem;
