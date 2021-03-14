import React from "react";
import Document, { Html, Head, Main, NextScript } from "next/document";

class MyDocument extends Document {
    static async getInitialProps(ctx) {
        const initialProps = await Document.getInitialProps(ctx);
        return { ...initialProps };
    }

    render() {
        return (
            <Html>
                <Head>
                    <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png" />
                    <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png" />
                    <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png" />
                    <link rel="manifest" href="/site.webmanifest" />
                    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#5bbad5" />
                    <meta name="apple-mobile-web-app-title" content="Royalafg" />
                    <meta name="application-name" content="Royalafg" />
                    <meta name="msapplication-TileColor" content="#da532c" />
                    <meta name="theme-color" content="#ffffff" />
                    <link rel="shortcut icon" type="image/x-icon" href="/favicon.ico" />
                    <meta charSet="utf-8" />
                    <meta name="keywords" content="Royalafg, Online Casino" />
                    <meta name="author" content="Jonas Schneider" />
                    <meta name="copyright" content="Jonas Schneider 2021" />
                    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap" />
                    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" />

                    <meta name="application-name" content="Royalafg" />
                    <meta name="apple-mobile-web-app-capable" content="yes" />
                    <meta name="apple-mobile-web-app-status-bar-style" content="default" />
                    <meta name="apple-mobile-web-app-title" content="Roaylafg" />
                    <meta name="description" content="The Royalafg Online Casino" />
                    <meta name="format-detection" content="telephone=no" />
                    <meta name="mobile-web-app-capable" content="yes" />
                    <meta name="msapplication-config" content="/browserconfig.xml" />
                    <meta name="msapplication-tap-highlight" content="no" />

                    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500" />

                    <meta name="twitter:card" content="summary" />
                    <meta name="twitter:url" content="https://royalafg.games" />
                    <meta name="twitter:title" content="Royalafg" />
                    <meta name="twitter:description" content="The Royalafg Online Casino" />
                    <meta name="twitter:image" content="https://royalafg.games/android-chrome-192x192.png" />
                    <meta name="twitter:creator" content="@JohnnyS318" />
                    <meta property="og:type" content="website" />
                    <meta property="og:title" content="Royalafg" />
                    <meta property="og:description" content="The Royalafg Online Casino" />
                    <meta property="og:site_name" content="Royalafg" />
                    <meta property="og:url" content="https://royalafg.games" />
                    <meta property="og:image" content="https://royalafg.games/apple-touch-icon.png" />
                </Head>
                <body>
                    <Main />
                    <NextScript />
                </body>
            </Html>
        );
    }
}

export default MyDocument;
