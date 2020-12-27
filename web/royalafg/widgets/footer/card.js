import React from "react";
import shortid from "shortid";
import PropTypes from "prop-types";

const title = (title) => {
    return <div className="font-sans text-base font-medium">{title.toUpperCase()}</div>;
};

const content = (items) => {
    const listItems = React.Children.map(items, (child) => {
        return (
            <li key={shortid.generate()}>
                <style jsx>{`
                li:before {
                    content: "-";
                    margin-right: 0.5rem;
                }
            .
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

const FooterCard = (props) => {
    return (
        <div className="mb-4">
            {title(props.title)}
            {content(props.children)}
        </div>
    );
};

FooterCard.propTypes = {
    title: PropTypes.string,
    children: PropTypes.element.isRequired
};

export default FooterCard;
