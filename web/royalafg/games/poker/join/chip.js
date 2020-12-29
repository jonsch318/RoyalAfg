import React from "react";
import PropTypes from "prop-types";

const Chip = ({ min, max, smallBlind, selected, onSelect }) => {
    return (
        <button
            style={{ background: selected ? "#ecc94b" : "#3182ce" }}
            className="bg-blue-600 text-white px-2 py-3 mx-4 rounded cursor-pointer hover:shadow-xl transition-shadow duration-200 ease-in-out "
            onMouseUp={onSelect}>
            <span>{`${min} - ${max}`}</span>
            <span className="drop">{smallBlind}</span>
        </button>
    );
};

Chip.propTypes = {
    min: PropTypes.number,
    max: PropTypes.number,
    smallBlind: PropTypes.number,
    onSelect: PropTypes.func,
    selected: PropTypes.bool
};

export default Chip;
