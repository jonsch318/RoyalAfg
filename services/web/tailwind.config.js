const colors = require("tailwindcss/colors");

module.exports = {
    content: ["./src/**/*.html", "./src/**/*.jsx", "./src/**/*.js", "./src/**/*.tsx", "./src/**/*.ts"],
    future: {
        removeDeprecatedGapUtilities: true,
        purgeLayersByDefault: true
    },
    theme: {
        fontFamily: { sans: ["Poppins", "sans-serif"] },
        colors: {
            transparent: "transparent",
            current: "currentColor",
            black: colors.black,
            white: colors.white,
            gray: {
                50: "#fafafa",
                100: "#f5f5f5",
                200: "#efefef",
                300: "#e0e0e0",
                400: "#d0d0d0",
                500: "#737373",
                600: "#525252",
                700: "#404040",
                800: "#262626",
                900: "#171717"
            },
            indigo: colors.indigo,
            red: colors.rose,
            blue: colors.blue,
            yellow: colors.amber
        },
        extend: {
            width: {
                fit: "fit-content",
                min: "min-content"
            },
            minWidth: {
                36: "9rem",
                40: "10rem",
                44: "11rem",
                48: "12rem",
                52: "13rem",
                56: "14rem",
                60: "15rem"
            }
        }
    },
    variants: {
        variantOrder: [
            "first",
            "last",
            "odd",
            "even",
            "visited",
            "checked",
            "group-hover",
            "group-focus",
            "focus-within",
            "hover",
            "focus",
            "focus-visible",
            "active",
            "disabled"
        ]
    }
};
