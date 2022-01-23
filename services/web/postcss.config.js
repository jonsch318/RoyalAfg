/*module.exports = {
    plugins: [
        require("postcss-import"),
        require("tailwindcss"),
        "postcss-preset-env"({
            features: {
                "focus-within-pseudo-class": false
            }
        }),
        require("autoprefixer")
    ]
};*/

module.exports = {
    plugins: {
        "postcss-import": {},
        tailwindcss: {},
        "postcss-preset-env": {
            features: {
                "focus-within-pseudo-class": false
            }
        },
        autoprefixer: {}
    }
};
