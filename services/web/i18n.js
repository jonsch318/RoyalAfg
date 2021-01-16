const locales = require("./content/locale/index");

export const GetMessage = (defaultLocal, locale, pathname) => {
    const localeMessages = locales[locale.replace("-", "")];

    try {
        return localeMessages[pathname];
    } catch (error) {
        try {
            return locales[defaultLocal.replace("-", "")][pathname];
        } catch (error) {
            return null;
        }
    }
};
