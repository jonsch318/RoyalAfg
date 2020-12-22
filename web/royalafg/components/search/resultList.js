import React from "react";
import Result from "./result";

const SearchResultList = (props) => {
  return (
    <ul className="py-4 md:py-0 bg-white md:bg-transparent">
      {props.results.map((result) => (
        <li key={result.name} className="px-5">
          <Result result={result} />
        </li>
      ))}
    </ul>
  );
};

export default SearchResultList;
