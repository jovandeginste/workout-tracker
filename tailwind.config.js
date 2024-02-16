/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./{assets,views}/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [require("tailwind-fontawesome")],
};
