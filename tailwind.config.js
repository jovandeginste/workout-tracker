/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./{assets,views}/**/*.{html,js}", "./pkg/templatehelpers/*.go"],
  theme: {
    extend: {},
  },
  plugins: [require("tailwind-fontawesome")],
  safelist: [
    {
      pattern: /text-(green|rose)-500/,
    },
  ],
  variants: {
    customPlugin: ["responsive", "hover"],
  },
};
