import React from "react";
import PropTypes from "prop-types";
import { Provider } from "next-auth/client";
import { IntlProvider } from "react-intl";
import "../../styles/globals.css";
import "../../styles/tailwind.css";

import { useRouter } from "next/router";
import { GetMessage } from "../../i18n";

function MyApp({ Component, pageProps }) {
    //const store = useStore(pageProps.initialReduxState);
    const router = useRouter();
    const { locale, defaultLocale, pathname } = router;

    return (
        <IntlProvider locale={locale} defaultLocale={defaultLocale} messages={GetMessage(defaultLocale, locale, pathname)}>
            <Provider session={pageProps.session}>
                <div className="main-container">
                    <Component {...pageProps} />
                </div>
            </Provider>
        </IntlProvider>
    );
}

MyApp.propTypes = {
    Component: PropTypes.elementType.isRequired,
    pageProps: PropTypes.object
};

export default MyApp;
