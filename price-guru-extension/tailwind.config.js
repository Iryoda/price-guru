/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')

export default {
    content: ['./src/**/*.{html,js,ts,jsx,tsx}'],
    theme: {
        fontFamily: {
            inter: ['Inter', 'sans-serif', ...defaultTheme.fontFamily.sans],
        },
        extend: {},
    },
    plugins: [],
}
