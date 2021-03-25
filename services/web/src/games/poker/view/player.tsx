import React, { FC, useEffect, useRef } from "react";
import { Container, Text } from "@inlet/react-pixi";
import { IPlayer } from "../models/player";
import { Rectangle } from "./rectangle";
import { useResize, useWidth } from "../../../hooks/dimensions";
import Board from "./board";
import { BorderRadius, CardHeight, CardWidth, LoadingRadius, PlayerPaddingX, PlayerPaddingY } from "./constants";
import Loading from "./loading";

export const PlayerWidth = 300;
export const PlayerHeight = 2 * LoadingRadius + 30 + CardHeight + 30;

const getAlpha = (playerIn: boolean, dealer: boolean): number => {
    let base = playerIn ? 0.12 : 0.35;
    base += dealer ? 0.2 : 0;
    return base;
};

type PlayerProps = {
    player: IPlayer;
    dealer: boolean;
    oneSelf: boolean;
    x?: number;
    y?: number;
};

const Player: FC<PlayerProps> = ({ x, y, player, dealer, oneSelf }) => {
    const ref = useRef<PIXI.Text>();
    const { width } = useResize(ref);

    useEffect(() => {
        console.log("Width Player: ", width);
    }, [width]);

    return (
        <Container x={x} y={y}>
            <Rectangle
                x={0}
                y={0}
                width={PlayerWidth}
                height={PlayerHeight}
                alpha={getAlpha(player.in, dealer)}
                radius={BorderRadius}
                fill={dealer ? 0x27611b : 0x000000}
                border={oneSelf}
            />
            {player.waiting && <Loading y={0} />}
            <Text
                text={player.username}
                x={(player.waiting ? 2 * LoadingRadius : 0) + PlayerPaddingX}
                y={PlayerPaddingY}
                style={{ fontSize: 17, align: "center" }}
            />
            <Text
                text={player.buyIn + " -> " + player.bet}
                x={PlayerWidth - PlayerPaddingX - width}
                ref={ref}
                y={PlayerPaddingY}
                style={{ fontSize: 17 }}
            />
            <Board cards={player.cards} forLength={2} x={0.5 * PlayerWidth - CardWidth - 2 * PlayerPaddingX} y={2 * LoadingRadius + PlayerPaddingY} />
        </Container>
    );
};

export default Player;
