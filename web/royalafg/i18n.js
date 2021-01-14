const NextI18Next = require("next-i18next").default;
require("next/serviceConfig").default().publicRuntimeConfig;
const path = require("path");

module.exports = new NextI18Next({
    otherLanguages: ["de"],
    localePath: path.resolve("./public/static/locales")
});
