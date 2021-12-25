import React, { FC } from "react";
import Header from "../widgets/header";
import Footer from "../widgets/footer";
type LayoutProps = {
    footerAbsolute?: boolean;
    disableNav?: boolean;
    enableAlternativeNav?: boolean;
    disableFooter?: boolean;
    headerAbsolute?: boolean;
    alternativeNav?: React.ReactNode;
};

const Layout: FC<LayoutProps> = (props) => {
    const foot = <Footer absolute={props.footerAbsolute} />;

    return (
        <div id="root" className="root h-full w-full">
            {!props.disableNav && <Header absolute={props.headerAbsolute} />}
            {props.enableAlternativeNav && props.alternativeNav}
            {props.children}
            {!props.disableFooter && foot}
        </div>
    );
};

Layout.defaultProps = {
    disableNav: false,
    enableAlternativeNav: false,
    disableFooter: false,
    headerAbsolute: false,
    footerAbsolute: false
};

export default Layout;
