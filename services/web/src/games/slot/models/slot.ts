import * as yup from "yup";

export const SlotGameSchema = yup.object().shape({
    numbers: yup.array().of(yup.number().min(0)).length(4),
    publicKey: yup.string().required(),
    proof: yup.string(),
    alpha: yup.string(),
    beta: yup.string(),
    algorithms: yup.string(),
    timeOfGame: yup.date(),
    seed: yup.string().required()
});

export type SlotGame = yup.InferType<typeof SlotGameSchema>;
