import { STATIC_STATUS_PAGES } from "next/dist/shared/lib/constants";
import { bool } from "prop-types";
import * as yup from "yup";

export const SlotGameDTOSchema = yup.object().shape({
    numbers: yup.array().of(yup.number().min(0)).length(4).required(),
    alpha: yup.string().required(),
    beta: yup.string().required(),
    proof: yup.string().required(),
    algorithms: yup.string(),
    timeOfGame: yup.number().required().min(1)
});

export type SlotGameDTO = yup.InferType<typeof SlotGameDTOSchema>;

export type SlotGameOptionsDTO = {
    doubleFactor: boolean;
};
