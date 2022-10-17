import * as yup from "yup";

export const CryptoInfoDTOSchema = yup.object().shape({
    publicKey: yup.string().required(),
    seed: yup.string().required()
});

export type CryptoInfoDTO = yup.InferType<typeof CryptoInfoDTOSchema>;
