import React from "react";
import { useForm } from "react-hook-form";
import Layout from "../../components/layout";
import { signIn } from "../../hooks/auth";
import PasswordBox from "../../components/form/passwordBox";

const Login = () => {
    console.log("URL: ", process.env.NEXT_PUBLIC_AUTH_HOST);
    const { register, handleSubmit, errors } = useForm();
    const onSubmit = async (data) => {
        await signIn({ username: data.username, password: data.password });
    };

    return (
        <Layout disableFooter>
            <div className="w-full md:h-screen grid items-center justify-center mt-24 md:mt-0 md:absolute md:inset-0">
                <div className="bg-gray-200 md:rounded-md shadow-md">
                    <div className="heading mx-16 my-8">
                        <h1 className="text-center font-sans font-semibold text-3xl">Sign into your Account</h1>
                    </div>
                    <div className="content md:px-24 px-4">
                        <form onSubmit={handleSubmit(onSubmit)}>
                            <section className="mb-8 font-sans text-lg font-medium">
                                <label htmlFor="username" className="mb-2 block">
                                    Username*:
                                </label>
                                <input
                                    className="block px-8 py-4 rounded w-full"
                                    ref={register({ required: true, maxLength: 100, minLength: 3 })}
                                    type="text"
                                    id="username"
                                    name="username"
                                    placeholder="Your Username"
                                    aria-describedby="username-constraints"
                                    required
                                />
                                {errors.username && (
                                    <span className="text-sm text-red-700" id="username-constraints">
                                        This field is required and can only be more than 3 and less than 100!
                                    </span>
                                )}
                            </section>
                            <PasswordBox errors={errors} register={register} />
                            <button
                                className="block w-full px-8 py-3  bg-blue-600 text-white rounded hover:bg-blue-700 transition-colors duration-150 font-sans font-medium cursor-pointer"
                                type="submit">
                                Log In
                            </button>
                            <span className="font-sans font-light text-sm mb-8">
                                Or{" "}
                                <a href="/register" className="font-sans text-blue-800">
                                    register
                                </a>{" "}
                                a new account
                            </span>
                            <span className="text-sm mb-8 font-sans font-light block text-ce">text fields with a * are required</span>
                        </form>
                    </div>
                </div>
            </div>
        </Layout>
    );
};

export default Login;
