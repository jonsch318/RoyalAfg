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
