import React, { FC, SetStateAction } from "react";
import { DefaultRegisterDto, RegisterDto } from "../../pages/auth/register";
import Checkbox from "@material-ui/core/Checkbox";
import moment from "moment";
import "react-datepicker/dist/react-datepicker.css";
import { useTranslation } from "next-i18next";
import { register as registerAccount } from "../../hooks/auth";
import { useSnackbar } from "notistack";
import { Controller, useForm } from "react-hook-form";
import { information } from "../../models/register";
import { yupResolver } from "@hookform/resolvers/yup";

type InformationProps = {
    setStep: (value: SetStateAction<number>) => void;
    dto: RegisterDto;
    setDto: React.Dispatch<React.SetStateAction<RegisterDto>>;
    csrf: string;
};

const isEmail = (str: string): boolean => {
    // eslint-disable-next-line no-control-regex
    const r = /(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])/;
    return r.test(str);
};

const isValidBirthdate = (date: Date): boolean => {
    return moment(date).isBefore(moment().subtract(16, "years")) && moment(date).isAfter(moment().subtract(100, "years"));
};

interface IFormInputs {
    email: string;
    birthdate: Date;
    fullName: string;
    acceptTerms: boolean;
}

const Information: FC<InformationProps> = ({ setStep, dto, setDto, csrf }) => {
    const { t } = useTranslation("auth");
    const { enqueueSnackbar } = useSnackbar();
    const {
        register,
        handleSubmit,
        formState: { errors },
        control
    } = useForm<IFormInputs>({
        resolver: yupResolver(information)
    });

    const handleReset = () => {
        setDto({
            ...dto,
            acceptTerms: true,
            birthdate: DefaultRegisterDto.birthdate,
            email: DefaultRegisterDto.email,
            fullName: DefaultRegisterDto.fullName
        });
    };

    const onSubmit = (data: IFormInputs): Promise<void> => {
        console.log("Register");
        return registerAccount(
            {
                username: dto.username,
                password: dto.password,
                email: dto.email,
                birthdate: moment(dto.birthdate).unix(),
                fullName: dto.fullName,
                acceptTerms: true //Can only press register with accepted terms
            },
            csrf
        ).then((res) => {
            if (res.ok) {
                //Successfully registered a new account
                enqueueSnackbar(t("Successfully Registered"), { variant: "success" });
            } else {
                enqueueSnackbar("Something went wrong! Error code [" + res.status + "] " + res.statusText, { variant: "error" });
                handleReset();
            }
        });
    };

    return (
        <div className="mx-16 my-6 font-sans">
            <form onSubmit={handleSubmit(onSubmit)}>
                <section className="mb-6 font-sans text-lg font-medium">
                    <label htmlFor="email" className="mb-2 block">
                        {t("Email*:")}
                    </label>
                    <input
                        className="block px-8 py-4 rounded w-full outline-none"
                        type="text"
                        name="email"
                        placeholder={t("Your email")}
                        {...register("email")}
                    />
                    <p>{errors.email?.message}</p>
                </section>
                <section className="mb-6 font-sans text-lg font-medium">
                    <label htmlFor="fullName" className="mb-2 block">
                        {t("Fullname*:")}
                    </label>
                    <input
                        className="block px-8 py-4 rounded w-full outline-none"
                        type="text"
                        name="fullName"
                        placeholder={t("Your name")}
                        {...register("fullName")}
                    />
                    <p>{errors.fullName?.message}</p>
                </section>
                <section className="mb-6 font-sans text-lg font-medium">
                    <label htmlFor="birthdate" className="mb-2 block">
                        {t("Birthdate*:")}
                    </label>
                    <input
                        className="block px-8 py-4 rounded w-full outline-none"
                        type="date"
                        name="birthdate"
                        placeholder="Your Birthdate"
                        {...register("birthdate")}
                    />
                    <p>{errors.birthdate?.message}</p>
                </section>
                <section>
                    <div className="mb-4 font-sans text-lg font-medium">
                        <Controller
                            name="birthdate"
                            control={control}
                            defaultValue=""
                            render={({ field }) => <Checkbox color="primary" {...field} />}
                        />
                        <span>
                            {t("I consent to the") + " "}
                            <a href="/legal/terms" className="font-sans text-blue-800">
                                {t("terms and conditions")}
                            </a>{" "}
                            {t("and our") + " "}
                            <a href="/legal/privacy" className="font-sans text-blue-800">
                                {t("privacy statement")}
                            </a>
                        </span>
                        <p>{errors.acceptTerms?.message}</p>
                    </div>
                </section>
                <div>
                    <button
                        className="w-full font-semibold text-xl py-4 bg-gray-700 hover:bg-gray-800 transition-colors duration-150 disabled:opacity-70 text-white my-2 rounded"
                        onClick={() => {
                            setStep((x) => x - 1);
                        }}>
                        {t("Back")}
                    </button>
                    <input
                        className="w-full font-semibold text-xl py-4 bg-blue-600 hover:bg-blue-500 transition-colors duration-150 disabled:opacity-50 disabled:cursor-not-allowed opacity-100 text-white my-2 rounded mb-8"
                        type="submit"
                        value={t("Register").toString()}
                    />
                </div>
            </form>
        </div>
    );
};

export default Information;
