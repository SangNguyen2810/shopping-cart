import { z } from 'zod';
import { fetchApi } from '@api/config';
import { ENDPOINTS, NOTIFICATIONS } from '@constants/index';
import { notifications } from '@mantine/notifications';
import { productSchema } from '@/types';

export const orderItemSchema = z.object({
  productId: z.string(),
  quantity: z.number().int().positive(),
});

export const orderSchema = z.object({
  items: z.array(orderItemSchema),
  discountCode: z.string().optional(),
});

const responseItemSchema = z.object({
  product: productSchema,
  quantity: z.number().int().positive(),
});

export const apiResponseSchema = z.object({
  orderId: z.string(),
  items: z.array(responseItemSchema),
});

export type OrderItem = z.infer<typeof orderItemSchema>;
export type Order = z.infer<typeof orderSchema>;
export type OrderResponseItem = z.infer<typeof responseItemSchema>;

export interface OrderConfirmationResult {
  orderId: string;
  items: OrderResponseItem[];
}

/**
 * Places an order with the provided items
 * @param items Array of cart items with product IDs and quantities
 * @param discountCode Optional discount code
 * @returns Promise with order ID and full product details
 */
export const placeOrder = async (items: OrderItem[], discountCode?: string): Promise<OrderConfirmationResult> => {
  try {
    const response = await fetchApi<unknown>(ENDPOINTS.ORDERS, {
      method: 'POST',
      body: JSON.stringify({ 
        items,
        discountCode
      }),
    });
    
    const parsedResponse = apiResponseSchema.parse(response);
    
    return {
      orderId: parsedResponse.orderId,
      items: parsedResponse.items
    };
  } catch (error) {
    console.error('Error placing order:', error);
    notifications.show({
      title: NOTIFICATIONS.ERROR.TITLE,
      message: error instanceof Error ? error.message : 'Failed to place order. Please try again.',
      color: NOTIFICATIONS.ERROR.COLOR,
    });
    throw error;
  }
}; 