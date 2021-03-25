import React, { FC, useState } from "react";
import Layout from "../../components/layout";
import { register as registerAccount } from "../../hooks/auth";
import Head from "next/head";
import { formatTitle } from "../../utils/title";
import Stepper from "@material-ui/core/Stepper";
import Step from "@material-ui/core/Step";
import StepLabel from "@material-ui/core/StepLabel";
import Credentials from "../../widgets/auth/credentials";
import Information from "../../widgets/auth/information";
import { useSnackbar } from "notistack";
import moment from "moment";
import { GetServerSideProps, InferGetServerSidePropsType } from "next";
import { getCSRF } from "../../hooks/auth/csrf";
import { useRouter } from "next/router";

export type RegisterDto = {
    username: string;
    password: string;
    birthdate: Date;
    email: string;
    fullName: string;
    acceptTerms: boolean;
};

function getSteps() {
    return ["Credentials", "Additional Information"];
}

function getStepContent(step: number, handleNext: () => void, handleBack: () => void, dto: RegisterDto, setDto: any) {
    switch (step) {
        case 0:
            return <Credentials handleNext={handleNext} dto={dto} setDto={setDto} />;
        case 1:
            return <Information handleNext={handleNext} handleBack={handleBack} dto={dto} setDto={setDto} />;
        default:
            return "Unknown step";
    }
}

const defaultDto = { username: "", password: "", email: "", fullName: "", birthdate: new Date(), acceptTerms: false };

export const getServerSideProps: GetServerSideProps = async (context) => {
    const csrf = await getCSRF(context);
    return {
        props: {
            csrf: csrf
        }
    };
};

const Register: FC = ({ csrf }: InferGetServerSidePropsType<typeof getServerSideProps>) => {
    const [activeStep, setActiveStep] = useState(0);
    const [skipped, setSkipped] = useState(new Set<number>());
    const steps = getSteps();
    const { enqueueSnackbar } = useSnackbar();
    const [dto, setDto] = useState<RegisterDto>(defaultDto);
    const router = useRouter();

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

        console.log(activeStep + " === " + steps.length);
        if (activeStep === steps.length - 1) {
            onSubmit().then();
        }
    };

    const handleBack = () => {
        setActiveStep((prevActiveStep) => prevActiveStep - 1);
    };

    const handleReset = () => {
        setActiveStep(0);
        setDto(defaultDto);
    };

    const onSubmit = (): Promise<void> => {
        console.log("Register");
        return registerAccount(
            {
                username: dto.username,
                password: dto.password,
                email: dto.email,
                birthdate: moment(dto.birthdate).format(),
                fullName: dto.fullName
            },
            csrf
        ).then((res) => {
            if (res.ok) {
                enqueueSnackbar("Successfully Registered", { variant: "success" });
                console.log("Refreshing: ", router.asPath);
                if (res.ok && typeof window !== undefined) {
                    window.location.href = "/";
                }
            } else {
                enqueueSnackbar("Something went wrong! Error code [" + res.status + "] " + res.statusText, { variant: "error" });
                handleReset();
            }
        });
    };

    return (
        <>
            <Head>
                <title>{formatTitle("Register a new Account")}</title>
            </Head>
            <Layout disableFooter>
                <form>
                    <div className="w-full md:h-screen flex items-center justify-center md:absolute md:inset-0">
                        <div className="bg-gray-200 md:rounded-md shadow-md ">
                            <Stepper activeStep={activeStep}>
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
                                    Registered Refreshing...
                                </span>
                            ) : (
                                getStepContent(activeStep, handleNext, handleBack, dto, setDto)
                            )}
                        </div>
                    </div>
                </form>
            </Layout>
        </>
    );
};

export default Register;
