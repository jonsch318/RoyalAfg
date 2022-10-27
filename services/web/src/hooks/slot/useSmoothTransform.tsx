import { MotionValue, useSpring, useTransform } from "framer-motion";

export function useSmoothTransform(value: MotionValue<number>, springOptions: unknown, transformer: (v: number) => number) {
    return useSpring(useTransform(value, transformer), springOptions);
}
