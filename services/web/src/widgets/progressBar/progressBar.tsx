import React, { useEffect } from "react";
import PrgBar from "@ramonak/react-progress-bar";

type ProgressBarProps = {
    callback: () => void;
    completed: boolean;
};

const ProgressBar: React.FC<ProgressBarProps> = ({ callback, completed }) => {
    const [progress, setProgress] = React.useState(0);

    useEffect(() => {
        setInterval(() => {
            if (completed) {
                callback();
            }
        }, 1200);
    }, [completed, callback]);

    useEffect(() => {
        if (completed) {
            setProgress(100);
        }
    }, [completed]);

    return (
        <>
            <PrgBar completed={progress} customLabel={"%"} width={"30rem"} animateOnRender transitionDuration="0.75s" />
        </>
    );
};

export default ProgressBar;
