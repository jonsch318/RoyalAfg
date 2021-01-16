import React from "react";
import PropTypes from "prop-types";

const ActionMenu = ({ children }) => {
    return <div className="bg-white p-16 rounded-xl">{children}</div>;
};

ActionMenu.propTypes = {
    children: PropTypes.oneOfType([PropTypes.element, PropTypes.array])
};

export default ActionMenu;
