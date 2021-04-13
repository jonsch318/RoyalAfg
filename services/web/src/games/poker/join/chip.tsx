import React, { FC } from "react";
import Dinero from "dinero.js";

type ChipProps = {
    lobbyClass: {
        min: number;
        max: number;
        blind: number;
    };
    onSelect: (e: any) => void;
    selected: boolean;
};

const Chip: FC<ChipProps> = ({ lobbyClass, selected, onSelect }) => {
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

export default Chip;
