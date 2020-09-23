import { useForm } from "react-hook-form";

export default function Login() {
    const { register, handleSubmit, watch, errors } = useForm();
    const onSubmit = data => console.log(JSON.stringify(data))

    console.log(watch("username"));

    return (
        <div className="w-full md:h-screen flex items-center justify-center md:absolute md:inset-0">
            <div className="bg-gray-300 md:rounded-md">
                <div className="heading mx-16 my-8">
                    <h1 className="text-center font-sans font-semibold text-3xl">Sign into your Account</h1>
                </div>
                <div className="content md:px-24 px-4">
                    <form onSubmit={handleSubmit(onSubmit)}>
                        <div className="mb-8 font-sans text-lg font-medium">
                            <label htmlFor="username" className="mb-2 block">
                                Username*:
                            </label>
                            <input className="block px-4 py-2 rounded w-full" ref={register({ required: true, maxLength: 100, minLength: 3 })} type="text" id="username" name="username" placeholder="Your Username" />
                            {errors.username && <span className="text-sm text-red-700" >This field is required and can only be more than 3 and less than 100!</span>}
                        </div>
                        <div className="mb-8 font-sans text-lg font-medium">
                            <label htmlFor="password" className="mb-2 block">
                                Passphrase*:
                            </label>
                            <input className="block px-4 py-2 rounded w-full" ref={register({ required: true, maxLength: 100, minLength: 3 })} type="password" id="password" name="password" placeholder="Your Password" />
                            {errors.password && <span className="text-sm text-red-700" >This field is required and can only be more than 3 and less than 100!</span>}
                        </div>
                        <input className="block w-full px-4 py-2  bg-blue-600 text-white rounded hover:bg-blue-700 transition-colors duration-150" type="submit" value="Log in" />
                        <span className="text-sm mb-8 font-sans font-light block text-ce">Textfields with a * are required</span>
                    </form>
                </div>
            </div>
        </div>
    )
}
