import { z } from 'zod';

const imageSchema = z.object({
  thumbnail: z.string(),
  mobile: z.string(),
  tablet: z.string(),
  desktop: z.string(),
});

export const productSchema = z.object({
  id: z.string(),
  name: z.string(),
  price: z.number(),
  image: imageSchema.optional(),
  category: z.string().optional(),
});

export type Product = z.infer<typeof productSchema>;

export const orderConfirmationSchema = z.object({
  orderId: z.string(),
  subtotal: z.number(),
  discount: z.number(),
  total: z.number(),
});

export type OrderConfirmation = z.infer<typeof orderConfirmationSchema>;