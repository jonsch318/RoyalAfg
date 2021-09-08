import React, { FC, forwardRef } from "react";
import { useFormContext } from "react-hook-form";

const DateInput: FC = () => {
    const ExampleCustomInput = forwardRef(({ value, onClick }, ref) => (
        <button className="example-custom-input" onClick={onClick} ref={ref}>
            {value}
        </button>
    ));

    const {
        register,
        formState: { errors }
    } = useFormContext();

    return (
        <section>
            <input name="birthdate" type="text" {...register("birthdate")} />
        </section>
    );
};
