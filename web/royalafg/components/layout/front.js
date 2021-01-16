import React from "react";
import PropTypes from "prop-types";

const Front = ({ children }) => {
    return <div className="bg-gray-200 md:px-10 py-28 font-sans text-5xl font-semibold text-center">{children}</div>;
};

Front.propTypes = {
    children: PropTypes.oneOfType([PropTypes.element, PropTypes.string, PropTypes.array])
};

export default Front;
