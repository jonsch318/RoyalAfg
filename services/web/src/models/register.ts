import * as yup from "yup";
import { getDateFromPast } from "../utils/time";

export type RegisterDto = {
    username: string;
    password: string;
    birthdate: Date;
    email: string;
    fullName: string;
    acceptTerms: boolean;
};

export const DefaultRegisterDto = {
    username: "",
    password: "",
    email: "",
    fullName: "",
    birthdate: new Date(),
    acceptTerms: false
};

export const credentials = yup.object().shape({
    username: yup.string().min(4).max(100).required(),
    password: yup.string().min(4).max(100).required()
});

export const information = yup.object().shape({
    email: yup.string().email().required("email is required"),
    birthdate: yup.date().min(getDateFromPast(150)).max(getDateFromPast(18)).required("Date is required"),
    acceptTerms: yup.boolean().required(),
    fullName: yup.string().required()
});
