import React from "react";
import { useForm } from "react-hook-form";
import FormItem from "../components/form/form-item";
import Layout from "../components/layout";

const Submit = (data) => {
    console.log(JSON.stringify(data));
};

const Register = () => {
    const { register, handleSubmit, errors } = useForm();

    return (
        <Layout disableFooter>
            <div className="w-full md:h-screen flex items-center justify-center md:absolute md:inset-0">
                <div className="bg-gray-300 md:rounded-md shadow-md">
                    <div className="heading mx-16 my-8">
                        <h1 className="text-center font-sans font-semibold text-3xl">
                            Register A New Account
                        </h1>
                    </div>
                    <div className="content md:px-24 px-4">
                        <form onSubmit={handleSubmit(Submit)}>
                            <FormItem>
                                <label htmlFor="username" className="mb-2 block">
                                    Username*:
                                </label>
                                <input
                                    className="block px-4 py-2 rounded w-full"
                                    ref={register({ required: true, maxLength: 100, minLength: 3 })}
                                    type="text"
                                    id="username"
                                    name="username"
                                    placeholder="Your Username"
                                />
                                {errors.username && (
                                    <span className="text-sm text-red-700">
                                        This field is required and can only be more than 3 and less
                                        than 100!
                                    </span>
                                )}
                            </FormItem>

                            <FormItem>
                                <label htmlFor="password" className="mb-2 block">
                                    Passphrase*:
                                </label>
                                <input
                                    className="block px-4 py-2 rounded w-full"
                                    ref={register({ required: true, maxLength: 100, minLength: 3 })}
                                    type="password"
                                    id="password"
                                    name="password"
                                    placeholder="Your Password"
                                />
                                {errors.password && (
                                    <span className="text-sm text-red-700">
                                        This field is required and can only be more than 3 and less
                                        than 100!
                                    </span>
                                )}
                            </FormItem>

                            <FormItem>
                                <label htmlFor="birthdate" className="mb-2 block">
                                    Birthdate*:
                                </label>
                                <input
                                    className="block px-4 py-2 rounded w-full"
                                    ref={register({ required: true })}
                                    type="date"
                                    id="birthdate"
                                    name="birthdate"
                                />
                                {errors.birthdate && (
                                    <span className="text-sm text-red-700">
                                        This field is required!
                                    </span>
                                )}
                            </FormItem>

                            <FormItem>
                                <label htmlFor="email">Email*:</label>
                                <input
                                    className="block px-4 py-2 rounded w-full"
                                    ref={register({
                                        required: true,
                                        minLength: "3",
                                        maxLength: "100"
                                    })}
                                    name="email"
                                    id="email"
                                    type="email"
                                    placeholder="Your Email"></input>
                                {errors.birthdate && (
                                    <span className="text-sm text-red-700">
                                        This field is required
                                    </span>
                                )}
                            </FormItem>

                            <FormItem>
                                <input
                                    type="checkbox"
                                    className="p-2 border-none form-checkbox mr-4 text-blue-700"
                                />
                                <span>
                                    I consent to the{" "}
                                    <a href="/legal/terms" className="font-sans text-blue-800">
                                        terms and conditions
                                    </a>{" "}
                                    and our{" "}
                                    <a href="/legal/privacy" className="font-sans text-blue-800">
                                        privacy statement
                                    </a>
                                </span>
                            </FormItem>

                            <input
                                className="block w-full px-4 py-2  bg-blue-600 text-white rounded hover:bg-blue-700 transition-colors duration-150 font-sans font-medium cursor-pointer"
                                type="submit"
                                value="Register"
                            />
                            <span className="font-sans font-light text-sm mb-8">
                                Or{" "}
                                <a href="/login" className="font-sans text-blue-800">
                                    login
                                </a>{" "}
                                if you already have an account
                            </span>
                            <span className="text-sm mb-8 font-sans font-light block text-ce">
                                Textfields with a * are required
                            </span>
                        </form>
                    </div>
                </div>
            </div>
        </Layout>
    );
};

export default Register;
