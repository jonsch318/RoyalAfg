module.exports = {
    root: true, // Make sure eslint picks up the config at the root of the directory
    parserOptions: {
        ecmaVersion: 2020, // Use the latest ecmascript standard
        sourceType: "module", // Allows using import/export statements
        ecmaFeatures: {
            jsx: true, // Enable JSX since we're using React
            modules: true
        }
    },
    settings: {
        react: {
            version: "detect" // Automatically detect the react version
        }
    },
    env: {
        es6: true,
        browser: true, // Enables browser globals like window and document
        amd: true, // Enables require() and define() as global variables as per the amd spec.
        node: true // Enables Node.js global variables and Node.js scoping.
    },
    extends: [
        "eslint:recommended",
        "plugin:react/recommended",
        "plugin:jsx-a11y/recommended",
        "plugin:prettier/recommended" // Make this the last element so prettier config overrides other formatting rules
    ],
    rules: {
        "prettier/prettier": ["error", {}, { usePrettierrc: true }], // Use our .prettierrc file as source,
        quotes: [2, "double"]
    },
    overrides: [
        {
            files: ["**/*.ts", "**/*.tsx"],
            env: { browser: true, es6: true, node: true },
            extends: [
                "eslint:recommended",
                "plugin:react/recommended",
                "plugin:jsx-a11y/recommended",
                "plugin:prettier/recommended",
                "plugin:@typescript-eslint/eslint-recommended",
                "plugin:@typescript-eslint/recommended"
            ],
            globals: { Atomics: "readonly", SharedArrayBuffer: "readonly" },
            parser: "@typescript-eslint/parser",
            plugins: ["@typescript-eslint"],
            rules: {
                "prettier/prettier": ["error", {}, { usePrettierrc: true }], // Use our .prettierrc file as source,
                quotes: [2, "double"],
                "react/prop-types": "off"
            }
        }
    ]
};
