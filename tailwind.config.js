/** @type {import('tailwindcss').Config} */
const { addDynamicIconSelectors } = require("@iconify/tailwind");

module.exports = {
  content: [
    "./internal/views/assets/**/*.{html,js}",
    "./internal/views/**/*.{html,js}",
    "./pkg/templatehelpers/*.go",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    // Iconify plugin
    addDynamicIconSelectors(),
  ],
  safelist: [
    {
      pattern: /text-(green|rose)-500/,
    },
  ],
  variants: {
    customPlugin: ["responsive", "hover"],
  },
};
