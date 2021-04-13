import React, { FC, useEffect, useRef } from "react";
import PropType from "prop-types";

const CardList: FC = ({ children }) => {
    const ref = useRef(null);
    useEffect(() => {
        console.log("Ref Current", ref.current);
        console.log("width", ref.current ? ref.current.offsetWidth : 0);
    }, [ref.current]);

    return (
        <div ref={ref} className="grid grid-cols-2 gap-24 px-24">
            {children}
        </div>
    );
};

CardList.propTypes = {
    children: PropType.arrayOf(PropType.element)
};

export default CardList;
