import type { Config } from "tailwindcss";
import daisyui from "daisyui"

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {

    },
    fontFamily: {
      'poppins': ['Poppins', 'sans-serif'],
      'lato': ['Lato', 'sans-serif'],
      'shantell': ['Shantell Sans', 'sans-serif'],
      'ubuntu': ['Ubutnu', 'sans-serif']
    }
  },
  plugins: [
    daisyui,
  ],
  daisyui: {
    themes: [
      "pastel",
      "sunset"
    ],
  },
  darkMode: ['class', '[data-theme="sunset"]']
};
export default config;
