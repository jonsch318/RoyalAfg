import React, { FC, useContext, useEffect, useState } from "react";
import { useApp } from "@saitonakamura/react-pixi";
import { URL } from "../view";
import { Texture } from "pixi.js";

export const TextureContext = React.createContext({});

export const useTexture = (name: string): Texture => {
    return useContext(TextureContext)[name];
};

type TextureProviderProps = {
    children: React.ReactNode;
};

const TextureProvider: FC<TextureProviderProps> = ({ children }) => {
    const [textures, setTexture] = useState({});

    const app = useApp();
    useEffect(() => {
        app.loader?.reset();
        app.loader?.add(URL).load((_, resource) => {
            setTexture(resource[URL].textures);
        });
    }, [app.loader]);

    return <TextureContext.Provider value={textures}>{children}</TextureContext.Provider>;
};
export default TextureProvider;
