import React, { FC } from "react";
import { RegisterDto } from "../../pages/auth/registerstepper";
import Checkbox from "@material-ui/core/Checkbox";
import DatePicker from "react-datepicker";
import moment from "moment";
import "react-datepicker/dist/react-datepicker.css";

type InformationProps = {
    handleNext: () => void;
    handleBack: () => void;
    dto: RegisterDto;
    setDto: React.Dispatch<React.SetStateAction<RegisterDto>>;
};

const isEmail = (str: string): boolean => {
    const r = /(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])/;
    return r.test(str);
};

const isValidBirthdate = (date: Date): boolean => {
    return moment(date).isBefore(moment().subtract(16, "years")) && moment(date).isAfter(moment().subtract(100, "years"));
};

const Information: FC<InformationProps> = ({ handleNext, handleBack, dto, setDto }) => {
    const shouldDisable = (): boolean => {
        return !isValidBirthdate(dto.birthdate) || dto.fullName == "" || dto.email == "" || !isEmail(dto.email) || !dto.acceptTerms;
    };

    return (
        <div className="mx-16 my-6">
            <section className="mb-6 font-sans text-lg font-medium">
                <label htmlFor="email" className="mb-2 block">
                    Email*:
                </label>
                <input
                    className="block px-8 py-4 rounded w-full outline-none"
                    style={{ border: dto.email == "" || !isEmail(dto.email) ? "2px solid rgb(190, 18, 60)" : "" }}
                    type="email"
                    id="email"
                    name="email"
                    placeholder="Your Email"
                    required
                    value={dto.email}
                    onChange={(e) => setDto({ ...dto, email: e.target.value })}
                />
            </section>
            <section className="mb-6 font-sans text-lg font-medium">
                <label htmlFor="fullName" className="mb-2 block">
                    Full Name*:
                </label>
                <input
                    className="block px-8 py-4 rounded w-full outline-none"
                    type={"text"}
                    id="fullName"
                    name="fullName"
                    placeholder="Your Full Name"
                    style={{ border: dto.fullName == "" ? "2px solid rgb(190, 18, 60)" : "" }}
                    value={dto.fullName}
                    onChange={(e) => setDto({ ...dto, fullName: e.target.value })}
                    required
                />
            </section>
            <section className="mb-6 font-sans text-lg font-medium">
                <label htmlFor="birthdate" className="mb-2 block">
                    Birthdate*:
                </label>
                <DatePicker
                    className="block px-8 py-4 rounded w-full outline-none"
                    type="date"
                    id="birthdate"
                    name="birthdate"
                    placeholder="Your Birthdate"
                    selected={dto.birthdate}
                    style={{ border: !isValidBirthdate(dto.birthdate) ? "2px solid rgb(190, 18, 60)" : "" }}
                    onChange={(e: Date) => setDto({ ...dto, birthdate: e })}
                    required
                />
            </section>
            <section>
                <div className="mb-4 font-sans text-lg font-medium">
                    <Checkbox
                        checked={dto.acceptTerms}
                        onChange={(e) => setDto({ ...dto, acceptTerms: e.target.checked })}
                        color="primary"
                        required></Checkbox>
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
                </div>
            </section>
            <button
                className="w-full font-sans font-semibold text-xl py-4 bg-gray-500 disabled:opacity-70 text-white my-2 rounded"
                onClick={() => {
                    handleBack();
                }}>
                Back
            </button>
            <button
                className="w-full font-sans font-semibold text-xl py-4 bg-blue-500 disabled:opacity-70 text-white my-2 rounded"
                disabled={shouldDisable()}
                onClick={() => {
                    handleNext();
                }}>
                Register
            </button>
        </div>
    );
};

export default Information;
