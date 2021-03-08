import { tokens } from "./token";
import { GetServerSidePropsContext } from "next";
import Cookies from "cookies";

const getCSRF = async (ctx: GetServerSidePropsContext): Promise<string> => {
    const csrfSecret = await tokens.secret();
    const csrfToken = tokens.create(csrfSecret);
    const cookies = new Cookies(ctx.req, ctx.res);
    cookies.set("xcsrf", csrfSecret + ":" + csrfToken, {
        hostOnly: true,
        overwrite: true
    });

    return csrfToken;
};

export { getCSRF };
