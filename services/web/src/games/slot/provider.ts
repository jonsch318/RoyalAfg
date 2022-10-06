import { CryptoDTO, CryptoInfoDTO, CryptoInfoDTOSchema } from "./dtos/crypto"
import { SlotGameDTO, SlotGameDTOSchema, SlotGameOptionsDTO} from  "./dtos/slot"
import { SlotGame } from "./models/slot"

export const RequestMode = process.env.CORS_ENABLED ? "cors" : "same-origin"
export const RequestCredentialMode = process.env.CREDENTIAL_MODE ? process.env.CREDENTIAL_MODE : "same-origin"

export async function RequestSlotSpin(options: SlotGameOptionsDTO, csrfToken: string, cryptoState: CryptoDTO): Promise<SlotGame | null> {
    
    
    const res = await fetch(`${process.env.SLOT_HOST}/api/slot/spin`, {
        method: "POST",
        mode: RequestMode,
        credentials: RequestCredentialMode,
        body: {
            ...options 
        },
        headers: {
            "Content-Type": "application/json",
            "X-CSRF-Token": csrfToken
        }
    })

    if (!res.ok) {
        console.log("Slot spin result not ok: ", res.status)
        return null
    }

    const gameDetails: SlotGameDTO = await SlotGameDTOSchema.validate(await res.json())


    //decode and verify game details
    //const gameDetails: SlotGameDto = await res.json();

    if (!gameDetails.verify()) {
        return null
    }

    //convert nano unix time to date
    const time = new Date(gameDetails.timeOfGame)

    return {
        algorithms: gameDetails.algorithms,
        alpha: gameDetails.alpha,
        beta: gameDetails.beta,
        numbers: gameDetails.numbers,
        proof: gameDetails.proof,
        timeOfGame: time,
        publicKey: cryptoState.publicKey,
        seed: cryptoState.seed,
    }

}

export async function getCryptoInfo(): Promise<CryptoInfoDTO> {
    const res = await fetch(`${process.env.SLOT_HOST}/api/slot/cryptoinfo`, {
        method: "GET",
        mode: RequestMode,
    })

    if (!res.ok) {
        console.log("Could not request crypto info: ", res.statusText)
        throw new Error(`Cpuld not request crypto info: ${res.statusText}`)
    }

    try {
        const CryptoInfo: CryptoInfoDTO = await CryptoInfoDTOSchema.validate(await res.json())
        
        return CryptoInfo
    } catch (error) {
        console.log("Error while parsing result data. Error: ", error)
        throw error
    }
}
