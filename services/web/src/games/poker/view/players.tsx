import React, { FC } from "react";
import { Container } from "@saitonakamura/react-pixi";
import Player, { PlayerHeight, PlayerWidth } from "./player";
import { IPoker } from "../models/poker";
import { TableRadiusX, TableRadiusY } from "./constants";

type PlayersProps = {
    poker: IPoker;
};

const GetAngle = (i: number, l: number): number => {
    return ((360 / l) * i * Math.PI) / 180;
};

const GetX = (i: number, l: number): number => {
    return TableRadiusX * Math.sin(GetAngle(i, l)) - PlayerWidth * 0.5;
};

const GetY = (i: number, l: number): number => {
    return TableRadiusY * Math.cos(GetAngle(i, l)) - PlayerHeight * 0.5;
};

const Players: FC<PlayersProps> = ({ poker }) => {
    const players = poker.players;

    return (
        <Container x={0} y={0}>
            {players.map((p, i) => {
                console.log("Player " + p.username + " angle: " + GetAngle(i, players.length));
                return (
                    <Player
                        player={p}
                        x={GetX(i, players.length)}
                        y={GetY(i, players.length)}
                        dealer={i === poker.dealer}
                        oneSelf={i === poker.index}
                        key={p.id}
                    />
                );
            })}
        </Container>
    );
};

export default Players;
