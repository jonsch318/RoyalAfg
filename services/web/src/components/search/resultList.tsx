import React, { FC } from "react";
import Result from "./result";
import { AnimatePresence, motion } from "framer-motion";
import { SearchResult } from "./search";
import { useTranslation } from "next-i18next";

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
    focused: number;
};

const SearchResultList: FC<SearchResultListProps> = ({ results, loading, focused }) => {
    const { t } = useTranslation("common");
    return (
        <AnimatePresence>
            <motion.ul className="py-4 md:py-2 md: px-4 bg-blue ">
                {!loading &&
                    results.map((result, i) => (
                        <motion.li variants={items} animate="animate" initial="initial" key={result.name} className="px-5">
                            <Result result={result} focused={focused == i} />
                        </motion.li>
                    ))}
                {results.length < 1 && (
                    <motion.li variants={items} animate="animate" initial="initial" className="px-5 text-black">
                        {loading ? t("Searching...") : t("Sorry nothing was found.")}
                    </motion.li>
                )}
            </motion.ul>
        </AnimatePresence>
    );
};

export default SearchResultList;
