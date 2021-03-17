import React, { FC, useState } from "react";
import Layout from "../../components/layout";
import Front from "../../components/layout/front";
import ActionMenu from "../../components/actionMenu";
import ActionMenuLink from "../../components/actionMenu/link";
import { useRouter } from "next/router";
import { useSession, refreshSession } from "../../hooks/auth";
import { GetServerSideProps } from "next";
import { getCSRF } from "../../hooks/auth/csrf";

export const getServerSideProps: GetServerSideProps = async (context) => {
    const csrf = await getCSRF(context);

    try {
        const res = await fetch(process.env.USER_HOST ? `${process.env.USER_HOST}/api/user` : "/api/user", {
            method: "GET",
            headers: {
                cookie: context.req.headers.cookie ?? ""
            }
        });
        const user = await res.json();
        return {
            props: {
                csrf: csrf,
                user: user
            }
        };
    } catch (e) {
        return {
            props: {
                csrf: csrf,
                user: {}
            }
        };
    }
)
};

type AccountProps = {
    csrf: string;
    user: {
        id: string;
        fullName: string;
        username: string;
        email: string;
    };
};

const Account: FC<AccountProps> = ({ csrf, user }) => {
    const [session, loading] = useSession();
    const router = useRouter();

    const [u, setUser] = useState({
        fullName: user.fullName,
        username: user.username,
        email: user.email
    });

    if (loading) return null;

    if (!loading && !session) {
        router.push("/auth/login").then();
        return <div>Not signed in... redirecting</div>;
    }

    const onSubmit = (e) => {
        e.preventDefault();
        fetch("/api/user", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "X-CSRF-Token": csrf
            },
            body: JSON.stringify({
                username: u.username,
                email: u.email,
                fullName: u.fullName
            })
        })
            .then((res) => {
                if (!res.ok) {
                    console.log("Error: ", res.status);
                } else {
                    console.log("Updated User");
                }
                refreshSession();
                //updating auth session
            })
            .catch((err) => {
                console.error("Error", err);
            });
    };

    return (
        <Layout>
            <div>
                <Front>{"Your Account " + u.fullName}</Front>
                <div className="px-10 pb-10 bg-gray-200">
                    <ActionMenu>
                        <ActionMenuLink href="/wallet">My Wallet</ActionMenuLink>
                    </ActionMenu>
                </div>
                {session && session.user ? session.user.name : <></>}
            </div>
            <div>
                <form onSubmit={onSubmit}>
                    <div className="px-10 pb-10 bg-gray-200">
                        <ActionMenu>
                            <div className="grid grid-cols-2 mx-10 my-5">
                                <label htmlFor="username">Username: </label>
                                <input
                                    type="text"
                                    name="username"
                                    value={u.username}
                                    onChange={(e) => {
                                        setUser({ ...u, username: e.target.value });
                                    }}
                                    className="bg-gray-300 px-3 py-1 rounded outline-none"
                                />
                            </div>
                            <div className="grid grid-cols-2 mx-10 my-5">
                                <label htmlFor="username">Email: </label>
                                <input
                                    type="email"
                                    name="email"
                                    value={u.email}
                                    onChange={(e) => {
                                        setUser({ ...u, email: e.target.value });
                                    }}
                                    className="bg-gray-300 px-3 py-1 rounded outline-none"
                                />
                            </div>
                            <div className="grid grid-cols-2 mx-10 my-5">
                                <label htmlFor="username">Full Name: </label>
                                <input
                                    type="text"
                                    name="fullName"
                                    value={u.fullName}
                                    onChange={(e) => {
                                        setUser({ ...u, fullName: e.target.value });
                                    }}
                                    className="bg-gray-300 px-3 py-1 rounded outline-none"
                                />
                            </div>
                            <div className="grid justify-end mt-14 mx-10">
                                <button type="submit" className="bg-black text-white px-8 py-1 rounded ">
                                    Save
                                </button>
                            </div>
                        </ActionMenu>
                    </div>
                </form>
            </div>
        </Layout>
    );
};

export default Account;
