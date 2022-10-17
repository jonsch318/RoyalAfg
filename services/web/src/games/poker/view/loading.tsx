import React, { FC, useState } from "react";
import { Container, PixiComponent, useTick, withFilters } from "@saitonakamura/react-pixi";
import { Graphics } from "pixi.js";
import * as PIXI from "pixi.js";
import { LoadingRadius, LoadingStrokeWidth } from "./constants";

interface SpinnerProps {
    radius: number;
    angle: number;
    x?: number;
    y?: number;
}

const Spinner = PixiComponent<SpinnerProps, Graphics>("Spinner", {
    create: () => new Graphics(),
    applyProps: (ins, _, props) => {
        ins.clear();
        ins.lineStyle(LoadingStrokeWidth, 0x000000);
        ins.arc(props.radius, props.radius, props.radius, 0, Math.PI);
        ins.position.set(props.radius, props.radius);
        ins.pivot.set(props.radius);
        ins.angle = props.angle;
    }
});

const Filters = withFilters(Container, {
    blur: PIXI.filters.BlurFilter
});

type LoadingProps = {
    x?: number;
    y?: number;
};

const Loading: FC<LoadingProps> = ({ x, y }) => {
    const [vel, setVel] = useState(0);
    const [velPos, setVelPos] = useState(false);
    const [angle, setAngle] = useState(0);

    useTick((delta) => {
        if (vel < 3.25) {
            setVelPos(true);
        } else if (vel > 12) {
            setVelPos(false);
        }
        if (velPos) {
            setVel(vel + 0.075);
        } else {
            setVel(vel - 0.075);
        }
        setAngle((angle + vel + (delta ?? 0)) % 360);
    });

    return (
        <Container x={(x ?? 0) + LoadingStrokeWidth} y={(y ?? 0) + LoadingStrokeWidth}>
            <Filters blur={{ blur: 0.25 }}>
                <Spinner radius={LoadingRadius} angle={angle} />
            </Filters>
        </Container>
    );
};
export default Loading;
