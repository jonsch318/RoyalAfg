import React from "react";
import Result from "./result";
import PropTypes from "prop-types";
import { AnimatePresence, motion } from "framer-motion";

const container = {};

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

const SearchResultList = ({ results }) => {
    return (
        <AnimatePresence>
            <motion.ul className="py-4 md:py-2 md: px-4 bg-blue ">
                {results.map((result) => (
                    <motion.li variants={items} animate="animate" initial="initial" key={result.name} className="px-5">
                        <Result result={result} />
                    </motion.li>
                ))}
                {results.length < 1 && (
                    <motion.li variants={items} animate="animate" initial="initial" className="px-5 text-black">
                        Sorry nothing was found.
                    </motion.li>
                )}
            </motion.ul>
        </AnimatePresence>
    );
};

SearchResultList.propTypes = {
    results: PropTypes.array
};

export default SearchResultList;
