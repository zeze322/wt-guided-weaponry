/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./**/*.html", "./**/*.templ", "./**/*.go","./node_modules/flowbite/**/*.js"],
  safelist: [],
  plugins: [require('daisyui'), require('flowbite/plugin')],
  daisyui: {themes: ["dark", "black", "dim", "nord"]},
}