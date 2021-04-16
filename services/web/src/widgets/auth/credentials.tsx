import { useTranslation } from "next-i18next";
import React, { FC, useState } from "react";
import { RegisterDto } from "../../pages/auth/registerstepper";

type Credentials = {
    username: string;
    password: string;
};

type CredentialsProps = {
    handleNext: () => void;
    dto: RegisterDto;
    setDto: React.Dispatch<React.SetStateAction<RegisterDto>>;
};

const Credentials: FC<CredentialsProps> = ({ handleNext, dto, setDto }) => {
    const { t } = useTranslation("auth");

    const [hidePassword, setHidePassword] = useState(true);

    const shouldDisable = (): boolean => {
        return dto.username == "" || dto.password == "";
    };

    return (
        <div className="mx-16 my-6">
            <section className="mb-6 font-sans text-lg font-medium">
                <label htmlFor="username" className="mb-2 block">
                    {t("Username*:")}
                </label>
                <input
                    className="block px-8 py-4 rounded w-full outline-none"
                    type="text"
                    id="username"
                    name="username"
                    placeholder={t("Your username")}
                    required
                    style={{ border: dto.username == "" ? "2px solid rgb(190, 18, 60)" : "" }}
                    value={dto.username}
                    onChange={(e) => setDto({ ...dto, username: e.target.value })}
                />
            </section>
            <section className="mb-6 font-sans text-lg font-medium">
                <label htmlFor="password" className="mb-2 block">
                    {t("Passphrase*:")}
                </label>
                <input
                    className="block px-8 py-4 rounded w-full outline-none"
                    type={hidePassword ? "password" : "text"}
                    id="password"
                    name="password"
                    autoComplete="current-password"
                    placeholder={t("Your password")}
                    aria-describedby="password-constraints"
                    style={{ border: dto.password == "" ? "2px solid rgb(190, 18, 60)" : "" }}
                    value={dto.password}
                    onChange={(e) => setDto({ ...dto, password: e.target.value })}
                    required
                />
                <button
                    type="button"
                    onClick={() => {
                        setHidePassword(!hidePassword);
                    }}
                    aria-label={hidePassword ? "Show password in plain text. This will show your password on screen." : "Hide Password."}>
                    {hidePassword ? t("Show password") : t("Hide password")}
                </button>
            </section>
            <button
                className="w-full font-sans font-semibold text-xl py-4 bg-blue-500 disabled:opacity-70 text-white my-2 rounded"
                disabled={shouldDisable()}
                onClick={() => {
                    handleNext();
                }}>
                {t("Next")}
            </button>
        </div>
    );
};

export default Credentials;
