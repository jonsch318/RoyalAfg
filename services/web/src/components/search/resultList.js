import React from "react";
import Result from "./result";
import PropTypes from "prop-types";

const SearchResultList = (props) => {
    return (
        <ul className="py-4 md:py-2 md: px-4 bg-white ">
            {props.results.map((result) => (
                <li key={result.name} className="px-5">
                    <Result result={result} />
                </li>
            ))}
            {props.results.length < 1 && <li className="text-black">Sorry nothing was found.</li>}
        </ul>
    );
};

SearchResultList.propTypes = {
    results: PropTypes.array
};

export default SearchResultList;
