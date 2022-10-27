import { CryptoInfoDTO, CryptoInfoDTOSchema } from "./dtos/crypto";
import { SlotGameDTO, SlotGameDTOSchema, SlotGameOptionsDTO } from "./dtos/slot";
import { SlotGame, SlotGameSchema } from "./models/slot";

export const RequestMode = process.env.CORS_ENABLED ? "cors" : "same-origin";

export async function requestSlotSpin(options: SlotGameOptionsDTO, csrfToken: string, cryptoState: CryptoInfoDTO): Promise<SlotGame | null> {
    /*     const res = await fetch(`${process.env.SLOT_HOST}/api/slot/spin`, {
        method: "POST",
        mode: RequestMode,
        credentials: "include",
        body: JSON.stringify({
            ...options
        }),
        headers: {
            "Content-Type": "application/json",
            "X-CSRF-Token": csrfToken
        }
    });

    if (!res.ok) {
        console.log("Slot spin result not ok: ", res.status);
        return null;
    } */

    try {
        //const gameDetails: SlotGameDTO = await SlotGameDTOSchema.validate(await res.json());
        const gameDetails: SlotGameDTO = await SlotGameDTOSchema.validate(await mockRequestSlotSpin());

        //decode and verify game details
        //const gameDetails: SlotGameDto = await res.json();

        //convert nano unix time to date
        const time = new Date(gameDetails.timeOfGame);

        return SlotGameSchema.validate({
            numbers: gameDetails.numbers,
            alpha: gameDetails.alpha,
            beta: gameDetails.beta,
            proof: gameDetails.proof,
            timeOfGame: time,
            algorithms: gameDetails.algorithms,
            publicKey: cryptoState.publicKey,
            seed: cryptoState.seed
        });
    } catch (error) {
        console.log("Slot spin result error: ", error);
        return null;
    }
}

export async function getCryptoInfo(): Promise<CryptoInfoDTO> {
    // const res = await fetch(`${process.env.SLOT_HOST}/api/slot/cryptoinfo`, {
    //     method: "GET",
    //     mode: RequestMode
    // });

    // if (!res.ok) {
    //     console.log("Could not request crypto info: ", res.statusText);
    //     throw new Error(`Cpuld not request crypto info: ${res.statusText}`);
    // }

    try {
        //const CryptoInfo: CryptoInfoDTO = await CryptoInfoDTOSchema.validate(await res.json());
        const CryptoInfo: CryptoInfoDTO = await CryptoInfoDTOSchema.validate(await mockGetCryptoInfo());

        return CryptoInfo;
    } catch (error) {
        console.log("Error while parsing result data. Error: ", error);
        throw error;
    }
}

async function mockGetCryptoInfo(): Promise<string> {
    const mock = {
        publicKey: "mockPublicKey",
        seed: "mockSeed"
    };

    return JSON.stringify(mock);
}

async function mockRequestSlotSpin(): Promise<string> {
    const numbers = [];
    for (let i = 0; i < 4; i++) {
        numbers.push(rand(0, 7));
    }

    const mock = {
        numbers: [0, 0, 0, 5],
        alpha: 1,
        beta: 2,
        proof: "mockProof",
        timeOfGame: 123456789,
        algorithms: "mockAlgorithms"
    };

    return JSON.stringify(mock);
}

function rand(min, max): number {
    return Math.floor(Math.random() * (max - min + 1) + min);
}
