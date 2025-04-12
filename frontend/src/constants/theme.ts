import { createTheme } from '@mantine/core';

export const COLORS = {
  PRIMARY: {
    LIGHTEST: '#FFF0F3',
    LIGHT: '#FFE0E9',
    BASE: '#FF4D6D',
    MEDIUM: '#FF748C',
    DARK: '#C83B0E',
    DARKER: '#B34727',
    DARKEST: '#A61E1E',
  },
  TEXT: {
    PRIMARY: '#333333',
    SECONDARY: '#6B6B6B',
    MUTED: '#999999',
    HIGHLIGHT: '#C66C50',
  },
  BACKGROUND: {
    LIGHT: '#FCF8F5',
    CARD: '#FFFFFF',
    ACCENT: '#F8F8F8',
    HIGHLIGHT: '#FFF8E7',
  },
  BORDER: {
    LIGHT: '#E8E8E8',
    MEDIUM: '#D1D1D1',
  }
};

export const mantineTheme = createTheme({
  primaryColor: 'rose',
  colors: {
    rose: [
      COLORS.PRIMARY.LIGHTEST, // 0
      COLORS.PRIMARY.LIGHT,    // 1
      '#FFD1DF',               // 2
      '#FFC2D4',               // 3
      '#FFB3CA',               // 4
      COLORS.PRIMARY.MEDIUM,   // 5
      COLORS.PRIMARY.BASE,     // 6
      COLORS.PRIMARY.DARK,     // 7 - Primary shade
      COLORS.PRIMARY.DARKER,   // 8
      COLORS.PRIMARY.DARKEST,  // 9
    ],
  },
  defaultRadius: 'md',
  primaryShade: 7,
  components: {
    Button: {
      defaultProps: {
        color: 'rose',
      },
      styles: {
        root: {
          transition: 'all 0.2s ease-in-out',
          '&:hover': {
            transform: 'translateY(-2px)',
            boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
          },
        },
      },
    },
    Modal: {
      defaultProps: {
        centered: true,
        radius: 'md',
      },
    },
    Paper: {
      defaultProps: {
        radius: 'md',
        shadow: 'sm',
      },
    },
    Text: {
      defaultProps: {
        color: COLORS.TEXT.PRIMARY,
      },
    },
  },
});

export default mantineTheme; 