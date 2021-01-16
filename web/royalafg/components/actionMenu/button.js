import React from "react";
import PropTypes from "prop-types";

const ActionMenuButton = ({ children, onClick }) => {
    return (
        <button className="bg-gray-200 px-16 py-16 rounded-xl hover:bg-gray-300 outline-none hover:outline-none" onClick={onClick}>
            {children}
        </button>
    );
};

ActionMenuButton.propTypes = {
    children: PropTypes.string,
    onClick: PropTypes.func
};

export default ActionMenuButton;
