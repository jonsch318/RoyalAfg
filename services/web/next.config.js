const { process } = require("postcss-preset-env");
const { i18n } = require("./next-i18next.config");

module.exports = {
    pwa: {
        dest: "public",
        disable: true
    },
    i18n,
    publicRuntimeConfig: {
        processEnv: process.env
    }
};
