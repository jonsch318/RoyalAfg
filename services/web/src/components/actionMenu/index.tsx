import React, { FC } from "react";

const ActionMenu: FC = ({ children }) => {
    return <div className="bg-white p-16 rounded-xl">{children}</div>;
};

export default ActionMenu;
