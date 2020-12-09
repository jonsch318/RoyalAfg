import Header from "../widgets/header";
import Footer from "../widgets/footer";

const Layout = (props) => {

    const foot = (<Footer absolute={props.footerAbsolute} />);

    return (
        <div id="root" className="root">
            {!props.disableNav && <Header />}
            {props.enableAlternativNav && props.alternativNav}
            {props.children}
            {!props.disableFooter && foot}
        </div>
    );
}

export default Layout