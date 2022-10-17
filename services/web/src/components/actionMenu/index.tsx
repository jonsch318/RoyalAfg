import React, { FC } from "react";

type ActionmenuProps = {
    children: React.ReactNode;
};

const ActionMenu: FC<ActionmenuProps> = ({ children }) => {
    return <div className="bg-white p-16 rounded-xl">{children}</div>;
};

export default ActionMenu;
