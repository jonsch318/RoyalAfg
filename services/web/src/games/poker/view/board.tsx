import React, { FC } from "react";
import { Container } from "@inlet/react-pixi";
import { Rectangle } from "./rectangle";
import { ICard } from "../models/card";
import { BoardPaddingX, BoardPaddingY, BorderRadius, CardHeight, CardWidth } from "./constants";
import Card from "./card";

type BoardProps = {
    cards: ICard[];
    forLength?: number;
    x?: number;
    y?: number;
};

const computeWidth = (cards: ICard[], forLength?: number) =>
    BoardPaddingX + BoardPaddingX * Math.max(cards.length, forLength ?? 0) + Math.max(cards.length, forLength ?? 0) * CardWidth;

const Board: FC<BoardProps> = ({ cards, forLength, x, y }) => {
    return (
        <Container x={x} y={y}>
            <Rectangle
                x={0}
                y={0}
                radius={BorderRadius}
                width={computeWidth(cards, forLength)}
                height={CardHeight + BoardPaddingY * 2}
                fill={0x000000}
                alpha={0.12}
            />
            {cards.map((c, i) => {
                return <Card card={c} x={BoardPaddingX + i * CardWidth + i * BoardPaddingX} y={BoardPaddingY} key={c.value * c.color} />;
            })}
        </Container>
    );
};

export default Board;
