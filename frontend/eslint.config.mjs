import withNuxt from "./.nuxt/eslint.config.mjs";

export default withNuxt({
  files: ["**/*.ts", "**/*.vue"],
  rules: {
    "no-console": "warn",
    "@typescript-eslint/no-explicit-any": "off",
    "@typescript-eslint/no-duplicate-enum-values": "off",
    "no-useless-escape": "off",
    "vue/html-self-closing": "off",
    "vue/valid-v-slot": [
      "error",
      {
        allowModifiers: true,
      },
    ],
  },
})
  .append
  // ...append other flat config items
  ()
  .prepend
  // ...prepend other flat config items before the base config
  ()
  // override a specific config item based on their name
  .override(
    "nuxt/typescript/rules", // specify the name of the target config, or index
    {
      rules: {
        // ...override the rules
        "@typescript-eslint/no-unsafe-assignment": "off",
      },
    },
  );
