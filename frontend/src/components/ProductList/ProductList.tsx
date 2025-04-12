import React from 'react';
import { Product } from '../../types';
import ProductItem from '@components/ProductItem';
import styles from './ProductList.module.css';
interface ProductListProps {
  products: Product[];
}

const ProductList: React.FC<ProductListProps> = ({ products }) => {

  return (
    <section className={styles.productList}>
      {products.map((product) => (
        <ProductItem key={product.id} product={product} />
      ))}
    </section>
  );
};

export default ProductList;