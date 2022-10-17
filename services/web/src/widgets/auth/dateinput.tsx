import React, { FC, forwardRef } from "react";
import { useFormContext } from "react-hook-form";

const DateInput: FC = () => {
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
