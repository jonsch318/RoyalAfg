import { useCallback, useEffect, useState } from "react";

export const useWidth = () => {
    const [width, setWidth] = useState(0);

    const ref = useCallback((node) => {
        if (node !== null) {
            setWidth(node.width);
        }
    }, []);

    return { width, ref };
};

export const useHeight = () => {
    const [height, setHeight] = useState(0);

    const ref = useCallback((node) => {
        if (node !== null) {
            setHeight(node.height);
        }
    }, []);

    return { height, ref };
};

export const useDim = (deps?: any[]) => {
    const [dim, setDim] = useState({ width: 0, height: 0 });

    const ref = useCallback(
        (node) => {
            if (node !== null) {
                setDim({ width: node.width, height: node.height });
            }
        },
        [deps]
    );

    return { dim, ref };
};

export const useResize = (ref: any) => {
    const [width, setWidth] = useState(0);
    const [height, setHeight] = useState(0);

    useEffect(() => {
        if (ref.current) {
            setWidth(ref.current.width);
            setHeight(ref.current.height);
        }
    }, [ref.current]);

    useEffect(() => {
        const handleResize = () => {
            setWidth(ref.current.width);
            setHeight(ref.current.height);
        };

        window.addEventListener("resize", handleResize);
        handleResize();
        return () => {
            window.removeEventListener("resize", handleResize);
        };
    }, [ref]);

    return { width, height };
};

//useResizeElement is a hook to get the dimensions of a element
export const useResizeElement = (ref: any) => {
    const [width, setWidth] = useState(0);
    const [height, setHeight] = useState(0);

    useEffect(() => {
        if (ref.current) {
            setWidth(ref.current.offsetWidth);
            setHeight(ref.current.offsetHeight);
        }
    }, [ref.current]);

    useEffect(() => {
        const handleResize = () => {
            setWidth(ref.current.offsetWidth);
            setHeight(ref.current.offsetHeight);
        };

        window.addEventListener("resize", handleResize);
        handleResize();
        return () => {
            window.removeEventListener("resize", handleResize);
        };
    }, [ref]);

    return { width, height };
};
