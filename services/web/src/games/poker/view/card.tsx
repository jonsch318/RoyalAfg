import React, { FC } from "react";
import { Sprite } from "@inlet/react-pixi";
import { useTexture } from "./textures";
import { ICard } from "../models/card";
import { CardHeight, CardWidth } from "./constants";

type CardProps = {
    card: ICard;
    x?: number;
    y?: number;
};

const FileName = (card: ICard): string => {
    if (card.value >= 0 && card.color >= 0) return card.value + "_" + card.color + ".png";
    return "back.png";
};

const Card: FC<CardProps> = ({ card, x, y }) => {
    const texture = useTexture(FileName(card));
    return <>{texture && <Sprite texture={texture} x={x} y={y} width={CardWidth} height={CardHeight} />}</>;
};

Card.defaultProps = {
    card: {
        value: -1,
        color: -1
    },
    x: 0,
    y: 0
};

export default Card;
