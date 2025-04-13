import React from "react";
import { Image } from "@mantine/core";
import { Product } from "@/types";

interface CartItemProps {
  product: Product;
  quantity: number;
  onRemove: () => void;
}

const CartItem = ({ product, quantity, onRemove }: CartItemProps) => {
  const itemTotal = product.price * quantity;

  return (
    <li className="flex justify-between">
      <section className="space-y-1">
        <h2 className="text-[15px] font-medium leading-tight text-text-highlightDimmedDark">
          {product.name}
        </h2>
        <div className="flex items-center gap-[6px] text-[15px] ">
          <span className="text-text-highlightDimmed pr-4">{quantity}x</span>
          <span className="text-text-secondary">@${product.price.toFixed(2)}</span>
          <span className="text-text-secondary">
            ${itemTotal.toFixed(2)}
          </span>
        </div>
      </section>
      <button
        onClick={onRemove}
        className="text-[#6B6B6B] hover:text-[#4A4A4A] text-2xl leading-none mt-[-4px]"
        aria-label={`Remove ${product.name} from cart`}
      >
        <Image src="/close.svg" alt="Close" width={24} height={24} />
      </button>
    </li>
  );
};

export default React.memo(CartItem);