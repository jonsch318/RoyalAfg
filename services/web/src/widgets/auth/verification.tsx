/* eslint-disable jsx-a11y/no-autofocus */
import React, { FC, useEffect, useState } from "react";
import { RegisterDto } from "../../pages/auth/register";
import useDigitInput from "react-digit-input";
import CircularProgress from "@mui/material/CircularProgress";
import { verifyUser } from "../../hooks/auth";
import { useDots } from "../../hooks/dots";

type VerficationProps = {
    handleNext: () => void;
    dto: RegisterDto;
    setDet: React.Dispatch<React.SetStateAction<RegisterDto>>;
    csrf: string;
};

const regex = new RegExp(/^[0-9]{6}$/);
const numVerify = (num: number): boolean => {
    return regex.test(num.toString());
};

const Verification: FC<VerficationProps> = ({ handleNext, dto, setDet, csrf }) => {
    const [val, setVal] = useState("");
    const [num, setNum] = useState(0);
    const [verified, setVerified] = useState(false);
    const dots = useDots();

    const digits = useDigitInput({
        acceptedCharacters: /^[0-9]$/,
        length: 6,
        value: val,
        onChange: setVal
    });

    useEffect(() => {
        setNum(parseInt(val));
    }, [val]);

    useEffect(() => {
        if (numVerify(num)) {
            console.log("Verify: ", num);

            setTimeout(() => {
                setVerified(true);
            }, 2000);
            verifyUser(num, csrf);
        }
    }, [num]);

    return (
        <div className="mx-16 flex flex-col justify-center items-center">
            {!numVerify(num) && !verified && (
                <div className="flex flex-col justify-center items-center my-8">
                    <h1 className="text-center font-sans text-2xl font-semibold py-2">We have send you an email!</h1>
                    <p className="text-center font-sans text-lg font-normal">Please check your inbox and enter the verification number.</p>
                    <div className="input-group flex flex-row items-center justify-center my-8">
                        <input
                            inputMode="decimal"
                            autoFocus
                            {...digits[0]}
                            className="w-12 h-20 font-sans font-semibold text-7xl bg-transparent border-black border-b-4 pt-4 text-center outline-none rounded-sm mx-2"
                        />
                        <input
                            inputMode="decimal"
                            {...digits[1]}
                            className="w-12 h-20 font-sans font-semibold text-7xl bg-transparent border-black border-b-4 pt-4 text-center outline-none rounded-sm mx-2"
                        />
                        <input
                            inputMode="decimal"
                            {...digits[2]}
                            className="w-12 h-20 font-sans font-semibold text-7xl bg-transparent border-black border-b-4 pt-4 text-center outline-none rounded-sm mx-2"
                        />
                        <input
                            inputMode="decimal"
                            {...digits[3]}
                            className="w-12 h-20 font-sans font-semibold text-7xl bg-transparent border-black border-b-4 pt-4 text-center outline-none rounded-sm mx-2"
                        />
                        <input
                            inputMode="decimal"
                            {...digits[4]}
                            className="w-12 h-20 font-sans font-semibold text-7xl bg-transparent border-black border-b-4 pt-4 text-center outline-none rounded-sm mx-2"
                        />
                        <input
                            inputMode="decimal"
                            {...digits[5]}
                            className="w-12 h-20 font-sans font-semibold text-7xl bg-transparent border-black border-b-4 pt-4 text-center outline-none rounded-sm mx-2"
                        />
                    </div>
                </div>
            )}
            {numVerify(num) && !verified && (
                <div className="flex flex-col justify-center items-center my-8">
                    <h2 className="font-sans font-semibold text-2xl my-4">Verifing</h2>
                    <div className="flex flex-row items-center justify-center ">
                        <CircularProgress variant="indeterminate" size="2rem" color="primary" />
                    </div>
                </div>
            )}
            {verified && <h1 className="font-sans text-2xl font-semibold my-8">Thank you for registering. You will be redirected soon {dots}</h1>}
        </div>
    );
};

export default Verification;
