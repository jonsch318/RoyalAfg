import React, { useState } from "react";

type PasswordBoxProps = {
    errors: {
        password?: any;
    };
    register: any;
};

const PasswordBox: React.FC<PasswordBoxProps> = ({ errors, register }: PasswordBoxProps) => {
    const [hidePassword, setHidePassword] = useState(true);
    return (
        <section className="mb-6 font-sans text-lg font-medium">
            <label htmlFor="password" className="mb-2 block">
                Passphrase*:
            </label>
            <input
                className="block px-8 py-4 rounded w-full"
                ref={register({ required: true, maxLength: 100, minLength: 3 })}
                type={hidePassword ? "password" : "text"}
                id="password"
                name="password"
                autoComplete="current-password"
                placeholder="Your Password"
                aria-describedby="password-constraints"
                required
            />
            <button
                type="button"
                onClick={() => {
                    setHidePassword(!hidePassword);
                }}
                aria-label={hidePassword ? "Show password in plain text. This will show your password on screen." : "Hide Password."}>
                {hidePassword ? "Show Password" : "Hide Password"}
            </button>
            {errors?.password && (
                <span className="text-sm text-red-700" id="password-constraints">
                    This field is required and can only be more than 3 and less than 100!
                </span>
            )}
        </section>
    );
};

export default PasswordBox;
