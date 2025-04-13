// tailwind.config.js
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          dark: '#1a1a1a',
          darker: '#000000',
          background: '#FCF8F5'
        },
        text: {
          highlight: 'var(--color-text-highlight)',
          highlightDimmed: 'var(--color-text-highlight-dimmed)',
          secondary: 'var(--color-text-secondary)',
          muted: 'var(--color-text-muted)',
          highlightDimmedDark: 'var(--color-text-highlight-dimmed-dark)',
          button: 'var(--color-button)',
          gray: 'var(--color-text-gray)',
          grayDark: '#837975',
        },
        border: {
          light: 'var(--color-border-light)'
        },
        bg: {
          accent: 'var(--color-bg-accent)'
        },
        button: {
          text: '#EFBAA2',
          success: '#4CAF50',
          successHover: '#3d8b40',
          primary: '#C63B0F',
          hover: '#FFF5F1'
        }
      }
    },
  },
  plugins: [],
};