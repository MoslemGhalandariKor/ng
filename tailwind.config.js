/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
     "./node_modules/flowbite/**/*.js",
     "./templates/**/*.templ",
     "./templates/**/*.go",
     "./templates/**/**/*.templ",
     "./templates/**/**/*.go",
     "./templates/**/**/**/*.templ",
     "./templates/**/**/**/*.go",
     "./templates/**/**/**/**/*.templ",
     "./templates/**/**/**/**/*.go",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('flowbite/plugin'),
    require('@tailwindcss/forms')({
          strategy: 'base', // only generate global styles

      }),
  ],
}

