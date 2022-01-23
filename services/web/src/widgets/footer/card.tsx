import React, { FC } from "react";
import shortid from "shortid";

type TitleProps = {
    title: string;
};
const Title: FC<TitleProps> = ({ title }) => {
    return <div className="font-sans text-base font-medium">{title.toUpperCase()}</div>;
};

const Content: FC = ({ children }) => {
    const listItems = React.Children.map(children, (child) => {
        return (
            <li key={shortid.generate()}>
                <style jsx>{`
                    li:before {
                        content: "-";
                        margin-right: 0.5rem;
                    }
                `}</style>
                {child}
            </li>
        );
    });
    return (
        <div>
            <ul className="pl-2">{listItems}</ul>
        </div>
    );
};

type FooterCardProps = {
    title?: string;
};

const FooterCard: FC<FooterCardProps> = ({ title, children }) => {
    return (
        <div className="mb-4">
            <Title title={title} />
            <Content>{children}</Content>
        </div>
    );
};

export default FooterCard;
