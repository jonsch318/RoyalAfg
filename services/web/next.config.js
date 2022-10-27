const { i18n } = require("./next-i18next.config");
const withTM = require("next-transpile-modules")(["three"]);

module.exports = withTM({
    i18n: i18n
});
