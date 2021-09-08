/* eslint-disable jsx-a11y/anchor-is-valid */
import React, { FC, useState } from "react";
import Layout from "../../components/layout";
import Head from "next/head";
import { formatTitle } from "../../utils/title";
import Stepper from "@material-ui/core/Stepper";
import Step from "@material-ui/core/Step";
import StepLabel from "@material-ui/core/StepLabel";
import Credentials from "../../widgets/auth/credentials";
import Information from "../../widgets/auth/information";
import { GetServerSideProps, InferGetServerSidePropsType } from "next";
import { getCSRF } from "../../hooks/auth/csrf";
import { serverSideTranslations } from "next-i18next/serverSideTranslations";
import { useTranslation } from "next-i18next";
import Link from "next/link";
import Verification from "../../widgets/auth/verification";

export const getServerSideProps: GetServerSideProps = async (context) => {
    const csrf = await getCSRF(context);
    return {
        props: {
            csrf: csrf,
            ...(await serverSideTranslations(context.locale, ["common", "auth"]))
        }
    };
};

export type RegisterDto = {
    username: string;
    password: string;
    birthdate: Date;
    email: string;
    fullName: string;
    acceptTerms: boolean;
};

function getSteps() {
    return ["Credentials", "Additional Information", "Verification"];
}

function getStepContent(step: number, handleNext: () => void, handleBack: () => void, dto: RegisterDto, setDto: any, csrf: string) {
    switch (step) {
        case 0:
            return <Credentials handleNext={handleNext} dto={dto} setDto={setDto} />;
        case 1:
            return <Information handleNext={handleNext} handleBack={handleBack} dto={dto} setDto={setDto} csrf={csrf} />;
        case 2:
            return <Verification handleNext={handleNext} dto={dto} setDet={setDto} csrf={csrf} />;
        default:
            return "Unknown step";
    }
}

export const DefaultRegisterDto = { username: "", password: "", email: "", fullName: "", birthdate: new Date(), acceptTerms: false };

const Register: FC = ({ csrf }: InferGetServerSidePropsType<typeof getServerSideProps>) => {
    const { t } = useTranslation("auth");

    const [activeStep, setActiveStep] = useState(0);
    const [skipped, setSkipped] = useState(new Set<number>());
    const steps = getSteps();
    const [dto, setDto] = useState<RegisterDto>(DefaultRegisterDto);

    const isStepSkipped = (step: number) => {
        return skipped.has(step);
    };

    const handleNext = () => {
        let newSkipped = skipped;
        if (isStepSkipped(activeStep)) {
            newSkipped = new Set(newSkipped.values());
            newSkipped.delete(activeStep);
        }

        setActiveStep((prevActiveStep) => prevActiveStep + 1);
        setSkipped(newSkipped);
    };

    const handleBack = () => {
        setActiveStep((prevActiveStep) => prevActiveStep - 1);
    };

    return (
        <Layout disableFooter>
            <Head>
                <title>{formatTitle(t("TitleRegister"))}</title>
            </Head>
            <div className="flex flex-col w-full h-full md:absolute items-center justify-center md:inset-0">
                <div className="bg-gray-200 md:rounded-md shadow-md">
                    <Stepper activeStep={activeStep} orientation={"horizontal"}>
                        {steps.map((label, index) => {
                            const stepProps: { completed?: boolean } = {};
                            const labelProps: { optional?: React.ReactNode } = {};
                            if (isStepSkipped(index)) {
                                stepProps.completed = false;
                            }
                            return (
                                <Step key={label} {...stepProps}>
                                    <StepLabel {...labelProps}>{label}</StepLabel>
                                </Step>
                            );
                        })}
                    </Stepper>
                    {activeStep === steps.length ? (
                        <span className="text-center font-sans font-semibold text-white py-6 px-20 grid" style={{ background: "#10B981" }}>
                            {t("Successfully Registered")}...
                        </span>
                    ) : (
                        getStepContent(activeStep, handleNext, handleBack, dto, setDto, csrf)
                    )}
                </div>
                <Link href="/auth/login">
                    <a className="text-center pt-5 block w-full text-blue-700 hover:text-blue-500 transition-colors duration-100">
                        {t("if you already have an account login here")}
                    </a>
                </Link>
            </div>
        </Layout>
    );
};

export default Register;
