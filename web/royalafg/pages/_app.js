import React from 'react';
import PropTypes from 'prop-types';

function MyApp({ Component, pageProps }) {
    return (
        <div className="main-container">
            <Component {...pageProps} />
        </div>
    );
}

MyApp.propTypes = {
    Component: PropTypes.elementType.isRequired,
    pageProps: PropTypes.object
};

export default MyApp;
