import React, { FC } from "react";
import { Container, Stage } from "@inlet/react-pixi";
import Players from "./view/players";
import useWindowDimensions from "../../hooks/windowSize";
import { usePoker } from "./provider";
import TextureProvider from "./view/textures";
import Board from "./view/board";
import Status from "./view/status";

export const URL = "/static/games/poker/textures/cards.json";

const View: FC = () => {
    const { width, height } = useWindowDimensions();

    const poker = usePoker();

    return (
        <Stage
            style={{ top: 60, height: height - 60 }}
            className="absolute right-0 left-0 bottom-0 w-screen"
            width={width}
            height={height - 60}
            options={{ backgroundColor: 0xffffff }}>
            <TextureProvider>
                <Status pot={poker.pot} bet={poker.bet} lobbyId={poker.lobbyInfo.lobbyId} appWidth={width} appHeight={height - 60} />
                <Container x={width * 0.5} y={(height - 60) * 0.5}>
                    <Players poker={poker} />
                </Container>
                <Container x={width * 0.5} y={(height - 60) * 0.5} anchor={{ x: 0, y: 0 }}>
                    <Board cards={poker.board} forLength={5} />
                </Container>
            </TextureProvider>
        </Stage>
    );
};

export default View;
