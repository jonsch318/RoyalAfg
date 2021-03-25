import React from "react";
import PropTypes from "prop-types";
import Dinero from "dinero.js";

const Chip = ({ lobbyClass, selected, onSelect }) => {
    return (
        <button
            style={{ background: selected ? "#f59e0b" : "#3182ce" }}
            className="bg-blue-600 text-white px-2 py-3 mx-4 rounded cursor-pointer hover:shadow-xl transition-shadow duration-200 ease-in-out "
            onMouseUp={onSelect}>
            <span>{`${Dinero({ amount: lobbyClass.min, currency: "USD" }).toFormat()} - ${Dinero({
                amount: lobbyClass.max,
                currency: "USD"
            }).toFormat()}`}</span>
        </button>
    );
};

Chip.propTypes = {
    lobbyClass: PropTypes.shape({
        min: PropTypes.number,
        max: PropTypes.number,
        blind: PropTypes.number
    }),
    onSelect: PropTypes.func,
    selected: PropTypes.bool
};

export default Chip;
