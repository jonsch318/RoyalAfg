import React, { FC, useEffect, useState } from "react";
import CurrencyInput from "react-currency-input-field";
import Tooltip from "@material-ui/core/Tooltip";
import { useRouter } from "next/router";

type RaiseProps = {
    onRaise: (amount: number) => void;
};

const Raise: FC<RaiseProps> = ({ onRaise }) => {
    const { locale } = useRouter();
    const [raise, setRaise] = useState(0.0);
    useEffect(() => {
        console.log("Raise: ", raise);
    }, [raise]);
    return (
        <div className="flex justify-center items-center rounded mx-4 h-full">
            <CurrencyInput
                name="amount"
                className="border-blue-600 flex justify-center items-center h-full px-3 outline-none py-1 flex w-80 font-sans placeholder-blue-300 text-black border-solid ml-4 mr-2 rounded"
                placeholder="Raise Amount"
                intlConfig={{ locale: locale, currency: "USD" }}
                autoComplete="off"
                defaultValue={0.0}
                onValueChange={(val: string | undefined) => {
                    if (val !== undefined) setRaise(parseFloat(val));
                }}
                allowNegativeValue={false}
            />
            <Tooltip title={!(raise > 0) ? "Specify a amount to raise" : "Raise"} aria-label="Raise">
                <button
                    className="bg-white px-3 flex justify-center items-center h-full rounded text-black overflow-hidden hover:bg-yellow-500 transition-colors ease-in-out duration-150 disabled:opacity-60"
                    onClick={() => {
                        if (raise > 0) onRaise(raise * 100);
                    }}>
                    RAISE
                </button>
            </Tooltip>
        </div>
    );
};

export default Raise;
