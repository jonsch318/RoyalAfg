import React from "react";
import PropTypes from "prop-types";

const FormItem = (props) => {
    return <div className="mb-8 font-sans text-lg font-medium">{props.children}</div>;
};

FormItem.propTypes = {
    children: PropTypes.element
};

export default FormItem;
