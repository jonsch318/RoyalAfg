import { Graphics } from "pixi.js";
import { PixiComponent } from "@saitonakamura/react-pixi";

interface RectangleProps {
    x: number;
    y: number;
    width: number;
    height: number;
    fill?: number;
    alpha?: number;
    radius?: number;
    border?: boolean;
}

export const Rectangle = PixiComponent<RectangleProps, Graphics>("Rectangle", {
    create: (props) => new Graphics(),
    applyProps: (instance, _, props) => {
        const { x, y, width, height, fill, alpha, radius, border } = props;

        instance.clear();
        instance.lineStyle(2, 0xf66742, border ? 1 : 0);
        instance.beginFill(fill, alpha ? alpha : 1);
        instance.drawRoundedRect(x, y, width, height, radius ? radius : 0);
        instance.endFill();
    }
});
