const { process } = require("postcss-preset-env");
const withPWA = require("next-pwa");

module.exports = withPWA({
    pwa: {
        dest: "public"
    },
    i18n: {
        locales: ["en-US", "de-DE"],
        defaultLocale: "en-US",
        localeDetection: false
    },
    publicRuntimeConfig: {
        processEnv: process.env
    }
});
