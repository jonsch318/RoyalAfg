import React from "react";
import Layout from "../components/layout";
import { useRouter } from "next/router";
import { useIntl } from "react-intl";
import { useSession } from "../hooks/auth";
import Head from "next/head";

export default function Index() {
    const [session] = useSession();
    const { formatMessage } = useIntl();
    const f = (id) => formatMessage({ id });
    const router = useRouter();
    const { locale, locales, defaultLocale } = router;

    return (
        <>
            <Head>
                <title>Royalafg Casino</title>
                <meta name="description" content="The Royalafg Casino is a online casino built for a special lerning archivement"></meta>
            </Head>
            <Layout footerAbsolute>
                <div>
                    <h1>{f("header")}</h1>
                    <h1>is logged in [{session ? "signed in " + session.user.name : "not signed in"}]</h1>
                    <p>Current locale: {locale}</p>
                    <p>Default locale: {defaultLocale}</p>
                    <p>Configured locales: {JSON.stringify(locales)}</p>
                </div>
            </Layout>
        </>
    );
}
