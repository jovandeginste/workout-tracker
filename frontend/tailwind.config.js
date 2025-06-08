/** @type {import('tailwindcss').Config} */
const { addDynamicIconSelectors } = require("@iconify/tailwind");

module.exports = {
  content: ["../views/**/*.{html,js,go,templ}", "./src/**/*.{html,js}"],
  darkMode: "selector",
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
