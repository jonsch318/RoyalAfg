import React from 'react';
import { useSelector } from 'react-redux';
import Layout from '../components/layout';

export default function Index() {
    const state = useSelector((state) => state.auth.isLoggedIn);

    return (
        <Layout footerAbsolute>
            <div>
                <h1>Hello</h1>
                <a href="/about">About</a>
                <h1>is logged in [{state ? 'signed in' : 'not signed in'}]</h1>
                <button>Sign in</button>
            </div>
        </Layout>
    );
}
