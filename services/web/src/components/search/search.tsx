import { faSearch } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
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
    const [focusedIndex, setFocusedIndex] = useState(-2);
    const ref = useRef<HTMLDivElement>();
    const inputRef = useRef<HTMLInputElement>(null);
    const inputDivRef = useRef<HTMLDivElement>(null);

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

    useEffect(() => {
        setInputWidth(inputDivRef.current?.getBoundingClientRect()?.width);
        //setInputWidth(inputDivRef.current ? inputDivRef.current?.offsetWidth : 0);
    }, [inputDivRef.current?.offsetWidth]);

    useEffect(() => {
        if (focusedIndex == -1) {
            inputRef.current?.focus();
        } else if (focusedIndex > -1) {
            inputRef.current?.blur();
        }
    }, [focusedIndex, inputRef.current]);

    useEffect(() => {
        const handler = (e: KeyboardEvent) => {
            if (focusedIndex >= -1 && e.key == "ArrowDown") {
                if (focusedIndex < results.length - 1) {
                    setFocusedIndex((x) => x + 1);
                }
                e.preventDefault();
            } else if (focusedIndex > -1 && e.key == "ArrowUp") {
                setFocusedIndex((x) => x - 1);
                e.preventDefault();
            }
        };

        document.addEventListener("keydown", handler);

        return () => {
            document.removeEventListener("keydown", handler);
        };
    }, [focusedIndex, results]);

    return (
        <div className="h-full px-4 font-sans">
            <div
                className="flex h-full bg-white rounded focus-within:rounded-b-none focus-within:border-2 border-black focus-within:border-b-0 justify-center items-center"
                ref={inputDivRef}>
                <span className="pl-3 pr-2 text-black" style={{ opacity: query ? "100%" : "20%" }}>
                    <FontAwesomeIcon icon={faSearch} />
                </span>

                <input
                    type="text"
                    autoComplete="off"
                    ref={inputRef}
                    className="relative md:py-0 py-2 w-full h-full text-black outline-none px-4 pl-0 bg-transparent "
                    id="global-search-input"
                    placeholder={t("Search")}
                    onChange={(e) => {
                        setQuery(e.target.value);
                        if (focusedIndex == -1 && e.target.value == "") {
                            setFocusedIndex(-2);
                        } else {
                            setFocusedIndex(-1);
                        }
                    }}
                    onFocus={() => {
                        setFocused(true);
                        setFocusedIndex(-1);
                    }}
                    onBlur={() => {
                        if (focusedIndex == -1) {
                            setFocusedIndex(-2);
                        }
                    }}
                />
            </div>
            {query && focused && (
                <div ref={ref} className="absolute bg-white rounded-b border-x-2 border-b-2 border-black">
                    <hr className=" bg-black h-4px" />
                    <div
                        className="md:popup"
                        style={{
                            width: inputWidth
                        }}>
                        <SearchResultList results={results} loading={loading} focused={focusedIndex} />
                    </div>
                </div>
            )}
        </div>
    );
};

export default SearchInput;
