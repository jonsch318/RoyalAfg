import React from 'react';
import PropTypes from 'prop-types';

const CardListItem = ({ children, header }) => {
    return (
        <div className="w-full min-w-full bg-white rounded py-10 shadow-none hover:shadow-inner transition-shadow opacity-150 font-sans text-center">
            <h2 className="mb-8 font-semibold text-xl">{header}</h2>
            <div className="mb-8">{children}</div>
        </div>
    );
};

CardListItem.propTypes = {
    header: PropTypes.string,
    children: PropTypes.oneOf(PropTypes.string, PropTypes.element, PropTypes.array)
};

export default CardListItem;
