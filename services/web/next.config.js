const { process } = require("postcss-preset-env");

module.exports = {
    //pwa: {
    //    dest: "public",
    //    disable: true
    //},
    i18n: {
        locales: ["en-US", "de-DE"],
        defaultLocale: "en-US",
        localeDetection: false
    },
    publicRuntimeConfig: {
        processEnv: process.env
    }
};
