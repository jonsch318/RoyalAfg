import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import { debounceTime } from "rxjs/operators";
import { BehaviorSubject } from "rxjs";

const Subj = new BehaviorSubject(0);

const Raise = (onRaise) => {
    const [raise, setRaise] = useState(0);
    useEffect(() => {
        const subscription = Subj.pipe(debounceTime(700)).subscribe((val) => {
            setRaise(val);
        });

        return () => {
            subscription.unsubscribe();
        };
    }, []);
    return (
        <div className="flex justify-center items-center rounded mx-4">
            <input
                className="border-blue-600 px-3 h-full outline-none py-1 flex"
                style={{ marginRight: "2px", width: "min-content" }}
                type="text"
                id="raiseInput"
                name="raiseInput"
                value={raise}
                onChange={(e) => {
                    setRaise(e.target.value);
                    Subj.next(parseFloat(e.target.value));
                }}
            />
            <button
                className="bg-white px-3 h-full rounded text-black overflow-hidden rounded-none hover:bg-yellow-500 transition-colors ease-in-out duration-150"
                onClick={() => onRaise({ action: 3, payload: raise * 100 })}>
                RAISE
            </button>
        </div>
    );
};

Raise.propTypes = {
    onRaise: PropTypes.func
};

export default Raise;
