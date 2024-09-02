/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/views/assets/**/*.{html,js}",
    "./internal/views/**/*.{html,js}",
    "./pkg/templatehelpers/*.go",
  ],
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
