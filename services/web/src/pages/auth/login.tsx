import React, { FC } from "react";
import { useForm } from "react-hook-form";
import Layout from "../../components/layout";
import { signIn } from "../../hooks/auth";
import PasswordBox from "../../components/form/passwordBox";
import { GetServerSideProps, InferGetServerSidePropsType } from "next";
import { getCSRF } from "../../hooks/auth/csrf";
import { useRouter } from "next/router";
import { serverSideTranslations } from "next-i18next/serverSideTranslations";
import { useTranslation } from "next-i18next";
import Head from "next/head";
import { yupResolver } from "@hookform/resolvers/yup";
import { credentials } from "../../models/register";

interface IFormInput {
    username: string;
    password: string;
}

const Login: FC = ({ csrf }: InferGetServerSidePropsType<typeof getServerSideProps>) => {
    const { t } = useTranslation("auth");

    console.log("URL: ", process.env.NEXT_PUBLIC_AUTH_HOST);
    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<IFormInput>({
        resolver: yupResolver(credentials)
    });
    const router = useRouter();
    const onSubmit = (data) => {
        signIn({ username: data.username, password: data.password }, csrf)
            .then((res) => {
                console.log("Refreshing: ", router.asPath);
                if (res.ok) {
                    window.location.href = "/";
                }
            })
            .catch(() => {
                console.log("Refreshing: ", router.asPath);
            });
    };

    return (
        <Layout disableFooter>
            <Head>
                <title>{t("TitleLogin")}</title>
            </Head>
            <div className="w-full md:h-screen grid items-center justify-center mt-24 md:mt-0 md:absolute md:inset-0">
                <div className="bg-gray-200 md:rounded-md shadow-md font-sans font-medium">
                    <div className="heading mx-16 my-8">
                        <h1 className="text-center font-sans font-semibold text-3xl">{t("Sign into your account")}</h1>
                    </div>
                    <div className="content md:px-24 px-4">
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
                                />
                                <p className="text-sm text-red-700" id="username-constraints">
                                    {errors.username && t("This field is required and can only be more than 3 and less than 100!")}
                                </p>
                            </section>
                            <section className="mb-6 text-lg">
                                <PasswordBox register={register} />
                                <p className="text-sm text-red-700" id="username-constraints">
                                    {errors.password && t("This field is required and can only be more than 3 and less than 100!")}
                                </p>
                            </section>
                            <input
                                className="block w-full px-8 py-3  bg-blue-600 text-white rounded hover:bg-blue-700 transition-colors duration-150 font-sans font-medium cursor-pointer"
                                type="submit"
                                value={t("Log In").toString()}
                            />
                            <span className="font-sans font-light text-sm mb-8">
                                {t("Or") + " "}
                                <a href="/register" className="font-sans text-blue-800">
                                    {t("register")}
                                </a>{" "}
                                {t("a new account")}
                            </span>
                            <span className="text-sm mb-8 font-sans font-light block text-ce">{t("Text fields with a * are required")}</span>
                        </form>
                    </div>
                </div>
            </div>
        </Layout>
    );
};

export const getServerSideProps: GetServerSideProps = async (context) => {
    const csrf = await getCSRF(context);
    return {
        props: {
            csrf: csrf,
            ...(await serverSideTranslations(context.locale, ["common", "auth"]))
        }
    };
};

export default Login;
