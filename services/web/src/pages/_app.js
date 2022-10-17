/* eslint-disable react/prop-types */
import React from "react";
import { Provider } from "../hooks/auth";
import { SnackbarProvider } from "notistack";

import "../../styles/globals.css";

import { appWithTranslation } from "next-i18next";

const MyApp = ({ Component, pageProps }) => (
    <Provider session={pageProps.session}>
        <SnackbarProvider>
            <div className="main-container">
                <Component {...pageProps} />
            </div>
        </SnackbarProvider>
    </Provider>
);

export default appWithTranslation(MyApp);
