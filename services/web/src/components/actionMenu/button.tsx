import React, { FC } from "react";

type ActionMenuButtonProps = {
    children: React.ReactNode;
    onClick: React.MouseEventHandler<HTMLButtonElement>;
};

const ActionMenuButton: FC<ActionMenuButtonProps> = ({ children, onClick }) => {
    return (
        <button className="bg-gray-200 px-16 py-16 rounded-xl hover:bg-gray-300 outline-none hover:outline-none" onClick={onClick}>
            {children}
        </button>
    );
};

export default ActionMenuButton;
