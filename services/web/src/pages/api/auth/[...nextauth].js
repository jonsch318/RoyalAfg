import NextAuth from "next-auth";
import Providers from "next-auth/providers";

export default (req, res) => {
    console.log(res.body);
    return NextAuth(req, res, {
        providers: [
            Providers.Credentials({
                name: "Credentials",
                credentials: {
                    username: { label: "username", type: "text", placeholder: "username" },
                    password: { label: "password", type: "password", placeholder: "password" }
                },
                authorize: async (credentials) => {
                    // Add logic here to look up the user from the credentials supplied
                    let user = { id: 1, name: "J Smith", email: "jsmith@example.com", avatar: "http://localhost:3000/public/pb.png" };
                    if (user) {
                        // Any object returned will be saved in `user` property of the JWT
                        return user;
                    } else {
                        // If you return null or false then the credentials will be rejected
                        return null;
                        // You can also Reject this callback with an Error or with a URL:
                        // return Promise.reject(new Error('error message')) // Redirect to error page
                        // return Promise.reject('/path/to/redirect')        // Redirect to a URL
                    }
                }
            }),
            Providers.GitHub({
                clientId: process.env.NEXTAUTH_GITHUB_ID,
                clientSecret: process.env.NEXTAUTH_GITHUB_SECRET
            }),
            Providers.Google({
                clientId: process.env.NEXTAUTH_GOOGLE_ID,
                clientSecret: process.env.NEXTAUTH_GOOGLE_SECRET
            })
        ],

        session: {
            jwt: true
        },

        jwt: {
            encryption: true,
            secret: process.env.SECRET
        },

        callbacks: {
            // signIn: async (user, account, metadata) => {
            //     if (account.provider == "github") {
            //         const githubUser = {
            //             id: metadata.id,
            //             login: metadata.login,
            //             name: metadata.name,
            //             avatar: user.image
            //         };
            //         user.accessToken = "sss";
            //         return true;
            //     } else if (account.provider == "google") {
            //         return true;
            //     } else if (account.provider == "Credentials") {
            //         return true;
            //     }
            //     return false;
            // },
            // jwt: async (token, user) => {
            //     if (user) {
            //         token = { accessToken: user.accessToken };
            //     }
            //     return token;
            // },
            // session: async (session, token) => {
            //     session.accessToken = token.accessToken;
            //     return session;
            // }
        }
    });
};
