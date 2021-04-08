/* eslint-disable jsx-a11y/anchor-is-valid */
import React, { FC, useEffect } from "react";
import Link from "next/link";
import { SearchResult } from "./search";

type SearchResultProps = {
    result: SearchResult;
};

const Result: FC<SearchResultProps> = ({ result }) => {
    useEffect(() => {
        console.log("URL of search res: ", result.name, " is: ", result.url);
    }, [result]);
    return (
        <Link href={result.url}>
            <a className="flex md:text-black text-white bg-gray-200 md:bg-gray-300 w-full z-50 mb-1 md:my-2 py-1 px-2 rounded hover:bg-gray-300 md:hover:bg-gray-400">
                {result.name}
            </a>
        </Link>
    );
};

export default Result;
