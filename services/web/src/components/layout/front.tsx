import React, { FC } from "react";

type FrontProps = {
    children: React.ReactNode;
};

const Front: FC<FrontProps> = ({ children }) => {
    return <div className="bg-gray-200 md:px-10 py-28 font-sans text-5xl font-semibold text-center">{children}</div>;
};

export default Front;
