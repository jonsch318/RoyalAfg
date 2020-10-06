import shortid from "shortid";

const title = (title) => {
    return (
        <div className="font-sans text-lg font-medium">
            {title.toUpperCase()}
        </div>
    );
}

const content = (items) => {
    const listItems = React.Children.map(items, child => {
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
        )
    })
    return (
        <div>

            <ul className="pl-2">
                {listItems}
            </ul>
        </div>
    );
}

const FooterCard = (props) => {

    return (
        <div>
            {title(props.title)}
            { content(props.children)}
        </div>
    )

}

export default FooterCard