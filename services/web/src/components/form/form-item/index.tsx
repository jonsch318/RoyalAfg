import React, { FC } from "react";

type FormItemProps = {
    children: React.ReactNode;
};

const FormItem: FC<FormItemProps> = ({ children }) => {
    return <div className="mb-6 font-sans text-lg font-medium">{children}</div>;
};

export default FormItem;
