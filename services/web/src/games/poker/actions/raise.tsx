import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import CurrencyInput from "react-currency-input-field";
import { useRouter } from "next/router";
import Tooltip from "@material-ui/core/Tooltip";

const Raise = ({ onRaise }) => {
    const { locale } = useRouter();
    const [raise, setRaise] = useState(0.0);
    useEffect(() => {
        console.log("Raise: ", raise);
    }, [raise]);
    return (
        <div className="flex justify-center items-center rounded mx-4">
            <CurrencyInput
                name="amount"
                className="border-blue-600 px-3 h-full outline-none py-1 flex w-80 font-sans placeholder-blue-300 text-black border-solid mx-4 rounded"
                placeholder="Raise Amount"
                intlConfig={{ locale: locale, currency: "USD" }}
                autoComplete="off"
                defaultValue={0.0}
                onValueChange={(val) => {
                    setRaise(parseFloat(val));
                }}
                allowNegativeValue={false}
            />
            <Tooltip title={!(raise > 0) ? "Specify a amount to raise" : "Raise"} aria-label="Raise">
                <button
                    className="bg-white px-3 h-full rounded text-black overflow-hidden hover:bg-yellow-500 transition-colors ease-in-out duration-150 disabled:opacity-60"
                    onClick={() => {
                        if (raise > 0) onRaise({ action: 3, payload: raise * 100 });
                    }}>
                    RAISE
                </button>
            </Tooltip>
        </div>
    );
};

Raise.propTypes = {
    onRaise: PropTypes.func
};

export default Raise;
