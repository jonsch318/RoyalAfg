import React, { FC } from "react";
import { Provider } from "../hooks/auth";
import { SnackbarProvider } from "notistack";

import "../../styles/globals.css";
import "../../styles/tailwind.css";

import { appWithTranslation } from "next-i18next";
import { AppProps } from "next/dist/next-server/lib/router/router";

const MyApp: FC<AppProps> = ({ Component, pageProps }) => {
    return (
        <Provider session={pageProps.session}>
            <SnackbarProvider>
                <div className="main-container">
                    <Component {...pageProps} />
                </div>
            </SnackbarProvider>
        </Provider>
    );
};

export default appWithTranslation(MyApp);
