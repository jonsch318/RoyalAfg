import React, { FC } from "react";
import { Provider } from "../hooks/auth";
import { SnackbarProvider } from "notistack";

import "../../styles/tailwind.css";
//import "../../styles/globals.css";

import { appWithTranslation } from "next-i18next";
import { AppProps } from "next/app";

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
