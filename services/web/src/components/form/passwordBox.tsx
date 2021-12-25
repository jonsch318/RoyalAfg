import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons";
import { useTranslation } from "next-i18next";
import React, { useState } from "react";

type PasswordBoxProps = {
    register: any;
    value?: string;
};

const PasswordBox: React.FC<PasswordBoxProps> = ({ register, value }: PasswordBoxProps) => {
    const { t } = useTranslation("auth");

    const [showPassword, setShowPassword] = useState(false);

    return (
        <>
            <label htmlFor="password" className="mb-2 block">
                {t("Passphrase")}
            </label>
            <div className="rounded w-full bg-white flex flex-row">
                <input
                    className="block w-full h-full px-8 py-4 bg-transparent outline-none"
                    type={showPassword ? "text" : "password"}
                    name="password"
                    autoComplete="current-password"
                    placeholder={t("Your password")}
                    aria-describedby="password-constraints"
                    value={value}
                    {...register("password", { required: true, maxLength: 100, minLength: 3 })}
                />
                <button
                    className="inline w-20 outline-none focus:outline-none"
                    type="button"
                    onClick={() => setShowPassword((x) => !x)}
                    aria-label={showPassword ? "Show password in plain text. This will show your password on screen." : "Hide Password."}>
                    {showPassword ? <FontAwesomeIcon icon={faEye} /> : <FontAwesomeIcon icon={faEyeSlash} />}
                </button>
            </div>
        </>
    );
};

export default PasswordBox;
