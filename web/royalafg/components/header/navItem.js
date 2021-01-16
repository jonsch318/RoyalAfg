import React from "react";
import PropTypes from "prop-types";

const HeaderNavItem = (props) => {
    return (
        <a
            className="nav-item block py-4 px-4 md:p-0 border-gray-300 border-b-2 border-solid md:border-none"
            href={props.href}>
            {props.children}
        </a>
    );
};

HeaderNavItem.propTypes = {
    href: PropTypes.string,
    children: PropTypes.oneOfType([PropTypes.string, PropTypes.element, PropTypes.array])
};

export default HeaderNavItem;
