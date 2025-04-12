import { z } from 'zod';
import { fetchApi } from '@api/config';
import { ENDPOINTS, NOTIFICATIONS } from '@constants/index';
import { notifications } from '@mantine/notifications';

// Schema for discount validation request
export const discountValidationRequestSchema = z.object({
  code: z.string(),
});

// Schema for discount validation response
export const discountValidationResponseSchema = z.object({
  valid: z.boolean(),
  discountRate: z.number().optional(),
  error: z.string().optional(),
});

export type DiscountValidationRequest = z.infer<typeof discountValidationRequestSchema>;
export type DiscountValidationResponse = z.infer<typeof discountValidationResponseSchema>;

/**
 * Validates a discount code with the backend
 * @param code The discount code to validate
 * @returns Promise with validation result
 */
export const validateDiscountCode = async (code: string): Promise<DiscountValidationResponse> => {
  try {
    const response = await fetchApi<unknown>(ENDPOINTS.DISCOUNT_VALIDATE, {
      method: 'POST',
      body: JSON.stringify({ code }),
    });
    
    return discountValidationResponseSchema.parse(response);
  } catch (error) {
    console.error('Error validating discount code:', error);
    notifications.show({
      title: NOTIFICATIONS.ERROR.TITLE,
      message: error instanceof Error ? error.message : 'Failed to validate discount code. Please try again.',
      color: NOTIFICATIONS.ERROR.COLOR,
    });
    return {
      valid: false,
      error: 'Failed to validate discount code',
    };
  }
}; 