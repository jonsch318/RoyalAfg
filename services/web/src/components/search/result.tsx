/* eslint-disable jsx-a11y/anchor-is-valid */
import React, { FC, useEffect, useRef } from "react";
import Link from "next/link";
import { SearchResult } from "./search";

type SearchResultProps = {
    result: SearchResult;
    focused: boolean;
};

const Result: FC<SearchResultProps> = ({ result, focused }) => {
    const ref = useRef<HTMLAnchorElement>(null);

    useEffect(() => {
        if (focused) {
            ref.current?.focus();
        }
    }, [focused, ref.current]);

    return (
        <Link href={result.url}>
            <a
                className="flex md:text-black text-white bg-gray-200 md:bg-gray-300 w-full z-50 mb-1 md:my-2 py-1 px-2 rounded hover:bg-gray-300 md:hover:bg-gray-100"
                style={{ backgroundColor: focused ? "rgba(245, 245, 245, var(--tw-bg-opacity))" : "" }}
                ref={ref}>
                {result.name}
            </a>
        </Link>
    );
};

export default Result;
