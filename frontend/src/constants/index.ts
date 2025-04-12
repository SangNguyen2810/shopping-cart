export const API_BASE_URL = 'http://localhost:8080/api';

export const ENDPOINTS = {
  PRODUCTS: '/products',
  ORDERS: '/orders',
  DISCOUNT_VALIDATE: '/discounts/validate',
} as const;

export const QUERY_CONFIG = {
  RETRY_COUNT: 1,
  STALE_TIME: 5 * 60 * 1000,
  REFETCH_ON_WINDOW_FOCUS: false,
} as const;

export const NOTIFICATIONS = {
  SUCCESS: {
    TITLE: 'Success',
    COLOR: 'green',
  },
  ERROR: {
    TITLE: 'Error',
    COLOR: 'red',
  },
} as const; 

export const DEFAULT_CATEGORY = 'Dessert';
export const DEFAULT_IMAGE = 'https://placehold.co/400x300?text=Dessert';
