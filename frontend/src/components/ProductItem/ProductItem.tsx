import React from 'react';
import { Image, Button, Group } from '@mantine/core';
import { Product } from '@/types';
import useCartStore from '@/store/cartStore';
import styles from './ProductItem.module.css';
import { getFallbackImage } from '@/utils/getFallbackImage';
import { DEFAULT_CATEGORY } from '@/constants';

interface ProductItemProps {
  product: Product;
}

const ProductItem = ({ product }: ProductItemProps) => {
  const addItem = useCartStore((state) => state.addItem);
  const removeItem = useCartStore((state) => state.removeItem);
  const items = useCartStore((state) => state.items);

  const itemCount = items.find((item) => item.productId === product.id)?.quantity ?? 0;
  const showCounter = itemCount > 0;

  const handleAdd = () => {
    addItem(product.id);
  };

  const handleRemove = () => {
    removeItem(product.id);
  };

  return (
    <article className="flex flex-col">
      <figure className="w-full mb-2 relative">
        <picture className={`rounded-lg block ${showCounter ? 'border-2 border-red-500' : ''}`}>
          {product.image && (
            <>
              <source media="(min-width: 1024px)" srcSet={product.image.desktop} />
              <source media="(min-width: 768px)" srcSet={product.image.tablet} />
              <source media="(min-width: 375px)" srcSet={product.image.mobile} />
            </>
          )}
          <Image
            src={product.image?.desktop}
            alt={product.name}
            fit="cover"
            fallbackSrc={getFallbackImage(product.category)}
            className="!rounded-lg"
          />
        </picture>
        <section className="absolute -bottom-6 left-0 right-0 grid justify-center">
          {!showCounter ? (
            <Button
              variant="default"
              onClick={handleAdd}
              className={styles.button}
              classNames={{ label: 'flex gap-x-[9px]' }}
            >
              <Image src="/cart.svg" alt="Cart" width={24} height={24} />
              <h3 className="text-sm font-medium">Add to Cart</h3>
            </Button>
          ) : (
            <Group className={styles.counter}>
              <Button variant="default" onClick={handleRemove} className={styles.counterButton}>
                -
              </Button>
              <p className="text-sm font-medium w-8 text-center content-center text-[#EAB096]">
                {itemCount}
              </p>
              <Button variant="default" onClick={handleAdd} className={styles.counterButton}>
                +
              </Button>
            </Group>
          )}
        </section>
      </figure>

      <footer className="text-left mt-10">
        <p className="text-xs font-medium mb-[9px] text-[#ACA09C]">
          {product.category ?? DEFAULT_CATEGORY}
        </p>
        <p className="text-sm font-medium mb-2">{product.name}</p>
        <p className="text-sm mb-1 text-[#BC7863]">${product.price.toFixed(2)}</p>
      </footer>
    </article>
  );
};

export default React.memo(ProductItem);
