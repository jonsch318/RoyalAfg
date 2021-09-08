import { yupResolver } from "@hookform/resolvers/yup";
import { useTranslation } from "next-i18next";
import React, { FC, useState } from "react";
import { useForm } from "react-hook-form";
import PasswordBox from "../../components/form/passwordBox";
import { credentials } from "../../models/register";
import { RegisterDto } from "../../pages/auth/register";

type CredentialsProps = {
    handleNext: () => void;
    dto: RegisterDto;
    setDto: React.Dispatch<React.SetStateAction<RegisterDto>>;
};

interface IFormInput {
    username: string;
    password: string;
}

const Credentials: FC<CredentialsProps> = ({ handleNext, dto, setDto }) => {
    const { t } = useTranslation("auth");

    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<IFormInput>({
        resolver: yupResolver(credentials)
    });

    const onSubmit = (data: IFormInput) => {
        setDto((x) => {
            return { ...x, username: data.username, password: data.password };
        });
        handleNext();
    };

    return (
        <div className="mx-16 my-6">
            <form onSubmit={handleSubmit(onSubmit)}>
                <section className="mb-8 text-lg">
                    <label htmlFor="username" className="mb-2 block">
                        {t("Username")}
                    </label>
                    <input
                        className="block px-8 py-4 rounded w-full outline-none"
                        {...register("username", { required: true, maxLength: 100, minLength: 3 })}
                        type="text"
                        name="username"
                        placeholder={t("Your username")}
                        aria-describedby="username-constraints"
                        value={dto.username}
                    />
                    <p className="text-sm text-red-700" id="username-constraints">
                        {errors.username && t("This field is required and can only be more than 3 and less than 100!")}
                    </p>
                </section>
                <section className="mb-6 text-lg">
                    <PasswordBox register={register} value={dto.password} />
                    <p className="text-sm text-red-700" id="username-constraints">
                        {errors.password && t("This field is required and can only be more than 3 and less than 100!")}
                    </p>
                </section>
                <input
                    className="w-full font-sans font-semibold text-xl py-4 bg-blue-600 hover:bg-blue-500 transition-colors duration-150 disabled:opacity-50 disabled:cursor-not-allowed text-white my-2 rounded mb-8"
                    type="submit"
                    value={t("Next").toString()}
                />
            </form>
        </div>
    );
};

export default Credentials;
