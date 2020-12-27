import React from "react";
import Result from "./result";
import PropTypes from "prop-types";

const SearchResultList = (props) => {
    return (
        <ul className="py-4 md:py-0 bg-white md:bg-transparent">
            {props.results.map((result) => (
                <li key={result.name} className="px-5">
                    <Result result={result} />
                </li>
            ))}
            {props.results.length < 1 && (
                <span className="text-black">Sorry nothing was found.</span>
            )}
        </ul>
    );
};

SearchResultList.propTypes = {
    results: PropTypes.array
};

export default SearchResultList;
