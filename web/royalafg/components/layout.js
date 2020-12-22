import React from 'react';
import Header from '../widgets/header';
import Footer from '../widgets/footer';
import PropTypes from 'prop-types';

const Layout = (props) => {
    const foot = <Footer absolute={props.footerAbsolute} />;

    return (
        <div id="root" className="root">
            {!props.disableNav && <Header />}
            {props.enableAlternativNav && props.alternativNav}
            {props.children}
            {!props.disableFooter && foot}
        </div>
    );
};

Layout.propTypes = {
    footerAbsolute: PropTypes.bool,
    disableNav: PropTypes.bool,
    enableAlternativNav: PropTypes.bool,
    alternativNav: PropTypes.elementType,
    children: PropTypes.element,
    disableFooter: PropTypes.bool
};

export default Layout;
