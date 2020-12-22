import React from 'react';
import PropTypes from 'prop-types';
import { Provider } from 'react-redux';
import { useStore } from '../redux/store';
import '../styles/globals.css';
import '../styles/tailwind.css';

function MyApp({ Component, pageProps }) {
    const store = useStore(pageProps.initialReduxState);

    return (
        <Provider store={store}>
            <div className="main-container">
                <Component {...pageProps} />
            </div>
        </Provider>
    );
}

MyApp.propTypes = {
    Component: PropTypes.elementType.isRequired,
    pageProps: PropTypes.object
};

export default MyApp;
