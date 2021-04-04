import React, { FC } from "react";
import Result from "./result";
import { AnimatePresence, motion } from "framer-motion";
import { SearchResult } from "./search";

const items = {
    initial: {
        y: 10,
        opacity: 0
    },
    animate: {
        y: 0,
        opacity: 1
    }
};

type SearchResultListProps = {
    results: SearchResult[];
    loading: boolean;
};

const SearchResultList: FC<SearchResultListProps> = ({ results, loading }) => {
    return (
        <AnimatePresence>
            <motion.ul className="py-4 md:py-2 md: px-4 bg-blue ">
                {!loading &&
                    results.map((result) => (
                        <motion.li variants={items} animate="animate" initial="initial" key={result.name} className="px-5">
                            <Result result={result} />
                        </motion.li>
                    ))}
                {!loading && results.length < 1 && (
                    <motion.li variants={items} animate="animate" initial="initial" className="px-5 text-black">
                        Sorry nothing was found.
                    </motion.li>
                )}
            </motion.ul>
        </AnimatePresence>
    );
};

export default SearchResultList;
