import React from "react";
import { useSelector } from "react-redux";
import Layout from "../components/layout";
import { useRouter } from "next/router";
import { useIntl } from "react-intl";
import { useSession } from "next-auth/client";

export default function Index() {
    const [session] = useSession();
    //const state = useSelector((state) => state.auth.isLoggedIn);
    const { formatMessage } = useIntl();
    const f = (id) => formatMessage({ id });
    const router = useRouter();
    const { locale, locales, defaultLocale } = router;

    console.log("session", session);

    return (
        <Layout footerAbsolute>
            <div>
                <h1>{f("header")}</h1>
                <a href="/about">About</a>
                <h1>is logged in [{session ? "signed in " + session.user.name : "not signed in"}]</h1>
                <button></button>
                <p>Current locale: {locale}</p>
                <p>Default locale: {defaultLocale}</p>
                <p>Configured locales: {JSON.stringify(locales)}</p>
            </div>
        </Layout>
    );
}
