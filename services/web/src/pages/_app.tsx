import React, { FC } from "react";
import { IntlProvider } from "react-intl";
import { Provider } from "../hooks/auth";
import { SnackbarProvider } from "notistack";

import "../../styles/globals.css";
import "../../styles/tailwind.css";

import { useRouter } from "next/router";
import { GetMessage } from "../../i18n";
import { AppProps } from "next/dist/next-server/lib/router/router";

const MyApp: FC<AppProps> = ({ Component, pageProps }) => {
    const router = useRouter();
    const { locale, defaultLocale, pathname } = router;

    return (
        <IntlProvider locale={locale} defaultLocale={defaultLocale} messages={GetMessage(defaultLocale, locale, pathname)}>
            <Provider session={pageProps.session}>
                <SnackbarProvider>
                    <div className="main-container">
                        <Component {...pageProps} />
                    </div>
                </SnackbarProvider>
            </Provider>
        </IntlProvider>
    );
};

export default MyApp;
