import { z } from 'zod';
import { productSchema, Product } from '@/types';
import { fetchApi } from '@api/config';
import { ENDPOINTS } from '@constants/index';

const productsSchema = z.array(productSchema);

export const fetchProducts = async (): Promise<Product[]> => {
  const data = await fetchApi<unknown>(ENDPOINTS.PRODUCTS);
  
  try {
    const parsedData = productsSchema.parse(data);
    
    return parsedData.filter(product => {
      return (
        product.id !== undefined && 
        product.name !== undefined && 
        product.price !== undefined
      );
    });
  } catch (error) {
    console.error('Error parsing products data:', error);
    return [];
  }
};

export const fetchProduct = async (id: string): Promise<Product> => {
  const data = await fetchApi<unknown>(`${ENDPOINTS.PRODUCTS}/${id}`);
  
  try {
    return productSchema.parse(data);
  } catch (error) {
    console.error(`Error fetching product ${id}:`, error);
    throw new Error(`Product ${id} not found or invalid`);
  }
};