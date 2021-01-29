module.exports = {
    plugins: {
        tailwindcss: {},
        "postcss-preset-env": {
            features: {
                "focus-within-pseudo-class": false
            }
        },
        autoprefixer: {}
    }
};
