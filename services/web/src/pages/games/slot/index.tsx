import React, { FC } from "react"
import { getCryptoInfo } from "../../../games/slot/provider"

export async function getStaticProps(context){
    const CryptoInfo = getCryptoInfo()
    
    return {
        props: {
            crypto: CryptoInfo
        },
        revalidate: 100
    }
}