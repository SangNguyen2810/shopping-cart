import React from 'react';
import { useQuery } from '@tanstack/react-query';
import { fetchProducts } from '@/api/product-api';
import ProductList from '@components/ProductList';
import Cart from '@components/Cart';
import type { Product } from '@/types';
import { QUERY_CONFIG } from '@constants/index';

import styles from './MainContent.module.css';

const MainContent: React.FC = () => {
  const {
    data: products,
    isLoading,
    error,
  } = useQuery<Product[]>({
    queryKey: ['products'],
    queryFn: fetchProducts,
    retry: QUERY_CONFIG.RETRY_COUNT,
    staleTime: QUERY_CONFIG.STALE_TIME,
    refetchOnWindowFocus: QUERY_CONFIG.REFETCH_ON_WINDOW_FOCUS,
  });

  if (isLoading) return <h2>Loading products...</h2>;
  if (error) {
    console.error('Error fetching products:', error);
    return <h2>Error fetching products. Please try again later.</h2>;
  }
  if (!products) return <h2>No products available.</h2>;

  return (
    <main className={styles.mainContent}>
      <h1 className="text-3xl font-bold mb-[30px]">Desserts</h1>

      <section className={styles.content}>
        <ProductList products={products} />
        <Cart products={products} />
      </section>
    </main>
  );
};

export default MainContent;
