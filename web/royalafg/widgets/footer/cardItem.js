const FooterCardItem = (props) => {
    if (props.href) {
        return (
            <a href={props.href} className="font-sans font-thin text-sm hover:opacity-75 transition-opacity duration-100 ease-out">
                {props.children}
            </a>
        );
    }
    return (
        <span className="font-sans font-thin text-sm">{props.children}</span>
    )

}

export default FooterCardItem