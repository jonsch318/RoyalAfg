import { useTranslation } from "next-i18next";
import React, { FC, useEffect, useRef, useState } from "react";

import useDebounce from "../../hooks/debounce";

import SearchResultList from "./resultList";

function useOnClickOutside(ref, handler: (e: MouseEvent) => void) {
    useEffect(() => {
        const listener = (event: MouseEvent) => {
            if (!ref.current || ref.current.contains(event.target)) {
                return;
            }
            handler(event);
        };
        document.addEventListener("mousedown", listener);
        document.addEventListener("touchstart", listener);
        return () => {
            document.removeEventListener("mousedown", listener);
            document.removeEventListener("touchstart", listener);
        };
    }, [ref, handler]);
}

export interface SearchResult {
    name: string;
    url: string;
    maxPlayers: number;
}

const Search = async (query: string): Promise<SearchResult[]> => {
    const queryString = `q=${query}`;

    try {
        const res = await fetch(`/api/search?${queryString}`);
        if (res.ok) {
            return await res.json();
        } else {
            return [];
        }
    } catch (e) {
        console.log("error during search: ", e);
        return [];
    }
};

const SearchInput: FC = () => {
    const { t } = useTranslation("common");
    const [query, setQuery] = useState("");
    const [results, setResults] = useState([]);
    const [loading, setLoading] = useState(false);
    const [inputWidth, setInputWidth] = useState(0);
    const [focused, setFocused] = useState(false);
    const ref = useRef();

    const debouncedQuery = useDebounce(query, 150);
    useOnClickOutside(ref, () => setFocused(false));

    useEffect(() => {
        const f = async () => {
            if (debouncedQuery) {
                const res = await Search(query);
                setResults(res);
            } else {
                setResults([]);
            }
            setLoading(false);
        };
        f();
    }, [debouncedQuery]);

    useEffect(() => {
        setLoading(true);
    }, [query]);

    const inputRef = useRef(null);
    useEffect(() => {
        setInputWidth(inputRef.current ? inputRef.current?.offsetWidth : 0);
    }, [inputRef.current?.offsetWidth]);

    return (
        <div className="relative h-full mx-4">
            <input
                type="text"
                autoComplete="off"
                ref={inputRef}
                className="relative font-sans md:py-0 py-2 bg-white w-full h-full text-black outline-none px-4 rounded"
                id="global-search-input"
                placeholder={t("Search")}
                onChange={(e) => setQuery(e.target.value)}
                onFocus={() => setFocused(true)}
            />
            {query && focused && (
                <div ref={ref}>
                    <hr className="md:hidden bg-black h-1px opacity-50" />
                    <div
                        className="md:popup"
                        style={{
                            width: inputWidth
                        }}>
                        <SearchResultList results={results} loading={loading} />
                    </div>
                </div>
            )}
        </div>
    );
};

export default SearchInput;
