/** @type {import('tailwindcss').Config} */
const { addDynamicIconSelectors } = require("@iconify/tailwind");

module.exports = {
  content: ["../views/**/*.{html,js,ts,go,templ}", "./src/**/*.{html,js,ts}"],
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
    "md:table-cell",
    "lg:table-cell",
    "xl:table-cell",
  ],
  variants: {
    customPlugin: ["responsive", "hover"],
  },
};
