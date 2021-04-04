import React, { FC } from "react";
import PropTypes from "prop-types";

type CardListItemProps = {
    header: string;
};

const CardListItem: FC<CardListItemProps> = ({ children, header }) => {
    return (
        <div className="w-full min-w-full bg-white rounded py-10 shadow-none hover:shadow-inner transition-shadow opacity-150 font-sans text-center">
            <h2 className="mb-8 font-semibold text-xl">{header}</h2>
            <div className="mb-8">{children}</div>
        </div>
    );
};

export default CardListItem;
